#!/bin/bash

podman run --rm \
	-v $(pwd):/go/src/ \
	-e "GOPATH=/go" \
	--privileged \
	tinygo/tinygo-dev:latest \
	bash -c "cd src && tinygo build -o ./assets/wasm.wasm --target=wasm --no-debug ./module/main.go"