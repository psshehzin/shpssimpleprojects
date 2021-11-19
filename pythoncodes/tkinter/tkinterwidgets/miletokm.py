from tkinter import *
def convert():
    valuemil=int(entry.get())
    valuekm=round(valuemil*1.609,2)
    label3=Label(text=f"{valuekm}")
    label3.grid(row=2,column=2)
window=Tk()
window.minsize(250,100)
window.title(" Mile to Km Converter")
window.config(padx=10,pady=10)
entry=Entry(width=20)
entry.grid(row=1,column=2)
label1=Label(text="Miles")
label1.config(padx=10)
label1.grid(row=1,column=3)
label2=Label(text="is equal to")
label2.grid(row=2,column=1)
label4=Label(text="KM")
label4.grid(row=2,column=3)
button=Button(text="calculate",command=convert)
button.grid(row=3,column=2)
window.mainloop()