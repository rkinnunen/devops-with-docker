FROM ubuntu:18.04

COPY script.sh .

RUN apt update && apt upgrade -y && apt install curl -y

CMD ./script.sh