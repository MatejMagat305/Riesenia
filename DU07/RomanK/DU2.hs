-- Roman Kurcu - Haskell DU 2
-- Martinka, vysledky nabehnu tak do minutky, hlavne kvoli ulohe 4 (kruh), ostatne su rychle

import Data.List

-- 1
u235 :: Int -> Integer
u235 n = snd (splitAt 2 (sort ( concat[ [2^i,3^i, 5^i] | i <- [0 .. 5000] ] ))) !! (n - 1)
-- (5000, rezerva)


-- 3
determinant :: [[Int]] -> Int
determinant x = det x
-- Laplace
det :: (Num a, Ord a) => [[a]] -> a
det [[]] = error "chyba 1"
det x
    | lfex /= lx = error "chyba 2"
    | lfex == 2 && (length s) == 2    = a*d - b*c
    | otherwise = foldl (\acc y -> let ind = index y in acc + (sign ind) * f!!ind * (det $ genLaplaceMatrix x ind)) 0 [0 .. lx-1]
    where
          [[a,b],[c,d]] = x
          f:s:xr = x
          lx      = length x
          lfex    = length f
          sign y  = (-1)^(1+lx-y)
          index y = lx-1-y

genLaplaceMatrix :: (Num a) => [[a]] -> Int -> [[a]]
genLaplaceMatrix (_:x) pivotCol = foldr (\y -> (conVecs y:)) [] [0..((length x)-1)]
    where
      parts vec = splitAt pivotCol vec
      conVecs row = let y = x!!row in (fst $ parts y) ++ (tail $ snd $ parts y)
-- END Laplace

-- 4
kruh :: [[Int]]
kruh = kruh2 [1..10]

kruh2 :: [Int] -> [[Int]]
kruh2 [] = [[]]
kruh2 arr = [ x:y | x <- arr, y <- kruh2(arr\\[x]), kruh_test(x:y) ]

kruh_test :: [Int] -> Bool -- test set
kruh_test arr = if length arr == 10 then kruh_test2 arr && kruh_test2 [head arr, last arr] else kruh_test2 arr

kruh_test2 :: [Int] -> Bool -- test
kruh_test2 [] = True
kruh_test2 [x] = True
kruh_test2 (x:y:xs) = ((mod (x+y) 3) * (mod (x+y) 5) * (mod (x+y) 7)) > 0 && kruh_test2 (y:xs)

-- 5
-- riesil som to asi trocha nestandardne, no toto ma narychlo napadlo...
-- najprv beriem hodnoty [1] podla toho ako sa to postupne odblokovava a nasledne k tomu pripacnem hodnoty [0] a nakoniec iba zmazem duplikaty
-- [[0,1][0,1][0,1]...]

odblokuj :: [[Int]] -> [Int]
odblokuj arr = removeDuplicates ( odblokuj1 [] arr ++ odblokuj0 arr )

odblokuj0 :: [[Int]] -> [Int] -- vyzbiera vsetky hodnoty [1] z dvojic
odblokuj0 [] = []
odblokuj0 (a:arr) = [(a !! 0)] ++ odblokuj0 arr

odblokuj1 :: [Int] -> [[Int]] -> [Int] -- vyzbiera hodnoty [0] v poradi akom sa odblkovavaju
odblokuj1 arr arr2 = if ( length arr == (length arr2) ) then arr else odblokuj1 (arr ++ odblokuj2 arr2 arr2) (odb4 (odblokuj2 arr2 arr2) arr2)

odblokuj2 :: [[Int]] -> [[Int]] -> [Int]
odblokuj2 [] arr2 = []
odblokuj2 (x:arr) arr2 = if ((odblokuj3 (x !! 1) arr2) && (x !! 1 /= 0)) then [x !! 1] ++ odblokuj2 arr arr2 else odblokuj2 arr arr2

odblokuj3 :: Int -> [[Int]] -> Bool -- test ci je hodnota uz odblokovana
odblokuj3 x [] = True
odblokuj3 x (y:arr) = ((x /= (y !! 0)) || (y !! 1) == 0) && (odblokuj3 x arr)

odb4 :: [Int] -> [[Int]] -> [[Int]] -- tu potom prepisujem povodne [1] na hodnotu "0" co funkcia "odblokuj3" povazuje za odomknute
odb4 arr [] = []
odb4 arr (x:arr2) = if ( elem (x !! 1) arr ) then [[x !! 0, 0]] ++ odb4 arr arr2 else [[x !! 0, x !! 1]] ++ odb4 arr arr2

removeDuplicates = foldl (\seen x -> if x `elem` seen then seen else seen ++ [x]) [] -- vymazanie duplikatov...

-- 7
magickyStvorec :: [[Int]] -> Bool
magickyStvorec stvorec = and (map (==sucet) (map sum stvorec)) && and (map (==sucet) (map sum (transponuj stvorec))) && sucet == sum [ symetrickyStvorec !! i !! i | i <- [0..(length stvorec)-1] ]
  where
  sucet = sum [ stvorec !! i !! i | i <- [0..(length stvorec)-1]]
  symetrickyStvorec = map reverse stvorec

transponuj [] = []
transponuj ([]:xss) = transponuj xss
transponuj ((x:xs):xss) = (x:(map head xss)):(transponuj (xs:(map tail xss)))

-- TESTY
main = do
    putStrLn ("1 | u235 1 = " ++ show (u235 1))
    putStrLn ("1 | u235 10 = " ++ show (u235 10))
    putStrLn ("1 | u235 100 = " ++ show (u235 100))
    putStrLn ("1 | u235 1000 = " ++ show (u235 1000))
    putStrLn ("1 | u235 10000 = " ++ show (u235 10000))
    putStrLn ("")

    putStrLn ("3 | determinant [[1,-2,5,3],[-4,2,-3,-1],[2,-3,-4,5],[-3,0,2,0]] = " ++ show ( determinant [[1,-2,5,3],[-4,2,-3,-1],[2,-3,-4,5],[-3,0,2,0]] ))
    putStrLn ("")

    putStrLn ("4 | kruh = " ++ show (kruh))
    putStrLn ("")

    putStrLn ("5 | odblokuj [[2,9],[6,8],[5,1],[1,6],[4,5],[3,7],[5,3],[7,8],[8,9]] = " ++ show (odblokuj [[2,9],[6,8],[5,1],[1,6],[4,5],[3,7],[5,3],[7,8],[8,9]]))
    putStrLn ("")

    putStrLn ("7 | magickyStvorec [[16,3,2,13],[5,10,11,8],[9,6,7,12],[4,15,14,1]] = " ++ show (magickyStvorec [[16,3,2,13],[5,10,11,8],[9,6,7,12],[4,15,14,1]]))
    putStrLn ("7 | magickyStvorec [[15,3,2,13],[5,10,11,8],[9,6,7,12],[4,15,14,1]] = " ++ show (magickyStvorec [[15,3,2,13],[5,10,11,8],[9,6,7,12],[4,15,14,1]]))
    putStrLn ("")
