.PHONY: walker numberDistribution

wasm_js:
	cp '$(shell go env GOROOT)/misc/wasm/wasm_exec.js' wasm_bin/

walker: walker/examples/main.go
	cd walker && \
	GOOS=js GOARCH=wasm go build -o ../wasm_bin/walker.wasm ./examples/main.go

numberDistribution: numberDistribution/main.go
	cd numberDistribution && \
	GOOS=js GOARCH=wasm go build -o ../wasm_bin/numberDistribution.wasm ./main.go

all: wasm_js walker numberDistribution

clean:
	rm -rf wasm_bin/*