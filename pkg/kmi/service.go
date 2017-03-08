// Package kmi provides functionality to handle kmi files
package kmi

// The Service interface describes the functions necessary for a KMI Service
type Service interface {
	// AddKMI resolves the path to a kmi file, extracts it and adds its contents to the database as a new kontainer module
	AddKMI(path string) (id uint, err error)

	// RemoveKMI removes the kontainer module information and all files related
	RemoveKMI(id uint) error

	// GetKMI retrieves kontainer module information for a specific module
	GetKMI(id uint) (*KMI, error)

	// KMI returns display information for all exisiting kontainer modules
	KMI() *[]KMDI
}

type service struct{}

func (s *service) AddKMI(path string) (uint, error) {
	//TODO: implement
	return 0, nil
}

func (s *service) RemoveKMI(id uint) error {
	//TODO: implement
	return nil
}

func (s *service) GetKMI(id uint) (*KMI, error) {
	//TODO: implement
	return nil, nil
}

func (s *service) KMI() *[]KMDI {
	//TODO: implement
	return nil
}

// NewService returns a new instance of a KMIService
func NewService() Service {
	return &service{}
}
