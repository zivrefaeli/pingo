# pingo

ping networking utility implemented in Go

## Use locally

You can run pingo using:

```bash
go run main.go
```

Or you can build it and then execute:

```bash
go build
./pingo
```

## Usage

Run `./pingo -h` to view help:

```
send ICMP ECHO_REQUEST to network hosts

Usage:
  pingo [TARGET_NAME] [flags]

Flags:
  -n, --count int     Number of echo requests to send. (default 4)
  -h, --help          help for pingo
  -l, --size uint16   Send buffer size. (default 32)
```
