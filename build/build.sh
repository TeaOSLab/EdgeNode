#!/usr/bin/env bash

function build() {
	ROOT=$(dirname $0)
	NAME="edge-node"
	VERSION=$(lookup-version "$ROOT"/../internal/const/const.go)
	DIST=$ROOT/"../dist/${NAME}"
	MUSL_DIR="/opt/homebrew/bin"
	SRCDIR=$(realpath "$ROOT/..")

	# for macOS users: precompiled gcc can be downloaded from https://github.com/messense/homebrew-macos-cross-toolchains
	GCC_X86_64_DIR="/usr/local/gcc/x86_64-unknown-linux-gnu/bin"
	GCC_ARM64_DIR="/usr/local/gcc/aarch64-unknown-linux-gnu/bin"

	OS=${1}
	ARCH=${2}
	TAG=${3}

	if [ -z "$OS" ]; then
		echo "usage: build.sh OS ARCH"
		exit
	fi
	if [ -z "$ARCH" ]; then
		echo "usage: build.sh OS ARCH"
		exit
	fi
	if [ -z "$TAG" ]; then
		TAG="community"
	fi

	echo "checking ..."
	ZIP_PATH=$(which zip)
	if [ -z "$ZIP_PATH" ]; then
		echo "we need 'zip' command to compress files"
		exit
	fi

	echo "building v${VERSION}/${OS}/${ARCH}/${TAG} ..."
	ZIP="${NAME}-${OS}-${ARCH}-${TAG}-v${VERSION}.zip"

	echo "copying ..."
	if [ ! -d "$DIST" ]; then
		mkdir "$DIST"
		mkdir "$DIST"/bin
		mkdir "$DIST"/configs
		mkdir "$DIST"/logs
		mkdir "$DIST"/data

		if [ "$TAG" = "plus" ]; then
			mkdir "$DIST"/scripts
			mkdir "$DIST"/scripts/js
		fi
	fi

	cp "$ROOT"/configs/api_node.template.yaml "$DIST"/configs
	cp "$ROOT"/configs/cluster.template.yaml "$DIST"/configs
	cp -R "$ROOT"/www "$DIST"/
	cp -R "$ROOT"/pages "$DIST"/

	# we support TOA on linux only
	if [ "$OS" == "linux" ] && [ -f "${ROOT}/edge-toa/edge-toa-${ARCH}" ]
	then
		if [ ! -d  "$DIST/edge-toa" ]
		then
			mkdir "$DIST/edge-toa"
		fi
		cp "${ROOT}/edge-toa/edge-toa-${ARCH}" "$DIST/edge-toa/edge-toa"
	fi

	echo "building ..."

	CC_PATH=""
	CXX_PATH=""
	CGO_LDFLAGS=""
	CGO_CFLAGS=""
	BUILD_TAG=$TAG
	if [[ "${OS}" == "linux" ]]; then
		if [ "${ARCH}" == "amd64" ]; then
			CC_PATH=$(command -v x86_64-linux-musl-gcc)
			CXX_PATH=$(command -v x86_64-linux-musl-g++)
			# if [ "$TAG" = "plus" ]; then
			# 	BUILD_TAG="plus,script,packet"
			# fi
		fi
		if [ "${ARCH}" == "arm64" ]; then
			CC_PATH=$(command -v aarch64-linux-musl-gcc)
			CXX_PATH=$(command -v aarch64-linux-musl-g++)
			# if [ "$TAG" = "plus" ]; then
			# 	BUILD_TAG="plus,script,packet"
			# fi
		fi
	fi

	# libpcap
	if [ "$OS" == "linux" ] && [[ "$ARCH" == "amd64" || "$ARCH" == "arm64" ]] &&  [ "$TAG" == "plus" ]; then
		CGO_LDFLAGS="-L${SRCDIR}/libs/libpcap/${ARCH} -lpcap -L${SRCDIR}/libs/libbrotli/${ARCH} -lbrotlienc -lbrotlidec -lbrotlicommon"
		CGO_CFLAGS="-I${SRCDIR}/libs/libpcap/src/libpcap -I${SRCDIR}/libs/libpcap/src/libpcap/pcap -I${SRCDIR}/libs/libbrotli/src/brotli/c/include"
	fi

	if [ -f $CC_PATH ]; then
		env CC=$CC_PATH \
		 CXX=$CXX_PATH GOOS="${OS}" \
		 GOARCH="${ARCH}" CGO_ENABLED=1 \
		 CGO_LDFLAGS="${CGO_LDFLAGS}" \
		 CGO_CFLAGS="${CGO_CFLAGS}" \
		 go build -trimpath -tags $BUILD_TAG -o "$DIST"/bin/${NAME} -ldflags "-linkmode external -extldflags -static -s -w" "$ROOT"/../cmd/edge-node/main.go
	else
		if [[ `uname` == *"Linux"* ]] && [ "$OS" == "linux" ] && [[ "$ARCH" == "amd64" || "$ARCH" == "arm64" ]] &&  [ "$TAG" == "plus" ]; then
			BUILD_TAG="plus,script,packet"
		fi

		env GOOS="${OS}" GOARCH="${ARCH}" CGO_ENABLED=1  CGO_LDFLAGS="${CGO_LDFLAGS}" CGO_CFLAGS="${CGO_CFLAGS}" go build -trimpath -tags $BUILD_TAG -o "$DIST"/bin/${NAME} -ldflags="-s -w" "$ROOT"/../cmd/edge-node/main.go
	fi

	if [ ! -f "${DIST}/bin/${NAME}" ]; then
		echo "build failed!"
		exit
	fi

	# delete hidden files
	find "$DIST" -name ".DS_Store" -delete
	find "$DIST" -name ".gitignore" -delete

	echo "zip files"
	cd "${DIST}/../" || exit
	if [ -f "${ZIP}" ]; then
		rm -f "${ZIP}"
	fi
	zip -r -X -q "${ZIP}" ${NAME}/
	rm -rf ${NAME}
	cd - || exit

	echo "OK"
}

function lookup-version() {
	FILE=$1
	VERSION_DATA=$(cat "$FILE")
	re="Version[ ]+=[ ]+\"([0-9.]+)\""
	if [[ $VERSION_DATA =~ $re ]]; then
		VERSION=${BASH_REMATCH[1]}
		echo "$VERSION"
	else
		echo "could not match version"
		exit
	fi
}

build "$1" "$2" "$3"
