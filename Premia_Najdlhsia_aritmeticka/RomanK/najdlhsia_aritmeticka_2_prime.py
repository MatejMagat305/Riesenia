import math

m = 0

i = 1
xx = 1

def isPrime(n) :
    if (n <= 1) : return False
    if (n <= 3) : return True
    if (n % 2 == 0 or n % 3 == 0) : return False
    i = 5
    while(i * i <= n) : 
        if (n % i == 0 or n % (i + 2) == 0) : return False
        i = i + 6
    return True

while ( True ):
    if not ( isPrime(xx) ):
        xx += 1
        i = 1
        continue

    arr = []
    x = xx

    while ( isPrime(x) ):
        arr.append(x)
        x += i

    if len(arr) > m:
        print( len(arr), "|", i, arr )
        m = len(arr)

    i += 1
    if ( i > 1000000 ):
        xx += 1
        i = 1
