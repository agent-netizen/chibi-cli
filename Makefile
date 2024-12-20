VERSION=1.0
BINARY_NAME=chibi
BUILD_DIR=build
DEB_PATH=chibi_debian
LDFLAGS="-s -w"

LINUX_BIN=${BINARY_NAME}_x64_linux
WIN_BIN=${BINARY_NAME}_x64_win
APPLE_INTEL_BIN=${BINARY_NAME}_darwin_intel
APPLE_SILICON_BIN=${BINARY_NAME}_darwin_silicon

.PHONY: all

all: clean compile pack_deb

clean:
	if [ -d ${BUILD_DIR} ]; then rm -rf ${BUILD_DIR}; fi
	mkdir ${BUILD_DIR}

compile:
	go mod tidy
	GOARCH=amd64 GOOS=darwin go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${APPLE_INTEL_BIN}
	GOARCH=arm64 GOOS=darwin go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${APPLE_SILICON_BIN}
	GOARCH=amd64 GOOS=windows go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${WIN_BIN}.exe
	GOARCH=amd64 GOOS=linux go build -ldflags=${LDFLAGS} -v -o ${BUILD_DIR}/${LINUX_BIN}

pack_deb:
	echo "Packing debian archive"
	mkdir -p ${DEB_PATH}/DEBIAN ${DEB_PATH}/usr/local/bin
	touch ${DEB_PATH}/DEBIAN/control
	cp ${BUILD_DIR}/${LINUX_BIN} ${DEB_PATH}/usr/local/bin/chibi

	@echo "Package: chibi" > ${DEB_PATH}/DEBIAN/control
	@echo "Version: ${VERSION}" >> ${DEB_PATH}/DEBIAN/control
	@echo "Section: base" >> ${DEB_PATH}/DEBIAN/control
	@echo "Priority: optional" >> ${DEB_PATH}/DEBIAN/control
	@echo "Architecture: amd64" >> ${DEB_PATH}/DEBIAN/control
	@echo "Maintainer: Cosmic Predator" >> ${DEB_PATH}/DEBIAN/control
	@echo "Description: Chibi for AniList" >> ${DEB_PATH}/DEBIAN/control
	@echo "    A lightweight anime & manga tracker CLI app powered by AniList, written in Go." >> ${DEB_PATH}/DEBIAN/control
	@echo "Build-Using: go-1.23 (= 1.23.4)" >> chibi_debian/DEBIAN/control
	@echo "Homepage: https://github.com/CosmicPredator/chibi-cli" >> ${DEB_PATH}/DEBIAN/control

	dpkg-deb --build ${DEB_PATH}
	mv chibi_debian.deb ${BUILD_DIR}/chibi_debian.deb

	rm -rf ${DEB_PATH}/