build:
	go build -gcflags='all=-N -l' -v  -o  ammtest

.all: build
