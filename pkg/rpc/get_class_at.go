package api

import (
	"context"

	"github.com/QinPengLin/starknet-go/pkg/data"
)

// GetClassAt -
func (api API) GetClassAt(ctx context.Context, block data.BlockID, contractAddress string, opts ...RequestOption) (*Response[data.Class], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		block, contractAddress,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[data.Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
