# contextによるキャンセル制御

このパッケージでは、Goのcontextパッケージを使ったキャンセル、タイムアウト、伝播の各パターンを示しています。

## context.WithCancel

WithCancelDemoは、キャンセル可能なcontextを作成し、goroutineがctx.Done()チャネルを監視してキャンセルを検知する基本パターンを実演しています。cancel関数を呼ぶと、ctx.Done()が閉じられ、ctx.Err()はcontext.Canceledを返します。

## context.WithTimeout

WithTimeoutDemoは、指定した時間が経過すると自動的にキャンセルされるcontextを実演しています。タイムアウト後、ctx.Err()はcontext.DeadlineExceededを返します。APIコールやデータベースクエリに制限時間を設ける場面で広く使われています。

## キャンセルの伝播

PropagateCancelは、親contextがキャンセルされると子contextにも自動的に伝播する仕組みを実演しています。contextはツリー構造を形成し、上位のキャンセルはすべての下位contextに波及します。ただし、子のキャンセルは親に影響を与えません。

## defer cancel()パターン

CleanupPatternは、context作成直後にdefer cancel()を呼ぶイディオムを実演しています。このパターンにより、関数がどの経路で終了してもcontextに紐づくリソースが確実に解放されます。go vetはcancel関数の呼び忘れを検出するため、常にdeferで呼ぶことが推奨されています。
