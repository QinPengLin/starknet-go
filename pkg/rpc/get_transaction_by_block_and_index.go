package api

import (
	"context"

	"github.com/QinPengLin/starknet-go/pkg/data"
)

// GetTransactionByBlockNumberAndIndex - Get the details of the transaction given by the identified block and index in that block. If no transaction is found, null is returned.
func (api API) GetTransactionByBlockNumberAndIndex(ctx context.Context, block data.BlockID, index uint64, opts ...RequestOption) (*Response[data.Transaction], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getTransactionByBlockIdAndIndex", []any{
		block, index,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[data.Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
