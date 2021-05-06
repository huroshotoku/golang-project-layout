# huroshotoku/golang-project-layout

本リポジトリは [golang-standards/project-layout](https://github.com/golang-standards/project-layout) のプロジェクトレイアウトをベースに、DDDを実践するためのパッケージレイアウト、およびサンプルコードを追加した。

サンプルコードは書籍: [ドメイン駆動設計入門](https://www.amazon.co.jp/dp/479815072X) から引用。

## Directories

オリジナルのREADMEは[これ]()で、基本的にはこれに従う（余分なものは適宜削れば良い）。

`/internal/app`ディレクトリ以下にDDDを実践したサンプルコードが追加されている。

```
$ tree -L 3 -d internal/app/
internal/app/
├── application         # アプリケーションサービス(WIP)
│   ├── circles         # Aggregateごとに分けているが、ユースケースで分ける方が適切かも
│   └── users
├── domain
│   ├── models
│   │   ├── circles     # circleドメインに関するコード
│   │   └── users         - `DomainModel`や`DomainService`、及び`repository`等のインターフェースを置く
│   ├── services        # 複数ドメインにまたがるサービスの場合、ここに置く
│   └── shared
└── infrastructure
    ├── mysql           # 技術基盤ごとにパッケージを分ける
    │   ├── circles     # Aggregateごとに`repository`や`factory`の実装を用意する
    │   └── users
    └── pub
```

### `domain`

WIP.

### `application`

WIP.

### `infrastructure`

WIP.
