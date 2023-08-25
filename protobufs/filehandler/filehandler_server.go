package filehandler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	common "github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"github.com/thomas-osgood/OGOR/protobufs/general"
	"google.golang.org/grpc/metadata"
)

// structure defining a Filehandler Server
type FHServer struct {
	filehandler.UnimplementedFileserviceServer
}

// function designed to download a file from the client's
// machine to the server.
//
// the filename will be transmitted via the rpc's metadata
// header information. the name cleansing will take place
// on the agent side so the server does not have to account
// for a different OS.
func (fhs *FHServer) DownloadFile(fsrv filehandler.Fileservice_DownloadFileServer) (err error) {
	var filename string
	var fptr *os.File
	var headerslice []string
	var ok bool
	var reqmeta metadata.MD

	// pull down all the header data from the incoming request.
	reqmeta, ok = metadata.FromIncomingContext(fsrv.Context())
	if !ok {
		return errors.New("unable to read download header")
	}

	// get name of file being uploaded by reading the
	// header data. if no filename is given, use
	// download_<timestamp> as the filename.
	headerslice = reqmeta.Get("filename")
	if (headerslice == nil) || (len(headerslice) < 1) {
		filename = fmt.Sprintf("download_%s", strings.ReplaceAll(time.Now().String(), " ", "_"))
	} else {
		filename = headerslice[0]
		if len(filename) < 1 {
			filename = fmt.Sprintf("download_%s", strings.ReplaceAll(time.Now().String(), " ", "_"))
		}
	}

	// open the local file for writing. create it if it does
	// not exist, overwrite the exiting file if it does.
	fptr, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error opening file to write: %s", err.Error())
	}

	// call the ReceiveFileBytes function to save the
	// data to the target local file. if there is an
	// error, the file will be closed and removed.
	if err = general.ReceiveFileBytes(fsrv, fptr); err != nil {
		fptr.Close()
		os.Remove(filename)
		return err
	}

	// make sure to send the acknowledgement that the
	// file has been uploaded successfully. if this
	// SendAndClose is not called, there will be communication
	// errors between the client and server.
	if err = fsrv.SendAndClose(&common.StatusMessage{Code: 0, Message: "file successfully uploaded"}); err != nil {
		return err
	}

	return nil
}

// function designed to upload a file from the server
// to the machine the client is running on.
func (fhx *FHServer) UploadFile(fr *filehandler.FileRequest, fsrv filehandler.Fileservice_UploadFileServer) (err error) {
	var filescanner *bufio.Reader
	var fptr *os.File

	fptr, err = os.Open(fr.Filename)
	if err != nil {
		return fmt.Errorf("unable to open file for upload: %s", err.Error())
	}
	defer fptr.Close()

	// create scanner to pass to TransmitFileBytes function.
	// this is what will read the file chunk-by-chunk.
	filescanner = bufio.NewReader(fptr)

	if err = general.TransmitFileBytes(fsrv, filescanner); err != nil {
		return err
	}

	return nil
}
