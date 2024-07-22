/***
We are going to encode and transmit, and then deode a transaction with gob
We are going to send a banking transaction from client to a server using a dummy network
The transaction is a struct that also has an embedded user struct. This is to show how easy it is to encode complex data

We will demonstrate the flexibility of the gob paockage/protocol
The client and server structs do not match in several ways
The client's user is a pointer, the server's user is not
The amounts are of different float types; client is float64, server is a *float32
Some fields are missing in the server type that are present in the client type

We will use the bytes package to store the encoded data
Once encoded, we can use gob to work with the binary data
***/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

// create a client side User struct
type UserClient struct {
	ID   string
	Name string
}

// create a client-side transaction, Tx is shorthand for transaction
type TxClient struct {
	ID          string
	User        *UserClient
	AccountFrom string
	AccountTo   string
	Amount      float64
}

// create server side User struct; does not match client model
type UserServer struct {
	ID string
}

// create struct for server-side transaction. User here is a not pointer, the amount is
type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32
}

// create sendToServer function; it takes in an io.Reader and returns a server side tx and an error
func sendToServer(net io.Reader) (*TxServer, error) {
	// create decoding target
	tx := &TxServer{}
	// create decoder with network as the source
	dec := gob.NewDecoder(net)
	// decode and capture any errors
	err := dec.Decode(tx)
	// return decoded data and errors captured
	return tx, err
}

func main() {
	// create dummy network; a buffer of bytes from the bytes package
	var net bytes.Buffer

	//create dummy data with client-side structs
	clientTx := &TxClient{
		ID: "1234567",
		User: &UserClient{
			ID:   "ABCDEF",
			Name: "Jake",
		},
		AccountFrom: "Neteyam",
		AccountTo:   "Loak",
		Amount:      9.99,
	}

	// encode the data. The target for the encoding is the dummy network
	enc := gob.NewEncoder(&net)
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("error encoding: ", err)
	}

	// send the encoded data to server
	serverTx, err := sendToServer(&net)
	if err != nil {
		log.Fatal("server error: ", err)
	}

	fmt.Printf("%#v\n", serverTx)
}
