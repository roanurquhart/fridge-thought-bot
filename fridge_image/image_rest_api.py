from flask import Flask
from flask import request
from flask import jsonify
import image_generator

app = Flask(__name__)
@app.route('/fridge/submit', methods=['POST', 'GET'])
def submit():

    # Validate the request body contains JSON
    if request.is_json:

        # Parse the JSON into a Python dictionary
        req = request.get_json()

        # Print the dictionary
        print(req)
        image_generator.render_image(req['submission'])
        # Return a string along with an HTTP status code
        return "JSON received!", 200
