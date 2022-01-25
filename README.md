# Fridge Thought Bot

## What is it?
A Discord-Bot which takes user input and dynamically generates input as a phrase of fridge letter-magnets. Uses Discord-Go for the golang backend logic, and Python Pillow / Flask for hosting a Python REST API to take input and generate the resulting image.

<img width="497" alt="Screen Shot 2022-01-25 at 1 39 03 AM" src="https://user-images.githubusercontent.com/24252598/150924895-93a321e4-4e87-419c-b358-45e3f38162bb.png">

## How does it work?
The bot generates a daily sequence of letters (seeded with more frequent vowels because it's better that way 🤷‍♂️ ... may change this letter to generate based off a standard letter frequency chart.) Users can submit entries which must consist of a subset of those letters only, otherwise it's not accepted.

Ex. Letters of the day are abdcy, baby is not a valid submission. But dab and cab are.

## Commands

`!fridge check` - Checks the letters of the day

`!fridge submit <submission>` - Submits an entry

## How to run?

It can be run locally using Python venv and go, but the preferred method is through a docker container.
```bash
 docker build -t fridgethoughtbot .
 docker run -dit -e BOT_TOKEN=${YOUR_TOKEN_HERE} fridge_bot
 ```
 
 Once the container is launched, execute the associated shell script to run everything.

 
