FROM golang:1.15.10-buster

EXPOSE 8080

ENV PORT=8080

WORKDIR /usr/src/app

RUN git init ./material-applications && cd material-applications && \
    git remote add -f origin https://github.com/docker-hy/material-applications.git && \
    git config core.sparseCheckouyt true && \
    echo "/example-backend" >> .git/info/sparse-checkout && \
    git pull origin main

WORKDIR ./material-applications/example-backend

RUN go build && \
    go test ./...

CMD ./server