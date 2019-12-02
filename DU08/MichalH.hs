import Data.List

-- ======================================================================
-- ==                            uloha 1
-- ======================================================================

najmensich::[Int]->Int 
najmensich xs = snd $ foldl minWithCount (999999,0) xs

-- (minumum, pocet) -> aktualny_prvok -> (nove_minimum, novy_pocet)
minWithCount :: (Int,Int) -> Int -> (Int,Int)
minWithCount (x, xc) y 
    | x == y = (x, xc + 1)
    | x > y  = (y, 1)
    | otherwise = (x, xc)

vacsiZaMensim :: [Int]->Int
vacsiZaMensim xs = fst $ foldl vacsiZaMensimCount (0,999999) xs

-- (pocet, predchadzajuci_prvok) -> aktualny_prvok -> (novy_pocet, aktualny_prvok)
vacsiZaMensimCount :: (Int,Int) -> Int -> (Int,Int)
vacsiZaMensimCount (c, x) y
    | x < y = (c + 1, y)
    | otherwise = (c, y)


-- ======================================================================
-- ==                            uloha 2
-- ======================================================================

geomR :: [Float]->Float
geomR xs = flip (**) (1.0 / (fromIntegral (length xs))) $ foldl (*) 1.0 xs

geomM :: [[Float]]->Float
geomM xs = geomR $ foldl (++) [] xs

harmoR :: [Float]->Float
harmoR xs = (fromIntegral (length xs)) / (foldr (\x invsum -> invsum + (1.0 / x)) 0 xs)

harmoM :: [[Float]]->Float
harmoM xs = harmoR $ foldr (++) [] xs


-- ======================================================================
-- ==                            uloha 5
-- ======================================================================

parnyNeparny::[a]->([a],[a])
parnyNeparny zoz = snd $ foldr (\x (i,(evens,odds)) -> if (even i) then (i+1,(x:evens,odds)) else (i+1,(evens,x:odds))) (if (even (length zoz)) then 1 else 0,([],[])) zoz
