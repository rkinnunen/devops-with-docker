FROM openjdk:8
EXPOSE 8080

RUN apt update && apt upgrade -y && apt install git -y
RUN git init ./material-applications && cd material-applications && \
    git remote add -f origin https://github.com/docker-hy/material-applications.git && \
    git config core.sparseCheckouyt true && \
    echo "/spring-example-directory" >> .git/info/sparse-checkout && \
    git pull origin main

WORKDIR /material-applications/spring-example-project
RUN ./mvnw package
CMD java -jar ./target/docker-example-1.1.3.jar