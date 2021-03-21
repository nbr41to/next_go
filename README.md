1. フロントエンド
   cd client->yarn->yarn dev
2. docker 起動
   cd client->yarn docker-up
3. go のコンテナに入り、main.go を実行
   cd client->yarn docker-go->go run main.go
   データベースは postgres
   user->root
   password->root
   dbname->next_go
   port->5555
   client->フロントエンド(Next.js)
   server->サーバーサイド(Go)
   docker->Docker の設定ファイル
