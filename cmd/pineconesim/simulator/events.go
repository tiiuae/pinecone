// Copyright 2021 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simulator

import "github.com/matrix-org/pinecone/router/events"

type SimEvent interface {
	isEvent()
}

type NodeAdded struct {
	Node string
}

// Tag NodeAdded as an Event
func (e NodeAdded) isEvent() {}

type NodeRemoved struct {
	Node string
}

// Tag NodeRemoved as an Event
func (e NodeRemoved) isEvent() {}

type PeerAdded struct {
	Node string
	Peer string
}

// Tag PeerAdded as an Event
func (e PeerAdded) isEvent() {}

type PeerRemoved struct {
	Node string
	Peer string
}

// Tag PeerRemoved as an Event
func (e PeerRemoved) isEvent() {}

type SnakeAscUpdate struct {
	Node string
	Peer string
	Prev string
}

// Tag SnakeAscUpdate as an Event
func (e SnakeAscUpdate) isEvent() {}

type SnakeDescUpdate struct {
	Node string
	Peer string
	Prev string
}

// Tag SnakeDescUpdate as an Event
func (e SnakeDescUpdate) isEvent() {}

type eventHandler struct {
	node string
	ch   <-chan events.Event
}

func (h eventHandler) Run(sim *Simulator) {
	for {
		event := <-h.ch
		switch e := event.(type) {
		case events.PeerAdded:
			sim.handlePeerAdded(h.node, e.PeerID, int(e.Port))
		case events.PeerRemoved:
			sim.handlePeerRemoved(h.node, e.PeerID, int(e.Port))
		case events.SnakeAscUpdate:
			sim.handleSnakeAscUpdate(h.node, e.PeerID)
		case events.SnakeDescUpdate:
			sim.handleSnakeDescUpdate(h.node, e.PeerID)
		default:
			sim.log.Println("Unhandled event!")
		}
	}
}