package brig

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the brig mixin in porter.yaml
// mixins:
// - brig:
//	  clientVersion: "v0.0.0"
type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

const dockerfileLinesTmpl = `RUN apt-get update && apt-get install wget -y && \
wget -O /usr/local/bin/brig https://github.com/brigadecore/brigade/releases/download/%s/brig-linux-amd64 && \
chmod +x /usr/local/bin/brig
`

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	// Create new Builder.
	var input BuildInput

	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	suppliedClientVersion := input.Config.ClientVersion

	if suppliedClientVersion != "" {
		m.ClientVersion = suppliedClientVersion
	}

	fmt.Fprintf(m.Out, dockerfileLinesTmpl, m.ClientVersion)

	return nil
}
