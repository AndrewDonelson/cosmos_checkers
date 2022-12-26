package main

import "fmt"

// CosmosAuthV1Beta1Accounts
// GET /cosmos/auth/v1beta1/accounts
// Accounts returns all the existing accounts
// Scope: Public
// Parameters:
//
//	none
//
// Returns:
//
// map[string]interface{}: The response as a map
// error: Any error that occurred
func consumeAPI_GetAccounts() (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/auth/v1beta1/accounts")
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}

// CosmosAuthV1Beta1Account
// GET /cosmos/auth/v1beta1/accounts/{address}
// Account returns the account details based on address
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
func consumeAPI_GetAccount(address string) (map[string]interface{}, error) {
	data, err := consumeAPI(GWFAPIHOST, "/cosmos/auth/v1beta1/accounts/"+address)
	if err != nil {
		return nil, fmt.Errorf("error consuming API: %v", err)
	}
	return data, nil
}
