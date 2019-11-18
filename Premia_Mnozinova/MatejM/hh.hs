import Data.List

vratNCislo :: Int -> Integer
vratNCislo y = (sort((p++(pom (p!!0))))!!(y-1-(fromInteger ((p!!1)!!0))) where p = vratNVelkePole [1] (y) 1 0



vratNVelkePole ::[Integer]->Int->Int->Int->[[Integer]]
vratNVelkePole xs i j k | (j>=i) = [xs, [toInteger k]]
                         | otherwise = vratNVelkePole (p) i (j+length p) j
							where p = (pom xs)
					 
					
pom :: [Integer]->[Integer] 					
pom xs = nub [y*x | y<-[2..9],x<-(xs), x*y `notElem` xs ]					
--1					
--1              2 3 4 5 6 7 8 9
--1 2 3 4 5 6 7 8 9              10 12 14 15 16 18 20 24 28 30.......81
--1 2 3 4 5 6 7 8 9 10 12 14 15 16 18 20 