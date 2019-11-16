import Data.List

--program a riesnie by bolo ale kocky nemam ;p

--  a
-- bcde
--  f
data List a = Nil | Cons a (List a) deriving (Eq)
data Farba = R | B | G | Y |W | O | C deriving (Show, Eq)

vpravo [a,b,c,d,e,f] 0 = [a,b,c,d,e,f]
vpravo [a,b,c,d,e,f] n = vpravo [a,e,b,c,d,f] (n-1)

dole [a,b,c,d,e,f] 0 = [a,b,c,d,e,f]
dole [a,b,c,d,e,f] n = dole [e,b,a,d,f,c] (n-1)

vsetky x = nub([vpravo i j| i<-[dole x a| a<-[0..3]],j<-[0..3]] ++ [vpravo i j | i<-[(dole (vpravo x 1) 3),(dole (vpravo x 1) 1)],j<-[0..3]])

test x = and[length (nub [i!!j| i<-x]) == (length x)|j<-[1..4]]

postav:: [[Farba]]->[[[Farba]]]
postav x| length x ==1  = [[i]|i<-(vsetky (x!!0))]
        | otherwise = [j : i| i<-(postav (tail x) ),j<-(vsetky (head x)), test (j : i) ]


--head $ postav [[B,R,R,R,G,Y],[R,G,Y,G,B,B],[R,B,G,R,Y,Y],[G,B,R,Y,G,Y]]
-- je:[[R,B,G,Y,R,R],[G,R,B,B,Y,G],[Y,Y,R,G,B,R],[B,G,Y,R,G,Y]] a pocet je 8

--head $ postav [[B,R,Y,R,G,Y],[G,G,Y,Y,G,B],[R,B,G,B,Y,R],[G,G,R,Y,G,Y]]
-- je:[[G,R,Y,R,B,Y],[G,Y,B,G,G,Y],[Y,B,R,B,R,G],[G,G,G,Y,Y,R]] a pocet je 128

-- head $ postav [[O,C,O,R,G,Y],[G,R,G,Y,O,C],[R,G,R,O,Y,C],[C,G,C,R,O,Y],[Y,R,Y,O,G,C]]
-- je: [[Y,O,C,G,R,O],[G,G,R,C,Y,O],[R,Y,G,R,O,C],[G,C,O,Y,C,R],[Y,R,Y,O,G,C]] a pocet je 24

-- head $ postav [[O,C,B,G,Y,R],[O,C,B,R,Y,G],[O,G,B,Y,R,C],[O,G,B,C,R,Y],[O,Y,B,R,C,G],[O,R,B,Y,G,C]]
-- je:[[B,G,O,C,R,Y],[C,O,Y,G,B,R],[Y,B,C,R,O,G],[R,Y,G,O,C,B],[G,C,R,B,Y,O],[O,R,B,Y,G,C]] a pocet je 88