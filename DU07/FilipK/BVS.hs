module BVS where
-- binarny vyhladavaci strom

data BVS t     = Nod (BVS t) t (BVS t) | Nil deriving (Eq, Show, Ord)

-- dobre je mat nejaku konstantu toho typu
b1 :: BVS Int
b1 = Nod (Nod Nil 3 Nil) 5 (Nod Nil 7 Nil) 
b2 :: BVS Int
b2 = Nod (Nod Nil 1 Nil) 1 (Nod Nil 4 Nil) 
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
delete  :: (Ord t) => t -> BVS t -> BVS t             
delete  x bvs@(Nod left value right)   | find x bvs ==False = bvs
                                       | x == value && left /=Nil  = Nod (delete (maxBVS left) left) (maxBVS left) right
                                       | x == value && right /=Nil  = Nod left (minBVS right) (delete (minBVS right) right)
                                       | x == value = Nil
                                       | x < value  = Nod (delete x left) value right
                                       | otherwise  = Nod left value (delete x right)

maxBVS       :: (Ord t) => BVS t -> t
maxBVS Nil   =  error "prazdny nema maximum"
maxBVS bvs@(Nod left value right) | right == Nil  = value
                                  | otherwise  = maxBVS right

minBVS       :: (Ord t) => BVS t -> t
minBVS Nil   = error "prazdny nema minimum"
minBVS bvs@(Nod left value right) | left == Nil  = value
                                  | otherwise  = minBVS left