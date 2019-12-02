module ParnyNeparny where

-- 5. HS 4 - foldový

parnyNeparny :: [a] -> ([a],[a])
parnyNeparny xs = foldl (\(z1, z2) x -> 
                    if (length z1 + length z2) `mod` 2 == 0 then (z1++[x], z2) 
                    else (z1, z2++[x])) ([], []) xs