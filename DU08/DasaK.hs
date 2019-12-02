--DU08

--01
najmensich::[Int]->Int
najmensich (x:xs) = snd $ foldl(\(a, count) -> \x -> 
                                if x < a then (x, 1) 
                                else if x == a then (a, count+1)
                                else (a, count)
                                ) (x, 1) xs  

najmensich'::[Int]->Int                              
najmensich' (x:xs) = snd $ foldr(\x -> \(a, count) ->
                                if x < a then (x, 1) 
                                else if x == a then (a, count+1)
                                else (a, count)
                                ) (x, 1) xs 
                                
vacsiZaMensim :: [Int]->Int
vacsiZaMensim (x:xs) = snd $ foldl(\(a, count) -> \x -> 
                                    if a < x then (x, count+1)
                                    else (x, count) ) (x, 0) xs
                                    
vacsiZaMensim' :: [Int]->Int
vacsiZaMensim' xs = snd $ foldr(\x -> \(a, count) ->
                                    if a > x then (x, count+1)
                                    else (x, count) ) (last xs, 0) xs
                                    
--02
geomR :: [Float]->Float
geomR xs = sum ** (1/count)
    where (sum, count) = foldr (\x -> \(a, count) -> (a*x, count+1)) (1, 0) xs
    
geomM :: [[Float]]->Float
geomM xs = sum ** (1/(fromIntegral count))
    where (sum, count) = foldr (\x -> \(a, count) -> (a*(product x), count+(length x)) ) (1, 0) xs

geomM' :: [[Float]]->Float
geomM' xs = uncurry (\sum -> \count -> sum ** (1/count))
    $ foldr (\x -> \acount -> 
      (foldr (\y -> \(b, c) -> (b*y, c+1)) acount x )
      ) (1, 0) xs

harmoR :: [Float]->Float
harmoR xs = count / sum
    where (sum, count) = foldr (\x -> \(a, count) -> (a+(1/x), count+1)) (0,0) xs

harmoM :: [[Float]]->Float
harmoM xs = uncurry (\sum -> \count -> count/sum) 
        $ foldr (\x -> \acount -> 
        (foldr (\y -> \(b, c) -> (b+(1/y), c+1)) acount x)
        ) (0,0) xs
    
--03
--zistuje ci pootocenia nahrdelnikov su rovnake
rovnake :: [String] -> [String] -> Bool
rovnake x y = or[x == ((drop  (n - i) y) ++ (take (n - i) y)) | i<-[0..(n-1)]] 
    where n = length y
     
--vytvori vsetky mozne nahrdelinky vstupnej dlzky a zo zadanych koralikov     
vso :: Int -> [String] -> [[String]]
vso 0 _ = [[]]
vso _ [] = []
vso k xs = [x:a | x <- xs, a <-vso (k-1) xs]

--pouzije "rovnake" a "vso" a vrati iba rozne nahrdelniky
nahrdelniky :: Int -> [String] -> [[String]]
nahrdelniky k xs = foldl (\a -> \x -> 
                            if or[(rovnake i x) | i<-a] then a
                            else a++[x]) [] (vso k xs)

--04
druhyNajvacsi :: [Int] -> Int
druhyNajvacsi (x:y:xs) = snd $ foldl(\(a, b) -> \x -> 
            if x > a then (x, a)
            else if x > b then (a, x)
            else (a, b) ) ((max x y), (min x y)) xs

--05
parnyNeparny::[a]->([a],[a])
parnyNeparny xs = foldl (\(a, b) -> \x -> 
                      if length a == length b then (a++[x], b)
                      else (a, b++[x])) ([], []) xs