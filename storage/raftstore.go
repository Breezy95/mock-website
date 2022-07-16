package storage

import (
	"log"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc/peer"
)

type raftNode struct {
	node   *raft.Raft
	logger log.Logger

	peerlst []*peer.Peer
}
