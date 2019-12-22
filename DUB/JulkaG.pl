%- ULOHA 1

%- magic2(S,X,Y) :- between(1, S, X), between(1, S, Y), abs(X-Y, ABS), S is (X + Y) + (X*Y) + ABS + (X//Y).

%- magic2(S,X,Y) :- between(1, S, X), between(1, S, Y), X1 is max(X,Y), Y1 is min(X,Y), S is (X + Y) + (X*Y) + (X1-Y1) + (X1//Y1).

magic2(S,X,Y) :- between(1, S, X), between(1, S, Y), X > Y, S is (X + Y) + (X*Y) + (X-Y) + (X//Y).

/*
POZNAMKA: 
Kedze som si nebola ista, ako sa to ma spravat pre celociselne delenie, spravila som 3 verzie.
1. verzia: Dovoluje, aby X < Y a deli ich v tomto poradi, takze vysledok celociselneho delenia moze byt aj 0.
2. verzia: Dovoluje, aby X < Y, ale deli vzdy vacsie mensim, takze vysledok celociselneho delenia nie je nikdy 0.
3. verzia: Definuje, ze X musi by vacsie ako Y, inak vrati false.
*/



%- ULOHA 2


algebrogram(D, O, N, A, L, G, E, R, B, T) :- 
	cifra(D), T is (2*D) mod 10, T \= D, Pr1 is (2*D)//10,
	cifra(L), L\=D, L\=T, R is (2*L + Pr1) mod 10, R\=D, R\=T, R\=L, Pr2 is (2*L + Pr1)//10,
	cifra(A), A\=D, A\=T, A\=L, A\=R, E is (2*A + Pr2) mod 10, E\=D, E\=T, E\=L, E\=R, E\=A, Pr3 is (2*A + Pr2)//10,
	cifra(N), N\=D, N\=T, N\=L, N\=R, N\=A, N\=E, B is (N + R + Pr3) mod 10, B\=D, B\=T, B\=L, B\=R, B\=A, B\=E, B\=N, Pr4 is (N + R + Pr3)//10,
	cifra(O), O\=D, O\=T, O\=L, O\=R, O\=A, O\=E, O\=N, O\=B, O is (O + E + Pr4) mod 10, Pr5 is (O + E + Pr4)//10,
	cifra(G), G\=D, G\=T, G\=L, G\=R, G\=A, G\=E, G\=N, G\=B, G\=O, R is (D + G + Pr5) mod 10, 0 is (D + G + Pr5)//10,
	write(' '),write(D),write(O),write(N),write(A),write(L),write(D), nl ,write('+'),write(G),write(E),write(R),write(A),write(L),write(D), nl ,write('='),write(R),write(O),write(B),write(E),write(R),write(T), nl ,cifra(0).  

cifra(X) :- between(0, 9, X).


%- ULOHA 3


mliecne(X) :- member(X, [jogurt,mlieko,syr,acidko,tvaroh,bryndza,cokolada]).

pocetMliecnych([], 0).
pocetMliecnych([X|XS], P) :- mliecne(X), pocetMliecnych(XS, P1), P is P1 + 1.
pocetMliecnych([X|XS], P) :- not(mliecne(X)), pocetMliecnych(XS, P).

kbo(0, _, []).
kbo(K, [X|Xs], [X|Ys]) :- K > 0, K1 is K-1, kbo(K1, Xs, Ys).
kbo(K, [_|Xs], Ys) :- K > 0, kbo(K, Xs, Ys).

kso(0, _, []).
kso(K, [X|Xs], [X|Ys]) :- K > 0, K1 is K-1, kso(K1, [X|Xs], Ys).
kso(K, [_|Xs], Ys) :- K > 0, kso(K, Xs, Ys).

vyklad(V) :- between(0,16,K), kbo(K, [jogurt,mlieko,chlieb,pecivo,sol,cukor,muka,syr,acidko,vino,pivo,palenka,mineralka,tvaroh,bryndza,cokolada],V), pocetMliecnych(V, M), length(V, L), L < 2*M. 

vyklad3(V) :- vyklad(V), length(V, L), L > 2.

jeho_vyklad(V) :- between(0,7,K), kso(K, [jogurt,mlieko,chlieb,pecivo,sol,cukor,muka,syr,acidko,vino,pivo,palenka,mineralka,tvaroh,bryndza,cokolada],V), pocetMliecnych(V, M), length(V, L), L < 2*M, L < 8. 

%- riesenie predpoklada, ze vyklad nemoze byt prazdny