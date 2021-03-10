FROM snyk/snyk-cli:npm

MAINTAINER Patrick Pierson patrick.c.pierson@gmail.com

RUN apt-get update && apt-get install wget -y
RUN wget https://golang.org/dl/go1.16.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz && rm -f go1.16.linux-amd64.tar.gz
ENV PATH "$PATH:/usr/local/go/bin"

