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
	Genesis SecretGenesisPraos `json:"genesis,omitempty"`
	Bft     SecretBft          `json:"bft,omitempty"`
}

// SecretGenesisPraos ...
type SecretGenesisPraos struct {
	SigKey string `json:"sig_key,omitempty"`
	VrfKey string `json:"vrf_key,omitempty"`
	NodeID string `json:"node_id,omitempty"`
}

// SecretBft ...
type SecretBft struct {
	SigningKey string `json:"signing_key,omitempty"`
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
