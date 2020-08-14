package domain

import "testing"

func TestConnectDB(t *testing.T) { // Complete.
	uri := "mongodb://localhost:27017"
	ConnectDB(uri)
}
