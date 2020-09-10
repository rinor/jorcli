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
  treasury: {{ .Treasury }}
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
  {{with .RewardConstraints -}}
  {{- if and .RewardDrawingLimitMax .PoolParticipationCapping .PoolParticipationCapping.Min .PoolParticipationCapping.Max }}
  reward_constraints:
	reward_drawing_limit_max: {{ .RewardDrawingLimitMax }}
	pool_participation_capping:
	  min: {{ .PoolParticipationCapping.Min }}
      max: {{ .PoolParticipationCapping.Max }}
  {{- end}}
  {{- end}}
  {{- if .FeesGoTo}}
  fees_go_to: {{ .FeesGoTo }}
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
    {{- with .PerVoteCertificateFees -}}
    {{- if or .CertificateVoteCast .CertificateVotePlan }}
    per_vote_certificate_fees:
      {{- if .CertificateVoteCast }}
      certificate_vote_cast: {{ .CertificateVoteCast }}
      {{- end}}
      {{- if .CertificateVotePlan }}
      certificate_vote_plan: {{ .CertificateVotePlan }}
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
  {{- if .Committees}}
  committees:
    {{- range .Committees}}
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
  {{- with .LegacyFund}}
  - legacy_fund:
      {{- range .}}
      - address: {{ .Address }}
        value: {{ .Value}}
      {{- end -}}
  {{- end -}}
{{- end}}
{{- end}}
`

// Block0Config Genesis config
type Block0Config struct {
	BlockchainConfiguration BlockchainConfig    `json:"blockchain_configuration"`
	Initial                 []BlockchainInitial `json:"initial"`
}

// BlockchainConfig ...
type BlockchainConfig struct {
	Discrimination                       string   `json:"discrimination"`
	Block0Consensus                      string   `json:"block0_consensus"`
	Block0Date                           int64    `json:"block0_date"`
	SlotDuration                         uint8    `json:"slot_duration"`                             // [1 - 255] def=5 seconds
	SlotsPerEpoch                        uint32   `json:"slots_per_epoch"`                           // [1 - 1_000_000] def=720
	EpochStabilityDepth                  uint32   `json:"epoch_stability_depth"`                     // def=102_400
	KesUpdateSpeed                       uint32   `json:"kes_update_speed"`                          // [60 - 365*24*3600] def=43_200 seconds (12 * 3600)
	BlockContentMaxSize                  uint32   `json:"block_content_max_size"`                    // def=102_400
	ConsensusGenesisPraosActiveSlotCoeff float64  `json:"consensus_genesis_praos_active_slot_coeff"` // [0.001 - 1.000] def=0.1
	ConsensusLeaderIds                   []string `json:"consensus_leader_ids"`                      // bft leaders

	// Fees
	LinearFees LinearFees `json:"linear_fees"`
	FeesGoTo   string     `json:"fees_go_to"` // default is "rewards", can be "treasury"

	// Treasury
	Treasury           uint64             `json:"treasury"`
	TreasuryParameters TreasuryParameters `json:"treasury_parameters"`

	// Rewards
	TotalRewardSupply uint64            `json:"total_reward_supply"`
	RewardParameters  RewardParameters  `json:"reward_parameters"`
	RewardConstraints RewardConstraints `json:"reward_constraints"`

	Committees []string `json:"committees"`
}

// LinearFees ...
type LinearFees struct {
	Certificate            uint64                 `json:"certificate"`
	Coefficient            uint64                 `json:"coefficient"`
	Constant               uint64                 `json:"constant"`
	PerCertificateFees     PerCertificateFees     `json:"per_certificate_fees"`
	PerVoteCertificateFees PerVoteCertificateFees `json:"per_vote_certificate_fees"`
}

// FIXME: PerCertificateFees - check/handle 0 values on config

// PerCertificateFees ...
type PerCertificateFees struct {
	CertificatePoolRegistration     uint64 `json:"certificate_pool_registration"`
	CertificateStakeDelegation      uint64 `json:"certificate_stake_delegation"`
	CertificateOwnerStakeDelegation uint64 `json:"certificate_owner_stake_delegation"`
}

// PerVoteCertificateFees ...
type PerVoteCertificateFees struct {
	CertificateVoteCast uint64 `json:"certificate_vote_cast"`
	CertificateVotePlan uint64 `json:"certificate_vote_plan"`
}

// FIXME: TreasuryParameters - check/handle 0 values on config

// TreasuryParameters ...
type TreasuryParameters struct {
	Fixed    uint64 `json:"fixed"`
	Ratio    string `json:"ratio"`
	MaxLimit uint64 `json:"max_limit"`
}

// RewardParameters ...
type RewardParameters struct {
	Formula    string `json:"-"` // `"halving"` or `"linear"`
	Constant   uint64 `json:"constant"`
	Ratio      string `json:"ratio"`
	EpochStart uint32 `json:"epoch_start"`
	EpochRate  uint32 `json:"epoch_rate"`
}

// RewardConstraints ...
type RewardConstraints struct {
	RewardDrawingLimitMax    string                   `json:"reward_drawing_limit_max"`
	PoolParticipationCapping PoolParticipationCapping `json:"pool_participation_capping"`
}

// PoolParticipationCapping ...
type PoolParticipationCapping struct {
	Min uint32 `json:"min"`
	Max uint32 `json:"max"`
}

// BlockchainInitial ...
type BlockchainInitial struct {
	Fund       []InitialFund `json:"fund"`
	Cert       string        `json:"cert"`
	LegacyFund []InitialFund `json:"legacy_fund"`
}

// InitialFund ...
type InitialFund struct {
	Address string `json:"address"`
	Value   uint64 `json:"value"`
}

// NewBlock0Config provides default block0 config
func NewBlock0Config() *Block0Config {
	var chainConfig BlockchainConfig

	chainConfig.Discrimination = "test"
	chainConfig.Block0Consensus = "genesis_praos"
	chainConfig.Block0Date = time.Now().Unix()
	chainConfig.SlotDuration = 5
	chainConfig.SlotsPerEpoch = 720
	chainConfig.EpochStabilityDepth = 102_400
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

	chainConfig.FeesGoTo = "rewards"

	chainConfig.LinearFees.Certificate = 0
	chainConfig.LinearFees.Coefficient = 0
	chainConfig.LinearFees.Constant = 0

	chainConfig.LinearFees.PerCertificateFees.CertificatePoolRegistration = 0
	chainConfig.LinearFees.PerCertificateFees.CertificateStakeDelegation = 0
	chainConfig.LinearFees.PerCertificateFees.CertificateOwnerStakeDelegation = 0

	chainConfig.LinearFees.PerVoteCertificateFees.CertificateVoteCast = 0
	chainConfig.LinearFees.PerVoteCertificateFees.CertificateVotePlan = 0

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

// AddInitialLegacyFund to block0 Initial config
func (block0Cfg *Block0Config) AddInitialLegacyFund(address string, value uint64) error {
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
			LegacyFund: []InitialFund{fundInit},
		},
	)

	return nil
}

// AddCommittee to block0 Initial config
func (block0Cfg *Block0Config) AddCommittee(cid string) error {
	// FIXME: check validity
	if cid == "" {
		return fmt.Errorf("parameter missing : %s", "cid")
	}

	block0Cfg.BlockchainConfiguration.Committees = append(
		block0Cfg.BlockchainConfiguration.Committees,
		cid,
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
