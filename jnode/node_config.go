package jnode

import (
	"bytes"
	"text/template"
)

const nodeConfigTemplate = `
{{- if .Storage -}}
storage: {{ .Storage }}
{{end}}

{{- with .Explorer}}
explorer:
  enabled: {{ .Enabled }}
{{- end}}

{{- if .Rest.Enabled }}
{{with .Rest}}
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
{{- end}}

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
  {{- if .MaxConnections}}
  max_connections: {{ .MaxConnections }}
  {{- end}}
  {{- if .AllowPrivateAddresses}}
  allow_private_addresses: {{ .AllowPrivateAddresses }}
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
	PublicAddress         string                 // `"public_address"`
	ListenAddress         string                 // `"listen_address"`
	TrustedPeers          []string               // `"trusted_peers"`
	TopicsOfInterest      ConfigTopicsOfInterest // `"topics_of_interest"`
	MaxConnections        uint                   // `"max_connections"`
	AllowPrivateAddresses bool                   // `"allow_private_addresses"`
}

// ConfigTopicsOfInterest ...
type ConfigTopicsOfInterest struct {
	Messages string // `"messages"`
	Blocks   string // `"blocks"`
}

// ConfigRest ...
type ConfigRest struct {
	Enabled bool       // custom addition
	Listen  string     // `"listen"`
	Pkcs12  string     // `"pkcs12"`
	Cors    ConfigCors // `"cors"`
}

// ConfigCors ...
type ConfigCors struct {
	AllowedOrigins []string // `"allowed_origins"`
	MaxAgeSecs     int      // `"max_age_secs"`true
}

// ConfigExplorer --enable-explorer
type ConfigExplorer struct {
	Enabled bool // `"enabled"`
}

// ConfigLog --log-level, --log-format, --log-output <log_output>
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

	nodeCfg.Explorer.Enabled = false

	nodeCfg.Rest.Enabled = false
	nodeCfg.Rest.Listen = "127.0.0.1:8443"
	nodeCfg.Rest.Cors.MaxAgeSecs = 0

	nodeCfg.P2P.PublicAddress = "/ip4/127.0.0.1/tcp/8299"
	nodeCfg.P2P.ListenAddress = "/ip4/127.0.0.1/tcp/8299"
	nodeCfg.P2P.TopicsOfInterest.Messages = "high"
	nodeCfg.P2P.TopicsOfInterest.Blocks = "high"
	nodeCfg.P2P.MaxConnections = 256

	nodeCfg.Log.Level = "trace"   // off, critical, error, warn, info, debug, trace
	nodeCfg.Log.Format = "plain"  // "json", "plain"
	nodeCfg.Log.Output = "stdout" // "stdout", "stderr", ...

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

	tmpl, err := template.New("nodeConfigTemplate").Parse(nodeConfigTemplate)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&cfgYaml, nodeCfg)
	if err != nil {
		return nil, err
	}

	return cfgYaml.Bytes(), nil
}

// AddSecretFile to node config
func (nodeCfg *NodeConfig) AddSecretFile(secretFile string) {
	nodeCfg.SecretFiles = append(nodeCfg.SecretFiles, secretFile)
}
