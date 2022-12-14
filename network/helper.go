package network

import (
	network_pb "github.com/theneverse/neverse-kit/network/pb"
)

func Message(data []byte) *network_pb.Message {
	return &network_pb.Message{
		Data: data,
	}
}
