package register

type RegistryClient interface {
	NewRegistryClient(host string, port int) (*RegistryClient, error)
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}
