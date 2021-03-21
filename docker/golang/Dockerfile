FROM golang:1.16.0-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
RUN mkdir /go/src/server
# ワーキングディレクトリの設定
WORKDIR /go/src/server
EXPOSE 8080
