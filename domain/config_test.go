package domain

import "testing"

func TestConnectDB(t *testing.T) {
	uri := "mongodb://localhost:27017"
	ConnectDB(uri)
}
