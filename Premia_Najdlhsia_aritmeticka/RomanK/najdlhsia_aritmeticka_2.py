import math

m = 0

i = 1
xx = 1

while ( True ):
    if not ( math.sqrt(xx) % 1 == 0 ):
        xx += 1
        i = 1
        continue


    arr = []
    x = xx

    while ( math.sqrt(x) % 1 == 0 ):
        arr.append(x)
        x += i

    if len(arr) > 2:
        print( len(arr), "|", i, arr )
        m = len(arr)

    i += 1
    if ( i > 100000 ):
        xx += 1
        i = 1
