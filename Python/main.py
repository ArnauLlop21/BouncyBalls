import pygame
import time
 
SCREEN = pygame.display.set_mode ((600, 400))
circle_x = 50
circle_y = 50
speed = 0
radius = 30
gravity = 0.1
coefficient_of_restitution = 0.9
 
while True :
    for event in pygame.event.get () :
        if event.type == pygame.QUIT :
            quit ()
 
    SCREEN.fill ((0, 0, 0)) # Background color to rgb 0,0,0 (Black)

    pygame.draw.circle (SCREEN, (255, 0, 0), (circle_x,circle_y), radius=radius)

    speed += gravity
    circle_y += speed
    #circle_y += speed

    pygame.display.update ()
    time.sleep(0.01)

    if circle_y + radius >= SCREEN.get_height():
        circle_y = SCREEN.get_height() - radius
        speed *= -coefficient_of_restitution  # reverse direction with some energy loss
    elif circle_y - radius <= 0:
        circle_y = radius
        speed *= -coefficient_of_restitution  # reverse direction with some energy loss
        # NEGATIVE speed means it goes up
        # POSITIVE speed means it goes down
        # :)



    