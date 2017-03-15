package kmi

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/lib/pq"
)

// Folder is a type used to abstract file hierarchie
type Folder struct {
	files   map[string]*[]byte
	folders FolderMap
}

// FolderMap is a map mapping a string to a pointer of type folder
type FolderMap map[string]*Folder

func (f FolderMap) walk(dir string, create bool) (*Folder, error) {
	var (
		folder    *Folder
		folderMap FolderMap
	)
	dirs := strings.Split(path.Clean(dir), "/")

	folderMap = f
	dirCount := len(dirs) - 1
	for i, dir := range dirs {
		_, exists := folderMap[dir]
		if !exists {
			if create {
				folderMap[dir] = &Folder{
					files:   make(map[string]*[]byte),
					folders: make(FolderMap),
				}
			} else {
				return nil, fmt.Errorf("folder %s does not exist", dir)
			}
		}
		if i == dirCount {
			folder = folderMap[dir]
		} else {
			folderMap = folderMap[dir].folders
		}
	}

	return folder, nil
}

// AddFile adds a file in the given path with the given name to the Content
func (f FolderMap) AddFile(p string, data *[]byte) {
	dir, name := path.Split(p)
	folder, _ := f.walk(dir, true)
	folder.files[name] = data
}

// GetFile get a file in path p or an error
func (f FolderMap) GetFile(p string) (*[]byte, error) {
	dir, name := path.Split(p)
	folder, err := f.walk(dir, false)
	if err != nil {
		return nil, err
	}

	file, ok := folder.files[name]
	if !ok {
		return nil, fmt.Errorf("file %s does not exist in %s", name, dir)
	}

	return file, nil
}

// Content is a struct type which may hold a byte array for every file which can be in the kmi file
type Content struct {
	Module     []byte
	modulePath string
	folders    FolderMap
}

// AddFile add a file to the FolderMap and set Module if applicable
func (c *Content) AddFile(dir, file string, data *[]byte) {
	c.folders.AddFile(path.Join(dir, file), data)
	if file == "module.json" {
		c.Module = *data
		c.modulePath = dir
	}
}

// GetFile add a file to the FolderMap and set Module if applicable
func (c *Content) GetFile(p string) (*[]byte, error) {
	dir, _ := path.Split(p)
	dir = path.Clean(dir)
	if dir == "" || dir == "." {
		p = path.Join(c.modulePath, p)
	}
	return c.folders.GetFile(p)
}

// NewContent initializes a new Content instance
func NewContent() *Content {
	return &Content{
		folders: make(FolderMap),
	}
}

// Extract is used to get the data from a kmi tar ball
func Extract(src string, k *Content) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	tarReader := tar.NewReader(f)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		data := make([]byte, header.Size)
		_, err = tarReader.Read(data)
		if err != nil {
			return err
		}
		dir, file := path.Split(header.Name)
		k.AddFile(dir, file, &data)
	}
	return nil
}

type moduleJSON struct {
	Name        string
	Version     string
	Description string
	Type        float64
	Imports     interface{}
	Frontend    interface{}
	Interfaces  interface{}
	Env         interface{}
	Cmd         interface{}
}

// ChooseSource fills src with outsrc if src is not the expected data kind
func ChooseSource(src interface{}, outsrc *Content, kind reflect.Kind, name string) error {
	v := reflect.ValueOf(src).Elem().Elem()
	k := v.Kind()
	if k == reflect.String {
		data, err := outsrc.GetFile(v.String())
		if err != nil {
			return err
		}

		err = json.Unmarshal(*data, src)
		if err != nil {
			return err
		}
		v = reflect.ValueOf(src).Elem().Elem()
		k = v.Kind()
	}

	if k != kind {
		return fmt.Errorf("%s malformatted", name)
	}

	return nil
}

// ExtractStringMap creates a map[string]interface{} based on a reflect.Value of Kind map,
// restriction can be used to exclude specific Kinds from beeing the value to a key in the map
func ExtractStringMap(value reflect.Value, dst map[string]interface{}, restriction map[reflect.Kind]bool) error {
	if value.Kind() != reflect.Map {
		return fmt.Errorf("src aint no map bra")
	}
	keys := value.MapKeys()
	for _, key := range keys {
		element := value.MapIndex(key).Elem()
		switch element.Kind() {
		case reflect.String:
			if restriction[reflect.String] {
				return fmt.Errorf("unexpected string")
			}
			dst[key.String()] = element.String()
		case reflect.Float64:
			if restriction[reflect.Int] {
				return fmt.Errorf("unexpected number")
			}
			dst[key.String()] = int(element.Float())
		default:
			return fmt.Errorf("unexpected %s", element.Kind().String())
		}
	}
	return nil
}

// GetStringMap takes a src and an outsrc, a destination map and a restriction map, decides which source to use to fill
// the destination map
func GetStringMap(src interface{}, outsrc *Content, dst map[string]interface{}, name string, restriction map[reflect.Kind]bool) error {
	err := ChooseSource(&src, outsrc, reflect.Map, name)
	if err != nil {
		return err
	}

	err = ExtractStringMap(reflect.ValueOf(src), dst, restriction)
	if err != nil {
		return err
	}
	return nil
}

// GetStringSlice takes a src and an outsrc and a destination slice, decides which source to use and to fill
// the destination slice
func GetStringSlice(src interface{}, outsrc *Content, dst *pq.StringArray, name string) error {
	err := ChooseSource(&src, outsrc, reflect.Slice, name)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(src)
	len := value.Len()

	for i := 0; i < len; i++ {
		element := value.Index(i).Elem()
		if element.Kind() != reflect.String {
			return fmt.Errorf("unexpected %T", element)
		}
		*dst = append(*dst, element.String())
	}

	return nil
}

// GetFrontend takes a src and an outsrc and a destination frontendModule slice, decides which source to use and extracts
// the frontendModule specific information to fill the destination slice
func GetFrontend(src interface{}, outsrc *Content, dst *FrontendArray) error {
	err := ChooseSource(&src, outsrc, reflect.Slice, "frontend")
	if err != nil {
		return err
	}

	value := reflect.ValueOf(src)
	len := value.Len()
	for i := 0; i < len; i++ {
		fm := frontendModule{}
		module := value.Index(i).Elem()
		keys := module.MapKeys()
		for i, key := range keys {
			switch key.String() {
			case "template":
				tpl := module.MapIndex(key).Elem()
				if tpl.Kind() != reflect.String {
					return fmt.Errorf("template is not of type string in module %d", i)
				}
				fm.template = tpl.String()
			case "parameters":
				params := module.MapIndex(key).Elem()
				if params.Kind() != reflect.Map {
					return fmt.Errorf("parameters are no json in module %d", i)
				}
				fm.parameters = make(map[string]interface{})
				err = ExtractStringMap(params, fm.parameters, nil)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unsupported property %s", key.String())
			}
		}
		*dst = append(*dst, fm)
	}

	return nil
}

// GetData is used to fill a KMI struct based on a Content struct
func GetData(kC *Content, k *KMI) error {
	m := &moduleJSON{}
	err := json.Unmarshal(kC.Module, m)
	if err != nil {
		return err
	}

	k.Name = m.Name
	k.Version = m.Version
	k.Description = m.Description
	k.Type = int(m.Type)

	intRestriction := make(map[reflect.Kind]bool)
	intRestriction[reflect.Int] = true

	k.Commands = make(map[string]interface{})
	err = GetStringMap(m.Cmd, kC, k.Commands, "commands", intRestriction)
	if err != nil {
		return err
	}

	k.Environment = make(map[string]interface{})
	err = GetStringMap(m.Env, kC, k.Environment, "environment", nil)
	if err != nil {
		return err
	}

	k.Interfaces = make(map[string]interface{})
	err = GetStringMap(m.Interfaces, kC, k.Interfaces, "interfaces", nil)
	if err != nil {
		return err
	}

	err = GetFrontend(m.Frontend, kC, &k.Frontend)
	if err != nil {
		return err
	}

	err = GetStringSlice(m.Imports, kC, &k.Imports, "imports")
	if err != nil {
		return err
	}

	return nil
}
