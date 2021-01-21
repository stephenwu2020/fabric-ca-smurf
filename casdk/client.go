package casdk

import (
	"bytes"
	. "fabric-ca-smurf/config"
	. "fabric-ca-smurf/logger"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/pkg/errors"
)

type CaSdk struct {
	ServerAddr  string
	ServerPort  int
	HomeDir     string
	CommandName string
	CommandPath string
	Admin       string
	Adminpw     string
}

func NewCaSdk() *CaSdk {
	addr := Configer.GetString("caserver.addr")
	port := Configer.GetInt("caserver.port")
	dir := Configer.GetString("caserver.homedir")
	admin := Configer.GetString("caserver.admin")
	adminpw := Configer.GetString("caserver.adminpw")
	pwd, _ := os.Getwd()
	cliPath := path.Join(pwd, "caclient", "fabric-ca-client")
	MyLogger.Info(cliPath)

	return &CaSdk{
		ServerAddr:  addr,
		ServerPort:  port,
		HomeDir:     dir,
		CommandPath: cliPath,
		Admin:       admin,
		Adminpw:     adminpw,
	}
}

func (c *CaSdk) Help() (string, error) {
	cmd := exec.Command(c.CommandName)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Path = c.CommandPath

	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func (c *CaSdk) GetCAInfo() (string, error) {
	cainfoDir := Configer.GetString("caclient.cainfodir")
	cmd := exec.Command(c.CommandName, "getcainfo", "-H", cainfoDir)
	MyLogger.Info(cmd.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Path = c.CommandPath

	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func (c *CaSdk) Enroll() (string, error) {
	adminDir := Configer.GetString("caclient.admindir")
	url := fmt.Sprintf("http://%s:%s@%s:%d", c.Admin, c.Adminpw, c.ServerAddr, c.ServerPort)
	cmd := exec.Command(c.CommandName, "enroll", "-u", url, "-H", adminDir)
	MyLogger.Info(cmd.String())
	var out, cmdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &cmdErr
	cmd.Path = c.CommandPath

	if err := cmd.Run(); err != nil {
		return "", errors.WithMessage(err, cmdErr.String())
	}
	return out.String(), nil
}

func (c *CaSdk) Register() error {
	return nil
}
