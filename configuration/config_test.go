package configuration

import (
	"net/http"
	"os"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
)

func Test(t *testing.T) { check.TestingT(t) }

// configStruct is a canonical example configuration, which should map to configYaml
var configStruct = Configuration{
	Clusters: map[string]Cluster{
		"dal": Cluster{
			Scheduler: Scheduler{
				"swarm": Parameters{
					"address":   "1.1.1.1:2376",
					"tlsverify": true,
					"tlscacert": "ca-swarm.pem",
					"tlscert":   "cert-swarm.pem",
					"tlskey":    "key-swarm.pem",
				},
			},
		},
		"wdc": Cluster{
			Scheduler: Scheduler{
				"swarm": Parameters{
					"address":   "2.2.2.2:2376",
					"tlsverify": true,
					"tlscacert": "ca-swarm.pem",
					"tlscert":   "cert-swarm.pem",
					"tlskey":    "key-swarm.pem",
				},
			},
		},
		"sjc": Cluster{
			Disabled: true,
			Scheduler: Scheduler{
				"marathon": Parameters{
					"address":   "3.3.3.3:8081",
					"tlsverify": true,
					"tlscacert": "ca-marathon.pem",
					"tlscert":   "cert-marathon.pem",
					"tlskey":    "key-marathon.pem",
				},
			},
		},
	},
	Notifications: Notifications{
		Endpoints: []Endpoint{
			{
				Name: "endpoint-1",
				URL:  "http://example.com",
				Headers: http.Header{
					"Authorization": []string{"Bearer <example>"},
				},
			},
		},
	},
}

// configYaml document representing configStruct
var configYaml = `
cluster:
  dal:
    scheduler:
      swarm:
        address: 1.1.1.1:2376
        tlsverify: true
        tlscacert: ca-swarm.pem
        tlscert: cert-swarm.pem
        tlskey: key-swarm.pem
  wdc:
    scheduler:
      swarm:
        address: 2.2.2.2:2376
        tlsverify: true
        tlscacert: ca-swarm.pem
        tlscert: cert-swarm.pem
        tlskey: key-swarm.pem
  sjc:
    disabled: true
    scheduler:
      marathon:
        address: 3.3.3.3:8081
        tlsverify: true
        tlscacert: ca-marathon.pem
        tlscert: cert-marathon.pem
        tlskey: key-marathon.pem
notifications:
  endpoints:
    - name: endpoint-1
      url:  http://example.com
      headers:
        Authorization: [Bearer <example>]
`

type ConfigSuite struct {
	expectedConfig Configuration
}

var _ = check.Suite(&ConfigSuite{})

func (suite *ConfigSuite) SetUpTest(c *check.C) {
	os.Clearenv()
	suite.expectedConfig = configStruct
}

// TestMarshalRoundtrip validates that configStruct can be marshaled and
// unmarshaled without changing any parameters
func (suite *ConfigSuite) TestMarshalRoundtrip(c *check.C) {
	configBytes, err := yaml.Marshal(suite.expectedConfig)
	c.Assert(err, check.IsNil)
	var config Configuration
	err = yaml.Unmarshal(configBytes, &config)
	c.Assert(err, check.IsNil)
	c.Assert(config, check.DeepEquals, suite.expectedConfig)
}

// TestParseSimple validates that configYamlV0_1 can be parsed into a struct
// matching configStruct
func (suite *ConfigSuite) TestParseSimple(c *check.C) {
	var config Configuration
	err := yaml.Unmarshal([]byte(configYaml), &config)
	c.Assert(err, check.IsNil)
	c.Assert(config, check.DeepEquals, suite.expectedConfig)
}


func (suite *ConfigSuite) TestParameters(c *check.C) {
	var params = Scheduler{
		"marathon": Parameters {
			"lala":"nono",
		},
	}
	
	c.Assert(params.Parameters(), check.DeepEquals, Parameters{"lala":"nono",})
}

func (suite *ConfigSuite) TestType(c *check.C) {
	
	var scheduler = Scheduler {
		"swarm": Parameters {
			"disable": "true",
			"key": "sdfs.key",
			"cert": "cert.ctr",
		},
	}
	
	var schedulerError = Scheduler {
		"swarm": Parameters {
			"disable": "true",
			"key": "sdfs.key",
			"cert": "cert.ctr",
		},
		"marathon": Parameters {
			"disable": "true",
			"key": "sdfs.key",
			"cert": "cert.ctr",
		},
	}
	
	c.Assert(scheduler.Type(), check.Equals, "swarm")
	c.Assert(schedulerError.Type(), check.Panics, "")
	
		
}
