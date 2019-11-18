module Horska where

-- 1. Prémia Horská

import Data.List

-- popis v readme.pdf

mergeUnique :: [Integer] -> [Integer] -> [Integer]
mergeUnique a  [] = a
mergeUnique [] b  = b
mergeUnique (a:as) (b:bs) = case compare a b of
                                EQ -> a: mergeUnique as bs
                                LT -> a: mergeUnique as (b:bs)
                                GT -> b: mergeUnique (a:as) bs

nums :: [Integer]
nums = 1 : mergeUnique (map (2*) nums) 
            (mergeUnique (map (3*) nums) 
                (mergeUnique (map (4*) nums) 
                    (mergeUnique (map (5*) nums) 
                        (mergeUnique (map (6*) nums) 
                            (mergeUnique (map (7*) nums) 
                                (mergeUnique (map (8*) nums) (map (9*) nums)))))))

horska :: Int -> Integer
horska n = nums !! (n-1)

{-
*Horska> horska 1000
385875
(0.01 secs, 2,500,808 bytes)

*Horska> horska 1000
385875
(0.00 secs, 117,664 bytes)

*Horska> horska 10000
63221760000
(0.05 secs, 22,332,088 bytes)

*Horska> horska 100000
123093144973968750000
(0.48 secs, 230,867,416 bytes)

*Horska> horska 1000000
4157409948433216829957008507500000000
(3.91 secs, 2,574,677,928 bytes)

*Horska> horska 10000000
1037754506929393275846369194817675812165052808572460991244140544000
(39.43 secs, 26,789,343,112 bytes)

*Horska> horska 5
5
(0.00 secs, 113,984 bytes)

*Horska> horska 55
140
(0.01 secs, 222,096 bytes)

*Horska> horska 555
46875
(0.01 secs, 1,307,496 bytes)

*Horska> horska 5555
1475789056
(0.03 secs, 12,406,472 bytes)

*Horska> horska 55555
154414312500000000
(0.31 secs, 124,490,208 bytes)

*Horska> horska 555555
28662368719582789632000000000000
(2.31 secs, 1,418,698,456 bytes)

*Horska> horska 5555555
686257960245984000000000000000000000000000000000000000000
(21.65 secs, 14,653,399,784 bytes)
-}