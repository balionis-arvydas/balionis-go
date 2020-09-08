package bago0

import (
	"github.com/pkg/errors"
	"io"
	"testing"
)

var mockConfig = `
server:
  name: MyTestName
  loggerConfig:
    logType: file
    logLevel: debug
    properties:
      file: bago0.log
`

func mockReadFile(filename string) ([]byte, error) {
	if filename == "sample.yaml" {
		return []byte(mockConfig), nil
	} else {
		return nil, io.EOF
	}
}

func init() {
	ioutilReadFile = mockReadFile
}

func TestServerConfig_Load(t *testing.T) {
	var data = []struct {
		filename   string
		err        error
		serverName string
	}{
		{
			filename:"sample.yaml",
			serverName: "MyTestName",
		},
		{
			filename: "fake.yaml",
			err: errors.Wrap(io.EOF, "failed to read fake.yaml"),
		},
	}
	t.Log("Given the need to test loading config files.")
	{
		for _, d := range data {
			t.Logf("\tWhen loading \"%s\" ", d.filename)
			{
				var config ServerConfig
				err := config.Load(d.filename)
				if (err == nil && d.err == nil) || err.Error() == d.err.Error() {
					t.Logf("\t\tShould expect to load file \"%s\"", d.filename)
				} else {
					t.Fatal("\t\tShould expect to load file " + d.filename, err)
					break
				}

				if config.Name == d.serverName {
					t.Logf("\t\tShould expect name \"%s\"", d.serverName)
				} else {
					t.Errorf("\t\tShould expect name \"%s\" but received \"%s\"", d.serverName, config.Name)
				}
			}
		}
	}
}
