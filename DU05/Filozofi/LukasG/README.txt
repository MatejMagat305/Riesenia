﻿Podobne ako s klobukmi, pokus o prehladnu, objektovu implementaciu vecerajucich filozofov. Simulacia je pre zaujimavost parametrizovatelna poctom filozofov.
Okrem poctu filozofov je druhym parametrom bool, ci sa ma pouzit false: naivna implementacia filozofa (mozny deadlock), alebo true: "sofistikovana".
Strategia je prevzata z Java appletu v poskytnutom .pdf. Hlavne vlakno caka na vstup od uzivatela, inak povedane, dvojity enter to zabije.

V Diners.go je konstanta DELAY, ktora spomaluje simulaciu, cize cim vacsia cislo, tym pomalsie. 

Ukazka naivneho pristupu v ktorom nastane DEADLOCK:
...
[450ms] Forks available: 2 
########
[461ms] 4 trying to pick LEFT fork... 
[471ms] 3 trying to pick RIGHT fork... 
[481ms] 2 put down both forks... 
[481ms] 2 is thinking... 
[481ms] 3 picked RIGHT fork...                                      <-- 3 ma pravu vidlicku
[481ms] 1 picked LEFT fork and is eating... 
########
[500ms] Forks available: 2 
########
[511ms] 3 trying to pick LEFT fork... 
########
[550ms] Forks available: 2 
########
[561ms] 2 trying to pick RIGHT fork... 
[561ms] 0 picked LEFT fork and is eating... 
[561ms] 1 put down both forks... 
[561ms] 1 is thinking... 
[561ms] 2 picked RIGHT fork...                                      <-- 2 ma pravu vidlicku
[591ms] 2 trying to pick LEFT fork... 
########
[600ms] Forks available: 2 
########
[631ms] 0 put down both forks... 
[631ms] 1 trying to pick RIGHT fork... 
[631ms] 1 picked RIGHT fork...                                      <-- 1 ma pravu vidlicku 
[631ms] 0 is thinking... 
[631ms] 4 picked LEFT fork and is eating... 
########
[650ms] Forks available: 2 
########
[662ms] 1 trying to pick LEFT fork... 
########
[700ms] Forks available: 1 
########
[701ms] 0 trying to pick RIGHT fork... 
[711ms] 4 put down both forks... 
[711ms] 4 is thinking... 
[711ms] 4 trying to pick RIGHT fork... 
[711ms] 4 picked RIGHT fork...                                      <-- 4 ma pravu vidlicku 
[711ms] 0 picked RIGHT fork...                                      <-- 0 ma pravu vidlicku 
[741ms] 0 trying to pick LEFT fork... 
[741ms] 4 trying to pick LEFT fork... 
########
[750ms] Forks available: 0 
########
...
A vsetci cakaju na lavu :(

