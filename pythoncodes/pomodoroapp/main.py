from tkinter import *
# ---------------------------- CONSTANTS ------------------------------- #
#hex color value
PINK = "#e2979c"
RED = "#e7305b"
GREEN = "#9bdeac"
YELLOW = "#f7f5dd"
FONT_NAME = "Courier"
WORK_MIN = 25
SHORT_BREAK_MIN = 5
LONG_BREAK_MIN = 20

# ---------------------------- TIMER RESET ------------------------------- # 

# ---------------------------- TIMER MECHANISM ------------------------------- # 

# ---------------------------- COUNTDOWN MECHANISM ------------------------------- # 

# ---------------------------- UI SETUP ------------------------------- #
window=Tk()
window.title("Pomodoro")
window.config(padx=100,pady=50,bg=YELLOW)
#got width and height from pixel size
tomatoimage=PhotoImage(file="tomato.png")
canvas=Canvas(width=200,height=223,bg=YELLOW,highlightthickness=0)
canvas.create_image(100,112,image=tomatoimage)
canvas.create_text(102,135,text="00:00",fill="white",font=(FONT_NAME,30,"bold"))
canvas.pack()
window.mainloop()