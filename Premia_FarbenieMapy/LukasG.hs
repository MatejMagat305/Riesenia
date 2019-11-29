{-
*PlanarGraphColoring> coloring europe
[("portugal",Red),("spain",Green),("andorra",Red),("france",Blue),("united_kingdom",Yellow),("ireland",Green),("monaco",Red),("italy",Yellow),("san_marino",Red),("switzerland",Green),("liechtenstein",Blue),("germany",Yellow),("belgium",Green),("netherlands",Red),("luxembourg",Red),("austria",Red),("slovenia",Green),("croatia",Blue),("bosnia",Yellow),("montenegro",Green),("albania",Blue),("greece",Red),("cyprus",Green),("macedonia",Green),("bulgaria",Blue),("serbia",Red),("romania",Green),("hungary",Yellow),("slovakia",Blue),("czech_republic",Green),("poland",Red),("denmark",Red),("sweden",Blue),("norway",Green),("finland",Red),("iceland",Red),("ukraine",Green),("moldova",Red),("belarus",Blue),("lithuania",Green),("estonia",Green),("latvia",Red)]
(1.49 secs, 450,782,288 bytes)
*PlanarGraphColoring>
-}


module PlanarGraphColoring where

type Country = String
type Neighbors = [Country]
type Vertex = (Country, Neighbors)
type Graph = [Vertex]

data Color = Red | Green | Blue | Yellow deriving (Show, Eq)

-- PETER, v definicii grafu su chyby! Napr. "spain",["portugal,,andorra","france"]), pozri prveho suseda - "portugal,,andorra"!
-- Musel som si to fixnut...

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

small :: Graph
small = [
            ("portugal",["spain"]),
            ("spain",["portugal","andorra","france"]),
            ("andorra",["spain","france"]),
            ("france",["spain","andorra","monaco","italy","switzerland","germany","luxembourg","belgium","united_kingdom"]),
            ("united_kingdom",["france", "belgium","netherlands","denmark","norway","iceland","ireland"])
        ]

coloring :: Graph -> [(Country, Color)]
coloring gs = head (backtrack gs)

backtrack :: Graph -> [[(Country, Color)]]
backtrack [v] = [[(fst v, c)] | c <- [Red, Green, Blue, Yellow]] 
--backtrack (bs@(v:gs)) = [((fst v, c):ps) | ps <- backtrack(gs), c <- [Red, Green, Blue, Yellow], (check bs ((fst v, c):ps) 0)]
backtrack (bs@(v:gs)) = [((fst v, c):ps) | ps <- backtrack(gs), c <- [Red, Green, Blue, Yellow], (checkOptimized bs ps (fst v, c))]

-- Update == 969.35 secs -> 2.60 secs!!!
checkOptimized :: Graph -> [(Country, Color)] -> (Country, Color) -> Bool
checkOptimized gs cs v = and[not (elem (n, snd v) cs) | n <- snd (gs !! (index gs (fst v)))] 
--

check :: Graph -> [(Country, Color)] -> Int -> Bool
check gs cs i 
  | i >= length(cs) = True
  | otherwise = and[not (elem (n, snd t) cs) | t <- cs, n <- snd (gs !! (index gs (fst t)))] && (check gs cs (i+1))


index :: Graph -> Country -> Int
index gs c = head [i | (v, i) <- zip gs [0..], fst v == c]

-- :set +s: meranie casu funkcii v ghci

-- coloring europe
--[("portugal",Red),("spain",Green),("andorra",Red),("france",Blue),("united_kingdom",Yellow),("ireland",Green),("monaco",Red),("italy",Yellow),("san_marino",Red),("switzerland",Green),("liechtenstein",Blue),("germany",Yellow),("belgium",Green),("netherlands",Red),("luxembourg",Red),("austria",Red),("slovenia",Green),("croatia",Blue),("bosnia",Yellow),("montenegro",Green),("albania",Blue),("greece",Red),("cyprus",Green),("macedonia",Green),("bulgaria",Blue),("serbia",Red),("romania",Green),("hungary",Yellow),("slovakia",Blue),("czech_republic",Green),("poland",Red),("denmark",Red),("sweden",Blue),("norway",Green),("finland",Red),("iceland",Red),("ukraine",Green),("moldova",Red),("belarus",Blue),("lithuania",Green),("estonia",Green),("latvia",Red)]
--(969.35 secs, 227,409,168,384 bytes) == +- 16min, na backtrack a celkom velky graf "not bad". Samozrejme, urcite sa da zlepsit...

--- !!!!! ----
-- UPDATE:
-- Po zoptimalizovani funkcie check na checkOptimized (vdaka za hint v komentari) sa cas a pamat znizili na (2.60 secs, 450,877,712 bytes) !!! 
--- !!!!! ----

-- Druha cast ulohy - SLABA KAVA, to bude uplny (plne prepojeny) graf o piatich vrcholoch :)
uplny :: Graph 
uplny = [
          ("v1", ["v2", "v3", "v4", "v5"]),
          ("v2", ["v1", "v3", "v4", "v5"]),
          ("v3", ["v1", "v2", "v4", "v5"]),
          ("v4", ["v1", "v2", "v3", "v5"]),
          ("v5", ["v1", "v2", "v3", "v4"])
        ]

-- coloring uplny *** Exception: Prelude.head: empty list
-- Je nejake krajsie riesenie? Zjavne nechcem najprv spocitat vsetky riesenia a az potom zistovat, ci tam je aspon jeden prvok, kedze to by
-- sme sa pri velkom grafe nedockali... Takze by to asi chcel nejaky try.. catch.

