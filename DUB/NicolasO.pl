
%1. uloha
magic2(S,X,Y):-
    S is (X+Y + X-Y + X*Y + X/Y).

%2.uloha
cifra(1).
cifra(2).
cifra(3).
cifra(4).
cifra(5).
cifra(6).
cifra(7).
cifra(8).
cifra(9).

s(D,O,N,A,L,G,E,R,B,T):-
    cifra0(D),
    T is (D+D) mod 10,
    T\=D,
    Pr1 is (D+D) // 10,
    cifra0(L),L\=D,L\=T,
    R is (L+L+Pr1) mod 10,
    L\=R,R\=T,R\=D,
    Pr2 is (L+L+Pr1) // 10,
    cifra0(A), A\=D,A\=T,A\=L,A\=R,
    E is (A+A+Pr2) mod 10,
    E\=A,E\=L,E\=D,E\=T,E\=R,
    Pr3 is (A+A+Pr2) // 10,
    cifra0(N), N\=D,N\=T,N\=L,N\=R,N\=A,N\=E,
    B is (N+R+Pr3) mod 10,
    B\=D,B\=T,B\=L,B\=R,B\=A,B\=E,B\=N,
    Pr4 is (N+R+Pr3) //10,
    cifra0(O), O\=D,O\=T,O\=L,O\=R,O\=A,O\=E,O\=B,O\=N,
    O is (O+E+Pr4) mod 10,
    Pr5 is (O+E+Pr4) // 10,
    cifra0(G), G\=0, G\=D, G\=T, G\=L, G\=R, G\=A, G\=E, G\=B, G\=O,G\=N,
    R is (D+G+Pr5),
    write(' '),write(D),write(O),write(N),write(A),write(L),write(D), nl,
    write('+'),write(G),write(E),write(R),write(A),write(L),write(D), nl,
    write(' '),write(R),write(O),write(B),write(E),write(R),write(T),nl.

cifra0(0).
cifra0(X):-cifra(X).

%5.uloha
%zdroj: https://github.com/paradigmy/Kod/blob/master/CV12/CvProlog3_Puzzle.pl

solve(Riesenie) :- Riesenie = [P0,P1,P2,P3,P4,P5,P6,P7,P8,P9],
	between(0,9,P0), check([P0]),
	between(0,9,P1), check([P0,P1]),
	between(0,9,P2), check([P0,P1,P2]),
	between(0,9,P3), check([P0,P1,P2,P3]),
	between(0,9,P4), check([P0,P1,P2,P3,P4]),
	between(0,9,P5), check([P0,P1,P2,P3,P4,P5]),
	between(0,9,P6), check([P0,P1,P2,P3,P4,P5,P6]),
	between(0,9,P7), check([P0,P1,P2,P3,P4,P5,P6,P7]),
	between(0,9,P8), check([P0,P1,P2,P3,P4,P5,P6,P7,P8]),
	between(0,9,P9), check([P0,P1,P2,P3,P4,P5,P6,P7,P8,P9]),
	count(0,Riesenie,P0),
	count(1,Riesenie,P1),
	count(2,Riesenie,P2),
	count(3,Riesenie,P3),
	count(4,Riesenie,P4),
	count(5,Riesenie,P5),
	count(6,Riesenie,P6),
	count(7,Riesenie,P7),
	count(8,Riesenie,P8),
	count(9,Riesenie,P9).

count(_,[],0).
count(X,[Y|Xs],N1):-X=Y->count(X,Xs,N),N1 is N+1;count(X,Xs,N1).

check(L):-check2(L,0,0).


check2([], _, _).
check2([H|T],C,I):-I1 is I+1,CC is C+I*H,CC=<10,check2(T,CC,I1).

%iba jedno riesenie [6,2,1,0,0,0,1,0,0,0].

%4.uloha

init(state([1,1,1],[1,1],l)).
final(state([0,0,0],[0,0],r)).

%v lodke je vzdy velka opica

next(state([1,C2,C3],[O2,O3],l),state([0,C2,C3],[O2,O3],r)).
next(state([C1,1,C3],[O2,O3],l),state([C1,0,C3],[O2,O3],r)).
next(state([C1,C2,1],[O2,O3],l),state([C1,C2,0],[O2,O3],r)).

next(state([C1,C2,C3],[1,O3],l),state([C1,C2,C3],[0,O3],r)).
next(state([C1,C2,C3],[O2,1],l),state([C1,C2,C3],[O2,0],r)).

%ak pojd velka opica sama v lodke
next(state([C1,C2,C3],[O2,O3],l),state([C1,C2,C3],[O2,O3],r)).

next(state(Ms,Ws,r),state(Ms1,Ws1,l)):-next(state(Ms1,Ws1,l),state(Ms,Ws,r)).

bankOk(state([M1,M2,M3],[W2,W3],_)):-M1+M2+M3 >= W2+W3.


rightBank(state([M1,M2,M3],[W2,W3],_),state([RM1,RM2,RM3],[RW2,RW3],_)):-
          RM1 is 1-M1,RM2 is 1-M2,RM3 is 1-M3,
          RW2 is 1-W2,RW3 is 1-W3.

banksOk(Left):-bankOk(Left),
               rightBank(Left,Right),
               bankOk(Right).

cesta(X,X,P,P).
cesta(X,Y,Visited,P):-
  next(X,Z),
  not(member(Z,Visited)), banksOk(Z),
  cesta(Z,Y,[Z|Visited],P).


hladaj(L):-init(Z),final(K),cesta(Z,K,[Z],P),length(P,L),reverse(P,P1), write(P1).














