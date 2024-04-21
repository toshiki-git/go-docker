FROM golang:1.21.3

# コンテナ内にディレクトリを作成
RUN mkdir /app

# ワーキングディレクトリの設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 依存関係のダウンロード
RUN go mod download

# アプリケーションをビルド
RUN go build -o myapp ./cmd/myapp/main.go

CMD ["./myapp"]