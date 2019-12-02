module Nahrdelniky where

-- 3. Náhrdelníky

import Data.List

vsetky :: Int -> [String] -> [[String]]
vsetky n xs | n == 0 = [[]]
            | otherwise = [color:others | color <- xs, others <- vsetky (n-1) xs]

rotuj :: Int -> [a] -> [a]
rotuj 0 xs = xs
rotuj n xs = rotuj (n-1) (last xs : init xs)

uzJe :: [String] -> [[String]] -> Bool
uzJe xs ys = or [rotuj n xs `elem` ys | n <- [0..length xs - 1]]

dajNerovnake :: [[String]] -> [[String]]
dajNerovnake [] = []
dajNerovnake (x:xs) = [x | not (uzJe x xs)] ++ dajNerovnake xs

nahrdelniky :: Int -> [String] -> [[String]]
nahrdelniky n xs = dajNerovnake $ vsetky n xs

