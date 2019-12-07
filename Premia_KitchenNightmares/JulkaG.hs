{-
*Main> length alergoFriendly
121080
(43.30 secs, 17,916,387,320 bytes)
-}

import Data.Bits
import Data.List

type Alergeny = Int   -- existuje 7 alergénov, ich bity sú 1,2,4,8,16,32,64, takže 9 = (1001) sú alergény 1 a 8 
type Jedlo  = (String, Alergeny) 
polievky    :: [Jedlo] 
polievky    = [("sosovicova",32+4+1),("pomodoro",8+2+1),("pohankova",0),("hokajdo",16+2),("gulasovka",8+4+2),("drzkova",32+2+1),("brokolicova",64+16+1)] 
hlavne      :: [Jedlo] 
hlavne      = [("spagety",16+8+4),("pizza",2+1),("gulas",32+8+4+2),("hambac",1+2+4+8+16+32+64),("dukatove",8+4+1),("perkelt",32+1),("ratatoille",16+2)] 

pocetAlergenov ((p,pa), (h,ha)) = popCount (pa .|. ha)

jedla :: [Jedlo] -> [Jedlo] -> Alergeny -> [Menu]
jedla polievky hlavne alergeny = [ (p,h) | p@(_,pa)<-polievky, pa .&. alergeny == 0, h@(_,ha)<-hlavne, ha .&. alergeny == 0]

type Menu = (Jedlo, Jedlo)
type Listok = [Menu]   -- jedálny lístok 
tyzdenny::[Listok]
tyzdenny = alls 5 polievky hlavne
            where
            alls :: Int -> [Jedlo] -> [Jedlo] -> [Listok]
            alls 0 _ _ = [[]]
            alls n polievky hlavne = [ (p,h):l |  (p,h)<-jedla polievky hlavne 0, l <- alls (n-1) (polievky\\[p]) (hlavne\\[h]) ]

lessThan5 :: Listok -> Bool
lessThan5 listok  = and [ pocetAlergenov denneMenu < 5 | denneMenu <- listok ]

freeDay :: Listok -> Bool
freeDay listok = and [ length [ 0 | ((_,pa), (_,ha)) <- listok, (pa .|. ha) .&. alergeny == 0 ] > 0 | alergeny <- [1,2,4,8,16,32,64] ]

alergoFriendly::[Listok]
alergoFriendly = [ ponuka | ponuka <- vsetky, lessThan5 ponuka, freeDay ponuka ]
                  where
                  vsetky = tyzdenny