# 1-12

The `Dockerfile`is based on Ubuntu and is more based on the instructions given. `Dockerfile.node` does the same thing, with the image directly based on Node LTS.

Both pull the `example-frontend` directory from GitHub, no local clone needed.

## Usage

### Build

```
docker build -f Dockerfile.ubuntu . -t example-frontend:ubuntu
docker build -f Dockerfile.node . -t example-frontend:node
```

### Running

```
docker run -p 5000:5000 example-frontend:<tag>
```
