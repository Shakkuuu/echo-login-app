# echo-login-app

Create実装
ログイン後の機能実装

{"status":500,"message":"UserService.Create: Error 1062 (23000): Duplicate entry '12345678' for key 'users.PRIMARY'"}

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

curl -X GET localhost:8081/
curl -X POST -H "Content-Type: application/json" -d '{"":"", }' localhost:8081/
curl -X DELETE localhost:8081/
curl -X PUT -H "Content-Type: application/json" -d '{"":"", }' localhost:8081/

既に登録されているユーザー名じゃないか判断するときに、全てのユーザー名をGETして、バックエンドで確認するのと、API側で確認して、エラーのレスポンスを返すのと、どっちの方が効率がいいのか。

サーバー性の計測
ab -n 100 -c 100 localhost:8081/
-n リクエスト総数 -c 同時接続数

ユーザー取得、ユーザー作成、名前やIDからユーザー取得はフリー
名前とパスワード変更、今後実装のアプリ系はToken必要にしたい

フロントのブラウザ(HTML)とバックエンド間はセッションで管理
バックエンドとAPIはTokenで管理

interface型の知識が浅かった
