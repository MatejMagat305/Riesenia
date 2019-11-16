import Data.List

type Kocka = [Int]

isRightHanded :: Kocka -> Bool 
isRightHanded [a,b,c,d,e,f] = case a of
  1 -> (assemble [b,c,d,e] 2) == [2, 3, 5, 4]
  2 -> (assemble [b,c,d,e] 1) == [1, 4, 6, 3]
  3 -> (assemble [b,c,d,e] 1) == [1, 2, 6, 5]

  4 -> (assemble [b,c,d,e] 1) /= [1, 2, 6, 5]
  5 -> (assemble [b,c,d,e] 1) /= [1, 4, 6, 3]
  6 -> (assemble [b,c,d,e] 2) /= [2, 3, 5, 4]

indexOf :: Kocka -> Int -> Int -> Int
indexOf k g i 
   | (k !! i) == g = i 
   | otherwise = indexOf k g (i+1)

assemble :: Kocka -> Int -> Kocka
assemble k g = (k !! i : [ k !! (j `mod` 4) | j <- [i+1..i+3]]) where i = indexOf k g 0

---

moznosti :: [Kocka]
moznosti = [[a,b,c,7-b,7-c,7-a] | a<-[1..6], b<-[1..6], b/=a, b/=7-a, c<-[1..6], allDifferent [a,b,7-a,7-b,c]]

remove :: Eq(t) => t -> [t] -> [t]
remove e xs = filter (/= e) xs

allDifferent :: Eq(t) => [t] -> Bool
allDifferent xs = and [(length(xs) - 1) == length(remove x xs) | x <- xs]

allRightHanded :: [Kocka]
allRightHanded = filter isRightHanded moznosti

kockySuOK :: [Kocka] -> Bool
kockySuOK ks = and [allDifferent [k!!i | k <- ks] | i <- [1..4]]

naSebe :: Int -> [Kocka] -> [[Kocka]]
naSebe 0 _ = [[]]
naSebe k polohy = [(p:m) | m <- naSebe (k-1) polohy, p <- polohy, kockySuOK(p:m)]


--------------------------------------------------

data Farba = R | B | G | Y | O | C deriving (Show, Eq)
type FKocka = [Farba]

type Farbenie = [Farba]

farbenie :: Kocka -> Farbenie -> FKocka
farbenie xs fb = [fb!!(i-1) | i <-xs]

farbySuOK :: [FKocka] -> Bool
farbySuOK ks = and [allDifferent [k!!i | k <- ks] | i <- [1..4]]

permutacie :: Eq(t) => [t] -> [[t]]
permutacie [] = [[]]
permutacie xs = [(x:ys)| x <- xs, ys <- permutacie (xs \\ [x])]

-- :set +s: meranie casu funkcii v ghci

farebneNaSebe :: [Int] -> [[FKocka]] -> [[FKocka]]
farebneNaSebe [] _ = [[]]
farebneNaSebe indexyKociek polohyKociek = [(k:z) | perm <- (permutacie indexyKociek), 
                                                   z <- farebneNaSebe (indexyKociek \\ [perm !! 0]) polohyKociek, 
                                                   k <- polohyKociek !! (perm !! 0), 
                                                   farbySuOK(k:z)]

polohyUloha :: [FKocka] -> [[FKocka]]
polohyUloha farbs = [[farbenie ks fs | ks <- allRightHanded] | fs <- farbs]

-- V ulohe NIE JE napisane, ze mame najst vsetky mozne postavenia, takze staci zobrat jedno riesenie z backtracku?...
-- Farbenie som z obrazkov zmenil tak, aby kocky sedeli na plast pravotocivej kocky [1,2,3,5,4,6]
-- aby sme mohli vychadzat z allRightHanded...

--- U1 ---
farbenieU1 = [[B,R,R,G,R,Y], [R,G,Y,B,G,B], [R,B,G,Y,R,Y], [G,B,R,G,Y,Y]]

-- farebneNaSebe [0..3] (polohyUloha farbenieU1)
-- (farebneNaSebe [0..3] (polohyUloha farbenieU1)) !! 0 == [[R,R,B,G,Y,R],[G,Y,R,B,B,G],[Y,B,Y,R,G,R],[B,G,G,Y,R,Y]] (prikladam fotky :))
-- length (farebneNaSebe [0..3] (polohyUloha farbenieU1)) == 2304; (72.83 secs, 14,815,030,696 bytes)
----------

--- U2 ---
farbenieU2 = [[B,R,Y,G,R,Y], [G,G,Y,G,Y,B], [R,B,G,Y,B,R], [G,G,R,G,Y,Y]]

-- (farebneNaSebe [0..3] (polohyUloha farbenieU2)) !! 0 == [[Y,Y,R,B,R,G],[G,B,Y,G,G,Y],[G,R,B,R,B,Y],[R,G,G,Y,Y,G]]
-- V tomto momente som sa rozhodol, ze render spravim rychlejsie ako olepim kocku farebnymi papierikmi... (prikladam rendery)
-- length (farebneNaSebe [0..3] (polohyUloha farbenieU2)) == 73728; (125.19 secs, 25,274,851,192 bytes)
----------

--- U3 ---
-- Vzhladom na to, ze max su 3 body, idem rovno na najvyssi kaliber...
----------

--- U4 ---
farbenieU4 = [[O,C,B,Y,G,R], [O,C,B,Y,R,G], [O,G,B,R,Y,C], [O,G,B,R,C,Y], [O,Y,B,C,R,G], [O,R,B,G,Y,C]]

-- (farebneNaSebe [0..5] (polohyUloha farbenieU4)) !! 0 == [[B,G,O,C,R,Y],[G,Y,R,B,C,O],[R,C,G,O,Y,B],[C,B,Y,R,O,G],[Y,O,C,G,B,R],[O,R,B,Y,G,C]] 
-- (kto neveri, nech pozrie render)
-- length (farebneNaSebe [0..5] (polohyUloha farbenieU4)) == ... To bude na dlho.
-- V ulohe NIE JE napisane, ze mame najst vsetky mozne postavenia, takze staci zobrat jedno riesenie z backtracku?...
----------
