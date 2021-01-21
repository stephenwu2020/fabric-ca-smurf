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
	CommandPath string
	CommandName string
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
		CommandName: "fabric-ca-client",
		Admin:       admin,
		Adminpw:     adminpw,
	}
}

func (c *CaSdk) Help() (string, error) {
	return c.runCommand()
}

func (c *CaSdk) GetCAInfo() (string, error) {
	cainfoDir := Configer.GetString("caclient.cainfodir")
	args := []string{
		"getcainfo",
		"-H",
		cainfoDir,
	}
	return c.runCommand(args...)
}

func (c *CaSdk) Enroll() (string, error) {
	adminDir := Configer.GetString("caclient.admindir")
	url := fmt.Sprintf("http://%s:%s@%s:%d", c.Admin, c.Adminpw, c.ServerAddr, c.ServerPort)
	args := []string{
		"enroll",
		"-u",
		url,
		"-H",
		adminDir,
	}
	return c.runCommand(args...)
}

func (c *CaSdk) Register() error {
	return nil
}

func (c *CaSdk) runCommand(args ...string) (string, error) {
	cmd := exec.Command("", args...)
	cmd.Path = c.CommandPath
	var cmdOut, cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	cmd.Stdout = &cmdOut
	MyLogger.Infof("Run command:%s", cmd.String())
	if err := cmd.Run(); err != nil {
		return "", errors.WithMessage(err, cmdErr.String())
	}
	return cmdOut.String(), nil
}
