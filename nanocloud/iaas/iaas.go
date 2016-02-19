package iaas

type Iaas interface {
	ListVM() ([]VM, error)
	GetVM(id string) (VM, error)
}

type Driver interface {
	Open(options map[string]string) (Iaas, error)
}

var (
	drivers = make(map[string]Driver)
)

type VM interface {
	GetName() string
	Stop() error
	Start() error
}

var iaasConn Iaas

func Register(name string, driver Driver) {
	drivers[name] = driver
}

func Open(driverName string, options map[string]string) error {
	conn, err := drivers[driverName].Open(options)
	iaasConn = conn
	return err
}

func GetVM(id string) (VM, error) {
	return iaasConn.GetVM(id)
}
func ListVM() ([]VM, error) {
	return iaasConn.ListVM()
}
