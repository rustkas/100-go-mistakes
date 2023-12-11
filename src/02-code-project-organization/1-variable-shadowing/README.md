Make module:

```
go mod init 100-go-mistakes 
```

Run code:

```
go run main.go -function listing1
go run main.go -function listing2
go run main.go -function listing3
go run main.go -function listing4

go run main.go -function listing1 && go run main.go -function listing2 && go run main.go -function listing3 && go run main.go -function listing4
```