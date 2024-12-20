BINARY_NAME=chibi
BUILD_DIR=build
LDFLAGS="-s -w"

create_build_dir:
	mkdir ${BUILD_DIR}

compile:
	GOARCH=amd64 GOOS=darwin go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${BINARY_NAME}_darwin_intel
	GOARCH=arm64 GOOS=darwin go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${BINARY_NAME}_darwin_silicon
	GOARCH=amd64 GOOS=windows go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${BINARY_NAME}_win.exe
	GOARCH=amd64 GOOS=linux go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${BINARY_NAME}_linux

run: create_build_dir compile