slova 0 _ = [[]]
slova k abc = [[]] ++ [i:j | i <-abc, j <- slova (k-1) abc]
