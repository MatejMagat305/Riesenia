import sys
sys.setrecursionlimit(1100) # limit pre rekurzie 

print("Umocnovanie ( a**b )")
a = int(input("Zadaj a (a < 1000): "))

def run():
    global a
    
    if ( a > 0 ):
        a -= 1
        run()
        a += 1
    else:
        a = int(input("Zadaj b: "))
        a = a * 1000

### pridana funkcia
def run2():
    global a

    if ( (a%1000) > 0 ):
        a -= 1
        run2()
        a = (a%1000) * (a//1000) * 1000 + (a%1000)
    else:
        a += a//1000
### ### ###

if ( a < 1000 ): 
    run()

    ### vami oznaceny problem
    print("Vysledok je:", (a%1000)**(a//1000)) 
    ### ### ###

    ### pridany vypis pre run2
    a = (a%1000)*1000 + (a//1000)
    run2()
    a = (a//1000) // (a%1000)
    print("Vysledok cez rekurziu je:", a)
    ### ### ###
else:
    print("CHYBA: Zadana hodnota je >= 1000...")
    
