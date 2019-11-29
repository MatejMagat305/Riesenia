{-
*PlanarGraphColoring>
*PlanarGraphColoring> coloring europe
[("portugal",Red),("spain",Green),("andorra",Red),("france",Blue),("united_kingdom",Yellow),("ireland",Green),("monaco",Red),("italy",Yellow),("san_marino",Red),("switzerland",Green),("liechtenstein",Blue),("germany",Yellow),("belgium",Green),("netherlands",Red),("luxembourg",Red),("austria",Red),("slovenia",Green),("croatia",Blue),("bosnia",Yellow),("montenegro",Green),("albania",Blue),("greece",Red),("cyprus",Green),("macedonia",Green),("bulgaria",Blue),("serbia",Red),("romania",Green),("hungary",Yellow),("slovakia",Blue),("czech_republic",Green),("poland",Red),("denmark",Red),("sweden",Blue),("norway",Green),("finland",Red),("iceland",Red),("ukraine",Green),("moldova",Red),("belarus",Blue),("lithuania",Green),("estonia",Green),("latvia",Red)]
(17.75 secs, 5,047,426,344 bytes)
*PlanarGraphColoring>




-}

module PlanarGraphColoring where

-- popis v coloring.pdf

import Data.List

type Country = String
type Neighbors = [Country]
type Vertex = (Country, Neighbors)
type Graph = [Vertex]
type Assignment = [(Country, Color)]

data Color = Red | Green | Blue | Yellow deriving (Show, Eq)
--data Color = Red | Green | Blue | Yellow | Purple deriving (Show, Eq)

countries :: Graph -> [Country]
countries g = [c| (c, _) <- g]

colors :: [Color]
colors = [Red, Green, Blue, Yellow]
--colors = [Red, Green, Blue, Yellow, Purple]

neighbors :: Graph -> Country -> Neighbors
neighbors g ctr = [n | (c, neighs) <- g, c == ctr, n <- neighs]

getNeighborsColors :: Graph -> Country -> Assignment -> [Color]
getNeighborsColors g ctr as = [col | (c, col) <- as, ctr `elem` neighbors g c]
   
possibleColors :: Graph -> Country -> Assignment -> [Color]
possibleColors g ctr as = [col | col <- colors, col `notElem` getNeighborsColors g ctr as]

allFourColors :: Graph -> Assignment -> Bool
allFourColors g as | lg == la = length (nub [col | (ctr, col) <- as]) == lc
                   | otherwise = True
                where lg = length g
                      la = length as
                      lc = length colors

search :: Graph -> [Country] -> [Assignment]
search g [] = [[]]
search g ctrs = [(ctr, col):rem | ctr <- ctrs, rem <- search g (ctrs \\ [ctr]), col <- possibleColors g ctr rem, allFourColors g ((ctr, col):rem)]

coloring :: Graph -> Assignment
coloring g = head $ search g (countries g)

australia :: Graph
australia = [
    ("Western Australia", ["Northern Territory", "South Australia"]),
    ("Northern Territory", ["Western Australia", "Queensland", "South Australia"]),
    ("South Australia", ["Northern Territory", "Queensland", "New South Wales", "Victoria", "Western Australia"]),
    ("Queensland", ["New South Wales", "South Australia", "Northern Territory"]),
    ("New South Wales", ["Queensland", "Victoria", "South Australia"]),
    ("Victoria", ["New South Wales", "South Australia", "Tasmania"]),
    ("Tasmania", ["Victoria"])
    ]

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


-- POSSIBLE RESULTS:

{-

*PlanarGraphColoring> coloring australia
[("Western Australia",Yellow),("Northern Territory",Red),("South Australia",Blue),
("Queensland",Green),("New South Wales",Red),("Victoria",Green),("Tasmania",Red)]
(0.01 secs, 302,272 bytes)

*PlanarGraphColoring> coloring europe
[("portugal",Red),("spain",Green),("andorra",Red),("france",Blue),
("united_kingdom",Yellow),("ireland",Green),("monaco",Red),("italy",Yellow),
("san_marino",Red),("switzerland",Green),("liechtenstein",Blue),("germany",Yellow),
("belgium",Green),("netherlands",Red),("luxembourg",Red),("austria",Red),
("slovenia",Green),("croatia",Blue),("bosnia",Yellow),("montenegro",Green),
("albania",Blue),("greece",Red),("cyprus",Green),("macedonia",Green),
("bulgaria",Blue),("serbia",Red),("romania",Green),("hungary",Yellow),
("slovakia",Blue),("czech_republic",Green),("poland",Red),("denmark",Red),
("sweden",Blue),("norway",Green),("finland",Red),("iceland",Red),
("ukraine",Green),("moldova",Red),("belarus",Blue),("lithuania",Green),
("estonia",Green),("latvia",Red)]
(22.30 secs, 5,047,617,680 bytes)

-}

pentagon :: Graph
pentagon = [
    ("A", ["B", "C", "D", "E"]),
    ("B", ["A", "C", "D", "E"]),
    ("C", ["A", "B", "D", "E"]),
    ("D", ["A", "B", "C", "E"]),
    ("E", ["A", "B", "C", "D"])
    ]