import Data.List
--import Data.Set 

-- uloha1 --
sudelitelne :: Int -> Int -> Bool
sudelitelne a 0 = True
sudelitelne 0 b = True
sudelitelne a b = length [ d | d <- [2..c], (mod a d) == 0, (mod b d) == 0 ] > 0
	where c = min (abs a) (abs b)
	
-- uloha 2 --
comb :: Int -> Int -> Int
comb n 0 = 1
comb n k = if n < 0 || k < 0 then 0 else (product [1..n]) `div` ((product [1..k]) * (product [1..(n-k)]))

-- uloha 3 --
jeGeom :: [Int] -> Bool
jeGeom [] = False 
jeGeom [x] = False
jeGeom [x,y] = x /= 0 && y /= 0		-- predpokladam, ze ani postupnost nul nie je geom, kedze nevieme jednoznacne urcit kvocient
jeGeom (x:s@(y:z:xs)) = x /= 0 && y /= 0 && fromIntegral y / fromIntegral x == fromIntegral z / fromIntegral y && jeGeom(s)
-- POZNAMKA: Nie som si ista, co to ma robit pre prazdny a jednoprvkovy 
--			zoznam - dala som ze false, lebo predpokladam, ze na to, aby
--			sa nieco vobec nezyvalo "postupnost" treba aspon 2 prvky 

-- uloha 4 --
jePrefix :: [Int] -> [Int] -> Bool 
jePrefix [] _ = True
jePrefix _ [] = False
jePrefix (x:xs) (y:ys) = x == y && jePrefix xs ys 

-- uloha 5 --
int2Bin :: Int -> [Int]
int2Bin 0 = []
int2Bin n = int2Bin (n `div` 2) ++ [mod n 2] 
-- POZNAMKA: Dufam, ze to staci riesit pre cisla > 0

-- uloha 6 --
median :: [Float] -> Float
median [x] = x
median xs = (sort xs)!!(length xs `div` 2)

-- uloha 7 --
rotuj :: Int -> [a] -> [a]
rotuj 0 xs = xs
rotuj _ [] = []
rotuj n xs = rotuj (n-1) ((last xs):(init xs))

-- uloha 8 --
treti :: [Int] -> Int
treti xs = if length xs < 3 then 0 else (sort xs)!!2
-- POZNAMKA: tato funkcia vrati treti prvok v poradi bez ohladu 
--				na to, ci sa hodnoty pred nim opakuju; ak ma byt
--				treti najvacsi len z hladiska hodnoty (cize
--				opakujuce sa prvky sa nezucastnuju usporiadania),
--				tato definicia to zabezpeci (treba odkomentovat 
--				import Data.Set):
--treti xs = if length xs < 3 then 0 else (sort (toList(fromList xs)))!!2		

-- uloha 9 --
prienik :: [Int] -> [Int] -> [Int]
prienik [] _ = []
prienik _ [] = []
prienik (x:xs) ys = if x `elem` ys then x:(prienik xs ys) else prienik xs ys
-- POZNAMKA: tato funckia predpoklada, ze prvky vo vstupnych zoznamoch sa 
--				neopakuju; ak sa mozu opakovat, tato definicia to zabezpeci
--				(treba odkomentovat import Data.Set):
--prienik (x:xs) ys = if x `elem` ys then toList(fromList(x:(prienik xs ys))) else prienik xs ys