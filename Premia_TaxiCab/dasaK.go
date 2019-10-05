/*
2: 50
3: 325
4: 1105
5: 8125
6: 5525
7: 105625
8: 27625
9: 71825
10: 138125
11: 5281250
12: 160225
13: 1221025
14: 2442050
15: 1795625
16: 801125
17: 446265625
18: 2082925
19: 41259765625
20: 4005625
21: 44890625
22: 30525625
23: 61051250
24: 5928325
25: 303460625
26: 53955078125
27: 35409725
28: 100140625
29: 1289367675781250
30: 52073125
31: 763140625
32: 29641625
33: 28056640625
34: 33721923828125
35: 7586515625
36: 77068225
37: 5158830625
38: 10317661250
39: 701416015625
40: 148208125
41: 2053764050
42: 1301828125
43: 8716125488281250
44: 62587890625
45: 885243125
46: 2356840332031250
47: 108951568603515632
48: 243061325
49: 476962890625


*/




    import java.util.ArrayList;
    import java.util.Arrays;
    import java.util.Collections;
    import java.util.HashSet;
    import java.util.LinkedList;
    import java.util.List;
    import java.util.Set;
     
    public class TaxiCab {
    	
    	//num = 2^a0 p1^a1 p2^a2 ... pn^an q1^b1 q2^b2 ... qm^bm
    	//pn su tvaru 4k + 3 a an su parne 
    	//qm su tvaru 4k + 1
    	//http://mathworld.wolfram.com/SumofSquaresFunction.html
    	
    	public static boolean isPrime(int num) {
    		if (num % 2 == 0 || num % 3 == 0 
    				|| num % 5 == 0 || num % 7 == 0) {
    			return false;
    		}
    		for (int i = 3; i < num; i += 2) {
    			if (num % i == 0) return false;
    		}
    		return true;
    	}
    	
    	public static boolean isPrime2(int num) {
    		if (num == 2) {
    			return true;
    		}
    		if (num % 2 == 0) {
    			return false;
    		}
    		for (int i = 3; i < num; i += 2) {
    			if (num % i == 0) return false;
    		}
    		return true;
    	}
    	
    	public static List<Integer> primeList(int count){
    		List<Integer> zoz = new ArrayList<Integer>();
    		zoz.add(5);
    		int last = 5;
    		while (zoz.size() < count) {
    			last += 4;
    			if (isPrime(last)) {
    				zoz.add(last);
    			}
    		}
    		return zoz;
    	}
    	
    	public static Set<List<Integer>> decomposition(int num){
    		Set<List<Integer>> mnoz = new HashSet<>();
    		mnoz.add(new LinkedList<>(Arrays.asList(num)));
    		if (isPrime2(num)) {
    			return mnoz;
    		}
    		//a*b = num
    		int a = 2;
    		int b = (int) num / a;
    		while (a <= b) {
    			//System.out.println(a + " " + b + " " + num);
    			if (a * b == num) {
    				mnoz.add(new LinkedList<>(Arrays.asList(b, a)));;
    				if (! isPrime2(a)) {
    					for (List<Integer> x: decomposition(a)) {
    						x.add(b);
    						Collections.sort(x);
    						Collections.reverse(x);
    						mnoz.add(x);
    					}
    				}
    				if (! isPrime2(b)) {
    					for (List<Integer> x: decomposition(b)) {
    						x.add(a);
    						Collections.sort(x);
    						Collections.reverse(x);
    						mnoz.add(x);
    					}
    				}
    			}
    			a++;
    			b = (int) num / a;
    		}
    		//System.out.println(num + " " + mnoz);
    		return mnoz;
    	}
    	
    	public static long smallest(List<Integer> primes, int n) {
    		long num = Long.MAX_VALUE;	
    		for (List<Integer> zoz: decomposition(n)) {
    			//System.out.println(zoz);
    			long pom = 1;
    			for (int i = 0; i < zoz.size(); i++) {
    				pom *= Math.pow(primes.get(i), zoz.get(i) - 1);
    			}
    			if (pom < num) {
    				num = pom;
    			}
    		}
    		return num;
    	}
    	
    	//number = 2^a0 p1^a1 p2^a2 ... pn^an q1^b1 q2^b2 ... qm^bm
    	//B = (b1+1)(b2+1)...(bm+1)
    	//B-even => n = B/2
    	//B-odd => n = 1/2 (B - (-1)^a0)
    	//a0 -> 0/1 (aby bolo co najmensie)
    	//prime: 2 | 3, 5, 7, 11, 13, 17, 19, 23, 29
    	//4k + 3 	3, 7, 11, 19, 23, ...
    	//4k + 1	5, 13, 17, 29, 31, 37, 41, 53, ...
    	public static long taxiCab(int n) {
    		n *= 2;
    		List<Integer> primes = primeList(n);
    		//System.out.println(primes);
    		//even 		
    		long num = smallest(primes, n);
    		//odd 
    		//n = B - (-1)^ao
    		//n + 1 -> a0 = 0
    		//n - 1 -> a0 = 1
    		num = Math.min(smallest(primes, n + 1), num);
    		long pom = smallest(primes, n - 1);
    		if (pom < num) {
    			num = Math.min(pom * 2, num);
    		}
    		return num;
    	}
     
    	public static void main(String[] args) {
    		for (int i = 2; i < 50; i++) {
    			System.out.println(i + ": " + taxiCab(i));
    		}
    	}
    }
     
     
    /*
    najdene cisla:
    2: 50
    3: 325
    4: 1105
    5: 8125
    6: 5525
    7: 105625
    8: 27625
    9: 71825
    10: 138125
    11: 5281250
    12: 160225
    13: 1221025
    14: 2442050
    15: 1795625
    16: 801125
    17: 446265625
    18: 2082925
    19: 41259765625
    20: 4005625
    21: 44890625
    22: 30525625
    23: 61051250
    24: 5928325
    25: 303460625
    ...
    */