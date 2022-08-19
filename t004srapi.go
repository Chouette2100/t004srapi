/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

*/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dustin/go-humanize"

	"github.com/Chouette2100/exsrapi"
	"github.com/Chouette2100/srapi"
)

/*
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

	$ curl -OL https://github.com/Chouette2100/t000srapi/archive/refs/tags/vn.n.n.tar.gz
	$ tar xvf v1.0.0.tar.gz
	$ mv t004srapi-1.0.0 t000srapi
	$ cd t004srapi

		以上4行は、Githubからソースをダウンロードしてます。vn.n.nのところは、ソースのバージョンを指定します。
		バージョンは、Githubのリリースページで確認してください。
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

	Ver. 0.0.0

*/

type Config struct {
	SR_acct string
	Roomid  []int
}

//	イベントに参加しているルームのルームIDを指定して、獲得ポイントなどを取得する。
func main() {

	//	ログファイルを設定する。
	logfile := exsrapi.CreateLogfile("", fmt.Sprintf("%d", os.Getpid()))
	defer logfile.Close()

	if len(os.Args) != 2 {
		//      引数が足りない(設定ファイル名がない)
		log.Printf("usage:  %s NameOfConfigFile\n", os.Args[0])
		return
	}

	//	設定ファイルを読み込む。設定ファイルには各レベルを達成するのに必要な視聴時間、ポイント、コメント数を書いてある。
	var config Config
	err := exsrapi.LoadConfig(os.Args[1], &config)
	if err != nil {
		log.Printf("LoadConfig: %s\n", err.Error())
		return
	}

	//	cookiejarがセットされたHTTPクライアントを作る
	client, jar, err := exsrapi.CreateNewClient(config.SR_acct)
	if err != nil {
		log.Printf("CreateNewClient() returned error %s\n", err.Error())
		return
	}
	//	すべての処理が終了したらcookiejarを保存する。
	defer jar.Save()


	log.Printf("roomid rank      points         gap eventname")
	for _, roomid := range config.Roomid {
		//	ルームIDを指定して、ルームのポイントを取得する。
    	point, rank , gap, _, _, eventname, err := srapi.GetPointByApi(client, roomid)
		if err != nil {
			log.Printf("GetPointByApi(): %s\n", err.Error())
			return
		}
		//	ルームのポイントを表示する。
		log.Printf("%-7d%4d%12s%12s %s\n", roomid, rank, humanize.Comma(int64(point)), humanize.Comma(int64(gap)), eventname)
	}
}
