import threading
import time
import random
import math
import tkinter
    
def filozof(i):
    while True:
        rozmyslaj()
        zober(i)
        zober((i+1)%5)
        jedz()
        pusti(i)
        pusti((i+1)%5)


class Filozof(threading.Thread):
    zijeme = True    
    def __init__(self, meno, mutex, index):
        threading.Thread.__init__(self)
        self.meno = meno
        self.mutex = mutex
        self.mozem = threading.Lock()
        self.mozem.acquire()            ## musi zacinat na 0
        self.stav = "rozmysla"
        self.eth = [0, 1, 0]    #eat, think, hungry
        ###
        x = round(math.sin(math.radians((360 / 5)*index))*200 + 300)
        y = round(math.cos(math.radians((360 / 5)*index))*200 + 300)
        self.obr_person = tkinter.PhotoImage(file='think.png')
        self.person = canvas.create_image(x, y, image=self.obr_person)
        canvas.create_text(x-35, y-40, text=str(index+1))
        self.pole = [None]*3
        for i in range(3):
            self.pole[i] = canvas.create_text(675+i*50, 175+index*50, \
                                              text=str(self.eth[i]), font='arial 15', fill='navy')
    def nastav_susedov(self,lavy,pravy):
        self.susedL = lavy
        self.susedP = pravy

    def run(self):
        while self.zijeme:
            self.rozmyslaj()
            self.zober()
            self.jedz()
            self.pusti()

    def zober(self):
        with self.mutex:
            self.stav = "hladny"
            self.eth[2] += 1
            canvas.itemconfig(self.pole[2], text = str(self.eth[2]))
            self.obr_person = tkinter.PhotoImage(file='hungry.png')
            canvas.itemconfig(self.person, image = self.obr_person)
            self.test()
        self.mozem.acquire()

    def pusti(self):
        with self.mutex:
            self.stav = "rozmysla"
            self.eth[1] += 1
            canvas.itemconfig(self.pole[1], text = str(self.eth[1]))
            self.obr_person = tkinter.PhotoImage(file='think.png')
            canvas.itemconfig(self.person, image = self.obr_person)
            self.susedL.test()
            self.susedP.test()

    def test(self):
        if self.stav == "hladny" and \
          self.susedL.stav != "je" and \
          self.susedP.stav != "je":
            self.stav = "je"
            self.eth[0] += 1
            canvas.itemconfig(self.pole[0], text = str(self.eth[0]))
            self.obr_person = tkinter.PhotoImage(file='eat.png')
            canvas.itemconfig(self.person, image = self.obr_person)
            #canvas.update()
            self.mozem.release()

    def rozmyslaj(self):
        #print('{} rozmysla\n'.format(self.meno),end = "")
        time.sleep(random.uniform(3,13))
        #print('{} je hladny\n'.format(self.meno),end = "")

    def jedz(self):
        #print('{} zacina jest\n'.format(self.meno),end = "")
        time.sleep(random.uniform(1,10))
        #print('{} dojedol\n'.format(self.meno),end = "")

def filozofuj():
    mutex = threading.Lock()
    mena = ('Aristoteles','Kant','Buddha','Marx', 'Russel')
 
    f = [Filozof(mena[i], mutex, i) for i in range(5)]
    for i in range(5):
        f[i].nastav_susedov(f[(i-1)%5],f[(i+1)%5])
    for i in range(6):
        for j in range(4):
            canvas.create_rectangle(600+j*50, 100+i*50, 650+j*50, 150+i*50)
            if j == 0 and i != 0:
                canvas.create_text(625, 125+50*i, text=str(i), font='calibri 15 italic')
            elif j == 0 and i == 0:
                canvas.create_text(625, 125+50*i, text="FILOZOF")
            elif i == 0:
                text = ""
                if j == 1:
                    text = "EAT"
                elif j == 2:
                    text = "THINK"
                elif j == 3:
                    text = "HUNGRY"
                canvas.create_text(625+j*50, 125, text=text)
    Filozof.zijeme = True
    for p in f: p.start()
    time.sleep(30)
    Filozof.zijeme = False
    for p in f: p.join()
    print('koniec')
                  

canvas = None # na začiatku canvas ešte neexistuje

def canvasloop():
    global canvas
    canvas = tkinter.Canvas(height = 600, width = 900)
    canvas.pack()
    canvas.mainloop()    
    
threading.Thread(target = canvasloop).start()

while not canvas: pass # počkajme, kým sa vytvorí canvas

canvas.create_oval(150, 150, 450, 450, fill = 'saddlebrown')
obr_food = tkinter.PhotoImage(file='food.png')
###
food = [None] * 5
stick = [None] * 5
uhol = 360 / 5
uhol2 = uhol/2
for i in range(5):
    x = round(math.sin(math.radians(uhol*i))*120 + 300)
    y = round(math.cos(math.radians(uhol*i))*120 + 300)
    food[i] = canvas.create_image(x, y, image=obr_food)
    ###        
    x1 = round(math.sin(math.radians(uhol2 + uhol*i))*100 + 300)
    y1 = round(math.cos(math.radians(uhol2 + uhol*i))*100 + 300)
    x2 = round(math.sin(math.radians(uhol2 + uhol*i))*130 + 300)
    y2 = round(math.cos(math.radians(uhol2 + uhol*i))*130 + 300)
    stick[i] = canvas.create_line(x1, y1, x2, y2, width = 3, fill = 'tan')

                


filozofuj()
