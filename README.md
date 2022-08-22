# t004srapi

	指定したルームのイベントでの順位、ポイント等を取得尻API /api/room/event_and_support のテストのためのサンプル

	以下簡単に説明しますが、詳細は以下の記事を参照してください。

		【Windows】Githubにあるサンプルプログラムの実行方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/e27fc9

		【Unix/Linux】Githubにあるサンプルプログラムの実行方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/220e38

		【Windows】SHOWROOMのAPI関連パッケージ部分を含めたビルドの方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/fe982a

			（ロードモジュールさえできればいいということでしたらコマンド一つでできます）

	$ cd ~/go/src

	$ curl -OL https://github.com/Chouette2100/t000srapi/archive/refs/tags/v0.0.0.tar.gz
	$ tar xvf v0.0.0.tar.gz
	$ mv t004srapi-0.0.0 t000srapi
	$ cd t004srapi

		以上4行は、Githubからソースをダウンロードしてます。v0.0.0のところは、ソースのバージョンを指定します。
		バージョンは、Githubのリリースページで確認してください（当面 v0.0.0 のままにしておくつもりですが）
		ダウンロードはどんな方法でも構わなくて、極端な話Githubのソースをコピペでエディターに貼り付けてもOKです。
		詳細は上に紹介した三つの記事にあります。

	$ go mod init
	$ go mod tidy
	$ go build t000srapi.go
	$ cat config.yml
	sr_acct: ${SRACCT}
	roomid:
	- 00000000					<== ここにルームIDを羅列します。ルームIDはプロフィールやファンルームのURLの最後にある数字です。
	- 11111111
	- 22222222
	$ export SRACCT=xxxxxxxxx	<== これはCookiejarを保存するファイル名に使用します。ログインは不要なので任意の文字列でもかまいません。
	$ ./t004srapi config.yml

	Windowsでの実行については run.bat を参考にしてください。

	go mod init	で不具合があるときは go mod init t004srapi.go を試してください（ソースの位置（ディレクトリ、ディレクトリ構成）を検討する）
