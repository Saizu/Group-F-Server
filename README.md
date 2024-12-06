# Group F Server

## Architecture

本リポジトリで管理されるプログラムはDockerコンテナとしてTengu712のマシン上にデプロイされる。

データベースへのリクエストを捌くサーバ(pkgs/back/)のエンドポイントを`http://skd-sv.skdassoc.work/`とする。
このエンドポイントへのリクエストはCloudflare TunnelによってTengu712のマシンのlocalhost:63245、さらに本Dockerコンテナの63245番ポートへ転送される。

管理画面を提供するサーバ(pkgs/front/)のエンドポイントはまだ未定である。

## Build and Run (Local)

1. Dockerをインストール
2. `docker compose up -d`

ローカルでサーバを実行する場合、エンドポイントは`http://localhost:63245`(データベースサーバ)や`http://localhost:8080`(管理画面サーバ)となる。
従って、クライアントのコードを`http://skd-sv.skdassoc.work/`に保ちながら開発を進めるためにはDNS設定を変える必要があることに注意せよ。

## Database Access (Local)

次のようにして、ローカル上のデータベースにログインできる。

1. `docker ps`を実行してコンテナ`server-gpfdb-*`のIDを取得
2. `docker exec -it <container-id> psql -U postgres`を実行
