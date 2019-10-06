int x = -2147483648;
//long x = -9223372036854775808L; //byte a short nefunguje
System.out.println( (x==-x) && (x!=0) );

long x = Long.MAX_VALUE;
double y = x; //funguje x az po x-511
long z = Long.MAX_VALUE-1;
System.out.println( (x==y) + " " + (y==z) + " " + (x==z) );
//System.out.println( x + " " + y + " " + z );
