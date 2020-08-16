# named-pipe-ipc-client
This is repository contains implementation for Named-Pipe Client in Golang(Go).

## Background
Recently, I came across a requirement that we needed to communicate between two different running executables where one executable was written in C# and second one executable was written in Golang. So I decided to use named-pipe as a way to implement IPC. This repo contains named-pipe client implementation purely written in Golang.

## Steps to run
- Clone the repo and run navigate to project repo directory.
- Open windows powershell(Admin) and run client using below command
```
go run client.go
```
- Note that above command only connect to the already existing named-pipe server. So it is necessarry to run a server before running the client using above command. 
- There is an another repo which creates named-pipe server for IPC. Thats server implementation is in C# and can be used to run named-pipe server. 
- Link - [Named-Pipe Server Repo](https://github.com/viv2793/named-pipe-ipc)

## Supported Platform/Version
- This implemtation is only for Windows currently.
- Go version used while writing - go 1.13


