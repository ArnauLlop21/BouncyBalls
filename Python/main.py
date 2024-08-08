import pygame
import time
import ball
 
SCREEN = pygame.display.set_mode ((600, 400))
MOUSE = pygame.mouse

gravity = 0.1
coefficient_of_restitution_y = 0.7
coefficient_of_restitution_x = 0.5
coefficient_of_friction = 0.1

b1 = ball.Ball(30,50,50,10,0)

alive = True

clock = pygame.time.Clock()
while True :
    for event in pygame.event.get () :
        if event.type == pygame.QUIT :
            quit ()
    click_state = MOUSE.get_pressed(num_buttons=3)
    if any(click_state) == True:
        b1.speed_x += 1 if (b1.speed_x >= 0) else -1
        b1.speed_y += -1
        alive = True
    if alive == True: 
        SCREEN.fill ((0, 0, 0)) # Background color to rgb 0,0,0 (Black)

        pygame.draw.circle (SCREEN, (255, 0, 0), (b1.circle_x,b1.circle_y), radius=b1.radius)

        b1.speed_y += gravity
        b1.circle_y += b1.speed_y
        b1.circle_x += b1.speed_x
        
        if -0.07 <= b1.speed_y <= 0.07:
            b1.speed_y = 0
            if b1.speed_x > 0:
                b1.speed_x -= coefficient_of_friction
            else:
                b1.speed_x += coefficient_of_friction

        if -0.07 <= b1.speed_x <= 0.07:
            b1.speed_x = 0

        pygame.display.update ()
        time.sleep(0.01)

        if (b1.speed_y == 0) and (b1.speed_x == 0):
            alive = False

        if b1.circle_y + b1.radius >= SCREEN.get_height():
            b1.circle_y = SCREEN.get_height() - b1.radius
            b1.speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
        elif b1.circle_y - b1.radius <= 0:
            b1.circle_y = b1.radius
            b1.speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
            # NEGATIVE speed_y means it goes up
            # POSITIVE speed_y means it goes down
            # :)
        if b1.circle_x + b1.radius > SCREEN.get_width():
            b1.circle_x = SCREEN.get_width() - b1.radius
            b1.speed_x *= -coefficient_of_restitution_x
        elif b1.circle_x - b1.radius <= 0:
            b1.circle_x = b1.radius
            b1.speed_x *= -coefficient_of_restitution_x
        
        clock.tick()
        print(clock.get_fps())
        
    else:
        time.sleep(0.01)





    