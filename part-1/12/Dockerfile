FROM ubuntu

EXPOSE 5000

RUN apt update && apt install -y git curl bash && \
    curl -sL https://deb.nodesource.com/setup_14.x | bash && \
    apt install nodejs && \
    git init ./material-applications && cd material-applications && \
    git remote add -f origin https://github.com/docker-hy/material-applications.git && \
    git config core.sparseCheckouyt true && \
    echo "/example-frontend" >> .git/info/sparse-checkout && \
    git pull origin main

WORKDIR ./material-applications/example-frontend

RUN npm install && npm install -g serve

# COPY ./example-frontend ./example-frontend

# # RUN apt update && apt upgrade -y && apt install -y curl && \
# #     curl -sL https://deb.nodesource.com/setup_14.x | bash && \
# #     apt install -y nodejs && \
# #     cd example-frontend && \
# #     npm install && \
# #     npm install -g serve


CMD serve -s -l 5000 build

