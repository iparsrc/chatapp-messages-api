package domain

import (
	"fmt"
	"testing"
)

const (
	uri = "mongodb://localhost:27017"
)

func TestCreate(t *testing.T) {
	if db == nil {
		ConnectDB(uri)
	}
	// Try to create some new messages.
	messages := getTestMessages()
	for _, msg := range messages {
		msg, restErr := Create(&msg)
		if restErr != nil {
			t.Error(restErr.Message)
		} else {
			fmt.Println("    ", *msg)
		}

	}
}

func TestRetrive(t *testing.T) {
	if db == nil {
		ConnectDB(uri)
	}
	// Try to retrive 40 messages of direct message.
	reciver := "B"
	msg, restErr := RetriveFourty("A", 1, &reciver)
	if restErr != nil {
		t.Error(restErr.Message)
	}
	for _, value := range msg {
		fmt.Println("    ", value)
	}
	fmt.Println("")
	// Try to retvie 40 messages of a group chat.
	msgx, restErrx := RetriveFourty("C", 0, nil)
	if restErrx != nil {
		t.Error(restErrx.Message)
	}
	for _, valuex := range msgx {
		fmt.Println("    ", valuex)
	}
}

func TestUpdate(t *testing.T) {
	if db == nil {
		ConnectDB(uri)
	}
	// Try to update an existing message.
	if restErr := Update("1", "Hello how are you B?"); restErr != nil {
		t.Error(restErr.Message)
	}
	// Try to update a message that doesn't exist.
	if restErr := Update("0", "Hi there."); restErr != nil {
		t.Error(restErr.Message)
	}
	// Try to update a message with the same content.
	if restErr := Update("1", "Hi there."); restErr != nil {
		t.Error(restErr.Message)
	}
}
