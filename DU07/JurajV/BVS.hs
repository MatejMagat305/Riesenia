module BVS where
-- binarny vyhladavaci strom

data BVS t     = Nod (BVS t) t (BVS t) | Nil deriving (Eq, Show, Ord)

-- dobre je mat nejaku konstantu toho typu
b1 :: BVS Int
b1 = Nod (Nod Nil 3 Nil) 5 (Nod Nil 7 Nil) 
b2 :: BVS Int
b2 = Nod (Nod Nil 1 Nil) 1 (Nod (Nod Nil 2 Nil) 4 Nil) 
b3 :: BVS Int
b3 = Nod b2 3 b1

--
-- find
find  :: (Ord t) => t -> BVS t -> Bool
find  _ Nil  = False
find  x (Nod left value right) | x == value  = True
                               | x < value  = find x left
                               | otherwise  = find x right

-- insert
insert  :: (Ord t) => t -> BVS t -> BVS t                 
insert  x  Nil  = Nod Nil x Nil
insert  x bvs@(Nod left value right) | x == value  = bvs
                                     | x < value  = Nod (insert x left) value right
                                     | otherwise  = Nod left value (insert x right)

-- riesenie DU7

-- https://static.javatpoint.com/ds/images/binary-search-tree.png
testTree1 :: BVS Int
testTree1 = Nod (Nod (Nod Nil 1 Nil) 3 (Nod (Nod Nil 4 Nil) 6 (Nod Nil 7 Nil))) 8 (Nod Nil 10 (Nod (Nod Nil 13 Nil) 14 Nil))

testTree2 :: BVS Char
testTree2 = Nod (Nod Nil 'h' (Nod (Nod Nil 'n' Nil) 'o' Nil)) 'p' (Nod (Nod Nil 't' Nil) 'y' Nil)

minBVS :: (Ord t) => BVS t -> t
minBVS Nil = error "prazdny nema minimum"
minBVS (Nod Nil val _) = val
minBVS (Nod left val right) = minBVS left

delete  :: (Ord t) => t -> BVS t -> BVS t    
delete  _ Nil = Nil             
delete  x (Nod left val right)  | x < val = Nod (delete x left) val right
                                | x > val = Nod left val (delete x right)
                                | x == val = decide (Nod left val right)

decide :: (Ord t) => BVS t -> BVS t 
decide (Nod Nil val Nil)   = Nil
decide (Nod left val Nil)  = left
decide (Nod Nil val right) = right
decide (Nod left val right) = Nod left (minBVS right) (delete (minBVS right) right)


-- helper
build :: (Ord t) => [t] -> BVS t
build [] = Nil
build xs = foldl (\accTree -> \x -> insert x accTree) Nil xs