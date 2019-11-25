module Kruh where

-- 4. HS 2 - Kruh

import Data.List

kruh :: [[Int]]
kruh = backtrack [1..10]

backtrack :: [Int] -> [[Int]]
backtrack [] = [[]]
backtrack nums = [n:rem | n <- nums, rem <- backtrack (nums \\ [n]), check360 (n:rem)]

check :: [Int] -> Bool
check [] = True
check [_] = True
check (a:b:xs) = all (\i -> (a+b) `mod` i /= 0) [3, 5, 7] && check (b:xs)

check360 :: [Int] -> Bool
check360 xs | n == 10 = all (\i -> s `mod` i /= 0) [3, 5, 7] && check xs
            | otherwise = check xs
            where s = last xs + head xs
                  n = length xs