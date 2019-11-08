module Auto where


nextStop :: Int -> Int -> Int
nextStop dist lit = lit - dist

tryStart :: [Float] -> [Float] -> Float -> Bool -> Bool
tryStart [] [] n True = True
tryStart a b c False = False
tryStart (x:xs) (y:ys) n a =  tryStart xs ys (y + n - x) ((y + n - x) >= 0)

transform :: [Float] -> Int -> [Float]
transform a i = (drop i a) ++ (take i a) 


whereToStart :: [Float] -> [Float] -> [Int]
whereToStart dist lit = [ i | i<-[0..(length dist) -1],(tryStart (transform dist i) (transform lit i) 0 True) == True]