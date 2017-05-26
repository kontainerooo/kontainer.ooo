package abstraction

// Deprecated: Will be removed soon

// DCli is an interface to abstract dockers engine api client
type DCli interface {
	NetworkCreate() error
	NetworkRemove() error
	NetworkConnect() error
	NetworkDisconnect() error
	NetworkInspect() error
}

type dcliAbstract struct{}

func (d dcliAbstract) NetworkCreate() error {
	return nil
}

func (d dcliAbstract) NetworkRemove() error {
	return nil
}

func (d dcliAbstract) NetworkConnect() error {
	return nil
}

func (d dcliAbstract) NetworkDisconnect() error {
	return nil
}

func (d dcliAbstract) NetworkInspect() error {
	return nil
}

// NewDCLI returns an new Wrapper instance
func NewDCLI() DCli {
	return &dcliAbstract{}
}
