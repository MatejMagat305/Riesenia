module Cube where

subcubes :: Int -> Int -> Int

--https://en.wikipedia.org/wiki/Hypercube
subcubes k n | k > n || 0 > n || 0 > k = 0
             | k == 0 = 2^n
             | otherwise = 2*(subcubes k (n-1)) + (subcubes (k-1) (n-1))

{-
sum [ subcubes k n | k <- [0..n]]
sucet v riadku pre N-kocku = 3^N
(sucet K-kociek takych ze K <= N)
-}