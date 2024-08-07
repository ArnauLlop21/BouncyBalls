import pygame
import time

 
SCREEN = pygame.display.set_mode ((600, 400))
MOUSE = pygame.mouse
circle_x = 50
circle_y = 50
speed_y = 0
speed_x = 10
radius = 30
gravity = 0.1
coefficient_of_restitution_y = 0.7
coefficient_of_restitution_x = 0.5
coefficient_of_friction = 0.1

alive = True

clock = pygame.time.Clock()
while True :
    for event in pygame.event.get () :
        if event.type == pygame.QUIT :
            quit ()
    click_state = MOUSE.get_pressed(num_buttons=3)
    if any(click_state) == True:
        speed_x += 1 if (speed_x >= 0) else -1
        speed_y += -1
        alive = True
    if alive == True: 
        SCREEN.fill ((0, 0, 0)) # Background color to rgb 0,0,0 (Black)

        pygame.draw.circle (SCREEN, (255, 0, 0), (circle_x,circle_y), radius=radius)

        speed_y += gravity
        circle_y += speed_y
        circle_x += speed_x
        
        if -0.07 <= speed_y <= 0.07:
            speed_y = 0
            if speed_x > 0:
                speed_x -= coefficient_of_friction
            else:
                speed_x += coefficient_of_friction

        if -0.07 <= speed_x <= 0.07:
            speed_x = 0

        pygame.display.update ()
        time.sleep(0.01)

        if (speed_y == 0) and (speed_x == 0):
            alive = False

        if circle_y + radius >= SCREEN.get_height():
            circle_y = SCREEN.get_height() - radius
            speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
        elif circle_y - radius <= 0:
            circle_y = radius
            speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
            # NEGATIVE speed_y means it goes up
            # POSITIVE speed_y means it goes down
            # :)
        if circle_x + radius > SCREEN.get_width():
            circle_x = SCREEN.get_width() - radius
            speed_x *= -coefficient_of_restitution_x
        elif circle_x - radius <= 0:
            circle_x = radius
            speed_x *= -coefficient_of_restitution_x
        
        clock.tick()
        print(clock.get_fps())
        
    else:
        time.sleep(0.01)





    