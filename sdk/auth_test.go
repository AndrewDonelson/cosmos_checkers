package main

import "testing"

func TestGetAccounts(t *testing.T) {
	data, err := consumeAPI_GetAccounts()
	Ok(t, err)
	Assert(t, data["total"] != nil, "height is nil")
}
