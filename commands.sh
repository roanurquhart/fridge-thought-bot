nohup go run . </dev/null &>/dev/null &
export BOT_TOKEN $TOKEN
FLASK_APP=image_rest_api.py FLASK_ENV=development flask run &
d /code/fridge_bot/ \
	&& go run . 
