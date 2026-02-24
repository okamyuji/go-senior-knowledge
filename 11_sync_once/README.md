# sync.Onceによる一度きりの初期化

## 概要

sync.Onceは、複数のgoroutineから同時に呼ばれても、指定した関数を正確に1回だけ実行することを保証します。共有リソースの遅延初期化や、コストの高い計算結果のキャッシュに利用します。

## sync.Onceの基本動作

sync.OnceのDoメソッドに渡した関数は、何度Doを呼んでも最初の1回しか実行されません。すべての呼び出し元は初期化が完了するまでブロックされ、同じ結果を受け取ります。

```go
var once sync.Once
var resource *Resource

once.Do(func() {
    resource = createExpensiveResource()
})
```

## OnceValueパターン

高コストな計算を1回だけ実行し、結果をキャッシュするクロージャを返すパターンがあります。Go 1.21ではsync.OnceValueとして標準ライブラリに追加されました。このサンプルでは手動実装を示しています。

```go
getter := OnceValue(func() string {
    return expensiveComputation()
})
// 何度呼んでも計算は1回だけ実行されます
result := getter()
```

## デッドロックの注意点

Doの内部から同じsync.OnceのDoを再帰的に呼び出すと、デッドロックが発生します。sync.Onceは内部でMutexを使用しており、同じgoroutineからの再入はロックの取得を永久に待ち続けます。

## ファイル構成

- sync_once.go: ResourceLoaderによるInitOnceパターンと、OnceValueパターンを実装しています
- sync_once_test.go: 100個のgoroutineで同時にInitOnceを呼び出し、初期化が1回だけ実行されることを検証しています
