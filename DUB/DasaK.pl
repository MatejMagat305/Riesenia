%- Domaca uloha B (DU11)

%- pr01
magic2(S,X,Y) :-    between(1, S, X), between(1, S, Y),
                    S1 is X + Y,
                    S2 is abs(X - Y),
                    S3 is X * Y, 
                    S4 is div(X, Y),
                    S =:= (S1 + S2 + S3 + S4).
                    

%- pr02
%- DONALD + GERALD = ROBERT

%- [List] -> cislo 
naCislo(X, N) :- reverse(X, Y), naCislo2(Y, 1, N).
naCislo2([], _, 0).
naCislo2([X|Xs], C, N) :- 
    C1 is C * 10, naCislo2(Xs, C1, N1), N is N1 + (C * X).
    
parne(X) :- 0 is mod(X, 2).

puzzle(X, Y):-
    X = [D, O, N, A, L, D], Y = [G, E, R, A, L, D],
    Z = [R, O, B, E, R, T],
    select(D, [0,1,2,3,4,5,6,7,8,9], P), 
    select(G, P, P1), D + G =< 9, 
    select(R, P1, P2),
    select(T, P2, P3), parne(T), T is mod(2*D, 10), 
    select(A, P3, P4),
    select(L, P4, P5),
    naCislo([L, D], LD), naCislo([R, T], RT),
    RT is mod(2*LD, 100),
    select(O, P5, P6), 
    select(N, P6, P7),
    select(E, P7, P8),
    select(B, P8, _),
    naCislo(X, K), naCislo(Y, KK), naCislo(Z, KKK), K + KK =:= KKK.

%- pr04
%- 3ludia, 1 VO, 2 MO
                    
%- pr05
%- ci sa pocet cifier v cisle rovna N
count([], _, 0). 
count([X|Xs], Cifra, N) :-
    X = Cifra -> count(Xs, Cifra, N1), N is N1+1;
    count(Xs, Cifra, N).
    
%- sucet: cifra * hodnota <= 10
check(X):- sucet(0, X, 0).
sucet(_, [], _).
sucet(K, [X|Xs], S) :- number(K), number(S),
    K1 is K + 1, S1 is (S + (K * X)), S1 =< 10, sucet(K1, Xs, S1).
  
%- hladanie cisla tak aby cislo splnalo podmienku, ze prvá cifra zľava 
%- určuje počet všetkých núl v čísle,  druhá cifra počet jednotiek, 
%- tretia počet dvojek, atd   
cislo(X) :- 
    X = [X0, X1, X2, X3, X4, X5, X6, X7, X8, X9],
    between(1, 6, X0), between(0, 5, X1), between(0, 4, X2),
    between(0, 3, X3), between(0, 2, X4), between(0, 1, X5), 
    between(0, 1, X6), between(0, 1, X7), between(0, 1, X8), 
    between(0, 1, X9), check(X),
    count(X, 0, X0), count(X, 1, X1), count(X, 2, X2),
    count(X, 3, X3), count(X, 4, X4), count(X, 5, X5),
    count(X, 6, X6), count(X, 7, X7), count(X, 8, X8),
    count(X, 9, X9).

%- cislo(X).
%- X = [6,2,1,0,0,0,1,0,0,0] 