# echo-login-app

Create実装
ログイン後の機能実装

{"status":500,"message":"UserService.Create: Error 1062 (23000): Duplicate entry '12345678' for key 'users.PRIMARY'"}

既に登録されているユーザー名じゃないか判断するときに、全てのユーザー名をGETして、バックエンドで確認するのと、API側で確認して、エラーのレスポンスを返すのと、どっちの方が効率がいいのか。

サーバー性の計測
ab -n 100 -c 100 localhost:8081/
-n リクエスト総数 -c 同時接続数

ユーザー取得、ユーザー作成、名前やIDからユーザー取得はフリー
名前とパスワード変更、今後実装のアプリ系はToken必要にしたい

フロントのブラウザ(HTML)とバックエンド間はセッションで管理
バックエンドとAPIはTokenで管理

interface型の知識が浅かった
