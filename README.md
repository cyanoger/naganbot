# naganbot
Shoot yourself in Russian roulette

## Configuring and startup
Copy the `.env.dist` file as `.env` and set the [environment variables](#Environment) according to your requirements.

To build and run the application, use the `make run` command, and for more commands `make help`.

## Environment
### .env
| Var                    | Description                                                                    |
|------------------------|--------------------------------------------------------------------------------|
| APP_DEBUG              | Enables display of additional log messages. Available values: `true`, `false`. |
| APP_TELEGRAM_BOT_TOKEN | Telegram Bot API access token.                                                 |
