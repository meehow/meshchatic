build:
	go build -ldflags="-s -w" -trimpath

deploy: build
	rsync -avzL web meshchatic meshchatic@vps:

proto:
	protoc -Iprotobufs --go_out=. protobufs/meshtastic/*.proto