import Data.List

-- 1. Foldový najmenší a najväčší
najmensich::[Int]->Int
najmensich xs = (length . filter (== foldr1 min xs)) xs

--vacsiZaMensim :: [Int]->Int
--vacsiZaMensim xs = foldl


-- 4. HS3 - Druhý najväčší cez foldl/r  
-- isiel by som po tomto prvom, ale trebalo nadefinovat foldl
--druhyNajvacsi :: [Int] -> Int	
--druhyNajvacsi xs = maximum $ filter (/= (maximum xs)) xs

druhyNajvacsi :: Ord a => [a] -> a
druhyNajvacsi xs = fst $ foldl searcher (h, h) xs
    where
        h = head xs
        searcher :: (Ord a) => (a, a) -> a -> (a, a)
        searcher (s, f) x = (min f (max s x), max f x)


-- 5. HS 4 - foldový
parnyNeparny::[a]->([a],[a])
parnyNeparny xs = foldr g ([],[]) xs
  where
  g x ~(as,bs) = (bs,x:as)