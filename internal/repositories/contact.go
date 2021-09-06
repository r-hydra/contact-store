package repositories

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

type Contact struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type ContactRepository struct {
	contacts []Contact
}

func (c *ContactRepository) List() []Contact {
	return c.contacts
}

func (c *ContactRepository) Store(contact Contact) {
	c.contacts = append(c.contacts, contact)
}

func (c *ContactRepository) Update(id string, contact Contact) {
	index := c.IndexOf(Contact{Id: id})
	if index < 0 {
		return
	}

	c.contacts[index] = contact
}

func (c *ContactRepository) Remove(id string) {
	index := c.IndexOf(Contact{Id: id})
	if index < 0 {
		return
	}

	c.contacts[index] = c.contacts[len(c.contacts)-1]
	c.contacts = c.contacts[:len(c.contacts)-1]
}

func (c *ContactRepository) IndexOf(contact Contact) int {
	index := -1

	for i, ct := range c.contacts {
		if ct.Id == contact.Id {
			index = i
			break
		}
	}

	return index
}

func (c *ContactRepository) Preload() {
	defer log.Println("Loaded: ", c.contacts)
	c.contacts = []Contact{}
	data, err := os.Open("state")
	if err != nil {
		_ = fmt.Errorf("error %s", err)
		return
	}
	d := gob.NewDecoder(data)

	// Decoding the serialized data
	err = d.Decode(&c.contacts)
	if err != nil {
		return
	}
}

func (c *ContactRepository) Commit() {
	defer log.Println("Cached!")
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(c.contacts)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("state", b.Bytes(), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
