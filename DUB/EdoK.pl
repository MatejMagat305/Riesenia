% # ----------------------------------------------------------------------------
% #1. Magic number
magic2(S,X,Y) :- 
	between(0,S, X),
	between(0,S,Y),
	S is X * Y.

magic2(S,X,Y) :- 
	between(0,S, X),
	between(0,S,Y),
	X > 0,
	X < 3000,
	Y > 0,
	Y < 300,
	S is X / Y.

magic2(S,X,Y) :- 
	between(0,S, X),
	between(0,S,Y),
	S is X - Y.

magic2(S,X,Y) :- 
	between(0,S, X),
	between(0,S,Y),
	S is X + Y.

% # ----------------------------------------------------------------------------
% # 3. Výklad - aplikovaná kombinatorika
:- use_module(library(lists)).

mliecne(jogurt).
mliecne(mlieko).
mliecne(syr).
mliecne(acidko).
mliecne(tvaroh).
mliecne(bryndza).

comb(0,_,[]).
comb(N,[X|T],[X|Comb]) :-N>0, N1 is N-1, comb(N1,T,Comb).
comb(N,[_|T],Comb) :-  N>0,  comb(N,T,Comb).

% # mam zo stackoverflow, ale link som zabudol
permutInListN(List,N,Result) :- comb(N,List,Comb), permutation(Comb,Result).

% # count milk elements in list
countMilk([],0).
countMilk([X|T],N) :- mliecne(X), countMilk(T,N1), N is N1 + 1.
countMilk([X|T],N) :- not(mliecne(X)), countMilk(T,N1), N is N1 + 0.

size([],0).
size([_|T],N):- size(T,M), N is M+1.

kso(0, _, []).
kso(K, [X|Xs], [X|Ys]) :- K > 0, K1 is K-1, kso(K1, [X|Xs], Ys).
kso(K, [_|Xs], Ys) :- K > 0, kso(K, Xs, Ys).

vyklad(V) :- between(0,16,X),
	permutInListN([jogurt,mlieko,chlieb,pecivo,sol,cukor,muka,syr,acidko,vino,pivo, palenka,mineralka,tvaroh,bryndza,cokolada],X , V),
	size(V, R1), countMilk(V , R2), R2 > R1/2.

vyklad3(V) :- between(0,16,X),
	permutInListN([jogurt,mlieko,chlieb,pecivo,sol,cukor,muka,syr,acidko,vino,pivo, palenka,mineralka,tvaroh,bryndza,cokolada],X , V),
	size(V, R1), R1 >= 3, countMilk(V , R2), R2 > R1/2.

jeho_vyklad(V) :- between(0,16,X),
	kso(X , [jogurt,mlieko,chlieb,pecivo,sol,cukor,muka,syr,acidko,vino,pivo, palenka,mineralka,tvaroh,bryndza,cokolada] , V),
	size(V, R1), 7 >= R1, countMilk(V , R2), R2 > R1/2.


