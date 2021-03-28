


.PHONY: zebra clean

zebra:
	CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o ./build/output/ -ldflags "-w -s" -v ./rib/ribd

clean:
	rm -f ./build/output/*
