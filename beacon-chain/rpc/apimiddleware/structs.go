package apimiddleware

import (
	"strings"

	"github.com/prysmaticlabs/prysm/v3/api/gateway/apimiddleware"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/rpc/eth/helpers"
	ethpbv2 "github.com/prysmaticlabs/prysm/v3/proto/eth/v2"
)

//----------------
// Requests and responses.
//----------------

type genesisResponseJson struct {
	Data *genesisResponse_GenesisJson `json:"data"`
}

type genesisResponse_GenesisJson struct {
	GenesisTime           string                  `json:"genesis_time" time:"true"`
	GenesisValidatorsRoot apimiddleware.HexString `json:"genesis_validators_root"`
	GenesisForkVersion    apimiddleware.HexString `json:"genesis_fork_version"`
}

// WeakSubjectivityResponse is used to marshal/unmarshal the response for the
// /eth/v1/beacon/weak_subjectivity endpoint.
type WeakSubjectivityResponse struct {
	Data *struct {
		Checkpoint *checkpointJson         `json:"ws_checkpoint"`
		StateRoot  apimiddleware.HexString `json:"state_root"`
	} `json:"data"`
}

type feeRecipientsRequestJSON struct {
	Recipients []*feeRecipientJson `json:"recipients"`
}

type stateRootResponseJson struct {
	Data                *stateRootResponse_StateRootJson `json:"data"`
	ExecutionOptimistic bool                             `json:"execution_optimistic"`
}

type stateRootResponse_StateRootJson struct {
	StateRoot apimiddleware.HexString `json:"root"`
}

type stateForkResponseJson struct {
	Data                *forkJson `json:"data"`
	ExecutionOptimistic bool      `json:"execution_optimistic"`
}

type stateFinalityCheckpointResponseJson struct {
	Data                *stateFinalityCheckpointResponse_StateFinalityCheckpointJson `json:"data"`
	ExecutionOptimistic bool                                                         `json:"execution_optimistic"`
}

type stateFinalityCheckpointResponse_StateFinalityCheckpointJson struct {
	PreviousJustified *checkpointJson `json:"previous_justified"`
	CurrentJustified  *checkpointJson `json:"current_justified"`
	Finalized         *checkpointJson `json:"finalized"`
}

type stateValidatorsResponseJson struct {
	Data                []*validatorContainerJson `json:"data"`
	ExecutionOptimistic bool                      `json:"execution_optimistic"`
}

type stateValidatorResponseJson struct {
	Data                *validatorContainerJson `json:"data"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type validatorBalancesResponseJson struct {
	Data                []*validatorBalanceJson `json:"data"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type stateCommitteesResponseJson struct {
	Data                []*committeeJson `json:"data"`
	ExecutionOptimistic bool             `json:"execution_optimistic"`
}

type syncCommitteesResponseJson struct {
	Data                *syncCommitteeValidatorsJson `json:"data"`
	ExecutionOptimistic bool                         `json:"execution_optimistic"`
}

type blockHeadersResponseJson struct {
	Data                []*blockHeaderContainerJson `json:"data"`
	ExecutionOptimistic bool                        `json:"execution_optimistic"`
}

type blockHeaderResponseJson struct {
	Data                *blockHeaderContainerJson `json:"data"`
	ExecutionOptimistic bool                      `json:"execution_optimistic"`
}

type blockResponseJson struct {
	Data *signedBeaconBlockContainerJson `json:"data"`
}

type blockV2ResponseJson struct {
	Version             string                            `json:"version" enum:"true"`
	Data                *signedBeaconBlockContainerV2Json `json:"data"`
	ExecutionOptimistic bool                              `json:"execution_optimistic"`
}

type blockRootResponseJson struct {
	Data                *blockRootContainerJson `json:"data"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type blockAttestationsResponseJson struct {
	Data                []*attestationJson `json:"data"`
	ExecutionOptimistic bool               `json:"execution_optimistic"`
}

type attestationsPoolResponseJson struct {
	Data []*attestationJson `json:"data"`
}

type submitAttestationRequestJson struct {
	Data []*attestationJson `json:"data"`
}

type attesterSlashingsPoolResponseJson struct {
	Data []*attesterSlashingJson `json:"data"`
}

type proposerSlashingsPoolResponseJson struct {
	Data []*proposerSlashingJson `json:"data"`
}

type voluntaryExitsPoolResponseJson struct {
	Data []*signedVoluntaryExitJson `json:"data"`
}

type submitSyncCommitteeSignaturesRequestJson struct {
	Data []*syncCommitteeMessageJson `json:"data"`
}

type identityResponseJson struct {
	Data *identityJson `json:"data"`
}

type peersResponseJson struct {
	Data []*peerJson `json:"data"`
}

type peerResponseJson struct {
	Data *peerJson `json:"data"`
}

type peerCountResponseJson struct {
	Data peerCountResponse_PeerCountJson `json:"data"`
}

type peerCountResponse_PeerCountJson struct {
	Disconnected  string `json:"disconnected"`
	Connecting    string `json:"connecting"`
	Connected     string `json:"connected"`
	Disconnecting string `json:"disconnecting"`
}

type versionResponseJson struct {
	Data *versionJson `json:"data"`
}

type syncingResponseJson struct {
	Data *helpers.SyncDetailsJson `json:"data"`
}

type beaconStateResponseJson struct {
	Data *beaconStateJson `json:"data"`
}

type beaconStateV2ResponseJson struct {
	Version             string                      `json:"version" enum:"true"`
	Data                *beaconStateContainerV2Json `json:"data"`
	ExecutionOptimistic bool                        `json:"execution_optimistic"`
}

type forkChoiceHeadsResponseJson struct {
	Data []*forkChoiceHeadJson `json:"data"`
}

type v2ForkChoiceHeadsResponseJson struct {
	Data []*v2ForkChoiceHeadJson `json:"data"`
}

type forkScheduleResponseJson struct {
	Data []*forkJson `json:"data"`
}

type depositContractResponseJson struct {
	Data *depositContractJson `json:"data"`
}

type specResponseJson struct {
	Data interface{} `json:"data"`
}

type dutiesRequestJson struct {
	Index []string `json:"index"`
}

type attesterDutiesResponseJson struct {
	DependentRoot       apimiddleware.HexString `json:"dependent_root"`
	Data                []*attesterDutyJson     `json:"data"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type proposerDutiesResponseJson struct {
	DependentRoot       apimiddleware.HexString `json:"dependent_root"`
	Data                []*proposerDutyJson     `json:"data"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type syncCommitteeDutiesResponseJson struct {
	Data                []*syncCommitteeDuty `json:"data"`
	ExecutionOptimistic bool                 `json:"execution_optimistic"`
}

type produceBlockResponseJson struct {
	Data *beaconBlockJson `json:"data"`
}

type produceBlockResponseV2Json struct {
	Version string                      `json:"version"`
	Data    *beaconBlockContainerV2Json `json:"data"`
}

type produceBlindedBlockResponseJson struct {
	Version string                           `json:"version"`
	Data    *blindedBeaconBlockContainerJson `json:"data"`
}

type produceAttestationDataResponseJson struct {
	Data *attestationDataJson `json:"data"`
}

type aggregateAttestationResponseJson struct {
	Data *attestationJson `json:"data"`
}

type submitBeaconCommitteeSubscriptionsRequestJson struct {
	Data []*beaconCommitteeSubscribeJson `json:"data"`
}

type beaconCommitteeSubscribeJson struct {
	ValidatorIndex   string `json:"validator_index"`
	CommitteeIndex   string `json:"committee_index"`
	CommitteesAtSlot string `json:"committees_at_slot"`
	Slot             string `json:"slot"`
	IsAggregator     bool   `json:"is_aggregator"`
}

type submitSyncCommitteeSubscriptionRequestJson struct {
	Data []*syncCommitteeSubscriptionJson `json:"data"`
}

type syncCommitteeSubscriptionJson struct {
	ValidatorIndex       string   `json:"validator_index"`
	SyncCommitteeIndices []string `json:"sync_committee_indices"`
	UntilEpoch           string   `json:"until_epoch"`
}

type submitAggregateAndProofsRequestJson struct {
	Data []*signedAggregateAttestationAndProofJson `json:"data"`
}

type produceSyncCommitteeContributionResponseJson struct {
	Data *syncCommitteeContributionJson `json:"data"`
}

type submitContributionAndProofsRequestJson struct {
	Data []*signedContributionAndProofJson `json:"data"`
}

type forkchoiceResponse struct {
	JustifiedCheckpoint           *checkpointJson         `json:"justified_checkpoint"`
	FinalizedCheckpoint           *checkpointJson         `json:"finalized_checkpoint"`
	BestJustifiedCheckpoint       *checkpointJson         `json:"best_justified_checkpoint"`
	UnrealizedJustifiedCheckpoint *checkpointJson         `json:"unrealized_justified_checkpoint"`
	UnrealizedFinalizedCheckpoint *checkpointJson         `json:"unrealized_finalized_checkpoint"`
	ProposerBoostRoot             apimiddleware.HexString `json:"proposer_boost_root"`
	PreviousProposerBoostRoot     apimiddleware.HexString `json:"previous_proposer_boost_root"`
	HeadRoot                      apimiddleware.HexString `json:"head_root"`
	ForkChoiceNodes               []*forkChoiceNodeJson   `json:"forkchoice_nodes"`
}

//----------------
// Reusable types.
//----------------

type checkpointJson struct {
	Epoch string                  `json:"epoch"`
	Root  apimiddleware.HexString `json:"root"`
}

type blockRootContainerJson struct {
	Root apimiddleware.HexString `json:"root"`
}

type signedBeaconBlockContainerJson struct {
	Message   *beaconBlockJson        `json:"message"`
	Signature apimiddleware.HexString `json:"signature"`
}

type beaconBlockJson struct {
	Slot          string                  `json:"slot"`
	ProposerIndex string                  `json:"proposer_index"`
	ParentRoot    apimiddleware.HexString `json:"parent_root"`
	StateRoot     apimiddleware.HexString `json:"state_root"`
	Body          *beaconBlockBodyJson    `json:"body"`
}

type beaconBlockBodyJson struct {
	RandaoReveal      apimiddleware.HexString    `json:"randao_reveal"`
	Eth1Data          *eth1DataJson              `json:"eth1_data"`
	Graffiti          apimiddleware.HexString    `json:"graffiti"`
	ProposerSlashings []*proposerSlashingJson    `json:"proposer_slashings"`
	AttesterSlashings []*attesterSlashingJson    `json:"attester_slashings"`
	Attestations      []*attestationJson         `json:"attestations"`
	Deposits          []*depositJson             `json:"deposits"`
	VoluntaryExits    []*signedVoluntaryExitJson `json:"voluntary_exits"`
}

type signedBeaconBlockContainerV2Json struct {
	Phase0Block    *beaconBlockJson          `json:"phase0_block"`
	AltairBlock    *beaconBlockAltairJson    `json:"altair_block"`
	BellatrixBlock *beaconBlockBellatrixJson `json:"bellatrix_block"`
	Signature      apimiddleware.HexString   `json:"signature"`
}

type beaconBlockContainerV2Json struct {
	Phase0Block    *beaconBlockJson          `json:"phase0_block"`
	AltairBlock    *beaconBlockAltairJson    `json:"altair_block"`
	BellatrixBlock *beaconBlockBellatrixJson `json:"bellatrix_block"`
}

type blindedBeaconBlockContainerJson struct {
	Phase0Block    *beaconBlockJson                 `json:"phase0_block"`
	AltairBlock    *beaconBlockAltairJson           `json:"altair_block"`
	BellatrixBlock *blindedBeaconBlockBellatrixJson `json:"bellatrix_block"`
}

type signedBeaconBlockAltairContainerJson struct {
	Message   *beaconBlockAltairJson  `json:"message"`
	Signature apimiddleware.HexString `json:"signature"`
}

type signedBeaconBlockBellatrixContainerJson struct {
	Message   *beaconBlockBellatrixJson `json:"message"`
	Signature apimiddleware.HexString   `json:"signature"`
}

type signedBlindedBeaconBlockBellatrixContainerJson struct {
	Message   *blindedBeaconBlockBellatrixJson `json:"message"`
	Signature apimiddleware.HexString          `json:"signature"`
}

type beaconBlockAltairJson struct {
	Slot          string                     `json:"slot"`
	ProposerIndex string                     `json:"proposer_index"`
	ParentRoot    apimiddleware.HexString    `json:"parent_root"`
	StateRoot     apimiddleware.HexString    `json:"state_root"`
	Body          *beaconBlockBodyAltairJson `json:"body"`
}

type beaconBlockBellatrixJson struct {
	Slot          string                        `json:"slot"`
	ProposerIndex string                        `json:"proposer_index"`
	ParentRoot    apimiddleware.HexString       `json:"parent_root"`
	StateRoot     apimiddleware.HexString       `json:"state_root"`
	Body          *beaconBlockBodyBellatrixJson `json:"body"`
}

type blindedBeaconBlockBellatrixJson struct {
	Slot          string                               `json:"slot"`
	ProposerIndex string                               `json:"proposer_index"`
	ParentRoot    apimiddleware.HexString              `json:"parent_root"`
	StateRoot     apimiddleware.HexString              `json:"state_root"`
	Body          *blindedBeaconBlockBodyBellatrixJson `json:"body"`
}

type beaconBlockBodyAltairJson struct {
	RandaoReveal      apimiddleware.HexString    `json:"randao_reveal"`
	Eth1Data          *eth1DataJson              `json:"eth1_data"`
	Graffiti          apimiddleware.HexString    `json:"graffiti"`
	ProposerSlashings []*proposerSlashingJson    `json:"proposer_slashings"`
	AttesterSlashings []*attesterSlashingJson    `json:"attester_slashings"`
	Attestations      []*attestationJson         `json:"attestations"`
	Deposits          []*depositJson             `json:"deposits"`
	VoluntaryExits    []*signedVoluntaryExitJson `json:"voluntary_exits"`
	SyncAggregate     *syncAggregateJson         `json:"sync_aggregate"`
}

type beaconBlockBodyBellatrixJson struct {
	RandaoReveal      apimiddleware.HexString    `json:"randao_reveal"`
	Eth1Data          *eth1DataJson              `json:"eth1_data"`
	Graffiti          apimiddleware.HexString    `json:"graffiti"`
	ProposerSlashings []*proposerSlashingJson    `json:"proposer_slashings"`
	AttesterSlashings []*attesterSlashingJson    `json:"attester_slashings"`
	Attestations      []*attestationJson         `json:"attestations"`
	Deposits          []*depositJson             `json:"deposits"`
	VoluntaryExits    []*signedVoluntaryExitJson `json:"voluntary_exits"`
	SyncAggregate     *syncAggregateJson         `json:"sync_aggregate"`
	ExecutionPayload  *executionPayloadJson      `json:"execution_payload"`
}

type blindedBeaconBlockBodyBellatrixJson struct {
	RandaoReveal           apimiddleware.HexString     `json:"randao_reveal"`
	Eth1Data               *eth1DataJson               `json:"eth1_data"`
	Graffiti               apimiddleware.HexString     `json:"graffiti"`
	ProposerSlashings      []*proposerSlashingJson     `json:"proposer_slashings"`
	AttesterSlashings      []*attesterSlashingJson     `json:"attester_slashings"`
	Attestations           []*attestationJson          `json:"attestations"`
	Deposits               []*depositJson              `json:"deposits"`
	VoluntaryExits         []*signedVoluntaryExitJson  `json:"voluntary_exits"`
	SyncAggregate          *syncAggregateJson          `json:"sync_aggregate"`
	ExecutionPayloadHeader *executionPayloadHeaderJson `json:"execution_payload_header"`
}

type executionPayloadJson struct {
	ParentHash    apimiddleware.HexString   `json:"parent_hash"`
	FeeRecipient  apimiddleware.HexString   `json:"fee_recipient"`
	StateRoot     apimiddleware.HexString   `json:"state_root"`
	ReceiptsRoot  apimiddleware.HexString   `json:"receipts_root"`
	LogsBloom     apimiddleware.HexString   `json:"logs_bloom"`
	PrevRandao    apimiddleware.HexString   `json:"prev_randao"`
	BlockNumber   string                    `json:"block_number"`
	GasLimit      string                    `json:"gas_limit"`
	GasUsed       string                    `json:"gas_used"`
	TimeStamp     string                    `json:"timestamp"`
	ExtraData     apimiddleware.HexString   `json:"extra_data"`
	BaseFeePerGas string                    `json:"base_fee_per_gas" uint256:"true"`
	BlockHash     apimiddleware.HexString   `json:"block_hash"`
	Transactions  []apimiddleware.HexString `json:"transactions"`
}

type executionPayloadHeaderJson struct {
	ParentHash       apimiddleware.HexString `json:"parent_hash"`
	FeeRecipient     apimiddleware.HexString `json:"fee_recipient"`
	StateRoot        apimiddleware.HexString `json:"state_root"`
	ReceiptsRoot     apimiddleware.HexString `json:"receipts_root"`
	LogsBloom        apimiddleware.HexString `json:"logs_bloom"`
	PrevRandao       apimiddleware.HexString `json:"prev_randao"`
	BlockNumber      string                  `json:"block_number"`
	GasLimit         string                  `json:"gas_limit"`
	GasUsed          string                  `json:"gas_used"`
	TimeStamp        string                  `json:"timestamp"`
	ExtraData        apimiddleware.HexString `json:"extra_data"`
	BaseFeePerGas    string                  `json:"base_fee_per_gas" uint256:"true"`
	BlockHash        apimiddleware.HexString `json:"block_hash"`
	TransactionsRoot apimiddleware.HexString `json:"transactions_root"`
}

type syncAggregateJson struct {
	SyncCommitteeBits      apimiddleware.HexString `json:"sync_committee_bits"`
	SyncCommitteeSignature apimiddleware.HexString `json:"sync_committee_signature"`
}

type blockHeaderContainerJson struct {
	Root      apimiddleware.HexString         `json:"root"`
	Canonical bool                            `json:"canonical"`
	Header    *beaconBlockHeaderContainerJson `json:"header"`
}

type beaconBlockHeaderContainerJson struct {
	Message   *beaconBlockHeaderJson  `json:"message"`
	Signature apimiddleware.HexString `json:"signature"`
}

type signedBeaconBlockHeaderJson struct {
	Header    *beaconBlockHeaderJson  `json:"message"`
	Signature apimiddleware.HexString `json:"signature"`
}

type beaconBlockHeaderJson struct {
	Slot          string                  `json:"slot"`
	ProposerIndex string                  `json:"proposer_index"`
	ParentRoot    apimiddleware.HexString `json:"parent_root"`
	StateRoot     apimiddleware.HexString `json:"state_root"`
	BodyRoot      apimiddleware.HexString `json:"body_root"`
}

type eth1DataJson struct {
	DepositRoot  apimiddleware.HexString `json:"deposit_root"`
	DepositCount string                  `json:"deposit_count"`
	BlockHash    apimiddleware.HexString `json:"block_hash"`
}

type proposerSlashingJson struct {
	Header_1 *signedBeaconBlockHeaderJson `json:"signed_header_1"`
	Header_2 *signedBeaconBlockHeaderJson `json:"signed_header_2"`
}

type attesterSlashingJson struct {
	Attestation_1 *indexedAttestationJson `json:"attestation_1"`
	Attestation_2 *indexedAttestationJson `json:"attestation_2"`
}

type indexedAttestationJson struct {
	AttestingIndices []string                `json:"attesting_indices"`
	Data             *attestationDataJson    `json:"data"`
	Signature        apimiddleware.HexString `json:"signature"`
}

type feeRecipientJson struct {
	ValidatorIndex string                  `json:"validator_index"`
	FeeRecipient   apimiddleware.HexString `json:"fee_recipient"`
}

type attestationJson struct {
	AggregationBits apimiddleware.HexString `json:"aggregation_bits"`
	Data            *attestationDataJson    `json:"data"`
	Signature       apimiddleware.HexString `json:"signature"`
}

type attestationDataJson struct {
	Slot            string                  `json:"slot"`
	CommitteeIndex  string                  `json:"index"`
	BeaconBlockRoot apimiddleware.HexString `json:"beacon_block_root"`
	Source          *checkpointJson         `json:"source"`
	Target          *checkpointJson         `json:"target"`
}

type depositJson struct {
	Proof []apimiddleware.HexString `json:"proof"`
	Data  *deposit_DataJson         `json:"data"`
}

type deposit_DataJson struct {
	PublicKey             apimiddleware.HexString `json:"pubkey"`
	WithdrawalCredentials apimiddleware.HexString `json:"withdrawal_credentials"`
	Amount                string                  `json:"amount"`
	Signature             apimiddleware.HexString `json:"signature"`
}

type signedVoluntaryExitJson struct {
	Exit      *voluntaryExitJson      `json:"message"`
	Signature apimiddleware.HexString `json:"signature"`
}

type voluntaryExitJson struct {
	Epoch          string `json:"epoch"`
	ValidatorIndex string `json:"validator_index"`
}

type syncCommitteeMessageJson struct {
	Slot            string                  `json:"slot"`
	BeaconBlockRoot apimiddleware.HexString `json:"beacon_block_root"`
	ValidatorIndex  string                  `json:"validator_index"`
	Signature       apimiddleware.HexString `json:"signature"`
}

type identityJson struct {
	PeerId             string        `json:"peer_id"`
	Enr                string        `json:"enr"`
	P2PAddresses       []string      `json:"p2p_addresses"`
	DiscoveryAddresses []string      `json:"discovery_addresses"`
	Metadata           *metadataJson `json:"metadata"`
}

type metadataJson struct {
	SeqNumber string                  `json:"seq_number"`
	Attnets   apimiddleware.HexString `json:"attnets"`
}

type peerJson struct {
	PeerId    string `json:"peer_id"`
	Enr       string `json:"enr"`
	Address   string `json:"last_seen_p2p_address"`
	State     string `json:"state" enum:"true"`
	Direction string `json:"direction" enum:"true"`
}

type versionJson struct {
	Version string `json:"version"`
}

type beaconStateJson struct {
	GenesisTime                 string                    `json:"genesis_time"`
	GenesisValidatorsRoot       apimiddleware.HexString   `json:"genesis_validators_root"`
	Slot                        string                    `json:"slot"`
	Fork                        *forkJson                 `json:"fork"`
	LatestBlockHeader           *beaconBlockHeaderJson    `json:"latest_block_header"`
	BlockRoots                  []apimiddleware.HexString `json:"block_roots"`
	StateRoots                  []apimiddleware.HexString `json:"state_roots"`
	HistoricalRoots             []apimiddleware.HexString `json:"historical_roots"`
	Eth1Data                    *eth1DataJson             `json:"eth1_data"`
	Eth1DataVotes               []*eth1DataJson           `json:"eth1_data_votes"`
	Eth1DepositIndex            string                    `json:"eth1_deposit_index"`
	Validators                  []*validatorJson          `json:"validators"`
	Balances                    []string                  `json:"balances"`
	RandaoMixes                 []apimiddleware.HexString `json:"randao_mixes"`
	Slashings                   []string                  `json:"slashings"`
	PreviousEpochAttestations   []*pendingAttestationJson `json:"previous_epoch_attestations"`
	CurrentEpochAttestations    []*pendingAttestationJson `json:"current_epoch_attestations"`
	JustificationBits           apimiddleware.HexString   `json:"justification_bits"`
	PreviousJustifiedCheckpoint *checkpointJson           `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  *checkpointJson           `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         *checkpointJson           `json:"finalized_checkpoint"`
}

type beaconStateAltairJson struct {
	GenesisTime                 string                    `json:"genesis_time"`
	GenesisValidatorsRoot       apimiddleware.HexString   `json:"genesis_validators_root"`
	Slot                        string                    `json:"slot"`
	Fork                        *forkJson                 `json:"fork"`
	LatestBlockHeader           *beaconBlockHeaderJson    `json:"latest_block_header"`
	BlockRoots                  []apimiddleware.HexString `json:"block_roots"`
	StateRoots                  []apimiddleware.HexString `json:"state_roots"`
	HistoricalRoots             []apimiddleware.HexString `json:"historical_roots"`
	Eth1Data                    *eth1DataJson             `json:"eth1_data"`
	Eth1DataVotes               []*eth1DataJson           `json:"eth1_data_votes"`
	Eth1DepositIndex            string                    `json:"eth1_deposit_index"`
	Validators                  []*validatorJson          `json:"validators"`
	Balances                    []string                  `json:"balances"`
	RandaoMixes                 []apimiddleware.HexString `json:"randao_mixes"`
	Slashings                   []string                  `json:"slashings"`
	PreviousEpochParticipation  EpochParticipation        `json:"previous_epoch_participation"`
	CurrentEpochParticipation   EpochParticipation        `json:"current_epoch_participation"`
	JustificationBits           apimiddleware.HexString   `json:"justification_bits"`
	PreviousJustifiedCheckpoint *checkpointJson           `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  *checkpointJson           `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         *checkpointJson           `json:"finalized_checkpoint"`
	InactivityScores            []string                  `json:"inactivity_scores"`
	CurrentSyncCommittee        *syncCommitteeJson        `json:"current_sync_committee"`
	NextSyncCommittee           *syncCommitteeJson        `json:"next_sync_committee"`
}

type beaconStateBellatrixJson struct {
	GenesisTime                  string                      `json:"genesis_time"`
	GenesisValidatorsRoot        apimiddleware.HexString     `json:"genesis_validators_root"`
	Slot                         string                      `json:"slot"`
	Fork                         *forkJson                   `json:"fork"`
	LatestBlockHeader            *beaconBlockHeaderJson      `json:"latest_block_header"`
	BlockRoots                   []apimiddleware.HexString   `json:"block_roots"`
	StateRoots                   []apimiddleware.HexString   `json:"state_roots"`
	HistoricalRoots              []apimiddleware.HexString   `json:"historical_roots"`
	Eth1Data                     *eth1DataJson               `json:"eth1_data"`
	Eth1DataVotes                []*eth1DataJson             `json:"eth1_data_votes"`
	Eth1DepositIndex             string                      `json:"eth1_deposit_index"`
	Validators                   []*validatorJson            `json:"validators"`
	Balances                     []string                    `json:"balances"`
	RandaoMixes                  []apimiddleware.HexString   `json:"randao_mixes"`
	Slashings                    []string                    `json:"slashings"`
	PreviousEpochParticipation   EpochParticipation          `json:"previous_epoch_participation"`
	CurrentEpochParticipation    EpochParticipation          `json:"current_epoch_participation"`
	JustificationBits            apimiddleware.HexString     `json:"justification_bits"`
	PreviousJustifiedCheckpoint  *checkpointJson             `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint   *checkpointJson             `json:"current_justified_checkpoint"`
	FinalizedCheckpoint          *checkpointJson             `json:"finalized_checkpoint"`
	InactivityScores             []string                    `json:"inactivity_scores"`
	CurrentSyncCommittee         *syncCommitteeJson          `json:"current_sync_committee"`
	NextSyncCommittee            *syncCommitteeJson          `json:"next_sync_committee"`
	LatestExecutionPayloadHeader *executionPayloadHeaderJson `json:"latest_execution_payload_header"`
}

type beaconStateContainerV2Json struct {
	Phase0State    *beaconStateJson          `json:"phase0_state"`
	AltairState    *beaconStateAltairJson    `json:"altair_state"`
	BellatrixState *beaconStateBellatrixJson `json:"bellatrix_state"`
}

type forkJson struct {
	PreviousVersion apimiddleware.HexString `json:"previous_version"`
	CurrentVersion  apimiddleware.HexString `json:"current_version"`
	Epoch           string                  `json:"epoch"`
}

type validatorContainerJson struct {
	Index     string         `json:"index"`
	Balance   string         `json:"balance"`
	Status    string         `json:"status" enum:"true"`
	Validator *validatorJson `json:"validator"`
}

type validatorJson struct {
	PublicKey                  apimiddleware.HexString `json:"pubkey"`
	WithdrawalCredentials      apimiddleware.HexString `json:"withdrawal_credentials"`
	EffectiveBalance           string                  `json:"effective_balance"`
	Slashed                    bool                    `json:"slashed"`
	ActivationEligibilityEpoch string                  `json:"activation_eligibility_epoch"`
	ActivationEpoch            string                  `json:"activation_epoch"`
	ExitEpoch                  string                  `json:"exit_epoch"`
	WithdrawableEpoch          string                  `json:"withdrawable_epoch"`
}

type validatorBalanceJson struct {
	Index   string `json:"index"`
	Balance string `json:"balance"`
}

type committeeJson struct {
	Index      string   `json:"index"`
	Slot       string   `json:"slot"`
	Validators []string `json:"validators"`
}

type syncCommitteeJson struct {
	Pubkeys         []apimiddleware.HexString `json:"pubkeys"`
	AggregatePubkey apimiddleware.HexString   `json:"aggregate_pubkey"`
}

type syncCommitteeValidatorsJson struct {
	Validators          []string   `json:"validators"`
	ValidatorAggregates [][]string `json:"validator_aggregates"`
}

type pendingAttestationJson struct {
	AggregationBits apimiddleware.HexString `json:"aggregation_bits"`
	Data            *attestationDataJson    `json:"data"`
	InclusionDelay  string                  `json:"inclusion_delay"`
	ProposerIndex   string                  `json:"proposer_index"`
}

type forkChoiceHeadJson struct {
	Root apimiddleware.HexString `json:"root"`
	Slot string                  `json:"slot"`
}

type v2ForkChoiceHeadJson struct {
	Root                apimiddleware.HexString `json:"root"`
	Slot                string                  `json:"slot"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type depositContractJson struct {
	ChainId string `json:"chain_id"`
	Address string `json:"address"`
}

type attesterDutyJson struct {
	Pubkey                  apimiddleware.HexString `json:"pubkey"`
	ValidatorIndex          string                  `json:"validator_index"`
	CommitteeIndex          string                  `json:"committee_index"`
	CommitteeLength         string                  `json:"committee_length"`
	CommitteesAtSlot        string                  `json:"committees_at_slot"`
	ValidatorCommitteeIndex string                  `json:"validator_committee_index"`
	Slot                    string                  `json:"slot"`
}

type proposerDutyJson struct {
	Pubkey         apimiddleware.HexString `json:"pubkey"`
	ValidatorIndex string                  `json:"validator_index"`
	Slot           string                  `json:"slot"`
}

type syncCommitteeDuty struct {
	Pubkey                        apimiddleware.HexString `json:"pubkey"`
	ValidatorIndex                string                  `json:"validator_index"`
	ValidatorSyncCommitteeIndices []string                `json:"validator_sync_committee_indices"`
}

type signedAggregateAttestationAndProofJson struct {
	Message   *aggregateAttestationAndProofJson `json:"message"`
	Signature apimiddleware.HexString           `json:"signature"`
}

type aggregateAttestationAndProofJson struct {
	AggregatorIndex string                  `json:"aggregator_index"`
	Aggregate       *attestationJson        `json:"aggregate"`
	SelectionProof  apimiddleware.HexString `json:"selection_proof"`
}

type signedContributionAndProofJson struct {
	Message   *contributionAndProofJson `json:"message"`
	Signature apimiddleware.HexString   `json:"signature"`
}

type contributionAndProofJson struct {
	AggregatorIndex string                         `json:"aggregator_index"`
	Contribution    *syncCommitteeContributionJson `json:"contribution"`
	SelectionProof  apimiddleware.HexString        `json:"selection_proof"`
}

type syncCommitteeContributionJson struct {
	Slot              string                  `json:"slot"`
	BeaconBlockRoot   apimiddleware.HexString `json:"beacon_block_root"`
	SubcommitteeIndex string                  `json:"subcommittee_index"`
	AggregationBits   apimiddleware.HexString `json:"aggregation_bits"`
	Signature         apimiddleware.HexString `json:"signature"`
}

type validatorRegistrationJson struct {
	FeeRecipient apimiddleware.HexString `json:"fee_recipient"`
	GasLimit     string                  `json:"gas_limit"`
	Timestamp    string                  `json:"timestamp"`
	Pubkey       apimiddleware.HexString `json:"pubkey"`
}

type signedValidatorRegistrationJson struct {
	Message   *validatorRegistrationJson `json:"message"`
	Signature apimiddleware.HexString    `json:"signature"`
}

type signedValidatorRegistrationsRequestJson struct {
	Registrations []*signedValidatorRegistrationJson `json:"registrations"`
}

type forkChoiceNodeJson struct {
	Slot                     string                  `json:"slot"`
	Root                     apimiddleware.HexString `json:"root"`
	ParentRoot               apimiddleware.HexString `json:"parent_root"`
	JustifiedEpoch           string                  `json:"justified_epoch"`
	FinalizedEpoch           string                  `json:"finalized_epoch"`
	UnrealizedJustifiedEpoch string                  `json:"unrealized_justified_epoch"`
	UnrealizedFinalizedEpoch string                  `json:"unrealized_finalized_epoch"`
	Balance                  string                  `json:"balance"`
	Weight                   string                  `json:"weight"`
	ExecutionOptimistic      bool                    `json:"execution_optimistic"`
	ExecutionPayload         apimiddleware.HexString `json:"execution_payload"`
	TimeStamp                string                  `json:"timestamp"`
}

//----------------
// SSZ
// ---------------

type sszRequestJson struct {
	Data string `json:"data"`
}

// sszResponse is a common abstraction over all SSZ responses.
type sszResponse interface {
	SSZVersion() string
	SSZData() string
}

type sszResponseJson struct {
	Data string `json:"data"`
}

func (ssz *sszResponseJson) SSZData() string {
	return ssz.Data
}

func (*sszResponseJson) SSZVersion() string {
	return strings.ToLower(ethpbv2.Version_PHASE0.String())
}

type versionedSSZResponseJson struct {
	Version string `json:"version"`
	Data    string `json:"data"`
}

func (ssz *versionedSSZResponseJson) SSZData() string {
	return ssz.Data
}

func (ssz *versionedSSZResponseJson) SSZVersion() string {
	return ssz.Version
}

// ---------------
// Events.
// ---------------

type eventHeadJson struct {
	Slot                      string                  `json:"slot"`
	Block                     apimiddleware.HexString `json:"block"`
	State                     apimiddleware.HexString `json:"state"`
	EpochTransition           bool                    `json:"epoch_transition"`
	ExecutionOptimistic       bool                    `json:"execution_optimistic"`
	PreviousDutyDependentRoot apimiddleware.HexString `json:"previous_duty_dependent_root"`
	CurrentDutyDependentRoot  apimiddleware.HexString `json:"current_duty_dependent_root"`
}

type receivedBlockDataJson struct {
	Slot                string                  `json:"slot"`
	Block               apimiddleware.HexString `json:"block"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type aggregatedAttReceivedDataJson struct {
	Aggregate *attestationJson `json:"aggregate"`
}

type eventFinalizedCheckpointJson struct {
	Block               apimiddleware.HexString `json:"block"`
	State               apimiddleware.HexString `json:"state"`
	Epoch               string                  `json:"epoch"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

type eventChainReorgJson struct {
	Slot                string                  `json:"slot"`
	Depth               string                  `json:"depth"`
	OldHeadBlock        apimiddleware.HexString `json:"old_head_block"`
	NewHeadBlock        apimiddleware.HexString `json:"old_head_state"`
	OldHeadState        apimiddleware.HexString `json:"new_head_block"`
	NewHeadState        apimiddleware.HexString `json:"new_head_state"`
	Epoch               string                  `json:"epoch"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
}

// ---------------
// Error handling.
// ---------------

// indexedVerificationFailureErrorJson is a JSON representation of the error returned when verifying an indexed object.
type indexedVerificationFailureErrorJson struct {
	apimiddleware.DefaultErrorJson
	Failures []*singleIndexedVerificationFailureJson `json:"failures"`
}

// singleIndexedVerificationFailureJson is a JSON representation of a an issue when verifying a single indexed object e.g. an item in an array.
type singleIndexedVerificationFailureJson struct {
	Index   int    `json:"index"`
	Message string `json:"message"`
}

type nodeSyncDetailsErrorJson struct {
	apimiddleware.DefaultErrorJson
	SyncDetails helpers.SyncDetailsJson `json:"sync_details"`
}

type eventErrorJson struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
