# 1-12

Both `Dockerfiles` pull the `example-frontend` directory from GitHub, no local clone needed.

## Usage

### Build

`docker build -f Dockerfile.ubuntu . -t example-frontend:ubuntu`

or

`docker build -f Dockerfile.go . -t example-frontend:go`

### Running

`docker run -p 8080:8080 example-frontend:<tag>`
