{-
*Main> length alergoFriendly
121080
(186.96 secs, 71,323,954,616 bytes)
*Main>
-}


import Data.Bits
import Data.List
type Alergeny = Int   -- existuje 7 alergénov, ich bity sú 1,2,4,8,16,32,64, takže 9 = (1001) sú alergény 1 a 8
type Jedlo  = (String, Alergeny)
polievky    :: [Jedlo]
polievky    = [("sosovicova",32+4+1),("pomodoro",8+2+1),("pohankova",0),("hokajdo",16+2),("gulasovka",8+4+2),("drzkova",32+2+1),("brokolicova",64+16+1)]
hlavne      :: [Jedlo]
hlavne      = [("spagety",16+8+4),("pizza",2+1),("gulas",32+8+4+2),("hambac",1+2+4+8+16+32+64),("dukatove",8+4+1),("perkelt",32+1),("ratatoille",16+2)]

type Menu  = (Jedlo, Jedlo)

pocetAlergenov::Menu->Int
pocetAlergenov ((a,b), (c,d))=sum [(shift (b .|. d) (-i)) .&. 1 | i<-[0..7]] 

jedla::[Jedlo]->[Jedlo]->Alergeny->[Menu]
jedla p h x=[(i,j) | i <- p, j <- h , (((snd i) .|. (snd j)) .&. x) ==0]

perm 0 xs =[[]]
perm n xs =[i:j | i <- xs , j <- (perm (n-1) xs )]

type Listok = [Menu]   -- jedálny lístok
tyzdenny     :: [Listok]
tyzdenny = tyzdennypom 5 polievky hlavne

tyzdennypom 0 p h =[[]] 
tyzdennypom n p h =[(i,j):x|i <- p, j <-h , x <- (tyzdennypom (n-1) (p \\ [i] ) (h \\ [j])) ]

alergoFriendly = filter (\c -> (and [pocetAlergenov i < 5 | i <- c]) && ((foldl (\h -> \((a,b),(x,y)) -> (h .|. (complement(b .|. y)))) 0 c) .&. 255  == 255)) tyzdenny
{-*Main Data.List> length alergoFriendly
121080
(176.43 secs, 71,323,953,608 bytes)
*Main Data.List>-}