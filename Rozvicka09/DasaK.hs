slova 0 = [[]]
slova k = [ ch : w | w <- slova (k-1), ch <- "abcdef" ]

slova' 0 = [[]]
slova' k = slova (k-1) ++ [ch : w | w <- slova (k-1), ch <-"abcdef"]

-- riesenie z tabule :)
slova' :: Int -> [String]
slova' 0 = [[]]
slova' k = slova' (k-1) ++ [ ch:w | w <-slova' (k-1), ch <- "ABCDEF"]
