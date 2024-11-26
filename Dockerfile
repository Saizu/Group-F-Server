FROM ubuntu:22.04

RUN apt-get update \
 && apt-get install -y --no-install-recommends gpg-agent software-properties-common \
 && add-apt-repository ppa:longsleep/golang-backports \
 && apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates golang-go \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

COPY pkgs /usr/local/pkgs

RUN cd /usr/local/ \
 && cd pkgs/back/ \
 && go get server \
 && go build -o /usr/local/bin/server \
 && cd ../../

CMD /usr/local/bin/server
