import Data.List
import Data.Maybe

-- uloha 1
mn235 n =  sort (mnozina n 4 1 1 [1,2,3,4,5,8,9,16,25]) !! (n) 

mnozina n i j k xs = if length xs > n then xs
    else mnozina n (snd pow2) (snd pow3) (k + 1) novePole
    where pow3 = threePowJ xs j (5 ** (k + 1))
          pow2 = twoPowI (fst pow3) i (5 ** (k + 1))
          novePole = ((5 ** (k + 1)):((fst pow2)))
		  
twoPowI xs i posledny = if 2 ** (i + 1) > posledny then (xs, i)
    else twoPowI ((2 ** (i + 1)):xs) (i + 1) posledny
	
threePowJ xs j posledny = if 3 ** (j + 1) > posledny then (xs, j)
    else threePowJ ((3 ** (j + 1)):xs) (j + 1) posledny

{- 
- n je dlzka pola a zaroven index do pola
- mn235 1000 = 9.989595361011175e145
- mn235 1000 = 9.795789456962752e1459
-}

-- uloha 3 - Determinant
miniMat mat stlpec = vynechajPrvok [(vynechajPrvok riadok stlpec) | riadok <- mat] 0
vynechajPrvok pole index = [pole!!i | i <- [0.. (length pole) - 1], i /= index]

type Matica = [[Int]]
determinant :: Matica -> Int
determinant [[]] = 0
determinant [[x]] = x
determinant [[x,y], [z,q]] = x*q - y*z
determinant xs = sum [((xs!!0)!!k) * ((-1)^k) * determinant (miniMat xs k) | k <- [0.. (length (xs) - 1)]]

-- uloha 4
kruh = backtrack [1..10]

nedel x = x `rem` 3 /= 0 && x `rem` 5 /= 0  && x `rem` 7 /= 0

--backtrack :: [Int] -> [[Int]]
backtrack [] = [[]]
backtrack cifry  = [[x] ++ y | 
            x <- cifry, 
            y <- backtrack (diff cifry [x]),
            jeOK ([x] ++ y)]

jeOK :: [Int] -> Bool
jeOK [] = True
jeOK [_] = True
jeOK (x:y:xs) = if ((length xs) + 2)  < 10 && nedel (x + y) then True --jeOK (y:xs)
    else if (length xs) + 2  == 10 && nedel (x + y) && nedel ((last xs) + x) then True --jeOK(y:xs)
    else False

diff :: (Eq a) => [a] -> [a] -> [a]
diff x y = [ z | z <- x, notElem z y] 

-- uloha 5 - EU OS
odblokuj blok = [sol | sol <- (backtrack2 (sort (zisti blok)) blok), skontroluj sol blok]!!0
-- riesenie je pole Int a vyberam prve riesenie

zisti proces = nub ([p!!0 | p <- proces] ++ [p!!1 | p <- proces])

backtrack2 [] blok = [[]]
backtrack2 cifry blok  = [(x:y) | 
            x <- cifry, 
            y <- backtrack2 (diff cifry [x]) blok]

skontroluj pole blokacie = and [(skorAko (b!!0) (b!!1) (pole)) | b <- blokacie ]

skorAko x y pole = (elemIndex x pole) > (elemIndex y pole)  

-- uloha 7
magickyStvorec :: [[Int]] -> Bool
magickyStvorec xs = length (nub ([sum (xs!!i) | i <- [0.. (length xs) - 1]] ++  -- pole suctov riadkov
    [sum [(xs!!i)!!(((length xs) - 1) - i) | i <- [0.. (length xs) - 1]]] ++ -- pole suctu uhlopriecky druhej
    [sum [xs!!i!!i | i <- [0.. (length xs) - 1]]] ++ -- pole suctu uhlopriecky prvej
    [scitajStlpec xs i | i <- [0.. (length xs) - 1]])) == 1 -- pole suctov stlpcov

scitajStlpec xs j = sum [xs!!i!!j | i <- [0.. (length xs) - 1]]