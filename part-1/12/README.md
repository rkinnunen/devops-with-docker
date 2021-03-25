# 1-12

Both `Dockerfile`s pull the `example-frontend` directory from GitHub, no local clone needed.

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
