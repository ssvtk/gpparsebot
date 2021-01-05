# gpparsebot
Telegram bot for parsing commment section of popular clothing web <a href="https://groupprice.ru/brands/dstrend/comments">store</a>.
Bot gets the element with HTTP requests with set timing and compares it with hashed data in DB, allerting telegram if it is new data (HTTP). 
API-key has to be replaced <a href="https://github.com/ssvtk/gpparstel/blob/8956229684689a5ad10f11c384e053f20273a72f/telegram/telegram.go#L10">telegram.go</a> <br>
<a href="https://github.com/ssvtk/gpparstel/blob/main/config.json">Config</a> has to be set accordingly
<br>

PostgreSQL (github.com/jackc/pgx)
<br>
Telegram API (github.com/go-telegram-bot-api/telegram-bot-api)
