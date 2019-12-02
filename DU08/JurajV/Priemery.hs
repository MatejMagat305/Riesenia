module Priemery where

-- 2. MID-HS-foldový

-- foldr vektor
geomR :: [Float] -> Float
geomR xs = p ** (1/n) 
        where (p, n) = foldr (\x (p, n) -> (p * x, n + 1)) (1,0) xs

-- foldr matica
geomM :: [[Float]] -> Float
geomM xs = p ** (1 / fromIntegral n)
        where (p, n) = foldr (\x (p, n) -> (p * product x, n + length x)) (1,0) xs

-- taka sugar verzia s flattnutou maticou
geomM' :: [[Float]] -> Float
geomM' xs = geomR (concat xs)

-- foldl vektor
harmoR :: [Float] -> Float
harmoR xs = uncurry (flip(/)) $ foldl (\(f,n) x -> (f+(1/x),n+1)) (0,0) xs

-- foldl matica
harmoM :: [[Float]] -> Float
harmoM xs = uncurry (flip(/)) $ foldl (\(fracs, n) x 
    -> (fracs + sum (map (1/) x), n + fromIntegral (length x))) (0,0) xs

-- taka sugar verzia s flattnutou maticou
harmoM' :: [[Float]] -> Float
harmoM' xs = harmoR (concat xs)