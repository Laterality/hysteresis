OUT := hyst

.PHONY: clean build

clean:
	rm -rf ./bin ./vendor Gopkg.lock

build-linux:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o $(OUT) *.go

build-mac:
	export GO111MODULE=on
	env GOOS=darwin go build -ldflags="-s -w" -o $(OUT) *.go