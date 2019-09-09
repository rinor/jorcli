package jnode

import (
	"fmt"
	"time"
)

// block0Config Genesis config
type block0Config struct {
	BlockchainConfiguration blockchainConfig    // `"blockchain_configuration"`
	Initial                 []blockchainInitial // `"initial"`
}

// blockchainConfig ...
type blockchainConfig struct {
	Discrimination                  string // `"discrimination"`
	Block0Consensus                 string // `"block0_consensus"`
	Block0Date                      int64  // `"block0_date"`
	SlotDuration                    int    // `"slot_duration"`
	SlotsPerEpoch                   int    // `"slots_per_epoch"`
	EpochStabilityDepth             int    // `"epoch_stability_depth"`
	KesUpdateSpeed                  int    // `"kes_update_speed"`
	MaxNumberOfTransactionsPerBlock int    // `"max_number_of_transactions_per_block"`

	BftSlotsRatio                        float64 // `"bft_slots_ratio"`
	ConsensusGenesisPraosActiveSlotCoeff float64 // `"consensus_genesis_praos_active_slot_coeff"`

	LinearFees struct {
		Certificate int // `"certificate"`
		Coefficient int // `"coefficient"`
		Constant    int // `"constant"`
	} // `"linear_fees"`

	ConsensusLeaderIds []string // `"consensus_leader_ids"`
}

// blockchainInitial ...
type blockchainInitial struct {
	Fund []initialFund // `"fund"`
	Cert string        // `"cert"`
}

// initialFund ...
type initialFund struct {
	Address string // `"address"`
	Value   uint64 // `"value"`
}

// NewBlock0Config provides default block0 config
func NewBlock0Config() *block0Config {
	var chainConfig blockchainConfig

	chainConfig.Discrimination = "test"
	chainConfig.Block0Consensus = "genesis_praos"
	chainConfig.Block0Date = time.Now().Unix()
	chainConfig.SlotDuration = 20
	chainConfig.SlotsPerEpoch = 30
	chainConfig.EpochStabilityDepth = 1
	chainConfig.KesUpdateSpeed = 43200
	chainConfig.BftSlotsRatio = 0.0
	chainConfig.ConsensusGenesisPraosActiveSlotCoeff = 0.1
	chainConfig.MaxNumberOfTransactionsPerBlock = 255
	chainConfig.LinearFees.Certificate = 0
	chainConfig.LinearFees.Coefficient = 0
	chainConfig.LinearFees.Constant = 0

	return &block0Config{
		BlockchainConfiguration: chainConfig,
	}
}

// AddConsensusLeader to block0 Blockchain Configuration
func (block0Cfg *block0Config) AddConsensusLeader(leaderPublicKey string) error {
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
func (block0Cfg *block0Config) AddInitialCertificate(cert string) error {
	// FIXME: check validity
	if cert == "" {
		return fmt.Errorf("parameter missing : %s", "cert")
	}

	block0Cfg.Initial = append(
		block0Cfg.Initial,
		blockchainInitial{Cert: cert},
	)

	return nil
}

// AddInitialFund to block0 Initial config
func (block0Cfg *block0Config) AddInitialFund(address string, value uint64) error {
	// FIXME: check validity
	if address == "" {
		return fmt.Errorf("parameter missing : %s", "address")
	}

	fundInit := initialFund{
		Address: address,
		Value:   value,
	}

	block0Cfg.Initial = append(
		block0Cfg.Initial,
		blockchainInitial{
			Fund: []initialFund{fundInit},
		},
	)

	return nil
}
