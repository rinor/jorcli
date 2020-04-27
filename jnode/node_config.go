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
      id: {{ .ID }}
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
          id: {{ .ID -}}
        {{end}}
      {{- end}}
  gossip_interval: {{ .GossipInterval }}
  {{- if .TopologyForceResetInterval}}
  topology_force_reset_interval: {{ .TopologyForceResetInterval }}
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
	Storage                            string           // `"storage"`
	SecretFiles                        []string         // `"secret_files"`
	HttpFetchBlock0Service             []string         // `"http_fetch_block0_service"`
	Explorer                           ConfigExplorer   // `"explorer"`
	Rest                               ConfigRest       // `"rest"`
	P2P                                ConfigP2P        // `"p2p"`
	Log                                ConfigLog        // `"log"`
	Mempool                            ConfigMempool    // `"mempool"`
	Leadership                         ConfigLeadership // `"leadership"`
	NoBlockchainUpdatesWarningInterval string           // `"no_blockchain_updates_warning_interval"`
	SkipBootstrap                      bool             // `"skip_bootstrap"`
	BootstrapFromTrustedPeers          bool             // `"bootstrap_from_trusted_peers"`
}

// ConfigP2P ...
type ConfigP2P struct {
	PublicAddress                        string                 // `"public_address"`
	ListenAddress                        string                 // `"listen_address"`
	PublicID                             string                 // `"public_id"`
	TrustedPeers                         []Peer                 // `"trusted_peers"`
	TopicsOfInterest                     ConfigTopicsOfInterest // `"topics_of_interest"`
	MaxConnections                       uint                   // `"max_connections"`
	MaxClientConnections                 uint                   // `"max_client_connections"`
	AllowPrivateAddresses                bool                   // `"allow_private_addresses"`
	Policy                               PolicyConfig           // `"policy"`
	MaxUnreachableNodesToConnectPerEvent uint                   // `"max_unreachable_nodes_to_connect_per_event"`
	GossipInterval                       string                 // `"gossip_interval"`
	TopologyForceResetInterval           string                 // `"topology_force_reset_interval"`
	MaxBootstrapAttempts                 int                    // `"max_bootstrap_attempts"`
	Layers                               Layers                 // `"layers"`
}

// Peer ...
type Peer struct {
	Address string // `"address"`
	ID      string // `"id"`
}

// ConfigTopicsOfInterest ...
type ConfigTopicsOfInterest struct {
	Messages string // `"messages"`
	Blocks   string // `"blocks"`
}

// PolicyConfig ...
type PolicyConfig struct {
	QuarantineDuration      string   // `"quarantine_duration"`
	MaxQuarantine           string   // `"max_quarantine"`
	MaxNumQuarantineRecords uint     // `"max_num_quarantine_records"`
	QuarantineWhitelist     []string // `"quarantine_whitelist"`
}

// Layers ...
type Layers struct {
	PreferredList PreferredList // `"preferred_list"`
}

// PreferredList ...
type PreferredList struct {
	ViewMax uint   // `"view_max"`
	Peers   []Peer // "`peers`"
}

// ConfigRest ...
type ConfigRest struct {
	Listen string     // `"listen"`
	TLS    ConfigTLS  // `"tls"`
	Cors   ConfigCors // `"cors"`
}

// ConfigTLS ...
type ConfigTLS struct {
	CertFile    string // `"cert_file"`
	PrivKeyFile string // `"priv_key_file"`
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
	PoolMaxEntries uint // `"pool_max_entries"`
	LogMaxEntries  uint // `"log_max_entries"`
}

// ConfigLeadership ...
type ConfigLeadership struct {
	LogsCapacity uint // `"logs_capacity"`
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

	nodeCfg.Mempool.PoolMaxEntries = 10_000
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
