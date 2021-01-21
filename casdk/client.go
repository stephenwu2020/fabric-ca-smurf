package casdk

import (
	"fmt"

	. "fabric-ca-smurf/config"

	"github.com/cloudflare/cfssl/log"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
)

type CaSdk struct {
	ServerAddr string
	ServerPort int
	HomeDir    string
	CaClient   *lib.Client
}

func NewCaSdk() *CaSdk {
	addr := Configer.GetString("caserver.addr")
	port := Configer.GetInt("caserver.port")
	dir := Configer.GetString("caserver.homedir")
	client := lib.Client{
		Config:  &lib.ClientConfig{URL: fmt.Sprintf("%s:%d", addr, port)},
		HomeDir: dir,
	}
	return &CaSdk{
		ServerAddr: addr,
		ServerPort: port,
		HomeDir:    dir,
		CaClient:   &client,
	}
}

func (c *CaSdk) Init() error {
	return c.CaClient.Init()
}

func (c *CaSdk) GetCAInfo() (*lib.GetCAInfoResponse, error) {
	cainfo, err := c.CaClient.GetCAInfo(&api.GetCAInfoRequest{})
	return cainfo, err
}

func (c *CaSdk) Enroll() error {
	req := api.EnrollmentRequest{}
	rsp, err := c.CaClient.Enroll(&req)
	log.Info("%+v", rsp)
	return err
}

func (c *CaSdk) Register() error {
	return nil
}
