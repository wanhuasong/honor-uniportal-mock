run:
	go run -trimpath main.go

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o bin/uniportal main.go

build-asset:
	go-bindata -pkg asset \
		-o asset/asset.go \
		views/...

start:
	nohup ./uniportal > ./uniportal.log 2>&1 &
	ps aux | grep genericfs | grep -v grep

log:
	tail -f ./uniportal.log

stop:
	ps aux | grep uniportal | grep -v grep | awk '{print $$2}' | xargs kill -9

upload:
	rsync -v --progress bin/uniportal new-marsdev:/tmp/uniportal/uniportal
	rsync -v --progress config.json new-marsdev:/tmp/uniportal/config.json