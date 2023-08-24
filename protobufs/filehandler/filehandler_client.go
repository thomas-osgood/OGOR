package filehandler

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/thomas-osgood/OGOR/misc/crosscompile"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"github.com/thomas-osgood/OGOR/protobufs/general"
	"google.golang.org/grpc/metadata"
)

// function deisgned to download a file from the
// server to the client machine..
func DownloadFile(client filehandler.FileserviceClient, targetfilename string, outfilename string) (err error) {
	var cancel context.CancelFunc
	var dlctx context.Context
	var fclnt filehandler.Fileservice_UploadFileClient
	var fptr *os.File

	dlctx, cancel = context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	// send request to the server, asking for it to
	// send an UploadFileClient object so the contents
	// of the target file down to it can be streamed to
	// the client.
	//
	// if this request does not complete within 10 seconds
	// a timeout error will be thrown. this timeout does
	// not apply to the actual streaming of data, only the
	// UploadFileClient object request.
	fclnt, err = client.UploadFile(dlctx, &filehandler.FileRequest{Filename: targetfilename})
	if err != nil {
		return err
	}

	// open file to save data in. if the file does not
	// exist, it will be created. if the file does exist,
	// it will be overwritten.
	fptr, err = os.OpenFile(outfilename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer fptr.Close()

	if err = general.ReceiveFileBytes(fclnt, fptr); err != nil {
		fptr.Close()
		os.Remove(outfilename)
		return err
	}

	err = fclnt.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

// function designed to upload a file from the client to
// the server.
func UploadFile(client filehandler.FileserviceClient, filename string) (err error) {
	var cancel context.CancelFunc
	var cleanfilename string
	var cutfilename []string
	var filescanner *bufio.Reader
	var fptr *os.File
	var headerdata metadata.MD = make(metadata.MD)
	var ulclient filehandler.Fileservice_DownloadFileClient
	var ulctx context.Context

	cutfilename = strings.Split(filename, crosscompile.PATH_SEPARATOR)
	cleanfilename = cutfilename[len(cutfilename)-1]

	fptr, err = os.Open(filename)
	if err != nil {
		return err
	}
	defer fptr.Close()

	ulctx, cancel = context.WithTimeout(context.Background(), time.Duration(general.DEFAULTTIMEOUT*int(time.Second)))
	defer cancel()

	// add filename to outgoing headers (ie metadata)
	// in the context attached to the upload request.
	headerdata.Set("filename", cleanfilename)

	ulclient, err = client.DownloadFile(ulctx)
	if err != nil {
		return err
	}

	filescanner = bufio.NewReader(fptr)

	err = general.TransmitFileBytes(ulclient, filescanner)
	if err != nil {
		return err
	}

	_, err = ulclient.CloseAndRecv()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}
