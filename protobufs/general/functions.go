package general

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"google.golang.org/grpc/status"
)

// function designed to read ByteStrings from a stream
// and save the data to a file.
func ReceiveByteStrings(receiver ByteReceiver, fptr *os.File) (err error) {
	var currenByteString *common.ByteString

	// read the stream ByteString-by-ByteString and
	// save the data to the file. If End-Of-File is
	// encountered, break the loop and return nil.
	for {
		currenByteString, err = receiver.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf(status.Convert(err).Message())
		}

		_, err = fptr.Write(currenByteString.GetChunk())
		if err != nil {
			return err
		}
	}

	return nil
}

// function designed to read bytes from a stream and
// save them to a file.
func ReceiveFileBytes(receiver Receiver, fptr *os.File) (err error) {
	var currentChunk *filehandler.FileChunk

	// read the stream chunk-by-chunk and save the data
	// to the local file.
	for {
		currentChunk, err = receiver.Recv()
		if err != nil {
			// the stream has been closed. there is no
			// more data to save. break the loop.
			if err == io.EOF {
				break
			}
			return fmt.Errorf(status.Convert(err).Message())
		}

		// write the chunk to the output file. trim the
		// null bytes to insure the file is written correctly
		// and no extraneous bytes are added to the end.
		_, err = fptr.Write(currentChunk.Chunk)
		if err != nil {
			return err
		}
	}

	return nil
}

// function designed to transmit bytes over the wire using
// ByteString objects.
func TransmitByteStrings(transmitter ByteTransmitter, byteReader *bytes.Reader) (err error) {
	var currentBlock []byte = make([]byte, CHUNKSIZE)

	// iterate through the slice of bytes and transmit
	// each buffer to the receiver. if an error is encountered,
	// return it unless it is an EOF error.
	for {
		_, err = byteReader.Read(currentBlock)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		err = transmitter.Send(&common.ByteString{Chunk: currentBlock})
		if err != nil {
			return fmt.Errorf("error sending bytes: %s", fmt.Errorf(status.Convert(err).Message()))
		}
	}

	return nil
}

// function designed to read a file and transmit its contents
// to another machine via a Transmitter object. this takes in
// a pointer to a bufio.Reader object which should already be
// pointing to a target file.
func TransmitFileBytes(transmitter Transmitter, scanner *bufio.Reader) (err error) {
	var currentChunk []byte = make([]byte, CHUNKSIZE)
	var bytesread int

	// loop through file, reading it chunk by chunk and
	// uploading it to the requester.
	for {
		// read the current chunk from the file.
		bytesread, err = scanner.Read(currentChunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error scanning file during upload: %s", err.Error())
		}

		// transmit the current chunk to the requester.
		err = transmitter.Send(&filehandler.FileChunk{Chunk: currentChunk[:bytesread]})
		if err != nil {
			return fmt.Errorf("error transmitting chunk: %s", fmt.Errorf(status.Convert(err).Message()))
		}
	}

	return nil
}
