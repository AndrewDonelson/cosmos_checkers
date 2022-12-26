package main

import "fmt"

// CosmosBaseTendermintV1Beta1NodeInfo
// GET /cosmos/base/tendermint/v1beta1/node_info
// Get the node info.
// Scope: Public
// Parameters:
//
//	none
//
// Returns:
//
// map[string]interface{}: The response as a map
// error: Any error that occurred
func consumeAPI_NodeInfo() (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/base/tendermint/v1beta1/node_info")
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}

// CosmosBaseTendermintV1Beta1GetLatestValidatorSet
// GET /cosmos/base/tendermint/v1beta1/validatorsets/latest
// Get the latest validator set.
// Scope: Public
// Parameters:
//
//	none
//
// Returns:
//
// map[string]interface{}: The response as a map
// error: Any error that occurred
func consumeAPI_GetLatestValidatorSet() (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/base/tendermint/v1beta1/validatorsets/latest")
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}

// CosmosBaseTendermintV1Beta1GetLatestBlock
// GET /cosmos/base/tendermint/v1beta1/blocks/latest
// Get the latest block.
// Scope: Public
// Parameters:
//
//	none
//
// Returns:
//
// map[string]interface{}: The response as a map
// error: Any error that occurred
func consumeAPI_GetLatestBlock() (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/base/tendermint/v1beta1/validatorsets/latest")
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}
