run:
	go run -trimpath main.go

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o uniportal main.go

build-asset:
	go-bindata -pkg asset \
		-o asset/asset.go \
		views/...

upload:
	rsync -v --progress uniportal new-marsdev:/tmp/uniportal/uniportal
	rsync -v --progress config.json new-marsdev:/tmp/uniportal/config.json