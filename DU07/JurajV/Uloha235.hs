module Uloha235 where

-- 1. HS 2 - 235

import Data.List

najmensi :: Int -> Integer
najmensi 0 = undefined
najmensi n =  sort (1 : [i ^ j | j <- [1..n], i <- [2, 3, 5]]) !! (n-1)