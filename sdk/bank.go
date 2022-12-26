package main

import "fmt"

// CosmosBankV1Beta1AllBalances
// GET /cosmos/bank/v1beta1/balances/{address}
// Get all balances for an address
// Scope: Public
// Parameters:
//
//	address: string
//		The address of the account
//
// Returns:
//
// map[string]interface{}: The response as a map
// error: Any error that occurred
func consumeAPI_GetBalance(address string) (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/bank/v1beta1/balances/"+address)
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}

// CosmosBankV1Beta1TotalSupply
// GET /cosmos/bank/v1beta1/supply
// Get the total supply of all coins
// Scope: Public
// Parameters:
//
//	none
//
// Returns:
//
// map[string]interface{}: The response as a map
// error: Any error that occurred
func consumeAPI_GetTotalSupply() (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/bank/v1beta1/supply")
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}
