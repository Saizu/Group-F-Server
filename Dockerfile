FROM ubuntu:22.04

RUN apt-get update \
 && apt-get install -y golang-go \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

COPY pkgs /usr/local/pkgs/

RUN cd /usr/local/ \
 && cd pkgs/back && go build -ldflags="-s -w" -trimpath -o /usr/local/bin/server && cd ../..

CMD /usr/local/bin/server
