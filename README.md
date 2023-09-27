# echo-login-app

コードの細かい解説はQiitaの[この](https://qiita.com/Shakku/items/2dbf455dc270cc514436)投稿をご覧ください。

## ディレクトリ構成

```c
echo-login-app/
├── api
│   ├── controller
│   │   ├── memo_controller.go
│   │   └── user_controller.go
│   ├── db
│   │   └── db.go
│   ├── entity
│   │   └── entity.go
│   ├── server
│   │   └── server.go
│   ├── service
│   │   ├── memo_service.go
│   │   └── user_service.go
│   ├── dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── app
│   ├── controller
│   │   ├── app_controller.go
│   │   ├── auth_controller.go
│   │   ├── memo_controller.go
│   │   └── user_controller.go
│   ├── entity
│   │   └── entity.go
│   ├── server
│   │   └── server.go
│   ├── service
│   │   ├── memo_service.go
│   │   └── user_service.go
│   ├── views
│   │   ├── index.html
│   │   ├── login.html
│   │   ├── memocreate.html
│   │   ├── memotop.html
│   │   ├── signup.html
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

`PUT /memo:id`

- IDからメモ削除

`DELETE /memo/:id`

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
- MemoのUser_IDは、そのメモを作成したユーザーのIDがForeignKeyとして保存されます。それにより、そのユーザーが作成したメモをUser_IDから全取得できます。

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

apiでは、`/memo`のURLにアクセスしようとすると、middleware.JWTが実行されて、Tokenの確認が行われる。

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

## めも

サーバー性の計測
ab -n 100 -c 100 localhost:8081/
-n リクエスト総数 -c 同時接続数

middleware.RemoveTrailingSlash()：こいつはGroupを使っているルータに使おうとすると、下の階層のGroupには反映されなかった。(なので実装見送り)

コインが0枚にならない(gormの性質上、ゼロ値がくると、Updateをスキップしてしまう。そのため、Qtyのintをポインタの*intに変更するとできる。[](https://qiita.com/cpp0302/items/3b1c0ca3adc698a79bc9))
