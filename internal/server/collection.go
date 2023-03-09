package server

import (
	"errors"

	"github.com/google/uuid"

	"github.com/nullentrypoint/tekken7-web-displayer/pkg/tekken"
)

type MessageID string

type Message struct {
	Id   MessageID         `json:"id,omitempty"`
	Info tekken.PlayerInfo `json:"info,omitempty"`
}

func (msg Message) ID() MessageID {
	return msg.Id
}

type Collection struct {
	hm    map[MessageID]*Message
	slice []*Message
}

// NewCollection creates and initialize an instance of Collection.
func NewCollection() (*Collection) {
	return &Collection{
		hm:    make(map[MessageID]*Message),
		slice: []*Message{},
	}
}

// All retreives all the items in MyCollecttion and returns them in a slice.
func (collection *Collection) All() ([]Message, error) {
	out := make([]Message, len(collection.slice))
	for _, value := range collection.slice {
		out = append(out, *value)
	}

	return out, nil
}

// Item retreives the item with the given ID from Collection and returns it.
func (collection *Collection) Item(ID MessageID) (Message, error) {
	out, ok := collection.hm[ID]
	if !ok {
		return Message{}, errors.New("not found")
	}

	return *out, nil
}

// Create adds the given item to the collection and returns it AFTER having set its id.
func (collection *Collection) Create(msg Message) (Message, error) {
	// Set item id
	msg.Id = MessageID(uuid.New().String())

	collection.hm[msg.Id] = &msg
	collection.slice = append(collection.slice, &msg)

	return msg, nil
}

// Update updates the given item in the collection.
func (collection *Collection) Update(msg Message) error {
	ptrItem, ok := collection.hm[msg.ID()]
	if ptrItem == nil || !ok {
		_, err := collection.Create(msg)
		return err
	}
	*ptrItem = msg

	return nil
}

// Delete deletes the given item from the collection.
func (_ *Collection) Delete(_ MessageID) error {
	return errors.New("not implement")
}
