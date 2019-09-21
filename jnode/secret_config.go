package jnode

import (
	"bytes"
	"text/template"
)

const secretConfigTemplate = `
{{- with .Genesis}}
genesis:
  {{- if .SigKey}}
  sig_key: {{ .SigKey }}
  {{- end}}
  {{- if .VrfKey}}
  vrf_key: {{ .VrfKey }}
  {{- end}}
  {{- if .NodeID}}
  node_id: {{ .NodeID }}
  {{- end}}
{{end}}
{{- with .Bft}}
bft:
  {{- if .SigningKey}}
  signing_key: {{ .SigningKey }}
  {{- end}}
{{end}}
`

// SecretConfig ...
type SecretConfig struct {
	Genesis SecretGenesisPraos // `"genesis"`
	Bft     SecretBft          // `"bft"`
}

// SecretGenesisPraos ...
type SecretGenesisPraos struct {
	SigKey string // `"sig_key"`
	VrfKey string // `"vrf_key"`
	NodeID string // `"node_id"`
}

// SecretBft ...
type SecretBft struct {
	SigningKey string // `"signing_key"`
}

// NewSecretConfig ...
func NewSecretConfig() *SecretConfig {
	return &SecretConfig{}
}

// ToYaml parses the config template and returns yaml
func (secretCfg *SecretConfig) ToYaml() ([]byte, error) {
	var secretYaml bytes.Buffer

	tmpl, err := template.New("secretConfigTemplate").Parse(secretConfigTemplate)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&secretYaml, secretCfg)
	if err != nil {
		return nil, err
	}

	return secretYaml.Bytes(), nil
}
