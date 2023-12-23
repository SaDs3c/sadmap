# SadMap

![SadMap Logo](./art.txt)

SadMap is a simple command-line tool for scanning ports on a target host or domain. It provides the ability to scan a single port, a range of ports, or a list of selected ports. Powered by sadsec.

Its all going to be dependent on the [golang net package](https://pkg.go.dev/net) currently scans TCP and UDP in range

## Features

- Scan a single port
- Scan a range of ports
- Scan selected ports
- Display results in a table format

## Usage
```
$ sadmap -t <target> -p <port(s)>
```

*<port(s)>*: Specify the port or ports to scan. Examples:

Single port: `-p 80`

Port range: `-p 1-25`

Multiple ports: `-p 80,22,443`

*<target>*: Specify the target host or domain. Example:

`-t example.com`
