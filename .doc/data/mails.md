**mails**

メールテーブル。ユーザーのメール情報を管理する。

**テーブル：mails**
| 物理名       |  主   | 型       | 必須  | 桁数  | 一意  |
| ------------ | :---: | -------- | :---: | :---: | :---: |
| id           |  〇   | int      |  〇   |       |  〇   |
| mail_address |       | varchar  |  〇   |  100  |  〇   |
| name         |       | varchar  |  〇   |  50   |       |
| created_at   |       | datetime |       |       |       |
| updated_at   |       | datetime |       |       |       |
| deleted_at   |       | datetime |       |       |       |