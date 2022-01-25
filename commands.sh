cd /code/fridge_bot
nohup go run . </dev/null &>/dev/null &
export BOT_TOKEN $TOKEN
FLASK_APP=image_rest_api.py FLASK_ENV=development flask run &
cd /code/fridge_bot/ \
	&& go run . 
