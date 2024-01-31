FROM golang:1.21.3
# コンテナ内にディレクトリを作成
RUN mkdir /go/src/app/
# ワーキングディレクトリの設定
WORKDIR /go/src/app/
# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY ./ /go/src/app/