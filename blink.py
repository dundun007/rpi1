from gpiozero import LED
from time import sleep

red = LED(6)

while True:
    red.on()
    print("red is on")
    sleep(1)
    red.off()
    sleep(1)
