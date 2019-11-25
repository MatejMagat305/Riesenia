module EuOS where

-- 5. HS 2 - EU OS

import Data.List

children :: Int -> [[Int]] -> [Int]
children n xs = [x !! 1 | x <- xs, head x == n]

graph :: [[Int]] -> [[Int]]
graph xs = nub [n:children n xs | n <- [head rel | rel <- xs]]

processesToKill :: [[Int]] -> [Int]
processesToKill xs = nub [x | x <- b, x `notElem` a]
                where a = [head rel | rel <- xs]
                      b = concat [tail rel | rel <- xs]

kill :: [[Int]] -> [Int]
kill [] = []
kill xs = concat (filter (\e -> length e == 1) rem) ++ next
        where rem = [x \\ processesToKill xs | x <- xs]
              next = kill (filter (\e -> length e > 1) rem)

odblokuj :: [[Int]] -> [Int]
odblokuj xs = processesToKill (graph xs) ++ kill (graph xs)