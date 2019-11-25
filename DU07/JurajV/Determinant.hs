module Determinant where

-- 3. HS 2 - Determinant

type Vector = [Int]
type Matrix = [[Int]]

val :: Matrix -> Int -> Int -> Int
val m row col = m !! row !! col

getRowWithoutColumn :: Matrix -> Int -> Int -> Vector
getRowWithoutColumn m row col = [val m row i | i <- [0..length m - 1], i /= col]

coMatrix :: Matrix -> Int -> Int -> Matrix
coMatrix m row col = [getRowWithoutColumn m i col | i <- [0..length m - 1], i /= row]

det :: Matrix -> Int
det m | n == 1 = val m 0 0
      | n == 2 =  val m 0 0  * val m 1 1 - val m 0 1 * val m 1 0
      | otherwise = sum [val m 0 col * (-1)^col * det (coMatrix m 0 col) | col <- c]
       where n = length m
             c = [0..n - 1]