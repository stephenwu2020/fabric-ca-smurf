package caclient

import (
	"fmt"

	. "fabric-ca-smurf/config"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
)

type CaClient struct {
}

func GetCAInfo() (*lib.GetCAInfoResponse, error) {
	serverAddr := Configer.GetString("caserver.addr")
	serverPort := Configer.GetInt("caserver.port")
	homeDir := Configer.GetString("caserver.homedir")
	client := lib.Client{
		Config:  &lib.ClientConfig{URL: fmt.Sprintf("%s:%d", serverAddr, serverPort)},
		HomeDir: homeDir,
	}
	cainfo, err := client.GetCAInfo(&api.GetCAInfoRequest{})
	return cainfo, err
}
