# Group F Server

## Architecture

本リポジトリで管理されるプログラムはDockerコンテナとしてTengu712のマシン上にデプロイされる。

サーバ(pkgs/back)のエンドポイントを`http://skd-sv.skdassoc.work/`とする。
このエンドポイントへのリクエストはCloudflare TunnelによってTengu712のマシンのlocalhost:63245、さらに本Dockerコンテナの63245番ポートへ転送される。

## Build and Run (Local)

ローカルでサーバを実行する場合、エンドポイントは`http://localhost:63245`となる。
従って、クライアントのコードを`http://skd-sv.skdassoc.work/`に保ちながら開発を進めるためにはDNS設定を変える必要がある。

### With Docker

1. Dockerをインストール
2. `docker compose up`

### Without Docker

1. Goをインストール
2. `cd pkgs/back && go build -ldflags="-s -w" -trimpath -o ../../bin/server && cd ../..`
3. `./bin/server`
