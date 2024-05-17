# Portarum

Basic port scanning tool like Nmap's. Implementing only a subset of functionalities of Nmap's port scanning
functionalities, Portarum will only implement simple port scanning for IP address (TCP connect scan technique).

Usage

```
go run exec/cmd/main.go -p 9300-9301 127.0.0.1
```
