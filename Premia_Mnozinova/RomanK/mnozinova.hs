import Data.List

runN :: Int -> [Integer] -> Integer
runN n arr | (n+n > length arr) = runN n (uni arr (map(2*)arr) (map(3*)arr) (map(4*)arr) (map(5*)arr)
                                                  (map(6*)arr) (map(7*)arr) (map(8*)arr) (map(9*)arr))
           | otherwise = (sort arr) !! (n-1)

uni a1 a2 a3 a4 a5 a6 a7 a8 a9 = nub (a1 ++ a2 ++ a3 ++ a4 ++ a5 ++ a6 ++ a7 ++ a8 ++ a9)

main = do
   putStrLn ("5. - " ++ show ( runN 5 [1] ) )
   putStrLn ("55. - " ++ show ( runN 55 [1] ) )
   putStrLn ("555. - " ++ show ( runN 555 [1] ) )
   putStrLn ("1000. - " ++ show ( runN 1000 [1] )  )
   putStrLn ("5555. - " ++ show ( runN 5555 [1] )  )
   putStrLn ("10000. - " ++ show ( runN 10000 [1] )  )
   putStrLn ("55555. - " ++ show ( runN 55555 [1] )  )
   putStrLn ("100000. - " ++ show ( runN 100000 [1] )  )
