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

{{- if .Rest.Listen }}
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
  {{- if and .TLS.CertFile .TLS.PrivKeyFile }}
  {{- with .TLS }}
  tls:
    cert_file: {{ .CertFile }}
    priv_key_file: {{ .PrivKeyFile }}
  {{end}}
  {{- end}}
{{end}}
{{- end}}

{{- with .P2P}}
p2p:
  {{- if .PublicAddress}}
  public_address: {{ .PublicAddress }}
  {{- end}}
  {{- if .ListenAddress}}
  listen_address: {{ .ListenAddress }}
  {{- end}}
  {{- with .TopicsOfInterest}}
  topics_of_interest:
    messages: {{ .Messages }}
    blocks: {{ .Blocks }}
  {{- end}}
  max_connections: {{ .MaxConnections }}
  max_client_connections: {{ .MaxClientConnections }}
  max_unreachable_nodes_to_connect_per_event: {{ .MaxUnreachableNodesToConnectPerEvent }}
  allow_private_addresses: {{ .AllowPrivateAddresses }}
  {{- if .TrustedPeers}}
  trusted_peers:
    {{- range .TrustedPeers}}
    - address: {{ .Address }}
      {{ if .ID }}
      id: {{ .ID -}}
      {{end}}
    {{- end}}
  {{- end}}
  {{- if .PublicID}}
  public_id: {{ .PublicID }}
  {{- end}}
  {{- with .Policy}}
  policy:
    quarantine_duration: {{ .QuarantineDuration }}
    {{- if .QuarantineWhitelist}}
    quarantine_whitelist:
      {{- range .QuarantineWhitelist}}
      - {{ . -}}
      {{end}}
    {{- end}}
  {{- end}}
  layers:
    preferred_list:
      view_max: {{ .Layers.PreferredList.ViewMax }}
      {{- if .Layers.PreferredList.Peers }}
      peers:
        {{- range .Layers.PreferredList.Peers}}
        - address: {{ .Address}}
          {{- if .ID }}
          id: {{ .ID -}}
          {{- end}}
        {{- end}}
      {{- end}}
  gossip_interval: {{ .GossipInterval }}
  {{- if .TopologyForceResetInterval}}
  topology_force_reset_interval: {{ .TopologyForceResetInterval }}
  {{- end}}
  {{- if .MaxBootstrapAttempts}}
  max_bootstrap_attempts: {{ .MaxBootstrapAttempts }}
  {{- end}}
{{end}}

{{- with .Log}}
log:
  - output: {{ .Output }}
    level:  {{ .Level }}
    format: {{ .Format }}
{{end}}

{{- with .Mempool}}
mempool:
    pool_max_entries: {{ .PoolMaxEntries }}
    log_max_entries: {{ .LogMaxEntries }}
{{end}}

{{- with .Leadership}}
leadership:
    logs_capacity: {{ .LogsCapacity }}
{{end}}

{{- if .SecretFiles}}
secret_files:
  {{- range .SecretFiles}}
  - {{ . }}
  {{- end}}
{{- end}}

no_blockchain_updates_warning_interval: {{ .NoBlockchainUpdatesWarningInterval }}

skip_bootstrap: {{ .SkipBootstrap }}

bootstrap_from_trusted_peers: {{ .BootstrapFromTrustedPeers }}

{{- if .HttpFetchBlock0Service}}
http_fetch_block0_service:
  {{- range .HttpFetchBlock0Service}}
  - {{ . }}
  {{- end}}
{{- end}}
`

// NodeConfig --config
type NodeConfig struct {
	Storage                            string           `json:"storage,omitempty"`
	SecretFiles                        []string         `json:"secret_files,omitempty"`
	HttpFetchBlock0Service             []string         `json:"http_fetch_block0_service,omitempty"`
	Explorer                           ConfigExplorer   `json:"explorer,omitempty"`
	Rest                               ConfigRest       `json:"rest,omitempty"`
	P2P                                ConfigP2P        `json:"p2p,omitempty"`
	Log                                ConfigLog        `json:"log,omitempty"`
	Mempool                            ConfigMempool    `json:"mempool,omitempty"`
	Leadership                         ConfigLeadership `json:"leadership,omitempty"`
	NoBlockchainUpdatesWarningInterval string           `json:"no_blockchain_updates_warning_interval,omitempty"`
	SkipBootstrap                      bool             `json:"skip_bootstrap,omitempty"`
	BootstrapFromTrustedPeers          bool             `json:"bootstrap_from_trusted_peers,omitempty"`
}

// ConfigP2P ...
type ConfigP2P struct {
	PublicAddress                        string                 `json:"public_address,omitempty"`
	ListenAddress                        string                 `json:"listen_address,omitempty"`
	PublicID                             string                 `json:"public_id,omitempty"`
	TrustedPeers                         []Peer                 `json:"trusted_peers,omitempty"`
	TopicsOfInterest                     ConfigTopicsOfInterest `json:"topics_of_interest,omitempty"`
	MaxConnections                       uint                   `json:"max_connections,omitempty"`
	MaxClientConnections                 uint                   `json:"max_client_connections,omitempty"`
	AllowPrivateAddresses                bool                   `json:"allow_private_addresses,omitempty"`
	Policy                               PolicyConfig           `json:"policy,omitempty"`
	MaxUnreachableNodesToConnectPerEvent uint                   `json:"max_unreachable_nodes_to_connect_per_event,omitempty"`
	GossipInterval                       string                 `json:"gossip_interval,omitempty"`
	TopologyForceResetInterval           string                 `json:"topology_force_reset_interval,omitempty"`
	MaxBootstrapAttempts                 int                    `json:"max_bootstrap_attempts,omitempty"`
	Layers                               Layers                 `json:"layers,omitempty"`
}

// Peer ...
type Peer struct {
	Address string `json:"address,omitempty"`
	// Deprecated - keep for compatibility
	ID string `json:"id,omitempty"`
}

// ConfigTopicsOfInterest ...
type ConfigTopicsOfInterest struct {
	Messages string `json:"messages"`
	Blocks   string `json:"blocks"`
}

// PolicyConfig ...
type PolicyConfig struct {
	QuarantineDuration      string   `json:"quarantine_duration,omitempty"`
	MaxQuarantine           string   `json:"max_quarantine,omitempty"`
	MaxNumQuarantineRecords uint     `json:"max_num_quarantine_records,omitempty"`
	QuarantineWhitelist     []string `json:"quarantine_whitelist,omitempty"`
}

// Layers ...
type Layers struct {
	PreferredList PreferredList `json:"preferred_list,omitempty"`
}

// PreferredList ...
type PreferredList struct {
	ViewMax uint   `json:"view_max,omitempty"`
	Peers   []Peer `json:"peers,omitempty"`
}

// ConfigRest ...
type ConfigRest struct {
	Listen string     `json:"listen,omitempty"`
	TLS    ConfigTLS  `json:"tls,omitempty"`
	Cors   ConfigCors `json:"cors,omitempty"`
}

// ConfigTLS ...
type ConfigTLS struct {
	CertFile    string `json:"cert_file,omitempty"`
	PrivKeyFile string `json:"priv_key_file,omitempty"`
}

// ConfigCors ...
type ConfigCors struct {
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
	MaxAgeSecs     int      `json:"max_age_secs,omitempty"`
}

// ConfigExplorer --enable-explorer
type ConfigExplorer struct {
	Enabled bool `json:"enabled,omitempty"`
}

// ConfigLog --log-level, --log-format, --log-output <log_output>
type ConfigLog struct {
	Level  string `json:"level,omitempty"`
	Format string `json:"format,omitempty"`
	Output string `json:"output,omitempty"`
	/*
		// FIXME: gelf through interface
		Output struct {
			Gelf ConfigGelf // `"gelf"`
		} // `"output"`
	*/
}

// ConfigGelf ...
type ConfigGelf struct {
	Backend string `json:"backend,omitempty"`
	LogID   string `json:"log_id,omitempty"`
}

// ConfigMempool ...
type ConfigMempool struct {
	PoolMaxEntries uint `json:"pool_max_entries,omitempty"`
	LogMaxEntries  uint `json:"log_max_entries,omitempty"`
}

// ConfigLeadership ...
type ConfigLeadership struct {
	LogsCapacity uint `json:"logs_capacity,omitempty"`
}

// NewNodeConfig ...
func NewNodeConfig() *NodeConfig {
	var nodeCfg NodeConfig

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Explorer.Enabled = false

	nodeCfg.Rest.Listen = "127.0.0.1:8443"
	nodeCfg.Rest.Cors.MaxAgeSecs = 0

	nodeCfg.P2P.PublicAddress = "/ip4/127.0.0.1/tcp/8299"
	nodeCfg.P2P.ListenAddress = "/ip4/127.0.0.1/tcp/8299"

	nodeCfg.P2P.TopicsOfInterest.Messages = "high"
	nodeCfg.P2P.TopicsOfInterest.Blocks = "high"

	nodeCfg.P2P.MaxConnections = 256
	nodeCfg.P2P.MaxClientConnections = 192

	nodeCfg.P2P.MaxUnreachableNodesToConnectPerEvent = 20
	nodeCfg.P2P.GossipInterval = "10s"

	nodeCfg.P2P.Policy.QuarantineDuration = "30m"

	// nodeCfg.P2P.MaxBootstrapAttempts = 0 // Fixme: unset and 0 have different effects.

	nodeCfg.P2P.Layers.PreferredList.ViewMax = 20

	nodeCfg.Log.Level = "trace"   // off, critical, error, warn, info, debug, trace
	nodeCfg.Log.Format = "plain"  // "json", "plain"
	nodeCfg.Log.Output = "stdout" // "stdout", "stderr", ...

	nodeCfg.Mempool.PoolMaxEntries = 100_000
	nodeCfg.Mempool.LogMaxEntries = 100_000

	nodeCfg.Leadership.LogsCapacity = 1_024

	nodeCfg.NoBlockchainUpdatesWarningInterval = "30m"

	nodeCfg.SkipBootstrap = false

	nodeCfg.BootstrapFromTrustedPeers = false

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

// AddTrustedPeer to node config
func (nodeCfg *NodeConfig) AddTrustedPeer(address string, id string) {
	nodeCfg.P2P.TrustedPeers = append(nodeCfg.P2P.TrustedPeers, Peer{address, id})
}

// AddHttpFetchBlock0Service to node config
func (nodeCfg *NodeConfig) AddHttpFetchBlock0Service(urlBlock0 string) {
	nodeCfg.HttpFetchBlock0Service = append(nodeCfg.HttpFetchBlock0Service, urlBlock0)
}

// AddPreferredList to node config
func (nodeCfg *NodeConfig) AddPreferredList(address string, id string) {
	nodeCfg.P2P.Layers.PreferredList.Peers = append(nodeCfg.P2P.Layers.PreferredList.Peers, Peer{address, id})
}

// AddQuarantineWhitelist to node config
func (nodeCfg *NodeConfig) AddQuarantineWhitelist(address string) {
	nodeCfg.P2P.Policy.QuarantineWhitelist = append(nodeCfg.P2P.Policy.QuarantineWhitelist, address)
}
