# GOMAXPROCS

GOMAXPROCSは、ユーザレベルのGoコードを同時に実行できるOSスレッドの数を制御します。
デフォルトでは利用可能な論理CPU数に設定されており、runtime.GOMAXPROCS関数で変更できます。

# GOMAXPROCSの取得

GetMaxProcs関数は、現在のGOMAXPROCSの値を返します。
runtime.GOMAXPROCSに0を渡すと、値を変更せずに現在の設定値だけを取得できます。

# GOMAXPROCSの一時的な変更

SetMaxProcs関数は、GOMAXPROCSを指定した値に変更し、関数を実行してから元の値に復元します。
deferを使って復元を保証しているため、渡された関数がパニックしても元の値に戻ります。
この関数は変更前の元のGOMAXPROCSの値を返します。

# 並列処理のデモンストレーション

ParallelismDemo関数は、指定したGOMAXPROCS値で複数のゴルーチンによるCPUバウンドな処理を実行します。
GOMAXPROCSが1の場合、ゴルーチンは逐次的に実行されるため合計の実行時間が長くなります。
GOMAXPROCSを増やすと、複数のCPUコアで並列に実行されるため壁時計時間が短縮されます。

# GOMAXPROCSの注意点

GOMAXPROCSはユーザのGoコードを実行するスレッド数を制限しますが、ランタイムはシステムコールやGCなどに追加のスレッドを使用します。
GOMAXPROCSを1に設定しても、複数のゴルーチンが協調的にスケジューリングされるため並行処理は可能です。
ただし、真の並列実行にはGOMAXPROCSを2以上に設定する必要があります。

# テストの実行方法

以下のコマンドでテストを実行できます。

```
go test -v ./25_gomaxprocs/...
```
