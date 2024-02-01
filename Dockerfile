FROM golang:1.21.3

# コンテナ内にディレクトリを作成
RUN mkdir /app

# ワーキングディレクトリの設定
WORKDIR /app

# Goの依存関係の管理ファイルをコピー
COPY go.mod go.sum ./

# 依存関係のダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .