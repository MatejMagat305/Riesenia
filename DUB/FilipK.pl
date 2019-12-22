magic2(S,X,Y):-between(1,S,X),between(1,S,Y),SU is X + Y, RO is abs(X - Y),SUC is X * Y, D is X div Y, SU + RO + SUC + D =:=S.

/*DONALD + GERALD = ROBERT*/
d2(X,Y,Z):-solve("DONALD","GERALD","ROBERT",X,Y,Z).
/*trva to 50s ale najde to a program je vseobecny pre lubovolny algebrogram staci zmenit na konci operciu +-*/
/*algebrogram solver*/
/********************************************/
strToChars(A,B,C,X,Y,Z):-string_to_list(A,X),string_to_list(B,Y),string_to_list(C,Z).

intToChars(A,B,C,X,Y,Z):-number_string(A,AA),number_string(B,BB),number_string(C,CC),strToChars(AA,BB,CC,X,Y,Z).
toInt([],0).
toInt([X|XS],A):-toInt(XS,B),length(XS,N),A is B + ((10^N)*X).

get(X,[X|_],[Y|_],Y).
get(X,[Z|XS],[_|YS],Y):-X\=Z, get(X,XS,YS,Y).

map([],_,_,[]).
map([X|XS],Y,Z,[XX|AKU]):-get(X,Y,Z,XX),map(XS,Y,Z,AKU).

%vbo(N, Base, V) toto som si pozical z cvika
vbo(0, _, []).    % variacia dlzky 0 z lubovolne mnoziny je []
vbo(N, Base, [X|Xs]) :- N>0, N1 is N-1, select(X,Base, Rest), vbo(N1, Rest, Xs).

/*Mapujem pismenka na cisla a potom skusam ci plati rovnost A+B=C*/
solve(D,G,R,A,B,C):-strToChars(D,G,R,DD,GG,RR),append(DD,GG,H),append(RR,H,HH),list_to_set(HH,PIS),length(PIS,N),vbo(N,[0,1,2,3,4,5,6,7,8,9],Z),map(DD,PIS,Z,AA),map(GG,PIS,Z,BB),map(RR,PIS,Z,CC),toInt(AA,A),toInt(BB,B),toInt(CC,C),A+B=:=C.
/*******************************************/
count(0,_,S,S).
count(X,Z,S,N):-X>0, X1 is div(X,10), XM  is mod(X,10),(Z=XM -> (S1 is S+1) ; (S1 is S )),count(X1,Z,S1,N).
count(X,Z,N):-count(X,Z,0,N).

sub(C,SO,_,9):-C is SO*10, count(C,9,0).
sub(C,SO,S,K):-K<9,between(0,9,X),10>=S+X,K1 is K+1,S1 is S+X,SO1 is SO *10 + X,SO1 > 0,sub(C,SO1,S1,K1), count(C,K,X).
