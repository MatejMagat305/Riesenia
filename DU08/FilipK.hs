--na zaaciatok nastavim ako najmensi prvy potom ho ked tak zmenim a vynulujem counter
najmensich::[Int]->Int
najmensich (x:xs) = snd$ foldl (\(a,b) -> \c->if (a>c) then (c,1) else (a,b+(fromEnum(a==c))) ) (x,1) xs

--mam zapamatany predchadzajuci prvok a counter ak je aktualny vacsi tak zvysim counter
vacsiZaMensim :: [Int]->Int
vacsiZaMensim (x:xs)=snd$ foldl (\(a,b) -> \c->(c,b+(fromEnum(a<c))) ) (x,0) xs

--ratam pocet a sucin prvkov nakoniec to odmocnim
geomR :: [Float]->Float
geomR xs = uncurry (\a -> \b -> (b**(1 / a))) $ foldl (\(a,b) -> \c -> (a+1,b*c) ) (0,1) xs

--to iste akurat sa vnorim po "riadkoch"
geomM :: [[Float]]->Float
geomM xs = uncurry (\a -> \b -> (b**(1 / a))) $ foldl (\a -> \b -> (foldl (\(x,y) -> \z -> (x+1,y*z) ) a b)) (0,1) xs

--ratam si pocet a sucet 1/x
harmoR :: [Float]->Float
harmoR xs = uncurry (\a -> \b -> (a/b)) $ foldl (\(a,b) -> \c -> (a+1,b+(1/c)) ) (0,0) xs

--rovnako ako v geomM sa vnaram do riadkov zapamatava sa to iste ako v harmoR 
harmoM :: [[Float]]->Float
harmoM xs = uncurry (\a -> \b -> (a/b)) $ foldl (\a -> \b -> (foldl (\(x,y) -> \z -> (x+1,y+(1/z)) ) a b)) (0,0) xs


--vyrabam ich rekurzivne s tym, ze funkcia vyber odfiltruje symetrie
vyber::Int->[[String]]->[[String]]->[[String]]
vyber n a [] = a
vyber n a (x:xs) = vyber n (x:a) [i|i<-xs,not(elem i ([(drop j a)++(take j a)|a<-[x],j<-[0..n]]++[(drop j a)++(take j a)|a<-[reverse x],j<-[0..n]])) ]

nahrdelniky :: Int -> [String] -> [[String]]
nahrdelniky 1 x = [[i]|i<-x]
nahrdelniky n (x:[])=[[x|i<-[0..n-1]]]
nahrdelniky n xs=vyber n [] [j:i|i<-(nahrdelniky(n-1) xs),j<-xs]

--pamatam si 2 najvacsie prvky ak najdem vacsi ako 1 z nich tak si ho zoberiem vratim druhy
druhyNajvacsi :: [Int] -> Int
druhyNajvacsi zoz = fst $ foldl (\(a,b) -> \c -> if c > a then (min b c,max b c) else (a,b) ) (head zoz, head zoz) zoz

--pamatam si 2 zoznamy a na striedacku do nich vkladam prvky
parnyNeparny::[a]->([a],[a])
parnyNeparny zoz = foldl (\(a,b) -> \c -> (if (length b==(length a)) then (a++[c],b)else (a,b++[c]))) ([],[]) zoz
