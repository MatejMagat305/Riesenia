module DruhyNajvacsi where

-- 4. HS3 - Druhý najväčší cez foldl/r

druhyNajvacsi :: [Int] -> Int
druhyNajvacsi (x:xs) = snd (foldl (\(naj1, naj2) p -> 
                                    if p > naj1 then (p, naj1)
                                    else if p > naj2 && p < naj1 then (naj1, p)
                                    else (naj1, naj2)) (x, x) xs)