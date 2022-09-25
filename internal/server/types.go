package server

import (
	"errors"

	"github.com/google/uuid"

	"github.com/nullentrypoint/tekken7-web-displayer/pkg/tekken"
)

type IdType string

type MyItem struct {
	Id   IdType            `json:"id,omitempty"`
	Info tekken.PlayerInfo `json:"info,omitempty"`
	// ...
}

func (myItem MyItem) ID() IdType {
	return myItem.Id
}

type MyCollection struct {
	// ...
	hm    map[IdType]*MyItem
	slice []*MyItem
}

// NewMyCollection creates and initialize an instance of MyCollection.
func NewMyCollection() (*MyCollection, error) {
	return &MyCollection{
		hm:    make(map[IdType]*MyItem),
		slice: []*MyItem{},
	}, nil
}

// All retreives all the items in MyCollecttion and returns them in a slice.
func (myColl *MyCollection) All() ([]MyItem, error) {
	out := make([]MyItem, len(myColl.slice))
	for _, value := range myColl.slice {
		out = append(out, *value)
	}

	return out, nil
}

// Item retreives the item with the given ID from MyCollection and returns it.
func (myColl *MyCollection) Item(ID IdType) (MyItem, error) {
	out, ok := myColl.hm[ID]
	if !ok {
		return MyItem{}, errors.New("not found")
	}

	return *out, nil
}

// Create adds the given item to the collection and returns it AFTER having set its id.
func (myColl *MyCollection) Create(myItem MyItem) (MyItem, error) {
	// ...
	// Set item id
	myItem.Id = IdType(uuid.New().String()) // ...

	myColl.hm[myItem.Id] = &myItem
	myColl.slice = append(myColl.slice, &myItem)
	// ...
	return myItem, nil
}

// Update updates the given item in the collection.
func (myColl *MyCollection) Update(myItem MyItem) error {

	ptrItem, ok := myColl.hm[myItem.ID()]
	if ptrItem == nil || !ok {
		_, err := myColl.Create(myItem)
		return err
	}
	*ptrItem = myItem

	return nil
}

// Delete deletes the given item from the collection.
func (myColl *MyCollection) Delete(ID IdType) error {
	return errors.New("not implement")
}
