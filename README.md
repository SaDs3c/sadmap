# sadmap
A network scanner and mapper 

Its all going to be dependent on the [golang net package](https://pkg.go.dev/net)

currently scans TCP and UDP in range


go run main.go -t example.com -p 80

go run main.go -t example.com -p 80-83

go run main.go -t example.com -p 22,25,80,443