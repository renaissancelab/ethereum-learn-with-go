// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package pss

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethersphere/swarm/pss/message"
)

var (
	rawTopic = message.Topic{}
)

// PssAddress is an alias for []byte. It represents a variable length address
type PssAddress []byte

// MarshalJSON implements the json.Marshaler interface
func (a PssAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(hexutil.Encode(a[:]))
}

// UnmarshalJSON implements the json.Marshaler interface
func (a *PssAddress) UnmarshalJSON(input []byte) error {
	b, err := hexutil.Decode(string(input[1 : len(input)-1]))
	if err != nil {
		return err
	}
	for _, bb := range b {
		*a = append(*a, bb)
	}
	return nil
}

// Signature for a message handler function for a Message
// Implementations of this type are passed to Pss.Register together with a topic,
type HandlerFunc func(msg []byte, p *p2p.Peer, asymmetric bool, keyid string) error

type handlerCaps struct {
	raw  bool
	prox bool
}

// Handler defines code to be executed upon reception of content.
type handler struct {
	f    HandlerFunc
	caps *handlerCaps
}

// NewHandler returns a new message handler
func NewHandler(f HandlerFunc) *handler {
	return &handler{
		f:    f,
		caps: &handlerCaps{},
	}
}

// WithRaw is a chainable method that allows raw messages to be handled.
func (h *handler) WithRaw() *handler {
	h.caps.raw = true
	return h
}

// WithProxBin is a chainable method that allows sending messages with full addresses to neighbourhoods using the kademlia depth as reference
func (h *handler) WithProxBin() *handler {
	h.caps.prox = true
	return h
}

// the stateStore handles saving and loading PSS peers and their corresponding keys
// it is currently unimplemented
type stateStore struct {
	values map[string][]byte
}

func (store *stateStore) Load(key string) ([]byte, error) {
	return nil, nil
}

func (store *stateStore) Save(key string, v []byte) error {
	return nil
}
