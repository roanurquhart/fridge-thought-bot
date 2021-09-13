from PIL import Image, ImageDraw, ImageFont
from random import randint
import string

make_color = lambda : (randint(50, 255), randint(50, 255), randint(50,255))

image = Image.new("RGB", (1920,1080), (0,0,0)) # scrap image
draw = ImageDraw.Draw(image)
image2 = Image.new("RGB", (1920, 1080), (0,0,0)) # final image

title_font = ImageFont.truetype('AlphaFridgeMagnets.ttf', 100)
# image_editable.text((15,15), title_text, (237, 230, 211), font=title_font)
# Filler used to determine upper and lower bounds of text
# Currently it fills up the sample image too fast with a high resolution
# Will have to either a
# chunk out the sentance and go vertically / make multiple canvases
fill = " Oj "
x = 0
w_fill, y = draw.textsize(fill, font=title_font)
x_draw, x_paste = 0, 0
for c in "The quick brown fox jumps over the lazy dog.":
    w_full = draw.textsize(fill+c, font=title_font)[0]
    w = w_full - w_fill     # the width of the character on its own
    draw.text((x_draw,0), fill+c, make_color(), font=title_font)
    iletter = image.crop((x_draw+w_fill, 0, x_draw+w_full, y))
    image2.paste(iletter, (x_paste, 0))
    x_draw += w_full
    x_paste += w
image2.show()
image.show()
