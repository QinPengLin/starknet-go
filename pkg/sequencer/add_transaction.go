package sequencer

import (
	"context"

	"github.com/QinPengLin/starknet-go/pkg/data"
)

// Invoke -
func (api API) Invoke(ctx context.Context, tx data.Invoke) (string, error) {
	body := map[string]any{
		"type":             data.TransactionTypeInvokeFunction,
		"contract_address": tx.ContractAddress,
		"calldata":         tx.Calldata,
		"signature":        tx.Signature,
		"nonce":            tx.Nonce,
		"max_fee":          tx.MaxFee,
		"version":          data.Version0,
	}

	var hash string
	err := api.postToGateway(ctx, "add_transaction", nil, body, &hash)
	return hash, err
}

// DeployAccount -
func (api API) DeployAccount(ctx context.Context, tx data.DeployAccount) (string, error) {
	body := map[string]any{
		"type":                  data.TransactionTypeDeployAccount,
		"contract_address_salt": tx.ContractAddressSalt,
		"constructor_calldata":  tx.ConstructorCalldata,
		"class_hash":            tx.ClassHash,
		"signature":             tx.Signature,
		"nonce":                 tx.Nonce,
		"max_fee":               tx.MaxFee,
		"version":               data.Version0,
	}

	var hash string
	err := api.postToGateway(ctx, "add_transaction", nil, body, &hash)
	return hash, err
}

// Declare -
func (api API) Declare(ctx context.Context, tx data.Declare) (string, error) {
	body := map[string]any{
		"type":           data.TransactionTypeDeclare,
		"contract_class": tx.ContractClass,
		"sender_address": tx.SenderAddress,
		"class_hash":     tx.ClassHash,
		"signature":      tx.Signature,
		"max_fee":        tx.MaxFee,
		"version":        data.Version0,
	}

	var hash string
	err := api.postToGateway(ctx, "add_transaction", nil, body, &hash)
	return hash, err
}
