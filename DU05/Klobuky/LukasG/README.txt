Spustenie simulácie klobúkov je pripravené priamo v Main.go. Úlohu som vylepšil a na základe prémie som ju parametrizoval N ľudmi, pričom 
vytvorená konfigurácia obsahuje maximálne (N-1) klobúkov (vtedy naša stratégia funguje). 

Úlohu som sa snažil naprogramovať pekne, objektovo (čo samozrejme vyžadovalo trocha sa vyhrať so smerníkmi, aby nám nevznikali nechcené kópie...). 
Život simulácie na najvyššej úrovni vidno v Hats.go - 1. vytvor simuláciu, 2. vypíš vygenerovanú konfiguráciu, 3. spusti simuláciu, 4. over správnosť.

######
Ukážka pre N=3 (pôvodné zadanie):
 
CONFIGURATION:
Agent id: 0, hat: WHITE 
Agent id: 1, hat: BLACK 
Agent id: 2, hat: WHITE 
-----
SIMULATION:
[0s] id: 1 - I see 0 BLACK hats and 2 WHITE hats 
[0s] id: 0 - I see 1 BLACK hats and 1 WHITE hats 
[0s] id: 2 - I see 1 BLACK hats and 1 WHITE hats 
[5s] id: 1 - No message, wait 5 seconds 
[5s] id: 0 - No message, wait 5 seconds 
[5s] id: 2 - No message, wait 5 seconds 
[10s] id: 2 - I see 1 BLACK hats and 1 WHITE hats 
[10s] id: 1 - I see 0 BLACK hats and 2 WHITE hats 
[10s] id: 2 - I have WHITE hat! 
[10s] id: 0 - I see 1 BLACK hats and 1 WHITE hats 
[10s] id: 0 - I have WHITE hat! 
[11s] id: 1 - Message from 0, he has WHITE hat 
[11s] id: 1 - I have BLACK hat! 
-----
VERIFICATION:
Agent id: 0, hat: WHITE, opinion: WHITE, correct: true 
Agent id: 1, hat: BLACK, opinion: BLACK, correct: true 
Agent id: 2, hat: WHITE, opinion: WHITE, correct: true 
Total faults: 0 
-----

Algoritmus živote jedného agenta korešponduje so stratégiou, ktorú som popísal v prémii "Klobúky z prednášky".
Pointa je:
 1. Pozriem, či náhodou nevidím maximálny možný počet čiernych klobúkov (pre N=3 na začiatku 2 po 10 sekundách 1 atd.). Ak áno, hlásim, že mám biely.
 2. Inak dám susedom priestor 5 sekúnd na to (kombinácia timeoutu a selectu), aby niečo zahlásili. Ak sa niekto ozve, znamená to, že si bol v predchádzajucom
    kroku istý, zatiaľ čo ja nie, takže musím mať čierny.
 3. Inak si dám sleep na 5 sekúnd (a so mnou aj ostatní), počas ktorých si premyslíme odpoveď na základnú otázku života, vesmíru a všetkého.

Ukážka pre N=7 (pre zaujímavosť):
######
CONFIGURATION:
Agent id: 0, hat: WHITE 
Agent id: 1, hat: WHITE 
Agent id: 2, hat: BLACK 
Agent id: 3, hat: BLACK 
Agent id: 4, hat: BLACK 
Agent id: 5, hat: WHITE 
Agent id: 6, hat: BLACK 
-----
SIMULATION:
[0s] id: 6 - I see 3 BLACK hats and 3 WHITE hats 
[0s] id: 0 - I see 4 BLACK hats and 2 WHITE hats 
[0s] id: 1 - I see 4 BLACK hats and 2 WHITE hats 
[0s] id: 2 - I see 3 BLACK hats and 3 WHITE hats 
[0s] id: 4 - I see 3 BLACK hats and 3 WHITE hats 
[0s] id: 3 - I see 3 BLACK hats and 3 WHITE hats 
[0s] id: 5 - I see 4 BLACK hats and 2 WHITE hats 
[5s] id: 1 - No message, wait 5 seconds 
[5s] id: 6 - No message, wait 5 seconds 
[5s] id: 0 - No message, wait 5 seconds 
[5s] id: 3 - No message, wait 5 seconds 
[5s] id: 2 - No message, wait 5 seconds 
[5s] id: 5 - No message, wait 5 seconds 
[5s] id: 4 - No message, wait 5 seconds 
[10s] id: 1 - I see 4 BLACK hats and 2 WHITE hats 
[10s] id: 6 - I see 3 BLACK hats and 3 WHITE hats 
[10s] id: 0 - I see 4 BLACK hats and 2 WHITE hats 
[10s] id: 5 - I see 4 BLACK hats and 2 WHITE hats 
[10s] id: 2 - I see 3 BLACK hats and 3 WHITE hats 
[10s] id: 4 - I see 3 BLACK hats and 3 WHITE hats 
[10s] id: 3 - I see 3 BLACK hats and 3 WHITE hats 
[15s] id: 6 - No message, wait 5 seconds 
[15s] id: 0 - No message, wait 5 seconds 
[15s] id: 3 - No message, wait 5 seconds 
[15s] id: 1 - No message, wait 5 seconds 
[15s] id: 4 - No message, wait 5 seconds 
[15s] id: 2 - No message, wait 5 seconds 
[15s] id: 5 - No message, wait 5 seconds 
[20s] id: 0 - I see 4 BLACK hats and 2 WHITE hats 
[20s] id: 0 - I have WHITE hat! 
[20s] id: 6 - I see 3 BLACK hats and 3 WHITE hats 
[20s] id: 2 - I see 3 BLACK hats and 3 WHITE hats 
[20s] id: 3 - I see 3 BLACK hats and 3 WHITE hats 
[20s] id: 4 - I see 3 BLACK hats and 3 WHITE hats 
[20s] id: 5 - I see 4 BLACK hats and 2 WHITE hats 
[20s] id: 5 - I have WHITE hat! 
[20s] id: 1 - I see 4 BLACK hats and 2 WHITE hats 
[20s] id: 1 - I have WHITE hat! 
[21s] id: 6 - Message from 0, he has WHITE hat 
[21s] id: 6 - I have BLACK hat! 
[21s] id: 2 - Message from 6, he has BLACK hat 
[21s] id: 2 - I have BLACK hat! 
[21s] id: 4 - Message from 0, he has WHITE hat 
[21s] id: 3 - Message from 0, he has WHITE hat 
[21s] id: 3 - I have BLACK hat! 
[21s] id: 4 - I have BLACK hat! 
-----
VERIFICATION:
Agent id: 0, hat: WHITE, opinion: WHITE, correct: true 
Agent id: 1, hat: WHITE, opinion: WHITE, correct: true 
Agent id: 2, hat: BLACK, opinion: BLACK, correct: true 
Agent id: 3, hat: BLACK, opinion: BLACK, correct: true 
Agent id: 4, hat: BLACK, opinion: BLACK, correct: true 
Agent id: 5, hat: WHITE, opinion: WHITE, correct: true 
Agent id: 6, hat: BLACK, opinion: BLACK, correct: true 
Total faults: 0 
-----