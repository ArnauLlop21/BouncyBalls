import pygame
import time
import ball
 
# DEFINITION OF CRITICAL OBJECTS
SCREEN = pygame.display.set_mode ((600, 400))
MOUSE = pygame.mouse
BALLS = []

# DEFINITION OF CONFIG VARIABLES
coefficient_of_restitution_y = 0.7
coefficient_of_restitution_x = 0.5
coefficient_of_friction = 0.1
gravity = 0.1

# DEFINITION OF AUXILIAR VARIABLES
alive = True

# INTANTIATION OF GAME COMPONENTS
b1 = ball.Ball(30,50,50,10,0,(255,0,0))
b2 = ball.Ball(20,40,35,5,5,(255,240,156))
BALLS.append(b1)
BALLS.append(b2)
clock = pygame.time.Clock()


while True :
    for event in pygame.event.get () :
        if event.type == pygame.QUIT :
            quit ()
    click_state = MOUSE.get_pressed(num_buttons=3)
    for curr_ball in BALLS:
        if any(click_state) == True:
            curr_ball.speed_x += 1 if (curr_ball.speed_x >= 0) else -1
            curr_ball.speed_y += -1
            alive = True
        if alive == True: 
            SCREEN.fill ((0, 0, 0)) # Background color to rgb 0,0,0 (Black)

            for i in BALLS:
                pygame.draw.circle (SCREEN, i.color, (i.circle_x,i.circle_y), radius=i.radius)

            curr_ball.speed_y += gravity
            curr_ball.circle_y += curr_ball.speed_y
            curr_ball.circle_x += curr_ball.speed_x

            if -0.07 <= curr_ball.speed_y <= 0.07:
                curr_ball.speed_y = 0
                if curr_ball.speed_x > 0:
                    curr_ball.speed_x -= coefficient_of_friction
                else:
                    curr_ball.speed_x += coefficient_of_friction

            if -0.07 <= curr_ball.speed_x <= 0.07:
                curr_ball.speed_x = 0

            pygame.display.update ()
    
            if (curr_ball.speed_y == 0) and (curr_ball.speed_x == 0):
                alive = False

            if curr_ball.circle_y + curr_ball.radius >= SCREEN.get_height():
                curr_ball.circle_y = SCREEN.get_height() - curr_ball.radius
                curr_ball.speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
            elif curr_ball.circle_y - curr_ball.radius <= 0:
                curr_ball.circle_y = curr_ball.radius
                curr_ball.speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
                # NEGATIVE speed_y means it goes up
                # POSITIVE speed_y means it goes down
                # :)
            if curr_ball.circle_x + curr_ball.radius > SCREEN.get_width():
                curr_ball.circle_x = SCREEN.get_width() - curr_ball.radius
                curr_ball.speed_x *= -coefficient_of_restitution_x
            elif curr_ball.circle_x - curr_ball.radius <= 0:
                curr_ball.circle_x = curr_ball.radius
                curr_ball.speed_x *= -coefficient_of_restitution_x

            clock.tick()
            print(clock.get_fps())
        
        
    
    time.sleep(0.01)





    