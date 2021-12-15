package relay

import eth2types "github.com/prysmaticlabs/eth2-types"

type ValidatorStatus int

const (
	ValidatorStatus_Unknown = iota
	ValidatorStatus_Deposited
	ValidatorStatus_Pending
	ValidatorStatus_Active
	ValidatorStatus_Exiting
	ValidatorStatus_Slashing
	ValidatorStatus_Exited
	ValidatorStatus_Invalid
	ValidatorStatus_PartiallyDeposited
)

type Duties struct {
	CurrentEpochDuties []*Duty
	NextEpochDuties    []*Duty
}

type Duty struct {
	Committee       []eth2types.ValidatorIndex `protobuf:"varint,1,rep,packed,name=committee,proto3" json:"committee,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
	CommitteeIndex  eth2types.CommitteeIndex   `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.CommitteeIndex"`
	AttesterSlot    eth2types.Slot             `protobuf:"varint,3,opt,name=attester_slot,json=attesterSlot,proto3" json:"attester_slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	ProposerSlots   []eth2types.Slot           `protobuf:"varint,4,rep,packed,name=proposer_slots,json=proposerSlots,proto3" json:"proposer_slots,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	PublicKey       []byte                     `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty" ssz-size:"48"`
	Status          ValidatorStatus            `protobuf:"varint,6,opt,name=status,proto3,enum=ethereum.eth.v1alpha1.ValidatorStatus" json:"status,omitempty"`
	ValidatorIndex  eth2types.ValidatorIndex   `protobuf:"varint,7,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
	IsSyncCommittee bool                       `protobuf:"varint,8,opt,name=is_sync_committee,json=isSyncCommittee,proto3" json:"is_sync_committee,omitempty"`
}
