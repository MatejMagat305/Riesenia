# idea nic extra
# snad som dobre pochopil zadanie
# funkcia minc generuje vysl -> vyslednu sumu a k tejto sume vytvara pouzite mince -> mince ktorymi by sa dala dana suma zaplatit
# vyuziva pritom najmensia_minca_pre_sumu(s) s -> je suma a funkcia vrati najvacsiu mincu ktora je mnesia ako suma s
# teda pre sumu 12 vrati mincu 10
# avsak kvoli moznosti ze dana suma moze byt vytvorena mensim poctom minci pri vydavku
# kontrolne spusta minc2() ktore obdrzi sumu a mince a skusa to opacne = platba vacsou mincou a nasledy vydaj
# okrem najmensia_minca_pre_sumu pouziva minc2 aj najvacsia_minca_pre_sumu vrati mincu z ktorej bude potrebne vydat
# teda najmensia_minca_pre_sumu pre sumu 18, vrati mincu 20
# ak minc2 zisti ze vydajom by bol pocet minci mensi, minc hlada dalej

# 22:45 oprava MIN2C...

# vloz pocet minci: 1
# Najmensia suma pre 1 mincu/e/i je: 1 a pouzite mince su: [1]
# vloz pocet minci: 2
# Najmensia suma pre 2 mincu/e/i je: 3 a pouzite mince su: [2, 1]
# vloz pocet minci: 3
# Najmensia suma pre 3 mincu/e/i je: 13 a pouzite mince su: [10, 2, 1]
# vloz pocet minci: 4
# Najmensia suma pre 4 mincu/e/i je: 33 a pouzite mince su: [20, 10, 2, 1]
# vloz pocet minci: 5
# Najmensia suma pre 5 mincu/e/i je: 83 a pouzite mince su: [50, 20, 10, 2, 1]
# vloz pocet minci: 6
# Najmensia suma pre 6 mincu/e/i je: 133 a pouzite mince su: [50, 50, 20, 10, 2, 1]
# vloz pocet minci: 7
# Najmensia suma pre 7 mincu/e/i je: 183 a pouzite mince su: [50, 50, 50, 20, 10, 2, 1]
# vloz pocet minci: 8
# Najmensia suma pre 8 mincu/e/i je: 233 a pouzite mince su: [50, 50, 50, 50, 20, 10, 2, 1]
# vloz pocet minci: 9
# Najmensia suma pre 9 mincu/e/i je: 283 a pouzite mince su: [50, 50, 50, 50, 50, 20, 10, 2, 1]
# vloz pocet minci:


# ak to bezi spravne
# z vypisu vidime ze sa nam mince opakuju, preto (ak to bezi spravne)\
#  by sa cyklom dalo poriesit pridavanie spravneho poctu 50kovzch minci a nasledne en prilepit 20, 10, 2, 1
#  nenakodene, cas...

# pokracovanie komentaru za cas...
# rozumnejsie by bolo generovat mince do baliku od konca
# naco cyklus, par podmienok to osefuje...

def mince_postupka(n):
    suma = 0
    pouzite_mince = []
    n = int(n)

    if n >= 1:
        pouzite_mince.append(1)
        suma += 1
    if n >= 2:
        pouzite_mince.append(2)
        suma += 2
    if n >= 3:
        pouzite_mince.append(10)
        suma += 10
    if n >= 4:
        pouzite_mince.append(20)
        suma += 20
    if n >= 5:
        pouzite_mince += [50]*(n-4)
        suma += 50*(n-4)

    print('postupka Najmensia suma pre {} mincu/e/i je: {} a pouzite mince su: {}'.format(n, suma, pouzite_mince))



import time
# 1,2,5,10,20,50
mince = [1,2,5,10,20,50,100]
kontrolny_vypis = False

def minc(n):


    finded = False
    vysl = 0 #int(n) * mince[0]
    suma = 0
    pocet_minci = int(n)
    pouzite_mince = []
    revers_better = False
    while not finded:
        vysl += 1
        while pocet_minci > 0:
            pocet_minci -= 1

            if vysl > suma:
                pouzite_mince.append(najmensia_minca_pre_sumu(vysl-suma))
                suma += najmensia_minca_pre_sumu(vysl-suma)

            else:
                pouzite_mince.append(najmensia_minca_pre_sumu(vysl))
                suma += najmensia_minca_pre_sumu(vysl)

            if pocet_minci == 0: #suma == vysl and

                if suma == vysl:
                    revers_better = minc2(suma, pouzite_mince)
                    if kontrolny_vypis: print('Testovana suma pre {} mincu/e/i je: {} a pouzite mince su: {}'.format(n, suma, pouzite_mince))
                    if not revers_better:

                        finded = True
                        print('Najmensia suma pre {} mincu/e/i je: {} a pouzite mince su: {}'.format(n, suma, pouzite_mince))



        # vysl += 1
        if kontrolny_vypis: print('velky presiel', vysl)
        if kontrolny_vypis: print('Testovana suma pre {} mincu/e/i je: {} a pouzite mince su: {}'.format(n, suma, pouzite_mince))
        suma = 0
        pocet_minci = int(n)
        pouzite_mince = []




def minc2(suma, mince=[]):
    suma = int(suma)
    pouzite_mince2 = []
    finded = False
    suma2 = 0

    while not finded:
        # if len(pouzite_mince2) == 0:
        #     pouzite_mince2.append(najvacsia_minca_pre_sumu(suma))
        #     suma2 += najvacsia_minca_pre_sumu(suma)
        if suma2 < suma:
            pouzite_mince2.append(najvacsia_minca_pre_sumu(suma-suma2))
            suma2 += najvacsia_minca_pre_sumu(suma-suma2)
        elif suma2 > suma:
             pouzite_mince2.append(najmensia_minca_pre_sumu(suma2-suma))
             suma2 -= najmensia_minca_pre_sumu(suma2-suma)
        elif suma2 == suma:
            if len(pouzite_mince2) < len(mince):
                return True
            else:
                return False
        # else:
        #     pouzite_mince2.append(najmensia_minca_pre_sumu(suma-suma2))
        #     suma2 += najmensia_minca_pre_sumu(suma-suma2)
        if kontrolny_vypis: print('MIN2 Testovana suma {} suma2 {} a pouzite mince su: {}'.format(suma, suma2, pouzite_mince2))



def najvacsia_minca_pre_sumu(s):
    if (s > mince[len(mince)-1]):
        return mince[len(mince)-1]
    for i in range(len(mince)):
        if int(s) <= mince[i]:
            # print(mince[i])
            return mince[i]
    return 0


def najmensia_minca_pre_sumu(s):
    if (s > mince[len(mince)-1]):
        return mince[len(mince)-1]
    for i in range(len(mince)-1, -1, -1):
        if int(s) >= mince[i]:
            # print(mince[i])
            return mince[i]
    return 0



print(minc(5))

for i in range(20):
    t = input('vloz pocet minci: ')
    c = time.time()
    minc(t)
    print("cas: ", time.time() - c)
    c = time.time()
    mince_postupka(t)
    print("cas: ", time.time() - c)
    # minc2(t)
    # s = input('vloz sumu: ')
    # najvacsia_minca_pre_sumu(s)
    # najmensia_minca_pre_sumu(s)


# Ako dopadnú výsledky s Euro-mincami, alebo s americkými mincami ?
# dopadnu dobre, v teoretickej hladine sa zmenia  len mince a s tym suvisiace vysledky, v praktickej
# funkcia mince_postupka bez uprav nebude fungovat, s upravami by to malo ist podobne riesit (za predpokladu ze mi to funguje spravne)
# funkcia minc() a jej spolupracujuce by po zmene globalnej mince = [1,2,5,10,20,50] na ine mince mali fungovat bez problemov dalej, neotestovane





# PROTOTYP / zaloha
# import time
# # 1,2,5,10,20,50
# mince = [1,2,5,10,20,50]
#
# def minc(n):
#     # print(n)
#     # if n == 1:
#     #     print('Najmensia suma pre {} mincu/e/i je: {}'.format(n, 1))
#
#     finded = False
#     vysl = 0 #int(n) * mince[0]
#     suma = 0
#     pocet_minci = int(n)
#     pouzite_mince = []
#     while not finded:
#         vysl += 1
#         while pocet_minci > 0:
#
#             pocet_minci -= 1
#
#             if vysl > suma:
#                 pouzite_mince.append(najmensia_minca_pre_sumu(vysl-suma))
#                 suma += najmensia_minca_pre_sumu(vysl-suma)
#             else:
#                 pouzite_mince.append(najmensia_minca_pre_sumu(vysl-suma))
#                 suma += najmensia_minca_pre_sumu(vysl)
#             print('pouzite mince', pouzite_mince)
#             if pocet_minci == 0: #suma == vysl and
#
#                 if suma == vysl:
#                     finded = True
#                     print('Najmensia suma pre {} mincu/e/i je: {}'.format(n, suma))
#
#
#         # vysl += 1
#         suma = 0
#         pocet_minci = int(n)
#         pouzite_mince = []
#         print('velky presiel', vysl)
#
#
#
#
# def najvacsia_minca_pre_sumu(s):
#     for i in range(len(mince)):
#         if int(s) <= mince[i]:
#             # print(mince[i])
#             return mince[i]
#
# def najmensia_minca_pre_sumu(s):
#     print(s)
#     for i in range(len(mince)-1, -1, -1):
#         if int(s) >= mince[i]:
#             # print(mince[i])
#             return mince[i]
#
#
#
#
# for i in range(5):
#     t = input('vloz pocet minci: ')
#     minc(t)
#     # s = input('vloz sumu: ')
#     # najvacsia_minca_pre_sumu(s)
#     # najmensia_minca_pre_sumu(s)
#
#
#
# # for i in range(len(mince)):
# #     if suma == mince[i]:
# #         finded = True
# #         print('Najmensia suma pre {} mincu/e/i je: {}'.format(n, suma))
# #     else:
# #         pass
#
#
#
# # for i in range(len(mince)):
# #     if vysl == suma:
# #         finded = True
# #         print('Najmensia suma pre {} mincu/e/i je: {}'.format(n, suma))
# #     else:
# #         suma += mince[i]
# #         print(suma, pocet_minci)
# #         pocet_minci -= 1
# #         vysl += 1
