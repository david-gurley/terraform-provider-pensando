package pensando

type Config struct {
	DSCgRPCClientConfigs []DSCgRPCClientConfig
}
type DSCgRPCClientConfig struct {
	Name    string
	Address string
	Port    int
}

type PensandoClient struct {
	DSCgRPCClients map[string]*sdn.DSCgRPCClient
}

func (config *Config) Client() (interface{}, error) {
	dscClients := make(map[string]*sdn.DSCCleint)
	client := PensandoClient{}
	for _, dscgRPCClientConfig := range config.DSCgRPCClientConfigs {
		dscgRPCClient, err := sdn.NewDSCgRPCClient(
			dscgRPCClientConfig.Address,
			dscgRPCClientConfig.Port,
		)
		if err != nil {
			return nil, err
		}
		client.DSCgRPCClients[DSCgRPCClientConfig.Name] = &dscgRPCClient
	}
	return &client, nil
}
