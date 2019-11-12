import Data.List

cartSucin3 k l m=[(x,y,z)|z<-nub m,y<-nub l,x<-nub k]