package zkvote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	zkvote "github.com/soroworks/sorovault/internal/codegen/testdata"
)

// TestGetProposalHappyPath confirms that a client built with the canonical
// contract address (the one the file was generated for) can encode
// GetProposal without error. The returned operation is non-nil and carries
// the invoke contract function call.
func TestGetProposalHappyPath(t *testing.T) {
	t.Parallel()

	client := zkvote.NewClient(zkvote.ContractID, zkvote.Network)
	op, err := client.GetProposal(0)
	require.NoError(t, err, "GetProposal with a valid contract ID must not fail")
	assert.NotNil(t, op, "GetProposal must return a non-nil operation")
}

// TestGetProposalInvalidContractID confirms that an invalid contract address
// is caught when the operation is built, not silently ignored.
func TestGetProposalInvalidContractID(t *testing.T) {
	t.Parallel()

	client := &zkvote.Client{ContractID: "not-a-valid-strkey", Network: zkvote.Network}
	op, err := client.GetProposal(0)
	require.Error(t, err, "GetProposal with an invalid contract ID must return an error")
	assert.Nil(t, op, "GetProposal must return nil operation on error")
}
