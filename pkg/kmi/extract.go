package kmi

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
)

// Content is a struct type which may hold a byte array for every file which can be in the kmi file
type Content struct {
	Module     []byte
	Frontend   []byte
	Interfaces []byte
	Env        []byte
	Cmd        []byte
	Imports    []byte
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

		_, name := path.Split(header.Name)

		data := make([]byte, header.Size)
		_, err = tarReader.Read(data)
		if err != nil {
			return err
		}

		switch name {
		case "module.json":
			k.Module = data
		case "frontend.json":
			k.Frontend = data
		case "interfaces.json":
		case "if.json":
			k.Interfaces = data
		case "env.json":
			k.Env = data
		case "cmd.json":
			k.Cmd = data
		case "imports.json":
			k.Imports = data
		default:
			continue
		}
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
func ChooseSource(src interface{}, outsrc []byte, kind reflect.Kind, name string) error {
	v := reflect.ValueOf(src).Elem().Elem()
	k := v.Kind()
	if k == reflect.String {
		err := json.Unmarshal(outsrc, src)
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
func GetStringMap(src interface{}, outsrc []byte, dst map[string]interface{}, name string, restriction map[reflect.Kind]bool) error {
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
func GetStringSlice(src interface{}, outsrc []byte, dst *[]string, name string) error {
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
func GetFrontend(src interface{}, outsrc []byte, dst *[]frontendModule) error {
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
	err = GetStringMap(m.Cmd, kC.Cmd, k.Commands, "commands", intRestriction)
	if err != nil {
		return err
	}

	k.Environment = make(map[string]interface{})
	err = GetStringMap(m.Env, kC.Env, k.Environment, "environment", nil)
	if err != nil {
		return err
	}

	k.Interfaces = make(map[string]interface{})
	err = GetStringMap(m.Interfaces, kC.Interfaces, k.Interfaces, "interfaces", nil)
	if err != nil {
		return err
	}

	err = GetFrontend(m.Frontend, kC.Frontend, &k.Frontend)
	if err != nil {
		return err
	}

	err = GetStringSlice(m.Imports, kC.Imports, &k.Imports, "imports")
	if err != nil {
		return err
	}

	return nil
}
