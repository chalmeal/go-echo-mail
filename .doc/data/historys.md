**historys**

メール送信履歴テーブル。送信したメールの履歴。

**historys**
| 物理名     |  主   | 型       | 必須  | 桁数  | 一意  |
| ---------- | :---: | -------- | :---: | :---: | :---: |
| id         |  〇   | int      |  〇   |       |  〇   |
| mail_id    |       | int      |  〇   |       |       |
| title      |       | varchar  |  〇   |  50   |       |
| detail     |       | varchar  |  〇   | 1000  |       |
| created_at |       | datetime |       |       |       |
| updated_at |       | datetime |       |       |       |
| deleted_at |       | datetime |       |       |       |