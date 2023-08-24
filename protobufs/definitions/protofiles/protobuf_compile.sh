#/bin/bash
#
# For reference:
#
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
#
# Above line was taken from: https://grpc.io/docs/languages/go/quickstart/
#
# The expected directory tree for this script to work is as follows:
#
# - protobufs:
# 	    |
# 	    |
# 	    common
# 	    |	
# 	    |	
# 	    filehandler
# 	    |
# 	    |
# 	    definitions (you are here)
#


# Step 1: Build the protobuf Golang files.
# 
# These files will be placed in the current directory. The movement
# of them will be handled in the next step.
protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative common.proto
protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative filehandler.proto


# Create variable for the target root directory the .pb.go  
# files will live in.
TGTROOT=".."

# Create directories if they do not already exist.
mkdir $TGTROOT/common/
mkdir $TGTROOT/filehandler/

# Step 2: Move protobuf Golang files.
#
# The target .pb.go files will be moved to the correct location
# in the controller directory based on what proto file they
# are related to/generated from.
mv common*.pb.go $TGTROOT/common/
mv filehandler*.pb.go $TGTROOT/filehandler/
