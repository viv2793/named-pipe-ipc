# named-pipe-ipc
This is repository contains implementation for Named-Pipe Client in Golang(Go).

Windows Named-pipe Documentation : https://docs.microsoft.com/en-us/windows/win32/ipc/named-pipes

## Background
Recently, I came across a requirement that we needed to communicate between two different running executables where one executable was written in C# and second one executable was written in Golang. So I decided to use named-pipe as a way to implement IPC. This repo contains named-pipe client and named-pipe server implementation purely written in Golang. 

There is an another repo which contains [named-pipe server](https://github.com/viv2793/named-pipe-server) implementation in C#.

## Steps to run
- Clone the repo and run navigate to project repo directory.
- Open windows powershell(Admin) and run server using below command
```
go run server.go
```
- Above command would start a named-pipe server and this server can send/receive message from any named-pipe client.

- Open windows powershell(Admin) and run client using below command
```
go run client.go
```
- Note that above command would only start a named-pipe client which can connect to the already existing named-pipe server. So it is necessarry to run the server before running the client.

- There is another repo which can be used for starting named-pipe server for IPC. This server is implemented in C#. Link - [Named-Pipe Server Repo](https://github.com/viv2793/named-pipe-server)

## Supported Platform/Version
- This implemtation is only for Windows currently.
- Go version used while writing - go 1.13

## Credits - 
- https://github.com/natefinch

