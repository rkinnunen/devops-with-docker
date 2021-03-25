FROM openjdk:8
EXPOSE 8080

WORKDIR /usr/src/app

ENV PROJECT=spring-example-project

RUN apt update && apt install git -y

RUN git init && \
    git remote add -f origin https://github.com/docker-hy/material-applications.git && \
    git fetch origin && \
    git checkout origin/main -- $PROJECT

WORKDIR ./$PROJECT

RUN ./mvnw package
CMD java -jar ./target/docker-example-1.1.3.jar