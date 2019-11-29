{-
*PlanarGraphColoring> coloring  europe
[("bulgaria",Green),("iceland",Green),("norway",Yellow),("denmark",Blue),("sweden",Green),("finland",Red),("estonia",Red),("belarus",Yellow),("latvia",Green),("lithuania",Red),("poland",Blue),("czech_republic",Green),("slovakia",Yellow),("ukraine",Red),("moldova",Green),("romania",Red),("hungary",Green),("slovenia",Red),("croatia",Yellow),("serbia",Blue),("bosnia",Green),("montenegro",Red),("albania",Yellow),("macedonia",Red),("greece",Blue),("cyprus",Red),("liechtenstein",Red),("austria",Blue),("switzerland",Yellow),("italy",Green),("san_marino",Red),("luxembourg",Green),("germany",Red),("belgium",Yellow),("netherlands",Green),("united_kingdom",Red),("ireland",Red),("monaco",Red),("france",Blue),("andorra",Red),("spain",Green),("portugal",Red)]
(0.04 secs, 1,824,128 bytes)
-}

module PlanarGraphColoring where
import Data.List
{-
vyfarbit sa nebude dat lubovolny graf ktory obsahuje K5(kompletny na 5 vrcholoch)
takze najmensi taky je ze mame vrcholy 1,2,3,4,5 a kazdy je sused s kazdym

P.S. povedal by som, ze naprogramovat farbenie mapy treba vediet urobit do 20min
vcelku mi trvalo 5hodin kym som zistil, ze v povodnej mape je chyba ;p  
-}
type Country = String
type Neighbors = [Country]
type Vertex = (Country, Neighbors)
type Graph = [Vertex]

data Color = Red | Green | Blue | Yellow deriving (Show, Eq)

sortState (a1, b1) (a2, b2)
  | (length b1) < length(b2) = LT
  | (length b1) > length(b2) = GT
  | otherwise = EQ

choose_wisely::[(Country,[Color])] -> Graph -> [Country] ->[Country]-> (Country,[Color])
choose_wisely possible mapa next used
        |not (null jednotky) = head jednotky
        |otherwise = head $ filter (\c -> fst c == stat) possible
        where jednotky = filter (\c -> length (snd c) ==1) possible
              zKoho= if length next >0 
                     then filter (\c -> elem (fst c) next ) mapa
                     else filter (\c -> not(elem (fst c) used )) mapa
              stat = if length zKoho > 0then fst $ head $ (sortBy sortState zKoho) else fst $ head $ possible 

neighbours x graph = concat [snd i|i<-graph,fst i == x ]
              
uprav_pos a b mapa= [if elem (fst i) s then (fst i, snd i \\ (snd a))else i|i<-b ]
        where s = neighbours (fst a) mapa 
        
backtrack :: [(Country, Color)]-> [(Country,[Color])] -> Graph -> [Country] -> [Country]->[(Country, Color)]
backtrack asigned [] mapa next used = asigned
backtrack asigned possible mapa next used
        |or[null (snd i)|i<-possible]=[]
        |otherwise = pod 
        where 
          x=choose_wisely possible mapa next used
          newpos=possible \\ [x]
          newnext=neighbours (fst x) mapa \\ used
          newused= (fst x):used
          back=[backtrack ((fst x, i):asigned) (uprav_pos (fst x,[i]) newpos mapa) mapa newnext newused| i<-(snd x)]
          pod = head ((filter (\c -> length c >0) back) ++ [[]])
              

coloring :: Graph -> [(Country, Color)]
coloring g = backtrack [] [(fst i,[Red,Green,Blue,Yellow])|i<-g] g [] []

xyz x = [fst i |i<-x] \\ (nub(concat[snd i|i<-x]))

australia :: Graph
australia = [
              ("WA",["NT","SA"]),
              ("NT",["WA", "Q", "SA"]),
              ("SA", ["NT", "Q", "NSW", "V", "WA"]),
              ("Q", ["NSW", "SA", "NT"]),
              ("NSW", ["Q", "V", "SA"]),
              ("V", ["NSW", "SA"]),
              ("T", [])
              ]

europe :: Graph
europe = [
            ("portugal",["spain"]),
            ("spain",["portugal","andorra","france"]),
            ("andorra",["spain","france"]),
            ("france",["spain,,andorra","monaco","italy","switzerland","germany","luxembourg","belgium","united_kingdom"]),
            ("united_kingdom",["france","belgium","netherlands","denmark","norway","iceland","ireland"]),
            ("ireland",["united_kingdom,,iceland"]),
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


--https://mapchart.net/europe.html
europe1 :: Graph
europe1 = [
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
        
