package caclient

import (
	"fmt"

	. "fabric-ca-smurf/config"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
)

type CaClient struct {
	ServerAddr string
	ServerPort int
	HomeDir    string
}

func NewCaClient() *CaClient {
	return &CaClient{
		ServerAddr: Configer.GetString("caserver.addr"),
		ServerPort: Configer.GetInt("caserver.port"),
		HomeDir:    Configer.GetString("caserver.homedir"),
	}
}

func (c *CaClient) GetCAInfo() (*lib.GetCAInfoResponse, error) {
	client := lib.Client{
		Config:  &lib.ClientConfig{URL: fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)},
		HomeDir: c.HomeDir,
	}
	cainfo, err := client.GetCAInfo(&api.GetCAInfoRequest{})
	return cainfo, err
}
