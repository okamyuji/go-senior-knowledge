# sync.Poolによる一時オブジェクトの再利用

## 概要

sync.Poolは、一時的なオブジェクトを再利用するための仕組みです。頻繁に確保と解放を繰り返すオブジェクトをプールしておくことで、GCの負荷を軽減し、アロケーション回数を減らすことができます。

## BufferPoolの使い方

bytes.Bufferをsync.Poolで管理する典型的なパターンを示しています。GetBufferでプールからバッファを取得し、使い終わったらPutBufferで返却します。

```go
buf := GetBuffer()
defer PutBuffer(buf)
buf.WriteString("data")
result := buf.String()
```

## New関数の役割

sync.PoolのNewフィールドには、プールが空のときに呼ばれる関数を設定します。プールにオブジェクトが残っていれば、Newは呼ばれずに既存のオブジェクトが返されます。

## sync.Poolはコネクションプールではありません

sync.Poolに格納されたオブジェクトは、GCのタイミングで通知なく削除される可能性があります。データベース接続やファイルハンドルなど、ライフサイクルの管理が必要なオブジェクトには使用しないでください。バイトバッファや一時的な構造体など、短命なオブジェクトの再利用に適しています。

## ファイル構成

- sync_pool.go: BufferPoolの定義、GetAndPutパターン、PoolWithNewによるNew関数の動作確認を実装しています
- sync_pool_test.go: プールからの取得と返却、New関数の呼び出し回数、並行アクセスの安全性を検証しています
