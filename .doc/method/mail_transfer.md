## 要件
* ユーザーはREST-APIのURIを叩いて問い合わせ内容を登録する
    * ユーザー自体は単純な問い合わせの動きを体感する
* ユーザーからの問い合わせをサーバーが受け取ると、サーバーは問い合わせを管理するメールアドレス（以下問い合わせ管理アドレス）宛にメールを送信する
* 問い合わせ管理アドレス上でのメールと、APIからの問い合わせ確認の両方が利用できる

## 仕様
* ユーザーは送信する宛先メールアドレス(複数可)、タイトル、内容をパラメタに指定する
* app.iniで設定しているHostとPortを基にメールを送信する
    * 開発環境上ではMailHogを利用
* メールの送信が行われる