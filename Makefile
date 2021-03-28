


.PHONY: zebra

zebra:
	go build -o ./build/output/ -ldflags "-w -s" -v ./cmd/ribd
