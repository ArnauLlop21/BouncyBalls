import pygame
import time
import ball
import math
 
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


# INTANTIATION OF GAME COMPONENTS
b1 = ball.Ball(30,50,50,10,0,(0,255,0))
b2 = ball.Ball(20,0,5,5,5,(255,240,156))
BALLS.append(b1)
BALLS.append(b2)
clock = pygame.time.Clock()

# NEGATIVE speed_y means it goes up
# POSITIVE speed_y means it goes down
# :)
def CircleCollisiondWithVerticalBounds(curr_ball:ball):
    if curr_ball.circle_y + curr_ball.radius >= SCREEN.get_height():
        curr_ball.circle_y = SCREEN.get_height() - curr_ball.radius
        curr_ball.speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss
    elif curr_ball.circle_y - curr_ball.radius <= 0:
        curr_ball.circle_y = curr_ball.radius
        curr_ball.speed_y *= -coefficient_of_restitution_y  # reverse direction with some energy loss

def CircleCollisiondWithHorizontalBounds(curr_ball:ball):
    if curr_ball.circle_x + curr_ball.radius > SCREEN.get_width():
        curr_ball.circle_x = SCREEN.get_width() - curr_ball.radius
        curr_ball.speed_x *= -coefficient_of_restitution_x
    elif curr_ball.circle_x - curr_ball.radius <= 0:
        curr_ball.circle_x = curr_ball.radius
        curr_ball.speed_x *= -coefficient_of_restitution_x

def PeriodicSpeedModifications(curr_ball:ball):
    curr_ball.speed_y += gravity
    curr_ball.circle_y += curr_ball.speed_y
    curr_ball.circle_x += curr_ball.speed_x

def BallNotJumping(curr_ball:ball, minimal_bound, maximal_bound):
    return minimal_bound <= curr_ball.speed_y <= maximal_bound

def BallNotMovingHoriz(curr_ball:ball, minimal_bound, maximal_bound):
    return minimal_bound <= curr_ball.speed_x <= maximal_bound

def BallTooLowSpeedMofications(curr_ball:ball):
    curr_ball.speed_y = 0
    if curr_ball.speed_x > 0:
        curr_ball.speed_x -= coefficient_of_friction
    else:
        curr_ball.speed_x += coefficient_of_friction

def BallAlive(curr_ball:ball):
    return (curr_ball.speed_y == 0) and (curr_ball.speed_x == 0)

def CalculateDistance(ball1:ball, ball2:ball):
    p = [ball1.circle_x,ball1.circle_y]
    q = [ball2.circle_x,ball2.circle_y]
    return math.dist(p,q)

def BallsTooClose(ball1:ball, ball2:ball):
    ball1.speed_x *= -1    
    ball1.speed_y *= -1    
    ball2.speed_x *= -1    
    ball2.speed_y *= -1    

while True :
    for event in pygame.event.get () :
        if event.type == pygame.QUIT :
            quit ()
    click_state = MOUSE.get_pressed(num_buttons=3)
    for i, curr_ball in enumerate(BALLS):
        if any(click_state) == True:
            curr_ball.speed_x += 1 if (curr_ball.speed_x >= 0) else -1
            curr_ball.speed_y += -1
            curr_ball.alive = True
        if curr_ball.alive == True: 
            SCREEN.fill ((0, 0, 0)) # Background color to rgb 0,0,0 (Black)

            for j in BALLS:
                pygame.draw.circle (SCREEN, j.color, (j.circle_x,j.circle_y), radius=j.radius)

            PeriodicSpeedModifications(curr_ball=curr_ball)

            if BallNotJumping(curr_ball=curr_ball, minimal_bound=-0.07, maximal_bound=0.07):
                BallTooLowSpeedMofications(curr_ball=curr_ball)

            if BallNotMovingHoriz(curr_ball=curr_ball, minimal_bound=-0.07, maximal_bound=0.07):
                curr_ball.speed_x = 0

            pygame.display.update ()

            if BallAlive(curr_ball=curr_ball):
                curr_ball.alive = False

            CircleCollisiondWithVerticalBounds(curr_ball=curr_ball)
            CircleCollisiondWithHorizontalBounds(curr_ball=curr_ball)
            
            clock.tick()
            #print(clock.get_fps())

            for other_ball in BALLS[i + 1:]:
                if CalculateDistance(other_ball, curr_ball) <= (curr_ball.radius + other_ball.radius):
                    BallsTooClose(curr_ball, other_ball)
                    
            
    
    time.sleep(0.01)

    