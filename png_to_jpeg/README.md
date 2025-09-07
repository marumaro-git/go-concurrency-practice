#　概要
このプログラムは、指定されたディレクトリ内のPNG画像をJPEG形式に変換するGo言語のサンプルコード。

参考：[手を動かしながら並行処理を学ぼう](https://zenn.dev/knowledgework/articles/9b9abd12e7c621)


## step01 ~ step05
- step01: シングルスレッドで順次変換
- step02: ゴルーチンを使用して並列変換
- step03: errgroupを使用して並行処理。最初に発生したエラーを取得し、各goroutineを終了。
- step04: conc/poolを使用して並行処理。全てのエラーを取得し、各goroutineの処理は止めない。
- step05: goroutineでpanicが発生した場合の処理を追加

# 実行
以下コマンドで実行
```bash
go run main.go <step> <ディレクトリ> 
# 例: step01を実行
go run main.go step01 testdata/*.png
```


## トレースの確認
実行するとtraceのファイルが生成されます。

生成されたファイルを以下のコマンドで確認できます。

```bash
go tool trace <ファイル名>>
```