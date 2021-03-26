# 1-14

Both `Dockerfiles` pull the `example-backend` directory from GitHub, no local clone needed.

## Usage

### Build

`docker build -f Dockerfile.ubuntu . -t example-backend:ubuntu`

or

`docker build -f Dockerfile.go . -t example-backend:go`

### Running

`docker run -p 8080:8080 example-backend:<tag>`
