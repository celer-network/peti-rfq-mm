package eth

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/celer-network/rfq-mm/sdk/bindings/multicall"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IMulticaller interface {
	TryBlockAndAggregate(requireSuccess bool, calls []multicall.IMulticall2Call) (*MulticallResponse, error)
}

// The quote functions in IQuoterV2 are a non-view functions. Manual low-level operations are needed to call it statically
type Multicaller struct {
	address common.Address

	ec  *ethclient.Client
	abi *abi.ABI
}

type MulticallResults = []struct {
	Success    bool    "json:\"success\""
	ReturnData []uint8 "json:\"returnData\""
}

type MulticallResponse struct {
	BlockNumber *big.Int
	BlockHash   Hash
	ReturnData  MulticallResults
}

func NewMulticaller(address common.Address, ec *ethclient.Client) (*Multicaller, error) {
	abi, err := abi.JSON(strings.NewReader(multicall.IMulticall2ABI))
	if err != nil {
		return nil, err
	}
	return &Multicaller{
		address: address,
		ec:      ec,
		abi:     &abi,
	}, nil
}

const methodTryBlockAndAggregate = "tryBlockAndAggregate"

func (m *Multicaller) TryBlockAndAggregate(requireSuccess bool, calls []multicall.IMulticall2Call) (*MulticallResponse, error) {
	data, err := m.abi.Methods[methodTryBlockAndAggregate].Inputs.Pack(requireSuccess, calls)
	if err != nil {
		return nil, fmt.Errorf("multicall pack request error: %s", err)
	}
	msg := ethereum.CallMsg{
		To:   &m.address,
		Data: append(m.abi.Methods[methodTryBlockAndAggregate].ID, data...),
	}
	res, err := m.ec.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("multicall CallContract to %x, data %x, error: %w", *msg.To, msg.Data, err)
	}
	results, err := m.abi.Methods[methodTryBlockAndAggregate].Outputs.Unpack(res)
	if err != nil {
		return nil, fmt.Errorf("multicall unpack response error: %s", err)
	}
	return &MulticallResponse{
		BlockNumber: results[0].(*big.Int),
		BlockHash:   common.Hash(results[1].([32]byte)),
		ReturnData:  results[2].(MulticallResults),
	}, nil
}
