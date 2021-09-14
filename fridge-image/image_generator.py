from PIL import Image, ImageDraw, ImageFont
from random import randint
import string
import time

# Globals
max_width = 600
max_height = 600
image = Image.new("RGB", (50000,1080), "#FFF") # scrap image
draw = ImageDraw.Draw(image)
image2 = Image.new("RGB", (max_width, max_height), "#FFF") # final image
font_size = 100
stroke_width = int(font_size / 20)

title_font = ImageFont.truetype('AlphaFridgeMagnets.ttf', font_size)

make_color = lambda : (randint(50, 255), randint(50, 255), randint(50,255))

gen_border = lambda color, shade: tuple(int((1-shade)*x) for x in color)
fill = " Oj "
w_fill, y_draw = draw.textsize(fill, font=title_font)



def render_char(c, x_draw, x_paste, w_fill, y_draw, y_paste):
    if randint(0,15) == 1:
        c = str(c.upper())
    letter_color = make_color()
    border_color = gen_border(letter_color, 0.3)
    # Stroke width 4, add 12
    w_full = draw.textsize(fill+c, font=title_font)[0] + stroke_width * 3
    print(w_full)
    w = w_full - w_fill    # the width of the character on its own
    draw.text((x_draw,0), fill+c, letter_color, font=title_font, stroke_width=stroke_width, stroke_fill=border_color)

    # Stroke width 4, minus 8 to [0] & [2]
    iletter = image.crop((x_draw+w_fill - 2 * stroke_width, 0, x_draw+w_full - 2 * stroke_width, y_draw))
    image2.paste(iletter, (x_paste, y_paste))
    return w_full

def render_word(word, x_draw, x_paste, y_draw, y_paste, max_width, max_height):
    w_word = draw.textsize(word, font=title_font, stroke_width=stroke_width)[0]

    # ... the + stroke_width * (len(word) + stroke_width) is magic number bullshit that happens to work very well
    # And the plus 2 is more magic bullshit that MAYBE accounts for the spaces before and after word
    # Extra 2 for chance of uppercase and extra 6 for randomly starting sentances
    # TODO: The 10 magic number works but is total bullshit and causes a lot of empty lines.
    # Maybe in the future work out a random number of spaces between words, and include them in the words
    # during generation
    while x_paste + w_word + stroke_width * (len(word) + stroke_width + 10) >= max_width:
        y_paste += y_draw
        x_paste = randint(0, max_width / 2 - font_size)
    for c in (word + " "):
        global w_full
        global w_fill
        w_full = render_char(c, x_draw, x_paste, w_fill, y_draw, y_paste)
        w = w_full - w_fill
        x_draw += w_full
        x_paste += w
    print(x_draw, x_paste, y_draw, y_paste)
    return x_draw, x_paste, y_draw, y_paste

def render_image(words):
    x = 0
    w_fill, y_draw = draw.textsize(fill, font=title_font)
    x_draw, x_paste = font_size, 0
    y_paste = 0

    for word in words.split():
        x_draw, x_paste, y_draw, y_paste = render_word(word, x_draw, x_paste, y_draw, y_paste, max_width, max_height)
        # If the letter width of a word goes past the resolution, render the entire word on the next line
        
    image2.save('fridge.jpeg')

if __name__ == '__main__':
    render_image("beep bop ;lkj ;lkj ;lkjdsaf lk;jsadf;lk ja")
