**go-echo-mail**

Golang + Echo + SendMail(MailHog)

## はじめに
本プロジェクトはGolangとEchoを利用したメール送信機能です。MailHogを利用してメール送信のテストを提供します。

## 構成
```
├── main.go
├── common
|    ├── db.go
|    └── response.go
├── config
|    ├── app.ini
|    ├── config.go
|    └── const.go
├── api
| 　　├── routers.go
| 　　├── support
| 　　|    └── mail_deliver.go
| 　　├── validator
| 　　|    └── validator.go
| 　　└── handlers
|          └── mail_handlers.go
└── domain
      ├── historys
      |      ├── model.go
      |      ├── stores.go
      |      └── validations.go
      ├── mails
      |      ├── model.go
      |      ├── stores.go
      |      └── validations.go
      └── setup_data.go
```
## ドキュメント
### データ定義
| テーブル名                                                                             | 概要                                             |
| -------------------------------------------------------------------------------------- | ------------------------------------------------ |
| [mails](https://github.com/chalmeal/go-echo-mail/blob/master/.doc/data/mails.md)       | メールテーブル。ユーザーのメール情報を管理する。 |
| [historys](https://github.com/chalmeal/go-echo-mail/blob/master/.doc/data/historys.md) | メール送信履歴テーブル。送信したメールの履歴。   |

### 仕様
| 書名           | 概要                   |
| -------------- | ---------------------- |
| [メール送信]() | メール送信機能について |

## セットアップ

### DB
* DBの環境は以下を想定します。
  * MySQL
  * GORM
* create schemaのみ行う必要があります。[DDL](.db/setup/ddl-create-chema.sql)
* [app.ini](config/app.ini)に対してDB接続情報を定義してください。
* テーブルはGORMが提供するAutoMigrateを利用します。
  * 各テーブルは初回API実行時に生成されます。
  * 本プロジェクトで定義されている詳細なデータ定義に関しては、各ドメインのdocを参照してください。
    * メール：[mails](https://github.com/chalmeal/go-echo-mail/blob/master/.doc/data/mails.md)
    * 履歴：[historys](https://github.com/chalmeal/go-echo-mail/blob/master/.doc/data/historys.md)

### Docker
MailHogの利用にDockerを利用しています。Dockerの導入を行ってください。

## アプリケーションスタート
アプリケーションのスタートはデバッガを推奨しています。

Run and Debugの`Run go-echo-mail`から実行してください。
