import Data.List
import Data.Set

horska n = build3 0 (fromList[1]) n 0
build3 last toDo n akt |  n==akt = last
                       | otherwise = build3 x (Data.Set.delete x b ) n (akt+1)
                        where x= elemAt 0 toDo
                              b= Data.Set.union toDo (fromList[x*i|i<-[2..9],(x*i > last)])
                              
                              
{-
*Main> :set +s
*Main> horska 1000
385875
(0.02 secs, 6,347,768 bytes)
*Main> horska 10000
63221760000
(0.17 secs, 65,307,216 bytes)
*Main> horska 100000
123093144973968750000
(1.88 secs, 695,421,096 bytes)
*Main> horska 5
5
(0.01 secs, 86,592 bytes)
*Main> horska 55
140
(0.00 secs, 380,888 bytes)
*Main> horska 555
46875
(0.01 secs, 3,506,552 bytes)
*Main> horska 5555
1475789056
(0.08 secs, 36,008,600 bytes)
*Main> horska 55555
154414312500000000
(0.91 secs, 370,630,016 bytes)
*Main> horska 555555
28662368719582789632000000000000
(10.14 secs, 4,233,234,288 bytes)
*Main> horska 5555555
686257960245984000000000000000000000000000000000000000000
(104.37 secs, 45,770,952,296 bytes)
-}