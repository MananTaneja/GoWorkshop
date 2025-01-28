package api_test

import (
	"testing"

	"manantaneja.com/go/crypto-masters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Did not error out when we sent empty currency value")
	}

}
