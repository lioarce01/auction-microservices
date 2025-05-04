package servicediscovery

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ConsulService struct {
	Client *api.Client
}

func NewConsulService() (*ConsulService, error) {
	config := api.DefaultConfig()
	config.Address = "consul:8500"

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error creating consul client: %v", err)
	}

	return &ConsulService{Client: client}, nil
}

func (c *ConsulService) RegisterService(serviceName, serviceID, address string, port int) error {
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: address,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	return c.Client.Agent().ServiceRegister(registration)
}
