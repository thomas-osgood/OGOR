package general

import (
	"github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
)

// a generic object designed to read ByteString data
// via a Recv function.
type ByteReceiver interface {
	Recv() (*common.ByteString, error)
}

// a generic object designed to transmit ByteString
// data via a Send function.
type ByteTransmitter interface {
	Send(*common.ByteString) error
}

// a generic object designed to read FileChunk
// data via a Recv function.
type Receiver interface {
	Recv() (*filehandler.FileChunk, error)
}

// a generic object designed to transmit FileChunk
// data via a Send function.
type Transmitter interface {
	Send(*filehandler.FileChunk) error
}
