package azure

import (
	"github.com/Nanocloud/community/nanocloud/iaas"
)

type vm struct {
	name string
}

func (v *vm) GetName() string {
	return v.name
}
func (v *vm) Stop() error {
	return nil
}

type conn struct {
}

func (c *conn) ListVM() ([]iaas.VM, error) {
	vm := vm{
		name: "Qemu VM",
	}
	vms := make([]iaas.VM, 0)
	vms = append(vms, &vm)
	return vms, nil
}

type driver struct {
}

func (d *driver) Open(options map[string]string) (iaas.Iaas, error) {
	return &conn{}, nil
}

func init() {
	iaas.Register("qemu", &driver{})
}
