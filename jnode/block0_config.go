package jnode

import (
	"bytes"
	"fmt"
	"text/template"
	"time"
)

// block0ConfigTemplate ...
const block0ConfigTemplate = `
{{- with .BlockchainConfiguration -}}
blockchain_configuration:
  block0_date: {{ .Block0Date }}
  discrimination: {{ .Discrimination }}
  block0_consensus: {{ .Block0Consensus }}
  slots_per_epoch: {{ .SlotsPerEpoch }}
  slot_duration: {{ .SlotDuration }}
  kes_update_speed: {{ .KesUpdateSpeed }}
  consensus_genesis_praos_active_slot_coeff: {{ .ConsensusGenesisPraosActiveSlotCoeff }}
  block_content_max_size: {{ .BlockContentMaxSize }}
  epoch_stability_depth: {{ .EpochStabilityDepth }}
  {{- if .Treasury}}
  treasury: {{ .Treasury }}
  {{- end}}
  {{with .TreasuryParameters -}}
  treasury_parameters:
    fixed: {{ .Fixed }}
    ratio: {{ .Ratio }}
    {{- if .MaxLimit }}
    max_limit: {{ .MaxLimit }}
    {{- end}}
  {{- end}}
  total_reward_supply: {{ .TotalRewardSupply }}
  {{with .RewardParameters -}}
  {{- if and .Formula .Constant .Ratio .EpochRate }}
  reward_parameters:
    {{ .Formula }}:
      constant: {{ .Constant }}
      ratio: {{ .Ratio }}
      epoch_start: {{ .EpochStart }}
      epoch_rate: {{ .EpochRate }}
  {{- end}}
  {{- end}}
  {{with .LinearFees -}}
  linear_fees:
    constant: {{ .Constant }}
    coefficient: {{ .Coefficient }}
    certificate: {{ .Certificate }}
    {{- with .PerCertificateFees -}}
    {{- if or .CertificatePoolRegistration .CertificateStakeDelegation .CertificateOwnerStakeDelegation }}
    per_certificate_fees:
      {{- if .CertificatePoolRegistration }}
      certificate_pool_registration: {{ .CertificatePoolRegistration }}
      {{- end}}
      {{- if .CertificateStakeDelegation }}
      certificate_stake_delegation: {{ .CertificateStakeDelegation }}
      {{- end}}
      {{- if .CertificateOwnerStakeDelegation }}
      certificate_owner_stake_delegation: {{ .CertificateOwnerStakeDelegation }}
      {{- end}}
    {{- end}}
    {{- end}}
  {{- end}}
  {{- if .ConsensusLeaderIds}}
  consensus_leader_ids:
    {{- range .ConsensusLeaderIds}}
    - {{ . -}}
    {{end}}
  {{end}}
{{end}}

{{- with .Initial -}}
initial:
{{- range .}}
  {{- with .Fund}}
  - fund:
      {{- range .}}
      - address: {{ .Address }}
        value: {{ .Value}}
      {{- end -}}
  {{- end -}}
  {{if .Cert}}
  - cert: {{ .Cert }}
  {{- end}}
{{- end}}
{{- end}}
`

// Block0Config Genesis config
type Block0Config struct {
	BlockchainConfiguration BlockchainConfig    // `"blockchain_configuration"`
	Initial                 []BlockchainInitial // `"initial"`
}

// BlockchainConfig ...
type BlockchainConfig struct {
	Discrimination                       string   // `"discrimination"`
	Block0Consensus                      string   // `"block0_consensus"`
	Block0Date                           int64    // `"block0_date"`
	SlotDuration                         uint8    // `"slot_duration"`
	SlotsPerEpoch                        uint32   // `"slots_per_epoch"`
	EpochStabilityDepth                  uint32   // `"epoch_stability_depth"`
	KesUpdateSpeed                       uint32   // `"kes_update_speed"`
	BlockContentMaxSize                  uint32   // `"block_content_max_size"`
	ConsensusGenesisPraosActiveSlotCoeff float64  // `"consensus_genesis_praos_active_slot_coeff"`
	ConsensusLeaderIds                   []string // `"consensus_leader_ids"`
	// Fees
	LinearFees LinearFees // `"linear_fees"`
	// Treasury
	Treasury           uint64             // `"treasury"`
	TreasuryParameters TreasuryParameters // `"treasury_parameters"`
	// Rewards
	TotalRewardSupply uint64           // `"total_reward_supply"`
	RewardParameters  RewardParameters // `"reward_parameters"`

}

// LinearFees ...
type LinearFees struct {
	Certificate        uint64             // `"certificate"`
	Coefficient        uint64             // `"coefficient"`
	Constant           uint64             // `"constant"`
	PerCertificateFees PerCertificateFees // `"per_certificate_fees"`
}

// FIXME: PerCertificateFees - check/handle 0 values on config

// PerCertificateFees ...
type PerCertificateFees struct {
	CertificatePoolRegistration     uint64 // `"certificate_pool_registration"`
	CertificateStakeDelegation      uint64 // `"certificate_stake_delegation"`
	CertificateOwnerStakeDelegation uint64 // `"certificate_owner_stake_delegation"`
}

// FIXME: TreasuryParameters - check/handle 0 values on config

// TreasuryParameters ...
type TreasuryParameters struct {
	Fixed    uint64 // `"fixed"`
	Ratio    string // `"ratio"`
	MaxLimit uint64 // `"max_limit"`
}

// RewardParameters ...
type RewardParameters struct {
	Formula    string // `"halving"` or `"linear"`
	Constant   uint64 // `"constant"`
	Ratio      string // `"ratio"`
	EpochStart uint32 // `"epoch_start"`
	EpochRate  uint32 // `"epoch_rate"`
}

// BlockchainInitial ...
type BlockchainInitial struct {
	Fund []InitialFund // `"fund"`
	Cert string        // `"cert"`
}

// InitialFund ...
type InitialFund struct {
	Address string // `"address"`
	Value   uint64 // `"value"`
}

// NewBlock0Config provides default block0 config
func NewBlock0Config() *Block0Config {
	var chainConfig BlockchainConfig

	chainConfig.Discrimination = "test"
	chainConfig.Block0Consensus = "genesis_praos"
	chainConfig.Block0Date = time.Now().Unix()
	chainConfig.SlotDuration = 2
	chainConfig.SlotsPerEpoch = 43_200
	chainConfig.EpochStabilityDepth = 10
	chainConfig.KesUpdateSpeed = 43_200
	chainConfig.ConsensusGenesisPraosActiveSlotCoeff = 0.1
	chainConfig.BlockContentMaxSize = 102_400

	chainConfig.Treasury = 0
	chainConfig.TreasuryParameters.Fixed = 0
	chainConfig.TreasuryParameters.Ratio = "0/1"
	chainConfig.TreasuryParameters.MaxLimit = 0

	chainConfig.TotalRewardSupply = 0
	chainConfig.RewardParameters.Formula = "linear"
	chainConfig.RewardParameters.Constant = 0
	chainConfig.RewardParameters.Ratio = "0/1"
	chainConfig.RewardParameters.EpochStart = 0
	chainConfig.RewardParameters.EpochRate = 0

	chainConfig.LinearFees.Certificate = 10_000
	chainConfig.LinearFees.Coefficient = 50
	chainConfig.LinearFees.Constant = 1_000

	chainConfig.LinearFees.PerCertificateFees.CertificatePoolRegistration = 10_000
	chainConfig.LinearFees.PerCertificateFees.CertificateStakeDelegation = 10_000
	chainConfig.LinearFees.PerCertificateFees.CertificateOwnerStakeDelegation = 10_000

	return &Block0Config{
		BlockchainConfiguration: chainConfig,
	}
}

// AddConsensusLeader to block0 Blockchain Configuration
func (block0Cfg *Block0Config) AddConsensusLeader(leaderPublicKey string) error {
	// FIXME: check validity
	if leaderPublicKey == "" {
		return fmt.Errorf("parameter missing : %s", "leaderPublicKey")
	}

	block0Cfg.BlockchainConfiguration.ConsensusLeaderIds = append(
		block0Cfg.BlockchainConfiguration.ConsensusLeaderIds,
		leaderPublicKey,
	)

	return nil
}

// AddInitialCertificate to block0 Initial config
func (block0Cfg *Block0Config) AddInitialCertificate(cert string) error {
	// FIXME: check validity
	if cert == "" {
		return fmt.Errorf("parameter missing : %s", "cert")
	}

	block0Cfg.Initial = append(
		block0Cfg.Initial,
		BlockchainInitial{Cert: cert},
	)

	return nil
}

// AddInitialFund to block0 Initial config
func (block0Cfg *Block0Config) AddInitialFund(address string, value uint64) error {
	// FIXME: check validity
	if address == "" {
		return fmt.Errorf("parameter missing : %s", "address")
	}

	fundInit := InitialFund{
		Address: address,
		Value:   value,
	}

	block0Cfg.Initial = append(
		block0Cfg.Initial,
		BlockchainInitial{
			Fund: []InitialFund{fundInit},
		},
	)

	return nil
}

// ToYaml parses the config template and returns yaml
func (block0Cfg *Block0Config) ToYaml() ([]byte, error) {
	var block0Yaml bytes.Buffer

	tmpl, err := template.New("block0ConfigTemplate").Parse(block0ConfigTemplate)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&block0Yaml, block0Cfg)
	if err != nil {
		return nil, err
	}

	return block0Yaml.Bytes(), nil
}
