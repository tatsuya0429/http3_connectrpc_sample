# Http3 Connect RPC Sample

以下の記事をもとに検証環境を実装

- [gRPC Over HTTP/3](https://kmcd.dev/posts/grpc-over-http3/)

## 事前準備

- secretsディレクトリに自己署名証明書を発行する

## 実行

サーバーを起動

```sh
go run main.go
```

クライアント側を実行

```sh
go run cmd/client/main.go
```

以下のようなログが出力される

```txt
2024/11/17 22:19:02 connect:  https://127.0.0.1:8080/greet.v1.GreetService/Greet
2024/11/17 22:19:02 send:  {"name": "world"}
2024/11/17 22:19:02 recv:  {"greeting":"Hello, world!"}
```
