# echo-login-app

プログラムの細かい解説はQiitaの[この](https://qiita.com/Shakku/items/2dbf455dc270cc514436)投稿をご覧ください。

追記:コイン,ガチャ,アイテム,シューティングゲーム機能が実装されました。

追記:上記のQiitaの記事は、Signup,Login,メモ機能まで実装した際に作成された記事のため、それ以降の機能については記載されていません。(このREADMEには一部書かれています。)

## ディレクトリ構成

```c
echo-login-app/
├── api
│   ├── controller
│   │   ├── coin_controller.go
│   │   ├── gacha_controller.go
│   │   ├── hasitem_controller.go
│   │   ├── item_controller.go
│   │   ├── memo_controller.go
│   │   ├── status_controller.go
│   │   └── user_controller.go
│   ├── db
│   │   ├── db.go
│   │   └── itemdata.go
│   ├── entity
│   │   └── entity.go
│   ├── server
│   │   └── server.go
│   ├── service
│   │   ├── coin_service.go
│   │   ├── gacha_service.go
│   │   ├── hasitem_service.go
│   │   ├── item_service.go
│   │   ├── memo_service.go
│   │   ├── status_service.go
│   │   └── user_service.go
│   ├── dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── app
│   ├── controller
│   │   ├── app_controller.go
│   │   ├── auth_controller.go
│   │   ├── coin_controller.go
│   │   ├── gacha_controller.go
│   │   ├── item_controller.go
│   │   ├── memo_controller.go
│   │   ├── shotgame_controller.go
│   │   └── user_controller.go
│   ├── entity
│   │   └── entity.go
│   ├── server
│   │   └── server.go
│   ├── service
│   │   ├── coin_service.go
│   │   ├── gacha_service.go
│   │   ├── hasitem_service.go
│   │   ├── item_service.go
│   │   ├── memo_service.go
│   │   ├── status_service.go
│   │   └── user_service.go
│   ├── views
│   │   ├── coin.html
│   │   ├── gachatop.html
│   │   ├── index.html
│   │   ├── itemtop.html
│   │   ├── login.html
│   │   ├── memocreate.html
│   │   ├── memotop.html
│   │   ├── shotgame.html
│   │   ├── signup.html
│   │   ├── status.html
│   │   ├── top.html
│   │   ├── userchangename.html
│   │   ├── userchangepassword.html
│   │   └── userpage.html
│   ├── dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── db
│   └── my.cnf
├── .env
├── .gitignore
├── docker-compose.yml
└── README.md
```

## 起動

.envファイルに以下の環境変数を記載しておく。

```env
ELA_ROOTPASS=dbのrootのpassword
ELA_DATABASE=dbのデータベース名
ELA_USERNAME=dbのユーザー名
ELA_USERPASS=dbのユーザーのパスワード
SESSION_KEY=appのセッション用キー
TOKEN_KEY=apiのToken用キー
```

起動

```docker
docker compose up -d
```

ポート

- api:8081
- app:8082
- db:3307

## api

### ping

`GET /`

### api ユーザー機能

- 全ユーザー取得

`GET /user`

- ユーザー作成

`POST /user`

- IDからユーザー取得

`GET /user/id/:id`

- ユーザー名からユーザー取得

`GET /user/username/:username`

- IDからユーザー更新

`PUT /user/:id`

- IDからユーザー削除

`DELETE /user/:id`

- ログイン

`POST /user/login`

### api メモ機能

- 全メモ取得

`GET /memo`

- メモ作成

`POST /memo`

- IDからメモ取得

`GET /memo/id/:id`

- ユーザーIDからメモ取得

`GET /memo/user_id/:user_id`

- IDからメモ更新

`PUT /memo/:id`

- IDからメモ削除

`DELETE /memo/:id`

### api コイン機能

- 全ユーザーのコイン取得

`GET /coin`

- ユーザーIDからコイン取得

`GET /coin/user_id/:user_id`

- ユーザーIDからコイン更新

`PUT /memo/:user_id`

### api アイテム機能

- 全アイテム取得

`GET /item`

- IDからアイテム取得

`GET /item/:id`

- IDからアイテム削除

`DELETE /item/:id`

### api ガチャ機能

  ガチャ実行と結果取得

`GET /gacha/:times`

### api 取得済みアイテムリスト機能

- 取得済みアイテムに追加

`POST /hasitem/:user_id`

- ユーザーIDから取得済みアイテムリスト取得

`GET /hasitem/user_id/:user_id`

- アイテムIDから取得済みアイテム削除

`DELETE /hasitem/:item_id`

### api シューティングゲーム用ステータス機能

- 全ユーザーのステータス取得

`GET /status`

- ユーザーIDからステータス取得

`GET /status/user_id/:user_id`

- ユーザーIDからステータス更新

`PUT /status/:user_id`

## app

### app ログイン機能

- indexページ

`GET /`

- SignUpページ

`GET /signup`

- Signup

`POST /signup`

- Loginページ

`GET /login`

- Login

`POST /login`

### app ユーザー設定

- Logout

`GET /logout`

- ユーザー名変更ページ

`GET /setting/changename`

- ユーザー名変更

`POST /setting/changename`

- パスワード変更ページ

`GET /setting/changepassword`

- パスワード変更

`POST /setting/changepassword`

- ユーザー削除

`DELETE /setting/delete`

### app ログイン後のアプリ

- topページ

`GET /app`

- ユーザーページ

`GET /app/userpage`

### app メモ機能

- 自分のメモ一覧表示ページ

`GET /app/memo`

- メモ作成ページ

`GET /app/memo/create`

- メモ作成

`POST /app/memo/create`

- メモの中身表示

`GET /app/memo/view/:id`

- メモ削除

`GET /app/memo/delete/:id`

- メモ中身更新

`POST /app/memo/change/:id`

### app コイン機能

- コイントップページ

`GET /app/coin`

- コイン増加

`POST /app/coin/add`

- コイン減少

`POST /app/coin/sub`

### app ガチャ機能

- ガチャトップページ

`GET /app/gacha`

- メモ作成

`POST /app/gacha/draw`

### app アイテム機能

- 所持アイテム一覧表示ページ

`GET /app/item`

### app シューティングゲーム機能

- シューティングゲームページ

`GET /app/game/shot`

- ステータス表示ページ

`GET /app/game/shot/status`

- ステータス強化

`POST /app/game/shot/status`

## entity

### User

```json
{
    "id":12345678,
    "name":"Shakku",
    "password":"$2a$10$nj.KCcTpJH.9bNrVkPho9.dTDlbXq1jyM7I5gEHmmv5Fu4J4Lpvr6",
    "createdat":"2023-05-21T12:34:56+09:00",
}
```

### Memo

```json
{
    "id":1,
    "title":"タイトル",
    "content":"内容",
    "createdat":"2023-05-21T12:34:56+09:00",
    "user_id":12345678,
}
```

### Coin

```json
{
    "id":1,
    "qty":3000,
    "user_id":12345678,
}
```

### レアリティ定数

```go
type Rarity string

const (
    RarityN   Rarity = "N"
    RarityR   Rarity = "R"
    RaritySR  Rarity = "SR"
    RaritySSR Rarity = "SSR"
    RarityUR  Rarity = "UR"
    RarityLR  Rarity = "LR"
)
```

### Item

```json
{
    "id":201,
    "name":"ダメージアップの素材",
    "rarity": "N",
    "raito":1000,
}
```

### HasItem

```json
{
    "items": "Item entityが複数入った配列",
    "user_id":12345678,
}
```

### HasItemList(api)

```json
{
    "id":1,
    "user_id":12345678,
    "item_id":201,
}
```

### ShowItems(app)

```json
{
    "item":"Item entity",
    "qty":12,
}
```

### Status

```json
{
    "id":1,
    "damage":1,
    "hp":10,
    "shotspeed":20,
    "enmcool":100,
    "score":1,
    "user_id":12345678,
}
```

### ResponseMessage

```json
{
    "status":200,
    "Message":"pong",
}
```

### Token

```json
{
    "token":"トークンがここに入ります",
}
```

### 補足

- UserのIDは99999999までのランダムな整数が自動で設定されます。
- UserのPasswordはbcryptによってハッシュ化されて、apiとのやりとりや、データベースへの保存に使用されます。
- UserのGetAllやMemoのGetAll,GetByUserIDなど、複数のデータを取得する際は、リスト形式のJSONがレスポンスとして帰ってきます。
- MemoなどのUser_IDは、そのメモを作成したユーザーのIDがForeignKeyとして保存されます。それにより、そのユーザーが作成したものをUser_IDから全取得できます。

```mysql
| memos | CREATE TABLE `memos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `memos_user_id_users_id_foreign` (`user_id`),
  CONSTRAINT `memos_user_id_users_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci |
```

- コインの枚数を0枚に更新するようDBに命令することができなかったため、CoinのQtyのintをポインタにする必要があった。
- Itemは、apiサーバーでのdb接続確認完了直後にitemdata.goによって、ガチャで排出されるアイテムが登録されている。
- apiサーバーの方では、HasItemList entityを使用したDBのテーブルにUserIDとItemIDを紐付けて登録しており、apiからappにそのユーザーの所持アイテム一覧を送るときは、UserIDからHasItemListのテーブルを検索して全て取得し、HasItem entityのItemsに詰めて送る。
- ShowItemsは、HasItemのItemsの配列を展開して、アイテムごとの個数を保存して表示するためのもの。

## 認証

フロントエンド(HTML)とバックエンド間の認証をSessionで行い、バックエンドとAPI間の認証をJWTで行った。

流れは以下の通り。

- フロントエンドのformにてユーザー名とパスワードを入力する。
- バックエンドでそのformを受け取り、ユーザー名が存在するかの確認を行ったのち、ユーザーIDとパスワードをAPIにPOSTする。
- APIでユーザーIDからユーザー情報を取り出し、POSTされたパスワードとbcrypt.CompareHashAndPasswordを使用してパスワードが一致するか確認する。
- パスワードが問題なければ、APIでTokenを発行し、バックエンドに送る。
- バックエンドで、受け取ったTokenとUserIDなどの入ったSessionを作成し、フロントエンドのCookieに保存する。
- ログイン完了。

appでは、`/app`以下や`/setting`のURLにアクセスしようとすると、Session確認用のミドルウェアが実行され、Sessionの確認を行う。

apiでは、`/memo`などのURLにアクセスしようとすると、middleware.JWTが実行されて、Tokenの確認が行われる。

## Middleware

### APIとバックエンド共通

- 突然のサーバーダウン時のリカバー用

```go
middleware.Recover()
```

- ログ出力のフォーマット化

```go
middleware.LoggerWithConfig()
```

```go
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "\n" + `time: ${time_rfc3339_nano}` + "\n" +
            `method: ${method}` + "\n" +
            `remote_ip: ${remote_ip}` + "\n" +
            `host: ${host}` + "\n" +
            `uri: ${uri}` + "\n" +
            `status: ${status}` + "\n" +
            `error: ${error}` + "\n" +
            `latency: ${latency}(${latency_human})` + "\n",
    }))
```

### api用

- JWTのToken確認

```go
middleware.JWT([]byte("tokenkey"))
```

### app用

- Session用キーの設定

```go
session.Middleware(sessions.NewCookieStore([]byte("sessionkey")))
```

- controller.AuthController.SessionCheck()オリジナルミドルウェア。

```go
auc.SessionCheck
```

Session内の"auth"がTrueかどうかで、このSessionが有効か無効化を判断する。

## HTMLテンプレート

text/templateとechoのRender機能にて、HTMLを使用した。

```go
type TemplateRender struct {
    templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func Init() {

    省略

    renderer := &TemplateRender{
        templates: template.Must(template.ParseGlob("./views/*.html")),
    }
    e.Renderer = renderer

    省略
}
```

```go
// GET indexページ表示
func (uc UserController) Index(c echo.Context) error {
    var us service.UserService

    // ユーザー全取得
    u, err := us.GetAll()
    if err != nil {
        log.Println("us.GetAll error")
    }
    return c.Render(http.StatusOK, "index.html", u)
}
```

```html
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>index</title>
</head>
<body>
    <h1>index</h1>
    {{range $v := .}}
        <div>
            <div>
                <p>{{$v.Name}}</p>
            </div>
        </div>
    {{end}}
    <p><a href="/login">ログイン</a></p>
    <p><a href="/signup">サインアップ</a></p>
    <p><a href="/app">メインページ(ログインが必要です)</a></p>
</body>
</html>
```

## 使用した技術

- Go言語
- Echo
- Docker
- Docker Compose
- mysql
- Token
- JWT
- Session
- Cookie
- bcrypt
- GORM

### 使用したパッケージ

- os
- net/http
- fmt
- log
- time
- strconv
- math/rand
- io
- text/template
- encoding/json
- bytes
- github.com/go-sql-driver/mysql
- github.com/golang-jwt/jwt
- github.com/jinzhu/gorm
- github.com/labstack/echo/v4
- github.com/labstack/echo/v4/middleware
- golang.org/x/crypto
- github.com/gorilla/sessions
- github.com/labstack/echo-contrib
- github.com/labstack/echo-contrib/session

## 参考

- [【gorm】数値のカラムを0で更新する Qiita](https://qiita.com/cpp0302/items/3b1c0ca3adc698a79bc9)
- シューティングゲームのプログラムの引用・参考サイト[javascriptでゲームを作ってみた - ぬるからの雑記帳](https://nullkara.jp/2020/10/25/javascriptmakegame001/)
- ガチャのアルゴリズム参考[Golangで重み付き乱択アルゴリズムを作成したので検証してみる Zenn](https://zenn.dev/koupro0204/articles/bf87dea2478b72)

## めも

サーバー性の計測
ab -n 100 -c 100 localhost:8081/
-n リクエスト総数 -c 同時接続数

middleware.RemoveTrailingSlash()：こいつはGroupを使っているルータに使おうとすると、下の階層のGroupには反映されなかった。(なので実装見送り)

コインが0枚にならない(gormの性質上、ゼロ値がくると、Updateをスキップしてしまう。そのため、Qtyのintをポインタの*intに変更するとできる。[参考](https://qiita.com/cpp0302/items/3b1c0ca3adc698a79bc9))

GORMのv2では、db.create()に構造体のリストを与えて一気にインサートすることができるが、v2にすると色々変わって対応が大変だったため、初期データを入れるところは一つ一つ入れている。

スロットゲームシューティングゲーム

シューティングゲーム[プログラムの引用・参考サイト](https://nullkara.jp/2020/10/25/javascriptmakegame001/)

アタッチメント(斜め方向、後ろ方向、横方向)とパワーアップ用素材が同時に出るガチャ(ダメージup,連写速度up,HPup,スコアアップ,敵レベルアップ間隔down)

アタッチメントの方は、既に持ってたらhasitemlistに追加しない、素材はそのまま追加していく。

そのリストをforで展開して、同じアイテムは加算して行って、個数表示

銃はめっちゃ貴重で所持未所持をどこかで保持、素材はスタック(素材でやすいから上がり幅低めで上限あり。上限いったら何か報酬？)

スコアによってコイン獲得

mysqlは配列型を扱えないことを今更知った。many2manyっていうやつを使ってやろうとしたが、結合テーブルで重複ができなかった。なので、別で作ることに。
