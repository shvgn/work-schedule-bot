# Work Schedule Bot

This is a telegram bot aimed to fill time in work schedule spreadsheet 
in Google Docs. This is specific solution that saved my time and calmness.

## Deployment

My name is hardcoded. Not very 12-factor, but I had no reason to parametrize it yet.

The bot expects `SPREADSHEET_ID` in env. 

Also, the bot expects the following secrets in files:

```
secrets/
├── credentials.json 	# google account credentials to access Spreadsheet API
└── telegram_token		# text file with the token of your telegram bot
```
