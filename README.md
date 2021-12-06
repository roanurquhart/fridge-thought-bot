# fridge-thought-bot
source /Users/speaker/Documents/fridge-thought-bot/fridge-image/.venv/bin/activate
go run . # In fridge-bot-go dir
set BOT_TOKEN $TOKEN
FLASK_APP=image_rest_api.py FLASK_ENV=development flask run