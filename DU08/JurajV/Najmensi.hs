module Najmensi where

-- 1. Foldový najmenší a najväčší

najmensich :: [Int] -> Int
najmensich xs = foldl (\c x -> if x == minim then c+1 else c) 0 xs
            where minim = foldr min 999999 xs

vacsiZaMensim :: [Int] -> Int
vacsiZaMensim (p:xs) = snd (foldl (\(a,c) x -> if x > a then (x,c+1)
                                               else (x,c))
                                               (p,0) xs)