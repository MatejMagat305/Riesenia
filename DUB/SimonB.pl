%1. Magic number - 

magic2(S,X,Y) :- between(0, S, X), between(1, S, Y), X > Y, S is X + Y + X - Y + X * Y + div(X, Y).

%2. DONALD + GERALD = ROBERT 
/* riesenie
D = 5,
O = 2,
N = 6,
A = 4,
L = 8,
G = 1,
E = 9,
R = 7,
B = 3,
T = 0 ;
*/

alldiff([]).
alldiff([X|Xs]) :- not(member(X, Xs)), alldiff(Xs).

sumCol(Prenos, Cifra1, Cifra2, Cifra, NovyPrenos) :-
NovyPrenos is (Cifra1+Cifra2+Prenos) // 10,
Cifra is (Cifra1+Cifra2+Prenos) mod 10.

cifra(X) :- between(0, 9, X).

puzzle(D, O, N, A, L, G, E, R, B, T):-
cifra(D), sumCol(0, D, D, T, Pr1), alldiff([D, T]),
cifra(L), alldiff([D, T, L]), sumCol(Pr1, L, L, R, Pr2),
cifra(A), alldiff([D, T, L, R, A]), sumCol(Pr2, A, A, E, Pr3),
cifra(N), alldiff([D, T, L, A, R, E, N]), sumCol(Pr3, N, R, B, Pr4),
cifra(O), alldiff([D, T, L, A, R, E, N, B, O]), sumCol(Pr4, O, E, O, Pr5),
cifra(G), alldiff([D, T, L, A, R, E, N, B, O, G]), sumCol(Pr5, D, G, R, 0).

