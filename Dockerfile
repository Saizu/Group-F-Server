FROM ubuntu:22.04

RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
 && apt-get update \
 && apt-get install -y --no-install-recommends gpg-agent software-properties-common \
 && add-apt-repository ppa:longsleep/golang-backports \
 && apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates golang-go \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY pkgs /usr/local/pkgs

RUN cd /usr/local/pkgs/back/sqlc/ \
 && ~/go/bin/sqlc generate \
 && cd ../ \
 && go get server \
 && go build -o ../../bin/server -ldflags="-s -w" -trimpath

CMD /usr/local/bin/server
