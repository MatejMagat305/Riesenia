import Data.List

data Farba = R | B | G | Y | W | P | O | C deriving (Show, Eq)

type Kocka = [Int]
type FKocka = [Farba]
type Farbenie = [Farba]

isRightHanded :: Kocka -> Bool
isRightHanded [a,b,c,d,e,f] | (a /= 1) && (b == 1) = isRightHanded[b,a,e,f,c,d]
                            | (a /= 1) && (c == 1) = isRightHanded[c,a,b,f,d,e]    
                            | (a /= 1) && (d == 1) = isRightHanded[d,a,c,f,e,b]        
                            | (a /= 1) && (e == 1) = isRightHanded[e,a,d,f,b,c]    
                            | (a /= 1) && (f == 1) = isRightHanded[f,b,e,d,c,a]                            
                            | b == 2 = (c == 3)
                            | b == 3 = (e == 2)
                            | b == 4 = (c == 2)
                            | b == 5 = (e == 3)
      
allDifferent :: Eq(t) => [t] -> Bool
allDifferent xs = length (nub xs) == length xs
      
moznosti :: [Kocka]
moznosti = [ [a,b,c,7-b,7-c,7-a] | a<-[1..6], b<-[1..6], b/=a, b/=7-a, c<-[1..6], 
                allDifferent [a,b,c,7-b,7-c,7-a] && isRightHanded  [a,b,c,7-b,7-c,7-a]]    

------------------------
--farbenie1 = [Y,B,R,R,Y,G]
--farbenie2 = [Y,G,G,Y,B,G]
--farbenie3 = [G,R,B,B,R,Y]
--farbenie4 = [R,G,G,Y,Y,G]
k1 = [[B,R,R,R,G,Y], [R,G,Y,G,B,B], [R,B,G,R,Y,Y], [G,B,R,Y,G,Y]]
k2 = [[B,R,Y,R,G,Y], [G,G,Y,Y,G,B], [R,B,G,B,Y,R], [G,G,R,Y,G,Y]]
k3 = [[O,B,O,P,G,Y], [G,P,G,Y,O,B], [P,G,P,O,Y,B], [B,G,B,P,O,Y], [Y,P,Y,O,G,B]]
k4 = [[O,C,B,G,Y,R], [O,C,B,R,Y,G], [O,G,B,Y,R,C], [O,G,B,C,R,Y], [O,Y,B,R,C,G], [O,R,B,Y,G,C]]
------------------------
steny = [0,1,2,4,3,5] 
farbenie :: Kocka -> Farbenie -> FKocka
farbenie xs fb = [ fb!!(steny!!(i-1)) | i <-xs]

kockySuOK :: [FKocka] -> Bool
kockySuOK ks = and [allDifferent (map (!!i) ks) | i <- [1..4]]

naSebe :: Int -> [FKocka] -> [[FKocka]]
naSebe 0 _ = [[]]
naSebe n fb = [ (k:m) | m <- naSebe  (n-1) fb, k <- polohy, kockySuOK (k:m) ]
    where polohy = [farbenie x (fb!!(n-1)) | x<-moznosti]


{--
*Main> length $ naSebe 4 k1
8
*Main> length $ naSebe 4 k2
256
*Main> length $ naSebe 5 k3
24
*Main> length $ naSebe 6 k4
88
--}