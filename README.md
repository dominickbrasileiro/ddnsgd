# ddnsgd
"ddnsgd" is a command line tool to update dynamic DNS records on Google Domains.
You can set an interval for it to fetch your IPv4 address and update the DDNS record.

## Usage
```
Usage:
  ddnsgd [flags]

Flags:
  -h, --help              help for ddnsgd
  -H, --hostname string   subdomain.yourdomain.com
  -i, --interval int      The interval in seconds to fetch the IPv4 address and update the DNS record.
  -p, --password string   The generated password associated with the host that is to be updated.
  -u, --username string   The generated username associated with the host that is to be updated.
```

### Example
```sh
$ ddnsgd --hostname subdomain.yourdomain.com --interval 30 --username GENERATED_USERNAME --password GENERATED_PASSWORD
```

## Installation
### Docker and Docker Compose
The "ddnsgd" docker image is available upon [ghcr.io](https://github.com/dominickbrasileiro/ddnsgd/pkgs/container/ddnsgd) registry.
You can pull it running the command below:
```sh
$ docker pull ghcr.io/dominickbrasileiro/ddnsgd
```

If you want to pull a specific version of "ddnsgd", use a tag:
```sh
$ docker pull ghcr.io/dominickbrasileiro/ddnsgd:1.1.0
```

To run the "ddnsgd" image, you'll need to provide the following environment variables:
- HOSTNAME
- INTERVAL
- USERNAME
- PASSWORD

Example using `docker run` command:
```sh
$ docker run \
-e HOSTNAME=subdomain.yourdomain.com \
-e INTERVAL=30 \
-e USERNAME=GENERATED_USERNAME \
-e PASSWORD=GENERATED_PASSWORD \
--name ddnsgd ghcr.io/dominickbrasileiro/ddnsgd
```

Example using `docker-compose.yml`:
```yml
version: "3.8"
services:
  ddnsgd:
    container_name: "ddnsgd"
    image: "ghcr.io/dominickbrasileiro/ddnsgd"
    restart: "always"
    environment:
      - "HOSTNAME=subdomain.yourdomain.com"
      - "INTERVAL=30"
      - "USERNAME=GENERATED _USERNAME"
      - "PASSWORD=GENERATED_PASSWORD"
```

### Building from source
#### Requirements
- [Go](https://go.dev/) (1.19 or later)

1. Using make and Makefile
```sh
$ git clone https://github.com/dominickbrasileiro/ddnsgd.git
$ cd ddnsgd
$ make build
```

2. Without Makefile
```sh
$ git clone https://github.com/dominickbrasileiro/ddnsgd.git
$ cd ddnsgd
$ go get -v ./...
$ go build -o ./bin/ddnsgd .
```

Now you have the "ddnsgd" binary in `./bin/ddnsgd`. If you want to use it globally, run the command below:
```sh
$ sudo mv ./bin/ddnsgd /usr/local/bin
```

---

Made with ❤️ by Dominick Brasileiro.

Feel free [to reach out](https://www.linkedin.com/in/dominickbrasileiro/)!

[![Linkedin Badge](https://img.shields.io/badge/-LinkedIn-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/dominickbrasileiro/)](https://www.linkedin.com/in/dominickbrasileiro/)
