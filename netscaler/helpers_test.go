package netscaler

import (
	"encoding/base64"
	"github.com/chiradeep/go-nitro/config/system"
	"github.com/chiradeep/go-nitro/netscaler"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
	"testing"
)

func uploadTestdataFile(t *testing.T, filename, targetDir string) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	// Get here path
	_, here_filename, _, _ := runtime.Caller(1)
	b, err := ioutil.ReadFile(path.Join(path.Dir(here_filename), "testdata", filename))

	if err != nil {
		return err
	}

	sf := system.Systemfile{
		Filename:     filename,
		Filecontent:  base64.StdEncoding.EncodeToString(b),
		Filelocation: targetDir,
	}
	_, err = nsClient.AddResource(netscaler.Systemfile.Type(), filename, &sf)
	if err != nil && strings.Contains(err.Error(), "File already exists") {
		args := map[string]string{"filelocation": "%2Fvar%2Ftmp"}
		err := nsClient.DeleteResourceWithArgsMap(netscaler.Systemfile.Type(), filename, args)
		if err != nil {
			return err
		}
		_, err = nsClient.AddResource(netscaler.Systemfile.Type(), filename, &sf)
		if err != nil {
			return err
		}
	}
	return nil
}
