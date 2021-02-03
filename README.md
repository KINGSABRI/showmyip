# showmyip
A very simple command line to get your public IP 

## Why?
I've built this greatly silly app to learn Go and I needed a command line to show and periodically log my public ip address during external pentest and red teaming engagements.

## Usage

- compile
```
$ git clone https://github.com/KINGSABRI/showmyip.git
$ cd showmyip
# go get -d ./...
$ go build main.go
$ sudo cp main /usr/bin/
```

- run it
```
showmyip -h
Usage of showmyip:
  -4    Show my public IPv4.
  -6    Show my public IPv6.
  -a    Show my public IPv4 & IPv6.
  -d int
        Run in loop to notify every X minutes. (default 10)
  -l string
        Log new IP address (use with '-n' flag).
  -n    Show desktop notifications. (Continues check every 10 min)
```

### Examples 
```
$ showmyip
$ showmyip -4 
$ showmyip -6 
$ showmyip -a
$ showmyip -a -n -l ip.log -d 10
```

## Run it as a service

### Linux (Systemd)

- Create a systemd service file

```
sudo nano /etc/systemd/system/showmyip.service
```

- Edit then Paste the following configuration file
```
[Unit]
Description = ShowMyIP Desktop Notifcation and Log
After = network.target
StartLimitIntervalSec = 0

[Service]
Type = simple
User = KING
ExecStart = /usr/bin/showmyip -n -4 -l /home/YOURUSER/.showmyip/ip.log -d 20
Restart = always
RestartSec = 1
StartLimitBurst = 5
StartLimitIntervalSec = 10
WantedBy=multi-user.target

[Install]
WantedBy=multi-user.target
```

```
$ sudo systemctl start showmyip.service
$ sudo systemctl enable showmyip.service
```
