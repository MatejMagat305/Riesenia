slova 0 = [[]]
slova k = ws ++ [ch:w | w<-ws, ch<-"abcdef",  (ch:w)  `notElem` ws ] where ws = slova (k-1) 