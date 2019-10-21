
# ---------------------------- SOURCES ----------------------------

a = 0

def my_input(prompt):
    global a
    a = int(input(prompt))

# ---------------------------- ULOHA 1 ----------------------------

# Zvladne 13x13

def zakoduj3():
    global a
    if a > 0:
        a -= 1
        zakoduj3()
        a *= 3
    else:
        my_input("Zadaj prirodzene cislo:")
        zakoduj5()


def zakoduj5():
    global a
    if a > 0:
        a -= 1
        zakoduj5()
        a *= 5
    else:
        a = 1


def zvysExponentDvojke():
    global a
    a *= 2


def log5():
    global a
    if a % 5 == 0:
        a /= 5
        zvysExponentDvojke()
        log5()
        a *= 5

def log3():
    global a
    if a % 3 == 0:
        log5()
        a /= 3
        log3()

def log2():
    global a
    if a % 2 == 0:
        a /= 2
        log2()
        a += 1
    else:
        a = 0

def vynasobCisla():
    my_input("Zadaj prirodzene cislo:")
    zakoduj3() # aj 5
    log3() # zakoduje vysledok do mocnin 2ky
    log2() # poskytne vysledok
    print(f"Ich sucin je: {a}")


# ---------------------------- RUN ----------------------------

if __name__ == '__main__':
    vynasobCisla()
