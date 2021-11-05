package slasher

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/bmatsuo/lmdb-go/lmdb"
	ssz "github.com/ferranbt/fastssz"
	"github.com/golang/snappy"
	"github.com/pkg/errors"
	types "github.com/prysmaticlabs/eth2-types"
	slashertypes "github.com/prysmaticlabs/prysm/beacon-chain/slasher/types"
	eth "github.com/prysmaticlabs/prysm/proto/eth/v1"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

const (
	// Signing root (32 bytes)
	attestationRecordKeySize = 32 // Bytes.
	signingRootSize          = 32 // Bytes.
)

// The schema will define how to store and retrieve data from the db.
// we can prefix or suffix certain values such as `block` with attributes
// for prefix-wide scans across the underlying BoltDB buckets when filtering data.
// For example, we might store attestations as shard + attestation_root -> attestation, making
// it easy to scan for keys that have a certain shard number as a prefix and return those
// corresponding attestations.
var (
	DBIEx lmdb.DBI
	// Slasher buckets.
	attestedEpochsByValidator  = []byte("attested-epochs-by-validator")
	attestationRecordsBucket   = []byte("attestation-records")
	attestationDataRootsBucket = []byte("attestation-data-roots")
	proposalRecordsBucket      = []byte("proposal-records")
	slasherChunksBucket        = []byte("slasher-chunks")
)

type store struct {
	env *lmdb.Env
}

func newStore() (*store, error) {
	ss := &store{}
	// create an environment and make sure it is eventually closed.
	env, err := lmdb.NewEnv()
	if err != nil {
		return nil, err
	}
	ss.env = env

	// configure and open the environment.  most configuration must be done
	// before opening the environment.
	err = env.SetMaxDBs(1)
	if err != nil {
		return nil, err
	}
	err = env.SetMapSize(1 << 30)
	if err != nil {
		return nil, err
	}
	err = env.Open("/tmp/myfastdb/", 0, 0644)
	if err != nil {
		return nil, err
	}

	// open a database that can be used as long as the enviroment is mapped.
	var dbi lmdb.DBI
	err = env.Update(func(txn *lmdb.Txn) (err error) {
		dbi, err = txn.CreateDBI("slasher")
		return err
	})
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (ss *store) Close() error {
	return ss.env.Close()
}

func (ss *store) SaveLastEpochWrittenForValidators(
	txn *lmdb.Txn, validatorIndices []types.ValidatorIndex, epoch types.Epoch,
) error {
	return nil
}

func (ss *store) SaveAttestationRecordsForValidators(
	txn *lmdb.Txn,
	attestations []*slashertypes.IndexedAttestationWrapper,
) error {
	return nil
}

func (ss *store) SaveSlasherChunks(
	txn *lmdb.Txn, kind slashertypes.ChunkKind, chunkKeys [][]byte, chunks [][]uint16,
) error {
	return nil
}

func (ss *store) SaveBlockProposals(
	txn *lmdb.Txn, proposal []*slashertypes.SignedBlockHeaderWrapper,
) error {
	return nil
}

func (ss *store) LastEpochWrittenForValidators(
	txn *lmdb.Txn, validatorIndices []types.ValidatorIndex,
) ([]*slashertypes.AttestedEpochForValidator, error) {
	attestedEpochs := make([]*slashertypes.AttestedEpochForValidator, 0)
	encodedIndices := make([][]byte, len(validatorIndices))
	for i, valIdx := range validatorIndices {
		encodedIndices[i] = encodeValidatorIndex(valIdx)
	}
	for i, encodedIndex := range encodedIndices {
		epochBytes, err := txn.Get(DBIEx, encodedIndex)
		if err != nil {
			return nil, err
		}
		if epochBytes != nil {
			var epoch types.Epoch
			if err := epoch.UnmarshalSSZ(epochBytes); err != nil {
				return nil, err
			}
			attestedEpochs = append(attestedEpochs, &slashertypes.AttestedEpochForValidator{
				ValidatorIndex: validatorIndices[i],
				Epoch:          epoch,
			})
		}
	}
	return attestedEpochs, nil
}

func (ss *store) AttestationRecordForValidator(
	txn *lmdb.Txn, validatorIdx types.ValidatorIndex, targetEpoch types.Epoch,
) (*slashertypes.IndexedAttestationWrapper, error) {
	return nil, nil
}

func (ss *store) BlockProposalForValidator(
	txn *lmdb.Txn, validatorIdx types.ValidatorIndex, slot types.Slot,
) (*slashertypes.SignedBlockHeaderWrapper, error) {
	return nil, nil
}

func (ss *store) CheckAttesterDoubleVotes(
	txn *lmdb.Txn, attestations []*slashertypes.IndexedAttestationWrapper,
) ([]*slashertypes.AttesterDoubleVote, error) {
	return nil, nil
}

func (ss *store) LoadSlasherChunks(
	txn *lmdb.Txn, kind slashertypes.ChunkKind, diskKeys [][]byte,
) ([][]uint16, []bool, error) {
	return nil, nil, nil
}

func (ss *store) CheckDoubleBlockProposals(
	txn *lmdb.Txn, proposals []*slashertypes.SignedBlockHeaderWrapper,
) ([]*eth.ProposerSlashing, error) {
	return nil, nil
}

func (ss *store) PruneAttestationsAtEpoch(
	txn *lmdb.Txn, maxEpoch types.Epoch,
) (numPruned uint, err error) {
	return 0, nil
}

func (ss *store) PruneProposalsAtEpoch(
	txn *lmdb.Txn, maxEpoch types.Epoch,
) (numPruned uint, err error) {
	return 0, nil
}

func (ss *store) HighestAttestations(
	txn *lmdb.Txn,
	indices []types.ValidatorIndex,
) ([]*ethpb.HighestAttestation, error) {
	return nil, nil
}

func suffixForAttestationRecordsKey(key, encodedValidatorIndex []byte) bool {
	encIdx := key[8:]
	return bytes.Equal(encIdx, encodedValidatorIndex)
}

// Disk key for a validator proposal, including a slot+validatorIndex as a byte slice.
func keyForValidatorProposal(slot types.Slot, proposerIndex types.ValidatorIndex) ([]byte, error) {
	encSlot, err := slot.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	encValidatorIdx := encodeValidatorIndex(proposerIndex)
	return append(encSlot, encValidatorIdx...), nil
}

func encodeSlasherChunk(chunk []uint16) ([]byte, error) {
	val := make([]byte, 0)
	for i := 0; i < len(chunk); i++ {
		val = append(val, ssz.MarshalUint16(make([]byte, 0), chunk[i])...)
	}
	if len(val) == 0 {
		return nil, errors.New("cannot encode empty chunk")
	}
	return snappy.Encode(nil, val), nil
}

func decodeSlasherChunk(enc []byte) ([]uint16, error) {
	chunkBytes, err := snappy.Decode(nil, enc)
	if err != nil {
		return nil, err
	}
	if len(chunkBytes)%2 != 0 {
		return nil, fmt.Errorf(
			"cannot decode slasher chunk with length %d, must be a multiple of 2",
			len(chunkBytes),
		)
	}
	chunk := make([]uint16, 0)
	for i := 0; i < len(chunkBytes); i += 2 {
		distance := ssz.UnmarshallUint16(chunkBytes[i : i+2])
		chunk = append(chunk, distance)
	}
	return chunk, nil
}

// Decode attestation record from bytes.
func encodeAttestationRecord(att *slashertypes.IndexedAttestationWrapper) ([]byte, error) {
	if att == nil || att.IndexedAttestation == nil {
		return []byte{}, errors.New("nil proposal record")
	}
	encodedAtt, err := att.IndexedAttestation.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	compressedAtt := snappy.Encode(nil, encodedAtt)
	return append(att.SigningRoot[:], compressedAtt...), nil
}

// Decode attestation record from bytes.
func decodeAttestationRecord(encoded []byte) (*slashertypes.IndexedAttestationWrapper, error) {
	if len(encoded) < signingRootSize {
		return nil, fmt.Errorf("wrong length for encoded attestation record, want 32, got %d", len(encoded))
	}
	signingRoot := encoded[:signingRootSize]
	decodedAtt := &ethpb.IndexedAttestation{}
	decodedAttBytes, err := snappy.Decode(nil, encoded[signingRootSize:])
	if err != nil {
		return nil, err
	}
	if err := decodedAtt.UnmarshalSSZ(decodedAttBytes); err != nil {
		return nil, err
	}
	return &slashertypes.IndexedAttestationWrapper{
		IndexedAttestation: decodedAtt,
		SigningRoot:        bytesutil.ToBytes32(signingRoot),
	}, nil
}

func encodeProposalRecord(blkHdr *slashertypes.SignedBlockHeaderWrapper) ([]byte, error) {
	if blkHdr == nil || blkHdr.SignedBeaconBlockHeader == nil {
		return []byte{}, errors.New("nil proposal record")
	}
	encodedHdr, err := blkHdr.SignedBeaconBlockHeader.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	compressedHdr := snappy.Encode(nil, encodedHdr)
	return append(blkHdr.SigningRoot[:], compressedHdr...), nil
}

func decodeProposalRecord(encoded []byte) (*slashertypes.SignedBlockHeaderWrapper, error) {
	if len(encoded) < signingRootSize {
		return nil, fmt.Errorf(
			"wrong length for encoded proposal record, want %d, got %d", signingRootSize, len(encoded),
		)
	}
	signingRoot := encoded[:signingRootSize]
	decodedBlkHdr := &ethpb.SignedBeaconBlockHeader{}
	decodedHdrBytes, err := snappy.Decode(nil, encoded[signingRootSize:])
	if err != nil {
		return nil, err
	}
	if err := decodedBlkHdr.UnmarshalSSZ(decodedHdrBytes); err != nil {
		return nil, err
	}
	return &slashertypes.SignedBlockHeaderWrapper{
		SignedBeaconBlockHeader: decodedBlkHdr,
		SigningRoot:             bytesutil.ToBytes32(signingRoot),
	}, nil
}

// Encodes an epoch into little-endian bytes.
func encodeTargetEpoch(epoch types.Epoch) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(epoch))
	return buf
}

// Encodes a validator index using 5 bytes instead of 8 as a
// client optimization to save space in the database. Because the max validator
// registry size is 2**40, this is a safe optimization.
func encodeValidatorIndex(index types.ValidatorIndex) []byte {
	buf := make([]byte, 5)
	v := uint64(index)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	return buf
}
