import Data.List
--ked vyrobim 1000 prvkov z kazdeho tak tam urcite je 1000ci najmensi
hs2_235 :: Int->Integer
hs2_235 n =(sort(1:[2^i|i<-[1..n]]++[3^i|i<-[1..n]]++[5^i|i<-[1..n]]))!!(n-1)
-- 10000 ci sa = hs2_235 10000

--determinant cez laplaceho
type Matica  = [ [ Int ] ] 
determinant :: Matica -> Int
determinant x   |length x == 1 = x!!0!!0
                |otherwise = sum[((-1)^(i))*(x!!0!!i)*(determinant (map (\a->[a!!j|j<-[0..(length a)-1],j/=i]) (tail x)))|i<-[0..(length x)-1]]



--10! nie je az tak vela ;p resp 9! kedze zafixujem 1 aby som obmedzil rotacie, takze len brute force vyrob vsetky permutacie a skontroluj ktore su dobere
perm [] = [[]]
perm xs = [ x:var | x <- xs, var <- perm (xs \\ [x])] 
kruh :: [[Int]]
kruh = [x|x<-(map(\a->1:a)(perm[2..10])), and[(((x!!i) + (x!!((i+1)`mod`(length x))))`mod` 3 /=0)  &&(((x!!i) + (x!!((i+1)`mod`(length x))))`mod` 5 /=0)&&(((x!!i) + (x!!((i+1)`mod`(length x))))`mod` 7 /=0)|i<-[0..(length x)-1] ]]

--postupne odbolokovavanie 
odblokuj :: [[Int]] -> [Int]
odblokuj x | length x == 0 = []
           | otherwise = c ++ (b \\ d) ++d
        where a=nub[i!!1|i<-x]
              b=nub[i!!0|i<-x]
              c=a\\b 
              d=odblokuj [i|i<-x, not (elem (i!!1) c)]


  
--najprv pozeram vramci rodinky (mama, otec, deti) a potom rekurzivne 
pribuzni :: Int -> Int -> [[Int]] -> Bool
pribuzni a b c | length [x|x<-c, (elem  a (take 2 x))&& (elem  b (drop 2 x))] > 0 = True
               | length [x|x<-c, (elem  b (take 2 x))&& (elem  a (drop 2 x))] > 0 = True
               | length [x|x<-c, (elem  b (drop 2 x))&& (elem  a (drop 2 x))] > 0 = True
               | otherwise = or[pribuzni i b [x|x<-c,not(x==detia)]|i<-(drop 2 detia)] || or[pribuzni i b [x|x<-c,not(x==rodiciaa)]|i<-(take 2 rodiciaa)] || or[pribuzni i b [x|x<-c,not(x==rodiciaa)]|i<-(drop 2 rodiciaa)]
               where detia = concat[x|x<-c, elem a (take 2 x)]
                     rodiciaa= concat[x|x<-c, elem a (drop 2 x)]

--staci len kontrolovat, ze sucet riadkov stlpcov a uhlopriecok je rovnaky
magickyStvorec :: [[Int]] -> Bool  
magickyStvorec a = and[sum i == b|i<-a] && and[sum i == b|i<-[[a!!k!!j|k<-[0..(length a)-1]]|j<-[0..(length a)-1]]] && (sum [a!!i!!i|i<-[0..(length a)-1]] == b) && (sum [a!!i!!(length a - i-1)|i<-[0..(length a)-1]] == b)
  where b = sum (a !! 0)

