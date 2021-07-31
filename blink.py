from gpiozero import LED
from time import sleep

red = LED(17)

while True:
    red.on()
    print("red is on")
    sleep(0.1)
    red.off()
    print("red is off")
    sleep(0.1)
