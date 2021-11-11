from turtle import Turtle, Screen
class Scoreboard(Turtle):
    def __init__(self):
        super().__init__()
        self.score=0
        self.hideturtle()
        self.penup()
        self.goto(0,295)
        self.pendown()
        self.color("yellow")
        self.write(f"SCORE : {self.score}", move=False, align="center", font=["Arial",20, 'normal'])
        
    def score_update(self):
        self.score=self.score+1
        self.clear()
        self.penup()
        self.goto(0,295)
        self.pendown()
        self.write(f"SCORE : {self.score}", move=False, align="center", font=["Arial",20, 'normal'])
    def game_over(self):
        self.penup()
        self.goto(0,0)
        self.pendown()
        self.write(f"GAME OVER SUCKER YOUR SCORE : {self.score}", move=False, align="center", font=["Arial",20, 'normal'])
score=Scoreboard()