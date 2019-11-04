determinant:
rozdelil som to na dve casti nad mainom je staré sekvenčné a pod je paralelné,
neviem prečo mi chanely vyhlasovali deadlock, obišiel som to, ale do teraz neviem ako a prečo......,
áno prakvapil ma výsledok pre menšie granularity, čakal som, že para bude zakaždým rýchlejšie.
zložitosť je o(n!) a ovela efektívneší je gausová eliminácia s o(n**3)   

Crawler:
môj crawler zvláda prechádzať paralelne stránky, využil som regulárny výraz,
o bol ukázaný v prednáške + volám modul net/url, že či tá stránka je validná url.parse.........,
ďalej som využil syncgroup, aby main počkal kým dobehnú goroutiny a nakoniec
som využil mutex pre zamknutie - synchonizáciu mapy, 
vypisuje stránky, ktoré navštívil, či sa mu podarilo spojiť, čo tam našiel
nezvláda otvoriť pdf, keď je na stránke.

matice:
obr. "servre" ukazujú ako som to spúštal, ak by bol potrebný výpis z konzoly, tak je vo "vypis.txt",
všetky triedy sú okopírované, zmenil som iba Worker.doit(v tej triede som priddal aj mutex)
a to iba tak, že pred tým výsledok dávala do chan,ktorý dostala ako param., 
teraz vracia returnom Result a teda už chan nebol treba....,
potom som už iba dopísal funkciu WorkerPool.mainLoop(pridal som aj mutex),
ktorá vyrobí n*n(pocet prvkov v matici) goroutin, 
ktoré riešia "iba" jeden "malý" problém dostať Task do servera a zapísať výsledok do chanálu, 
robil som to cez Worker.doit a využíval som oba mutexy(postupne).....

ku všetkým:
nažil som sa urobiť všetko(všetky experimenty, čo tam sú[s kódom]) ak som niečo opomenul, 
tak sa ospravedlnujem, ale po 15 hod(do hromady) sa mi už na to nechce ani pozrieť............   
   


