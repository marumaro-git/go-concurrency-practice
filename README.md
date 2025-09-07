# はじめに
Goの並行処理を学ぶためのサンプルコード

# 並行と並列

## このセクションは、https://go.dev/blog/waza-talk の内容を抜粋したものです。
並行性（Concurrency）と並列性（Parallelism）は似て非なる概念です。Go公式サイトでも以下のように定義されています：

- **並行性（Concurrency）**：独立して実行できる「プロセス（あるいは処理）」を構成することや設計上それらを同時に扱えるようにすること。
- **並列性（Parallelism）**：実際に複数の計算が物理的に同時に実行されること。

> "Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once."
（並行性は「多数のことを一度に扱うこと」、並列性は「多数のことを一度に実行すること」）

Goにおける並行処理の特徴これらは混同されがちですが、並行性は「設計上の概念」であり、並列性は「実行上の現象」である点が異なります。

Go言語は、並行処理を簡潔かつ効率的に実現するための機能を提供しています。具体的には以下のような特徴があります：

- **goroutines**: 軽量なスレッドのようなもので、`go`キーワードを使って簡単に並行処理を開始できます。
- **チャネル（channels）**: goroutines間でデータを安全にやり取りするための仕組みです。
- **select文**: 複数のチャネル操作を待機し、どれか一つが準備できたらそれを実行する構文です。

これらの機能により、Goは「並行性」を設計の中心に据えています。しかし、Goが提供するのはあくまで並行処理のための構造であり、実際にそれが並列に実行されるかどうかは、CPUコア数やランタイムのスケジューリングに依存します。

このように、Goは「並行性」と「並列性」の違いを明確にし、開発者が並行処理を直感的に記述できるように設計されています。

# 参考
- [Goでの並行処理を徹底解剖！
](https://zenn.dev/hsaki/books/golang-concurrency)
- [Go by Example: Worker Pools](https://oohira.github.io/gobyexample-jp/worker-pools.html)
- [pond](https://github.com/alitto/pond)
- [手を動かしながら並行処理を学ぼう](https://zenn.dev/knowledgework/articles/9b9abd12e7c621)
