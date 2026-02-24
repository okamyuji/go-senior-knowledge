# Go Senior Knowledge

Goの上級者向けトピックを38のテーマに分けて整理した学習用リポジトリです。
各ディレクトリにはGoソースコード、テスト、日本語のREADMEが含まれています。

すべてのコードは標準ライブラリのみを使用しており、外部依存はありません。

## トピック一覧

| # | ディレクトリ | テーマ |
|---|---|---|
| 01 | [01_memory_model](./01_memory_model/) | Goのメモリモデル |
| 02 | [02_happens_before](./02_happens_before/) | happens-before関係 |
| 03 | [03_data_races](./03_data_races/) | データ競合 |
| 04 | [04_goroutine_scheduling](./04_goroutine_scheduling/) | ゴルーチンスケジューリング |
| 05 | [05_preemption](./05_preemption/) | プリエンプション |
| 06 | [06_channel_internals](./06_channel_internals/) | チャネルの内部動作 |
| 07 | [07_select_behavior](./07_select_behavior/) | select文の動作 |
| 08 | [08_closing_channels](./08_closing_channels/) | チャネルのクローズ |
| 09 | [09_nil_channels](./09_nil_channels/) | nilチャネルの活用 |
| 10 | [10_context_cancellation](./10_context_cancellation/) | contextによるキャンセル制御 |
| 11 | [11_sync_once](./11_sync_once/) | sync.Onceによる一度きりの初期化 |
| 12 | [12_sync_pool](./12_sync_pool/) | sync.Poolによる一時オブジェクトの再利用 |
| 13 | [13_escape_analysis](./13_escape_analysis/) | エスケープ解析とスタック/ヒープ割り当て |
| 14 | [14_pointer_vs_value](./14_pointer_vs_value/) | ポインタレシーバと値レシーバ |
| 15 | [15_interfaces_internals](./15_interfaces_internals/) | インターフェースの内部構造 |
| 16 | [16_empty_interface](./16_empty_interface/) | 空インターフェース(any)と型アサーション |
| 17 | [17_defer_internals](./17_defer_internals/) | deferの内部動作 |
| 18 | [18_panic_recover](./18_panic_recover/) | panicとrecoverの仕組み |
| 19 | [19_map_internals](./19_map_internals/) | mapの内部動作 |
| 20 | [20_map_concurrency](./20_map_concurrency/) | mapの並行アクセス |
| 21 | [21_slice_internals](./21_slice_internals/) | スライスの内部構造 |
| 22 | [22_string_bytes](./22_string_bytes/) | 文字列とバイトスライスの関係 |
| 23 | [23_gc_basics](./23_gc_basics/) | GCの基本 |
| 24 | [24_write_barrier](./24_write_barrier/) | ライトバリア |
| 25 | [25_gomaxprocs](./25_gomaxprocs/) | GOMAXPROCS |
| 26 | [26_atomic_vs_mutex_vs_channel](./26_atomic_vs_mutex_vs_channel/) | atomic, Mutex, channelの比較 |
| 27 | [27_memory_alignment](./27_memory_alignment/) | メモリアラインメントとフォールスシェアリング |
| 28 | [28_pprof_basics](./28_pprof_basics/) | pprofによるプロファイリング基礎 |
| 29 | [29_cgo_overhead](./29_cgo_overhead/) | cgoのオーバーヘッド |
| 30 | [30_http_transport](./30_http_transport/) | http.Transportのカスタマイズ |
| 31 | [31_http_client_reuse](./31_http_client_reuse/) | http.Clientの再利用 |
| 32 | [32_json_pitfalls](./32_json_pitfalls/) | JSON処理の落とし穴 |
| 33 | [33_generics](./33_generics/) | ジェネリクス |
| 34 | [34_init_order](./34_init_order/) | パッケージ初期化順序 |
| 35 | [35_build_tags](./35_build_tags/) | ビルドタグ |
| 36 | [36_embedding](./36_embedding/) | 構造体埋め込み |
| 37 | [37_error_handling](./37_error_handling/) | エラーハンドリング |
| 38 | [38_race_detector](./38_race_detector/) | レースディテクタ |

## 実行方法

全テストを実行する場合は以下のコマンドを使用します。

```bash
go test ./...
```

静的解析を実行するには以下のコマンドを使用します。

```bash
go vet ./...
staticcheck ./...
golangci-lint run ./...
```

## 前提条件

- Go 1.23以上
- staticcheck (任意)
- golangci-lint (任意)
