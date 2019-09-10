package jnode

import (
	"bytes"
	"text/template"
)

const nodeConfigTemplate = `
{{- if .Storage}}
storage: {{ .Storage }}
{{end}}

{{- with .Explorer}}
explorer:
  enabled: {{ .Enabled }}
{{end}}

{{- with .Rest}}
rest:
  listen: {{ .Listen }}
  {{- with .Cors}}
  cors:
    {{- if .AllowedOrigins}}
    allowed_origins:
      {{- range .AllowedOrigins}}
      - {{ . }}
      {{- end}}
    {{- end}}
    max_age_secs: {{ .MaxAgeSecs }}
  {{- end}}
  {{- if .Pkcs12}}
  pkcs12: {{ .Pkcs12 }}
  {{- end}}
{{end}}

{{- with .P2P}}
p2p:
  public_address: {{ .PublicAddress }}
  {{- if .ListenAddress}}
  listen_address: {{ .ListenAddress }}
  {{- end}}
  {{- with .TopicsOfInterest}}
  topics_of_interest:
    messages: {{ .Messages }}
    blocks: {{ .Blocks }}
  {{- end}}

  {{- if .TrustedPeers}}
  trusted_peers:
    {{- range .TrustedPeers}}
	- {{ . }}
    {{- end}}
  {{- end}}
{{end}}

{{- with .Log}}
log:
  level:  {{ .Level }}
  format: {{ .Format }}
  output: {{ .Output }}
{{end}}

{{- with .Mempool}}
mempool:
    fragment_ttl: {{ .FragmentTTL }}
    log_ttl: {{ .LogTTL }}
    garbage_collection_interval: {{ .GarbageCollectionInterval }}
{{end}}

{{- with .Leadership}}
leadership:
    log_ttl: {{ .LogTTL }}
    garbage_collection_interval: {{ .GarbageCollectionInterval }}
{{end}}

{{- if .SecretFiles}}
secret_files:
  {{- range .SecretFiles}}
  - {{ . }}
  {{- end}}
{{- end}}
`

var _ = nodeConfigTemplate

// NodeConfig --config
type NodeConfig struct {
	Storage     string           // `"storage"`
	SecretFiles []string         // `"secret_files"`
	Explorer    ConfigExplorer   // `"explorer"`
	Rest        ConfigRest       // `"rest"`
	P2P         ConfigP2P        // `"p2p"`
	Log         ConfigLog        // `"log"`
	Mempool     ConfigMempool    // `"mempool"`
	Leadership  ConfigLeadership // `"leadership"`
}

// ConfigP2P ...
type ConfigP2P struct {
	PublicAddress    string                 // `"public_address"`
	ListenAddress    string                 // `"listen_address"`
	TrustedPeers     []string               // `"trusted_peers"`
	TopicsOfInterest ConfigTopicsOfInterest // `"topics_of_interest"`
}

// ConfigTopicsOfInterest ...
type ConfigTopicsOfInterest struct {
	Messages string // `"messages"`
	Blocks   string // `"blocks"`
}

// ConfigRest ...
type ConfigRest struct {
	Listen string     // `"listen"`
	Pkcs12 string     // `"pkcs12"`
	Cors   ConfigCors // `"cors"`
}

// ConfigCors ...
type ConfigCors struct {
	AllowedOrigins []string // `"allowed_origins"`
	MaxAgeSecs     int      // `"max_age_secs"`
}

// ConfigExplorer ...
type ConfigExplorer struct {
	Enabled bool // `"enabled"`127.0.0.2
}

// ConfigLog ...
type ConfigLog struct {
	Level  string // `"level"`
	Format string // `"format"`
	Output string // `"output"`
	/*
		// FIXME: gelf through interface
		Output struct {
			Gelf ConfigGelf // `"gelf"`
		} // `"output"`
	*/
}

// ConfigGelf ...
type ConfigGelf struct {
	Backend string // `"backend"`
	LogID   string // `"log_id"`
}

// ConfigMempool ...
type ConfigMempool struct {
	FragmentTTL               string // `"fragment_ttl"`
	LogTTL                    string // `"log_ttl"`
	GarbageCollectionInterval string // `"garbage_collection_interval"`
}

// ConfigLeadership ...
type ConfigLeadership struct {
	LogTTL                    string // `"log_ttl"`
	GarbageCollectionInterval string // `"garbage_collection_interval"`
}

// NewNodeConfig ...
func NewNodeConfig() *NodeConfig {
	var nodeCfg NodeConfig

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Explorer.Enabled = true

	nodeCfg.Rest.Listen = "127.0.0.1:8443"
	// nodeCfg.Rest.Pkcs12 = "00000000000000"
	nodeCfg.Rest.Cors.MaxAgeSecs = 0
	// nodeCfg.Rest.Cors.AllowedOrigins = append(nodeCfg.Rest.Cors.AllowedOrigins, "*")

	nodeCfg.P2P.PublicAddress = "/ip4/127.0.0.1/tcp/8299"
	// nodeCfg.P2P.ListenAddress = "/ip4/127.0.0.1/tcp/8299"
	nodeCfg.P2P.TopicsOfInterest.Messages = "high"
	nodeCfg.P2P.TopicsOfInterest.Blocks = "high"
	// nodeCfg.P2P.TrustedPeers = append(nodeCfg.P2P.TrustedPeers, "/ip4/127.0.0.2/tcp/8299")
	// nodeCfg.P2P.TrustedPeers = append(nodeCfg.P2P.TrustedPeers, "/ip4/127.0.0.3/tcp/8299")

	// nodeCfg.SecretFiles = append(nodeCfg.SecretFiles, "secret_01.key")
	// nodeCfg.SecretFiles = append(nodeCfg.SecretFiles, "secret_02.key")

	nodeCfg.Log.Level = "trace"
	nodeCfg.Log.Format = "yaml"
	nodeCfg.Log.Output = "stdout"

	nodeCfg.Mempool.FragmentTTL = "30m"
	nodeCfg.Mempool.LogTTL = "1h"
	nodeCfg.Mempool.GarbageCollectionInterval = "15m"

	nodeCfg.Leadership.LogTTL = "1h"
	nodeCfg.Leadership.GarbageCollectionInterval = "15m"

	return &nodeCfg
}

// ToYaml parses the config template and returns yaml
func (nodeCfg *NodeConfig) ToYaml() ([]byte, error) {
	var cfgYaml bytes.Buffer

	t := template.Must(template.New("nodeConfigTemplate").Parse(nodeConfigTemplate))

	err := t.Execute(&cfgYaml, nodeCfg)
	if err != nil {
		return nil, err
	}

	return cfgYaml.Bytes(), nil
}
