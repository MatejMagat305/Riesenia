module PlanarGraphColoring where

import Data.List
import Data.Function

type Country = String
type Neighbors = [Country]
type Vertex = (Country, Neighbors) 
type Graph = [Vertex]

data Color = Red | Green | Blue | Yellow deriving (Show, Eq)
colors = [Red, Green, Blue, Yellow]

---------------------------------------------------
coloring :: Graph -> [(Country, Color)]
coloring xs = backtrack [(x, colors) | x <- xs] []

--ci dane farbenie na danom grafe je OK -> susedia su roznej farby
isOk :: Graph -> [(Country, Color)] -> Bool
isOk xs col = and[and[d /= b1 | n <- b, (c, d) <- col, c == n] |
                (a, b) <- xs, (a1, b1) <- col, a == a1]

--usporiada staty tak, ze stat s najmensim poctom moznych farieb bude prvy
reorder :: [(Vertex, [Color])] -> [(Vertex, [Color])]
reorder xs = sortBy (compare `on` (length . snd)) xs

--najde prve mozne farbenie
backtrack :: [(Vertex, [Color])] -> [(Country, Color)] -> [(Country, Color)]
backtrack [] col = col
--xs -> [(Vertex, [Color])], Vertex = (Country, Neighbors)
--pridelujem jednotlivym statom postupne farby a volam backtracking
--backtracking mi vrati dobre ofarbenie jedine vtedy ked jeho dlzka je:
-- dlzka vstupu ofarbenia + pocet neofarbenych statov (xs) + 1,
-- ak to neplati dane farbenia nie je spravne, vraciam [] inak prve najdene farbenie
backtrack (x:xs) col 
        | null bck = []
        | otherwise = bck!!0
    where   v = fst (fst x)
            nei = snd (fst x)
            pom = [backtrack (reorder [(a, 
                              if (elem (fst a) nei) then (delete z b) else b) 
                              | (a, b) <- xs]) col++[(v, z)] | z <- (snd x)]
            bck =  [p | p <- pom, length p == (length xs) + 1 + (length col)]
                           
          

--https://mapchart.net/europe.html
europe :: Graph
europe = [
            ("portugal",["spain"]),
            ("spain",["portugal","andorra","france"]),
            ("andorra",["spain","france"]),
            ("france",["spain","andorra","monaco","italy","switzerland","germany","luxembourg","belgium","united_kingdom"]),
            ("united_kingdom",["france","belgium","netherlands","denmark","norway","iceland","ireland"]),
            ("ireland",["united_kingdom","iceland"]),
            ("monaco",["france"]),
            ("italy",["france","greece","albania","montenegro","croatia","slovenia","austria","switzerland","san_marino"]),
            ("san_marino",["italy"]),
            ("switzerland",["france","italy","austria","germany","liechtenstein"]),
            ("liechtenstein",["switzerland","austria"]),
            ("germany",["france","switzerland","austria","czech_republic","poland","sweden","denmark","netherlands","belgium","luxembourg"]),
            ("belgium",["france","luxembourg","germany","netherlands"]),
            ("netherlands",["belgium","germany","united_kingdom"]),
            ("luxembourg",["france","germany","belgium"]),
            ("austria",["italy","slovenia","hungary","slovakia","czech_republic","germany","switzerland","liechtenstein"]),
            ("slovenia",["italy","croatia","hungary","austria"]),
            ("croatia",["italy","montenegro","bosnia","serbia","hungary","slovenia"]),
            ("bosnia",["croatia","montenegro","serbia"]),
            ("montenegro",["croatia","italy","albania","serbia","bosnia"]),
            ("albania",["italy","greece","macedonia","serbia","montenegro"]),
            ("greece",["italy","cyprus","bulgaria","macedonia","albania"]),
            ("cyprus",["greece"]),
            ("macedonia",["albania","greece","bulgaria","serbia"]),
            ("bulgaria",["macedonia","greece","romania","serbia"]),
            ("serbia",["montenegro","albania","macedonia","bulgaria","romania","hungary","croatia","bosnia"]),
            ("romania",["serbia","bulgaria","hungary","moldova"]),
            ("hungary",["slovenia","croatia","serbia","romania","slovakia","austria","ukraine"]),
            ("slovakia",["austria","hungary","poland","czech_republic","ukraine"]),
            ("czech_republic",["germany","austria","slovakia","poland"]),
            ("poland",["germany","czech_republic","slovakia","sweden","ukraine","lithuania","belarus"]),
            ("denmark",["united_kingdom","germany","sweden","norway"]),
            ("sweden",["norway","denmark","germany","poland","finland"]),
            ("norway",["united_kingdom","denmark","sweden","finland","iceland"]),
            ("finland",["sweden","norway"]),
            ("iceland",["ireland","united_kingdom","norway"]),
            ("ukraine",["slovakia","moldova","poland","belarus","hungary"]),
            ("moldova",["ukraine","romania"]),
            ("belarus",["poland","ukraine","lithuania","latvia"]),
            ("lithuania",["poland","belarus","latvia"]),
            ("estonia",["latvia"]),
            ("latvia",["estonia","belarus","lithuania"])
        ]
        
-- +hranica ukraine-romania
europe2 :: Graph
europe2 = [
            ("portugal",["spain"]),
            ("spain",["portugal","andorra","france"]),
            ("andorra",["spain","france"]),
            ("france",["spain","andorra","monaco","italy","switzerland","germany","luxembourg","belgium","united_kingdom"]),
            ("united_kingdom",["france","belgium","netherlands","denmark","norway","iceland","ireland"]),
            ("ireland",["united_kingdom","iceland"]),
            ("monaco",["france"]),
            ("italy",["france","greece","albania","montenegro","croatia","slovenia","austria","switzerland","san_marino"]),
            ("san_marino",["italy"]),
            ("switzerland",["france","italy","austria","germany","liechtenstein"]),
            ("liechtenstein",["switzerland","austria"]),
            ("germany",["france","switzerland","austria","czech_republic","poland","sweden","denmark","netherlands","belgium","luxembourg"]),
            ("belgium",["france","luxembourg","germany","netherlands"]),
            ("netherlands",["belgium","germany","united_kingdom"]),
            ("luxembourg",["france","germany","belgium"]),
            ("austria",["italy","slovenia","hungary","slovakia","czech_republic","germany","switzerland","liechtenstein"]),
            ("slovenia",["italy","croatia","hungary","austria"]),
            ("croatia",["italy","montenegro","bosnia","serbia","hungary","slovenia"]),
            ("bosnia",["croatia","montenegro","serbia"]),
            ("montenegro",["croatia","italy","albania","serbia","bosnia"]),
            ("albania",["italy","greece","macedonia","serbia","montenegro"]),
            ("greece",["italy","cyprus","bulgaria","macedonia","albania"]),
            ("cyprus",["greece"]),
            ("macedonia",["albania","greece","bulgaria","serbia"]),
            ("bulgaria",["macedonia","greece","romania","serbia"]),
            ("serbia",["montenegro","albania","macedonia","bulgaria","romania","hungary","croatia","bosnia"]),
            ("romania",["serbia","bulgaria","hungary","moldova", "ukraine"]),
            ("hungary",["slovenia","croatia","serbia","romania","slovakia","austria","ukraine"]),
            ("slovakia",["austria","hungary","poland","czech_republic","ukraine"]),
            ("czech_republic",["germany","austria","slovakia","poland"]),
            ("poland",["germany","czech_republic","slovakia","sweden","ukraine","lithuania","belarus"]),
            ("denmark",["united_kingdom","germany","sweden","norway"]),
            ("sweden",["norway","denmark","germany","poland","finland"]),
            ("norway",["united_kingdom","denmark","sweden","finland","iceland"]),
            ("finland",["sweden","norway"]),
            ("iceland",["ireland","united_kingdom","norway"]),
            ("ukraine",["slovakia","moldova","poland","belarus","hungary", "romania"]),
            ("moldova",["ukraine","romania"]),
            ("belarus",["poland","ukraine","lithuania","latvia"]),
            ("lithuania",["poland","belarus","latvia"]),
            ("estonia",["latvia"]),
            ("latvia",["estonia","belarus","lithuania"])
        ]
        
australia :: Graph
australia = [
            ("wa", ["nt","sa"]),
            ("nt", ["wa","sa","q"]),
            ("sa", ["wa","nt","q","nsw","v"]),
            ("q", ["nt","sa","nsw"]),
            ("nsw", ["q","sa","v"]),
            ("v", ["sa","nsw"]),
            ("t", [])
            ]
{-
*PlanarGraphColoring> coloring australia
[("t",Red),("v",Red),("nsw",Green),("q",Red),("sa",Blue),("nt",Green),("wa",Red)]
(0.01 secs, 152,176 bytes)
-}