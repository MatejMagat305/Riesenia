slova 0 = [[]]
slova k = slova (k-1) ++ [ ch:w | w <- slova (k-1), length w == (k-1), ch <- "abcdef"]

