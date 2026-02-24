# panicとrecoverの仕組み

Goにはtry-catchのような例外機構がなく、代わりにpanicとrecoverの組み合わせで致命的なエラーからの回復を行います。
panicが発生すると現在の関数の実行が停止し、deferされた関数が順に実行された後、呼び出し元に伝播していきます。

# SafeExecuteによる安全な関数実行

SafeExecute関数は任意の関数を受け取り、deferされた関数内でrecoverを呼び出してpanicを捕捉します。
panicが発生した場合はその値をerrorに変換して返し、正常に完了した場合はnilを返します。
ライブラリやサーバーのハンドラなど、panic伝播を防ぎたい場面で活用できます。

# recoverが機能する条件

recoverはdeferされた関数の中で直接呼び出した場合にのみ機能します。
RecoverInDefer関数では、defer func()の中でrecoverを呼んでいるためpanicを正しく捕捉できます。

RecoverOutsideDefer関数では、deferの外でrecoverを呼んでいるため、常にnilを返してpanicは捕捉されません。
この場合、panicはそのまま伝播してプログラムがクラッシュします。

# panic値のerrorへの変換

PanicToError関数は、recoverで取得したpanic値を型スイッチで分類してerrorに変換します。
panic値がerror型の場合はそのまま返し、string型の場合はfmt.Errorfでラップし、それ以外の型はfmt.Sprintfで文字列化します。

# 別のゴルーチンのpanicについて

recoverは同一ゴルーチン内のpanicのみを捕捉できます。
別のゴルーチンで発生したpanicはそのゴルーチンのスタックを巻き戻すだけで、他のゴルーチンからrecoverで捕捉することはできません。
ゴルーチン内でpanicが発生しrecoverされなかった場合、プログラム全体がクラッシュします。

# テストの実行方法

以下のコマンドでテストを実行できます。

```
go test -v ./18_panic_recover/...
```
