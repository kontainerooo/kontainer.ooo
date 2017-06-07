// +build linux

package container

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/kontainerooo/kontainer.ooo/pkg/util"
	"github.com/opencontainers/runc/libcontainer"
	"github.com/opencontainers/runc/libcontainer/configs"
	"golang.org/x/net/context"
)

// Service Container Service
type Service interface {
	// CreateContainer instanciates a container for a User with the id refID and returns its id
	CreateContainer(refID uint, kmiID uint, name string) (id string, err error)

	// RemoveContainer is used to remove a container instance by id
	RemoveContainer(refID uint, id string) error

	// Instances returns a list of container instances of a user by id
	Instances(refID uint) []Container

	// StopContainer stops a container
	StopContainer(refID uint, id string) error

	// Execute executes a command in a given container
	Execute(refID uint, id string, cmd string, env map[string]string) (string, error)

	// GetEnv returns the value to a given environment variable setting. Returns the whole
	// environment as string if key is empty
	GetEnv(refID uint, id string, key string) (string, error)

	// SetEnv sets an environment variable for the container
	SetEnv(refID uint, id string, key string, value string) error

	// IDForName returs the containerID for a given container name and user
	IDForName(refID uint, name string) (string, error)

	// GetContainerKMI returns the KMI for a given container
	GetContainerKMI(containerID string) (kmi.KMI, error)

	// SetLink links a container's interface into a container
	SetLink(refID uint, containerID string, linkID string, linkName string, linkInterface string) error

	// RemoveLink links a container's interface into a container
	RemoveLink(refID uint, containerID string, linkID string, linkName string, linkInterface string) error

	// GetLinks returns all links a container has
	GetLinks(refID uint, containerID string) (map[string][]string, error)
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
	Update(interface{}, ...interface{}) error
}

type service struct {
	db        dbAdapter
	libcnt    libcontainer.Factory
	kmiClient *kmi.Endpoints
	logger    log.Logger
	config    util.ConfigFile
}

const (
	// ContainerTermVariable is the term variable
	ContainerTermVariable = "xterm"
)

// InitializeDatabases sets up the container service's database
func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&CKMI{}, &Container{})
}

func (s *service) checkAndCreate(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0755)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

// InitPaths initializes all paths this service needs
func (s *service) InitPaths() error {
	err := s.checkAndCreate(s.config.RootfsPath)
	if err != nil {
		return err
	}

	// TODO: When there is a reasonable root filesystem and a
	// server then check for a rootfs and if not present download
	// the image
	_, err = os.Stat(path.Join(s.config.RootfsPath, "rootfs.tar"))
	if err != nil {
		return errors.New("no rootfs present")
	}

	err = s.checkAndCreate(s.config.CustomerPath)
	if err != nil {
		return err
	}

	_, err = os.Stat(s.config.NetNSPath)
	if err != nil {
		return errors.New("netns binary not present, please go get github.com/jessfraz/netns")
	}

	return nil
}

func (s *service) CreateContainer(refID uint, kmiID uint, name string) (id string, err error) {
	kmi, err := s.getTemplateKMI(kmiID)
	if err != nil {
		return "", err
	}

	// Compute the container id - consisting of userID + imagename + timestamp
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%d%d%s", refID, kmi.ID, time.Now().Format("20060102150405")))
	containerID := fmt.Sprintf("%x", h.Sum(nil))

	s.initRootfs(refID, kmi.ProvisionScript, containerID, kmi)

	netnsCmd := configs.NewCommandHook(configs.Command{
		Path: s.config.NetNSPath,
		Args: []string{"netns", "-ipfile", path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID), containerID, ".ip")},
		Dir:  path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID), containerID),
	})

	mountCfg := []*configs.Mount{&configs.Mount{
		Source:      "proc",
		Destination: "/proc",
		Device:      "proc",
		Flags:       defaultMountFlags,
	}, &configs.Mount{
		Source:      "tmpfs",
		Destination: "/dev",
		Device:      "tmpfs",
		Flags:       syscall.MS_NOSUID | syscall.MS_STRICTATIME,
		Data:        "mode=755",
	}, &configs.Mount{
		Source:      "devpts",
		Destination: "/dev/pts",
		Device:      "devpts",
		Flags:       syscall.MS_NOSUID | syscall.MS_NOEXEC,
		Data:        "newinstance,ptmxmode=0666,mode=0620,gid=5",
	}, &configs.Mount{
		Source:      "mqueue",
		Destination: "/dev/mqueue",
		Device:      "mqueue",
		Flags:       defaultMountFlags,
	}, &configs.Mount{
		Source:      "sysfs",
		Destination: "/sys",
		Device:      "sysfs",
		Flags:       defaultMountFlags | syscall.MS_RDONLY,
	}}

	// TODO: Create a reasonable configuration
	cu, err := s.libcnt.Create(containerID, &configs.Config{
		NoPivotRoot:       false,
		ParentDeathSignal: 9,
		Rootfs:            path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID), containerID, "rootfs"),
		Readonlyfs:        false,
		RootPropagation:   0,
		// TODO: Don't give the container all Capabilities (obviously...)
		Capabilities: &allCaps,
		Mounts:       mountCfg,
		Devices:      configs.DefaultAutoCreatedDevices,
		Namespaces: []configs.Namespace{
			{Type: configs.NEWNS},
			{Type: configs.NEWUTS},
			{Type: configs.NEWIPC},
			{Type: configs.NEWPID},
			{Type: configs.NEWNET},
		},
		Networks: []*configs.Network{
			{
				Type:    "loopback",
				Address: "127.0.0.1/0",
				Gateway: "localhost",
			},
		},
		Routes: nil,
		Cgroups: &configs.Cgroup{
			Name:   containerID,
			Parent: "system",
			Resources: &configs.Resources{
				MemorySwappiness: nil,
				AllowAllDevices:  nil,
				AllowedDevices:   configs.DefaultAllowedDevices,
			},
		},
		Rootless: false,
		Hostname: name,
		Hooks: &configs.Hooks{
			Prestart: []configs.Hook{
				netnsCmd,
			},
		},
	})
	if err != nil {
		return "", err
	}

	p := &libcontainer.Process{
		Args: []string{"/bin/echo", "Done!"},
	}

	if err = cu.Run(p); err != nil {
		return "", err
	}

	_, err = p.Wait()
	if err != nil {
		return "", err
	}

	c := Container{
		RefID:         refID,
		ContainerName: name,
		ContainerID:   containerID,
	}

	ckmi := CKMI{
		KMI:   kmi,
		Links: abstraction.NewJSONFromMap(make(map[string]string)),
	}
	ckmi.ID = 0

	s.db.Begin()

	err = s.db.Create(&ckmi)
	if err != nil {
		s.db.Rollback()
		return "", err
	}

	c.KMIID = ckmi.ID

	err = s.db.Create(&c)
	if err != nil {
		s.db.Rollback()
		return "", err
	}

	s.db.Commit()
	return containerID, nil
}

func (s *service) RemoveContainer(refID uint, id string) error {
	err := s.StopContainer(refID, id)
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(s.config.CustomerPath, string(refID), id))
	if err != nil {
		return err
	}

	c := Container{
		ContainerID: id,
	}

	err = s.db.Delete(&c)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Instances(refID uint) []Container {
	s.db.Where("ref_id = ?", refID)

	cs := []Container{}
	err := s.db.Find(&cs)
	if err != nil {
		return []Container{}
	}

	for k := range cs {
		cKMI := CKMI{}
		err = s.db.First(&cKMI, "id = ?", cs[k].KMIID)
		if err != nil {
			return []Container{}
		}

		cs[k].KMI = cKMI
	}

	return cs
}

func (s *service) StopContainer(refID uint, id string) error {
	container, err := s.libcnt.Load(id)
	if err != nil {
		return err
	}

	err = container.Signal(os.Kill, true)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Execute(refID uint, id string, cmd string, env map[string]string) (string, error) {
	container, err := s.libcnt.Load(id)
	if err != nil {
		return "", err
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	cKMI, err := s.getCKMI(id)
	if err != nil {
		return "", err
	}

	execEnv := s.createEnvironmentMap(refID, cKMI, env)

	// Make env string array
	envString := []string{}
	for k, v := range execEnv {
		// Replace spaces in ENV variable key
		key := strings.Replace(k, " ", "_", -1)
		envString = append(envString, fmt.Sprintf("%s=%s", key, v))
	}

	p := &libcontainer.Process{
		Args:   []string{"/bin/sh", "-c", cmd},
		Env:    envString,
		Stdout: buf,
	}

	err = container.Run(p)
	if err != nil {
		return "", err
	}

	_, err = p.Wait()
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (s *service) GetEnv(refID uint, id string, key string) (string, error) {
	cKMI, err := s.GetContainerKMI(id)
	if err != nil {
		return "", err
	}

	env := cKMI.Environment.ToStringMap()

	// When key is empty the whole environment is returned as string
	if key == "" {
		envString := ""
		for k, v := range env {
			// Replace spaces in ENV variable key
			envKey := strings.Replace(k, " ", "_", -1)
			envString = fmt.Sprintf("%s, %s=%s", envString, envKey, v)
		}
		return envString, nil
	}

	val, ok := env[key]
	if !ok {
		return "", errors.New("variable does not exist")
	}

	return val, nil
}

func (s *service) SetEnv(refID uint, id string, key string, value string) error {
	cKMI, err := s.GetContainerKMI(id)
	if err != nil {
		return err
	}

	env := cKMI.Environment.ToStringMap()

	env[key] = value

	s.db.Begin()
	c := &Container{
		ContainerID: id,
		KMI: CKMI{
			KMI: kmi.KMI{
				Environment: abstraction.NewJSONFromMap(env),
			},
		},
	}

	err = s.db.Update(&Container{}, c)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()

	return nil
}

func (s *service) IDForName(refID uint, name string) (string, error) {
	c := &Container{}
	err := s.db.First(c, "ref_id = ? AND container_name = ?", refID, name)
	if err != nil {
		return "", err
	}
	return c.ContainerID, nil
}

func (s *service) createEnvironmentMap(refID uint, cKMI CKMI, env map[string]string) map[string]string {

	// Prefix interfaces to create the environment variables
	execEnv := make(map[string]string)
	for k, v := range cKMI.Interfaces.ToStringMap() {
		execEnv[fmt.Sprintf("KROO_%s", strings.ToUpper(k))] = v
	}

	// Prefix commands to create command environment variables
	for k, v := range cKMI.Commands.ToStringMap() {
		execEnv[fmt.Sprintf("KROO_CMD_%s", strings.ToUpper(k))] = v
	}

	// Create link environment variables
	links := cKMI.Links.ToStringArrayMap()

	for k, v := range links {
		containerID, err := s.IDForName(refID, k)
		if err != nil {
			continue
		}

		linkKMI, err := s.getCKMI(containerID)
		if err != nil {
			continue
		}

		// Get link container's IP
		ip := s.getContainerIP(refID, containerID)
		if ip == "" {
			continue
		}

		// set ip addres
		execEnv[fmt.Sprintf("KROO_LINK_%s_IP", strings.ToUpper(k))] = ip

		// Get the ports of the link interfaces
		for _, ports := range v {
			p, ok := linkKMI.Interfaces.ToStringMap()[ports]
			if !ok {
				continue
			}

			// set port
			execEnv[fmt.Sprintf("KROO_LINK_%s_%s_PORT", strings.ToUpper(k), strings.ToUpper(ports))] = p
		}

	}

	// Add the KMIs configured environment, possibly overriding
	for k, v := range cKMI.Environment.ToStringMap() {
		execEnv[k] = v
	}

	// Construct environment and replace global with exec specific env
	for gK := range execEnv {
		for xK, xV := range env {
			if strings.ToLower(gK) == strings.ToLower(xK) {
				execEnv[gK] = xV
				// Delete previously found key to reduce loop time
				// <- is that even true? Hopefully ¯\_(ツ)_/¯
				delete(env, xK)
			}
		}
	}

	// Append to path variable if there is one given
	_, ok := execEnv["PATH"]
	if !ok {
		execEnv["PATH"] = s.config.StandardPathVariable
	} else {
		execEnv["PATH"] = fmt.Sprintf("%s:%s", execEnv["PATH"], s.config.StandardPathVariable)
	}

	// Set or override TERM variable if there is one
	execEnv["TERM"] = ContainerTermVariable

	return execEnv
}

func (s *service) getTemplateKMI(kmiID uint) (kmi.KMI, error) {
	if s.kmiClient == nil {
		return kmi.KMI{}, errors.New("No KMI client")
	}

	kmiResponse, err := s.kmiClient.GetKMIEndpoint(context.Background(), kmi.GetKMIRequest{
		ID: kmiID,
	})
	if err != nil {
		return kmi.KMI{}, err
	}
	if kmiResponse.(kmi.GetKMIResponse).Error != nil {
		return kmi.KMI{}, kmiResponse.(kmi.GetKMIResponse).Error
	}

	kmi := kmiResponse.(kmi.GetKMIResponse).KMI

	return *kmi, nil
}

func (s *service) GetContainerKMI(containerID string) (kmi.KMI, error) {
	// TODO: cache container KMIs
	cKMI, err := s.getCKMI(containerID)
	if err != nil {
		return kmi.KMI{}, err
	}

	return kmi.KMI(cKMI.KMI), nil
}

func (s *service) getCKMI(containerID string) (CKMI, error) {
	c := Container{}
	err := s.db.First(&c, "container_id = ?", containerID)
	if err != nil {
		return CKMI{}, err
	}

	cKMI := CKMI{}
	err = s.db.First(&cKMI, "id = ?", c.KMIID)
	if err != nil {
		return CKMI{}, err
	}

	return cKMI, nil
}

func (s *service) getContainerIP(refID uint, containerID string) string {
	ipPath := path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID), containerID, ".ip")
	ip, err := ioutil.ReadFile(ipPath)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s", ip)
}

func (s *service) SetLink(refID uint, containerID string, linkID string, linkName string, linkInterface string) error {
	containerKMI, err := s.getCKMI(containerID)
	if err != nil {
		return err
	}

	// construct links
	links := containerKMI.Links.ToStringArrayMap()

	// check if container already has a link
	ar, ok := links[linkName]
	if !ok {
		ar = []string{}
	}
	// check if this interface already exists
	// TODO: make this more performant
	var exists bool
	for _, v := range ar {
		if v == linkInterface {
			exists = true
			break
		}
	}
	if !exists {
		ar = append(ar, linkInterface)
	}

	links[linkName] = ar

	s.db.Begin()
	containerKMI.Links = abstraction.NewJSONFromMapArray(links)

	err = s.db.Update(&CKMI{}, containerKMI)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) RemoveLink(refID uint, containerID string, linkID string, linkName string, linkInterface string) error {
	containerKMI, err := s.getCKMI(containerID)
	if err != nil {
		return err
	}

	// construct links
	links := containerKMI.Links.ToStringArrayMap()

	// check if link exists
	ar, ok := links[linkName]
	if !ok {
		return errors.New("link does not exist")
	}
	// check if this interface exists
	var exists bool
	for i, v := range ar {
		if v == linkInterface {
			ar = append(ar[:i], ar[i+1:]...)
			exists = true
			break
		}
	}
	if !exists {
		return errors.New("link interface does not exist")
	}

	links[linkName] = ar

	s.db.Begin()
	saveCKMI := CKMI{}
	saveCKMI.Links = abstraction.NewJSONFromMapArray(links)

	err = s.db.Update(&CKMI{}, saveCKMI)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) GetLinks(refID uint, containerID string) (map[string][]string, error) {
	containerKMI, err := s.getCKMI(containerID)
	if err != nil {
		return make(map[string][]string), err
	}

	links := containerKMI.Links.ToStringArrayMap()

	return links, nil
}

func (s *service) initRootfs(refID uint, provisionScript string, id string, cKMI kmi.KMI) error {
	imagePath := path.Join(s.config.RootfsPath, "rootfs.tar")
	_, err := os.Stat(imagePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	cPath := path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID))
	_, err = os.Stat(cPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	mPath := path.Join(cPath, id, "rootfs")
	err = os.MkdirAll(mPath, 0755)
	if err != nil {
		return nil
	}

	err = s.untar(mPath, imagePath)
	if err != nil {
		return err
	}

	err = s.provisionRootfs(cKMI, mPath, provisionScript)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) provisionRootfs(cKMI kmi.KMI, rfsPath string, provisionScript string) error {
	conf := provisionConfig
	conf.Rootfs = rfsPath
	conf.Hooks.Prestart[0] = configs.CommandHook{Command: configs.Command{Path: s.config.NetNSPath}}

	name := fmt.Sprintf("provision-%s", time.Now().Format("20060102150405"))

	container, err := s.libcnt.Create(name, conf)
	if err != nil {
		return err
	}

	conf.Cgroups.Name = name

	// Create provision script
	err = ioutil.WriteFile(path.Join(rfsPath, "provision.sh"), []byte(provisionScript), 0755)
	if err != nil {
		return err
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	e := []byte{}
	errBuf := bytes.NewBuffer(e)

	execEnv := cKMI.Environment.ToStringMap()

	// Append to path variable if there is one given
	_, ok := execEnv["PATH"]
	if !ok {
		execEnv["PATH"] = s.config.StandardPathVariable
	} else {
		execEnv["PATH"] = fmt.Sprintf("%s:%s", execEnv["PATH"], s.config.StandardPathVariable)
	}

	// Set or override TERM variable if there is one
	execEnv["TERM"] = ContainerTermVariable

	// Make env string array
	envString := []string{}
	for k, v := range execEnv {
		envString = append(envString, fmt.Sprintf("%s=%s", k, v))
	}

	p := &libcontainer.Process{
		Args:   []string{"/bin/sh", "/provision.sh"},
		Env:    envString,
		Stdout: buf,
		Stderr: errBuf,
	}

	err = container.Run(p)
	if err != nil {
		return err
	}

	// TODO: make clean up more... clean
	go func() error {
		time.Sleep(120 * time.Second)
		s.logger.Log("Timed out. Destroying provision container")
		container.Signal(os.Kill, true)
		container.Destroy()
		return nil
	}()

	_, err = p.Wait()
	if err != nil {
		return err
	}

	// Write output to log file
	err = ioutil.WriteFile(path.Join(rfsPath, "../provision.log"), b, 0644)
	if err != nil {
		return err
	}

	// Remove provision script
	err = os.Remove(path.Join(rfsPath, "provision.sh"))
	if err != nil {
		return err
	}

	err = container.Signal(os.Kill, true)
	if err != nil {
		return err
	}

	err = container.Destroy()
	if err != nil {
		return err
	}

	return nil
}

func (s *service) untar(dst string, src string) error {
	cmd := exec.Command("/bin/tar", "xzf", src, "-C", dst)
	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a new container service with necessary dependencies
func NewService(lc libcontainer.Factory, db dbAdapter, ke *kmi.Endpoints, l log.Logger) (Service, error) {
	conf, err := util.GetConfig()
	if err != nil {
		return &service{}, err
	}

	s := &service{
		libcnt:    lc,
		db:        db,
		kmiClient: ke,
		logger:    l,
		config:    conf,
	}

	err = s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	err = s.InitPaths()
	if err != nil {
		return s, err
	}

	return s, nil
}
