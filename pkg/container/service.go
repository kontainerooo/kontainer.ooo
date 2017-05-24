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
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/opencontainers/runc/libcontainer"
	"github.com/opencontainers/runc/libcontainer/configs"
	"golang.org/x/net/context"
)

// Service Container Service
type Service interface {
	// CreateContainer instanciates a container for a User with the id refid and returns its id
	CreateContainer(refid uint, kmiID uint, name string) (id string, err error)

	// RemoveContainer is used to remove a container instance by id
	RemoveContainer(refid uint, id string) error

	// Instances returns a list of container instances of a user by id
	Instances(refid uint) []Container

	// StopContainer stops a container
	StopContainer(refid uint, id string) error

	// IsRunning checks if a container is running
	IsRunning(refid uint, id string) bool

	// Execute executes a command in a given container
	Execute(refid uint, id string, cmd string) (string, error)
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
	Update(interface{}, ...interface{}) error
}

type service struct {
	db        dbAdapter
	libcnt    libcontainer.Factory
	kmiClient *kmi.Endpoints
	logger    log.Logger
}

const (
	// ContainerRootfsPath is the path were the container base images are stored
	ContainerRootfsPath = "/var/lib/kontainerooo/images"

	// ContainerCustomerPath is the path were customer data is stored
	ContainerCustomerPath = "/var/lib/kontainerooo/customers/"

	// ContainerNetNSPath is the path where the netns binary is stored
	ContainerNetNSPath = "/var/go/bin/netns"

	// ContainerStandardPathVariable is the standard path environment variable
	ContainerStandardPathVariable = "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games"

	// ContainerTermVariable is the term variable
	ContainerTermVariable = "TERM=xterm"
)

// InitializeDatabases sets up the container service's database
func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Container{})
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
	err := s.checkAndCreate(ContainerRootfsPath)
	if err != nil {
		return err
	}

	// TODO: When there is a reasonable root filesystem and a
	// server then check for a rootfs and if not present download
	// the image
	_, err = os.Stat(path.Join(ContainerRootfsPath, "rootfs.tar"))
	if err != nil {
		return errors.New("no rootfs present")
	}

	err = s.checkAndCreate(ContainerCustomerPath)
	if err != nil {
		return err
	}

	_, err = os.Stat(ContainerNetNSPath)
	if err != nil {
		return errors.New("netns binary not present, please go get github.com/jessfraz/netns")
	}

	return nil
}

func (s *service) CreateContainer(refid uint, kmiID uint, name string) (id string, err error) {
	kmi, err := s.getKMI(kmiID)
	if err != nil {
		return "", err
	}

	// Compute the container id - consisting of userID + imagename + timestamp
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%d%d%s", refid, kmi.ID, time.Now().Format("20060102150405")))
	containerID := fmt.Sprintf("%x", h.Sum(nil))

	s.initRootfs(refid, kmi.ProvisionScript, containerID)

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
		Rootfs:            path.Join(ContainerCustomerPath, fmt.Sprintf("%d", refid), containerID, "rootfs"),
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
				configs.CommandHook{Command: configs.Command{Path: ContainerNetNSPath}},
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
		RefID:       refid,
		ContainerID: containerID,
		KMI:         kmi,
		Running:     false,
	}
	err = s.db.Create(&c)
	if err != nil {
		return "", err
	}

	return containerID, nil
}

func (s *service) RemoveContainer(refid uint, id string) error {
	err := s.StopContainer(refid, id)
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(ContainerCustomerPath, string(refid), id))
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

func (s *service) Instances(refid uint) []Container {
	s.db.Where("refid = ?", refid)

	cs := []Container{}
	err := s.db.Find(&cs)
	if err != nil {
		return []Container{}
	}

	return cs
}

func (s *service) StopContainer(refid uint, id string) error {
	container, err := s.libcnt.Load(id)
	if err != nil {
		return err
	}

	c := Container{
		ContainerID: id,
		Running:     false,
	}

	err = s.db.Update(&Container{}, &c)
	if err != nil {
		return err
	}

	err = container.Signal(os.Kill, true)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) IsRunning(refid uint, id string) bool {
	s.db.Where("refid = ? AND id = ?", refid, id)

	c := Container{}
	err := s.db.Find(&c)
	if err != nil {
		return false
	}

	return c.Running
}

func (s *service) Execute(refid uint, id string, cmd string) (string, error) {
	container, err := s.libcnt.Load(id)
	if err != nil {
		return "", err
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	p := &libcontainer.Process{
		Args:   []string{"/bin/sh", "-c", cmd},
		Env:    []string{ContainerStandardPathVariable, ContainerTermVariable},
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

func (s *service) getKMI(kmiID uint) (kmi.KMI, error) {
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

func (s *service) initRootfs(refid uint, provisionScript string, id string) error {
	imagePath := path.Join(ContainerRootfsPath, "rootfs.tar")
	_, err := os.Stat(imagePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	cPath := path.Join(ContainerCustomerPath, fmt.Sprintf("%d", refid))
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

	err = s.provisionRootfs(mPath, provisionScript)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) provisionRootfs(rfsPath string, provisionScript string) error {
	conf := provisionConfig
	conf.Rootfs = rfsPath

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

	p := &libcontainer.Process{
		Args:   []string{"/bin/sh", "/provision.sh"},
		Env:    []string{ContainerStandardPathVariable, ContainerTermVariable},
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
	s := &service{
		libcnt:    lc,
		db:        db,
		kmiClient: ke,
		logger:    l,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	err = s.InitPaths()
	if err != nil {
		return s, err
	}

	return s, nil
}
