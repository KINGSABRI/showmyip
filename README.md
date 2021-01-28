# showmyip
A very simple command line to get your public IP 

## Why?
I've built this greatly silly app to learn Go and I needed a command line to show my public ip address.

## Usage

- compile
```
$ git clone https://github.com/KINGSABRI/showmyip.git
$ cd showmyip
$ go build main.go
$ sudo cp main /usr/bin/
```

- run it
```
showmyip -h
Usage of showmyip:
  -a    Show my public IPv4 & IPv6.
  -v4
        Show my public IPv4. (default true)
  -v6
        Show my public IPv6.
```
