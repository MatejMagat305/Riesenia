module Auto where
import Data.List
whereToStart :: [Float] -> [Float] -> [Int]
whereToStart dist lit = [i|i<-[0..(length dist)-1],and[(sum(take j ((snd(splitAt i dist))++(fst(splitAt i dist)))))<=(sum(take j ((snd(splitAt i lit))++(fst(splitAt i lit)))))|j<-[0..(length dist)-1]]]