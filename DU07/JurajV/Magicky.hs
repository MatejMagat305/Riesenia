module Magicky where

-- 7. HS3 - Magický štvorec

-- https://github.com/paradigmy/Kod/blob/master/PR07/haskell2.hs

transponuj :: [[Int]] -> [[Int]]
transponuj [] = []
transponuj ([]:xss) = transponuj xss
transponuj ((x:xs):xss) = (x:(map head xss)):(transponuj $ xs:(map tail xss))

magickyStvorec :: [[Int]] -> Bool
magickyStvorec xs = all (== head total) (tail total)
                where n = length xs
                      rows = [sum (xs !! row) | row <- [0..n-1]]
                      cols = [sum (transponuj xs !! row) | row <- [0..n-1]]
                      diag1 = [sum [xs !! row !! row | row <- [0..n-1]]]
                      diag2 = [sum [xs !! row !! (n-1-row) | row <- [0..n-1]]]
                      total = rows ++ cols ++ diag1 ++ diag2
