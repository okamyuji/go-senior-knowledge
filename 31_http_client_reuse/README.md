# http.Clientの再利用

## 概要

Goのnet/httpパッケージでHTTPリクエストを送信する際、http.Clientの使い方がパフォーマンスとリソース管理に大きな影響を与えます。このパッケージでは、http.Clientをパッケージレベルで共有する正しいパターンと、リクエストごとにクライアントを生成するアンチパターンを比較します。

## http.Clientを再利用すべき理由

http.Clientは内部にhttp.Transportを保持しており、このTransportがTCPコネクションプールを管理しています。同じクライアントを使い回すことで、TCPコネクションが再利用され、接続のオーバーヘッドが削減されます。

リクエストごとに新しいhttp.Clientを生成すると、毎回新しいTransportとコネクションプールが作られます。使い終わったクライアントのコネクションはプールに戻されず、アイドルタイムアウトまでファイルディスクリプタとgoroutineを占有し続けます。これがコネクションリークの原因になります。

## SharedClientパターン

パッケージレベルでhttp.Clientを定義し、Timeoutを明示的に設定します。デフォルトのhttp.DefaultClientにはタイムアウトが設定されていないため、本番コードでは必ず独自のクライアントを用意するべきです。

```go
var SharedClient = &http.Client{
    Timeout: 10 * time.Second,
}
```

## レスポンスボディのクローズ

HTTPレスポンスを受け取った後は、必ずresp.Body.Close()をdeferで呼び出す必要があります。ボディをクローズしないと、TCPコネクションがプールに返却されず、コネクションリークが発生します。

```go
resp, err := client.Get(url)
if err != nil {
    return "", err
}
defer resp.Body.Close()
```

エラーチェックの直後にdeferを置くのが定番のパターンです。

## アンチパターン

以下のようにリクエストごとに新しいクライアントを生成するのは避けてください。

```go
func BadPattern(url string) (string, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(url)
    // ...
}
```

この書き方では、呼び出しのたびにTransportが新規作成されます。高負荷環境ではgoroutineリークやファイルディスクリプタの枯渇につながります。

## テストでの扱い

httptest.NewServerを使うと、テスト用のHTTPサーバーを簡単に立ち上げることができます。テストサーバーのClientメソッドが返すhttp.Clientは、そのサーバー専用のTransportが設定されているため、テスト間の干渉を防ぐことができます。
