from PIL import Image, ImageDraw, ImageFont
from random import randint
import string
import time
from pathlib import Path

##TODO: If word is too long, doesn't print and gets stuck
##TODO: 

# Globals
max_width = 2400
max_height = 1300

# Image offset
offset_x = 800
offset_y = 300
#RN it generates a scrach image and final image to paste the text into to get the colors
#Maybe look into https://pillow.readthedocs.io/en/stable/reference/Image.html#PIL.Image.Image.transform
image = Image.new("RGBA", (50000,1080), (255, 0, 0, 0)) # scrap image
draw = ImageDraw.Draw(image)
image2 = None # = Image.new("RGBA", (max_width, max_height), "#FAC") # final image
font_size = 175 
stroke_width = int(font_size / 15)

title_font = ImageFont.truetype('AlphaFridgeMagnets.ttf', font_size)

make_color = lambda : (randint(50, 255), randint(50, 255), randint(50,255))

gen_border = lambda color, shade: tuple(int((1-shade)*x) for x in color)
fill = " Oj "
w_fill, y_draw = draw.textsize(fill, font=title_font)



def render_char(c, x_draw, x_paste, w_fill, y_draw, y_paste):
    global image
    global image2
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
    iletter = image.crop((x_draw+w_fill - 2 * stroke_width, 0, x_draw+w_full - 2 * stroke_width, y_draw)).convert("RGBA")
    image2.paste(iletter, (x_paste, y_paste))

    return w_full

def render_word(word, x_draw, x_paste, y_draw, y_paste, max_width, max_height):
    global image, image2, offset_x, offset_y
    w_word = draw.textsize(word, font=title_font, stroke_width=stroke_width)[0]

    
    if  w_word + stroke_width * (len(word) + stroke_width + font_size // 2) >= max_width:
        word_1, word_2 = word[:len(word)//2], word[len(word)//2:]
        x_draw, x_paste, y_draw, y_paste = render_word(word_1,x_draw, x_paste, y_draw, y_paste, max_width, max_height)

        return render_word(word_2,x_draw, x_paste, y_draw, y_paste, max_width, max_height)
    # ... the + stroke_width * (len(word) + stroke_width) is magic number bullshit that happens to work very well
    # And the plus 2 is more magic bullshit that MAYBE accounts for the spaces before and after word
    # Extra 2 for chance of uppercase and extra 6 for randomly starting sentances
    # TODO: The 10 magic number works but is total bullshit and causes a lot of empty lines.
    # Maybe in the future work out a random number of spaces between words, and include them in the words
    # during generation
    overflowed = False
    while x_paste + w_word + stroke_width * (len(word) + stroke_width + 10) >= max_width:
        overflowed = True
        print(str(offset_x) + " is x offset" + str(max_width / 4 - font_size))
        x_paste = randint(offset_x, max_width // 2)
    
    if overflowed:
        y_paste += randint(y_draw, int(3 * y_draw))

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
    global image2, image, draw
    image = Image.new("RGBA", (50000,1080), (255, 0, 0, 0)) # scrap image
    draw = ImageDraw.Draw(image)
    image3 = Image.open("refrigerator-door-closeup-high-res.png").convert("RGBA")
    image2 = Image.new("RGBA", (2560, 1680), (255, 0, 0, 0)) # final image

    x = 0
    w_fill, y_draw = draw.textsize(fill, font=title_font)
    x_draw, x_paste = font_size, offset_x
    y_paste = offset_y

    for word in words.split():
        x_draw, x_paste, y_draw, y_paste = render_word(word, x_draw, x_paste, y_draw, y_paste, max_width, max_height)
        # If the letter width of a word goes past the resolution, render the entire word on the next line
    # if 
    background = image3
    foreground = image2

    Image.alpha_composite(background, foreground).save("fridge.png")

    # image2.save('fridge.png')

if __name__ == '__main__':
    render_image("beep bop ;lkj ;lkj ;lkjdsaf lk;jsadf;lk ja")
