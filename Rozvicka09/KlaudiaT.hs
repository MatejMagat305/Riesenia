slova 0 = [[]]
slova k = [ ch : w | w <- slova (k-1), ch <- "abcdef" ]

slova' 0 pole = pole
slova' k pole = slova' (k-1) (pole ++ slova k)

totobolotrebadefinovat k = length $ slova' k [[]]
