a = 0
#vnara sa a nacita exponent do 5^exponent
def nacitaj_pom() :
        global a
        if a != 0 :
                a-=1
                nacitaj_pom()
                a *= 5
        else :
                a = 1

#vnara sa podla zakladu ak je 0 zavola nacitat exponent
def nacitaj():
        global a
        if a == 0:
                a=int(input("zadaj exponent"))
                nacitaj_pom()
        else:
                a-=1
                nacitaj()
                a*=3
def skopiruj2do7():
        global a
        while a%7 == 0 :
        	a //= 7
        if a%2 == 0:
        	a //= 2
        	skopiruj2do7()
        	a *= 2
        	a *=7
def pripocitaj3():
        global a
        if a%3 == 0 :
        	a //= 3
        	pripocitaj3()
        	a *= 3
        	a *= 2
def vynasob2_3():
        global a
        skopiruj2do7()
        while a%2 == 0:
        	a //= 2
        while a%7 == 0:
                a //= 7
                pripocitaj3()
def umocni():
        global a
        a *= 2
        while a%5 == 0:
                a //=5
                vynasob2_3()
def zapis():
        global a
        if a%2 == 0:
        	a //= 2
        	zapis()
        	a+=1
        else:
        	a = 0
def mocnina():
        global a
        a=int(input("zadaj zaklad"))
        nacitaj()
        umocni()
        zapis()
        print("mocnina je:"+ str(a))

