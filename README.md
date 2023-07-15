# echo-login-app

Create実装
ログイン後の機能実装

{"status":500,"message":"UserService.Create: Error 1062 (23000): Duplicate entry '12345678' for key 'users.PRIMARY'"}

既に登録されているユーザー名じゃないか判断するときに、全てのユーザー名をGETして、バックエンドで確認するのと、API側で確認して、エラーのレスポンスを返すのと、どっちの方が効率がいいのか。

セッション管理はIDで行って、ユーザー名変更の時は、そのIDを使ってユーザー名変更する。

ユーザー名存在しない時
