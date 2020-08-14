package domain

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	uri = "mongodb://localhost:27017"
)

func TestCreate(t *testing.T) { // Complete.
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
	// Try to create a message with an id that already exist.
	msg := Message{
		ID:      "1",
		Sender:  "A",
		Reciver: "B",
		Message: "Hi there, you are doin fine.",
	}
	res, restErr := Create(&msg)
	if restErr == nil {
		t.Error("Creating a message with an id that exists, must not give a nil err.")
	}
	if res != nil {
		t.Error("Creating a message with an id that exists, must return nil as message.")
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

func TestUpdate(t *testing.T) { // Complete.
	if db == nil {
		ConnectDB(uri)
	}
	// Try to update an existing message.
	if restErr := Update("1", "Hello how are you B?"); restErr != nil {
		t.Error(restErr.Message)
	}
	// Try to update a message that doesn't exist.
	if restErr := Update("0", "Hi there."); restErr == nil {
		t.Error("Updating a message that doesn't exist, must not give a nil err.")
	}
	// Try to update a message with the same content.
	if restErr := Update("1", "Hello how are you B?"); restErr == nil {
		t.Error("Updating a message that is already up-to-date, must not give a nil err.")
	}
}

func TestMakeSeen(t *testing.T) { // Complete.
	// Try to make an unseen message to seen.
	if restErr := MakeSeen("1"); restErr != nil {
		t.Error(restErr.Message)
	}
	// Try to make a message seen that doesn't exist.
	if restErr := MakeSeen("0"); restErr == nil {
		t.Error("Making a message seen that is doesn't exist, must not give a nil err.")
	}
	// Try to make a message that is already seen.
	if restErr := MakeSeen("1"); restErr == nil {
		t.Error("Making a message seen that is already seen, must not give a nil err.")
	}
}

func TestDelete(t *testing.T) { // Complete.
	// Try to delete all messages that are created by getTestMessages func.
	for i := 1; i <= 44; i++ {
		id := strconv.Itoa(i)
		if restErr := Delete(id); restErr != nil {
			t.Error(restErr.Message)
		}
	}
	// Try to delete a message that doesn't exist.
	if restErr := Delete("0"); restErr == nil {
		t.Error("Deleting a message that doesn't exist, must not give a nil err.")
	}
}
