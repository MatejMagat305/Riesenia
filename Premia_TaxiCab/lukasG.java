/*
N: 1 , Number: 2, Time: 0.0 seconds
N: 2 , Number: 50, Time: 0.0 seconds
N: 3 , Number: 325, Time: 0.0 seconds
N: 4 , Number: 1105, Time: 0.0 seconds
N: 5 , Number: 5525, Time: 0.0 seconds
N: 6 , Number: 5525, Time: 0.0 seconds
N: 7 , Number: 27625, Time: 0.01 seconds
N: 8 , Number: 27625, Time: 0.001 seconds
N: 9 , Number: 71825, Time: 0.009 seconds
N: 10 , Number: 138125, Time: 0.011 seconds
N: 11 , Number: 160225, Time: 0.019 seconds
N: 12 , Number: 160225, Time: 0.012 seconds
N: 13 , Number: 801125, Time: 0.077 seconds
N: 14 , Number: 801125, Time: 0.072 seconds
N: 15 , Number: 801125, Time: 0.12 seconds
N: 16 , Number: 801125, Time: 0.079 seconds
N: 17 , Number: 2082925, Time: 0.19 seconds
N: 18 , Number: 2082925, Time: 0.199 seconds
N: 19 , Number: 4005625, Time: 0.442 seconds
N: 20 , Number: 4005625, Time: 0.43 seconds
N: 21 , Number: 5928325, Time: 0.729 seconds
N: 22 , Number: 5928325, Time: 0.791 seconds
N: 23 , Number: 5928325, Time: 0.74 seconds
N: 24 , Number: 5928325, Time: 0.721 seconds
N: 25 , Number: 29641625, Time: 5.644 seconds
N: 26 , Number: 29641625, Time: 5.063 seconds
N: 27 , Number: 29641625, Time: 5.375 seconds
N: 28 , Number: 29641625, Time: 4.906 seconds
N: 29 , Number: 29641625, Time: 6.249 seconds

*/


    import java.util.HashMap;
    import java.util.Map;
     
    public class TaxiCab {
        // Na zaciatok jednoducha O(Long.MAX_VALUE^3) implementacia s jemne optimalizovanymi cyklami.
        // Ala "Som zvedavy, co dostanem za minimalnu namahu".
        // Sranda konci niekde okolo n = 12.
     
        public static void main(String[] args) {
            for (int i = 1; i < 20; i++) {
                long start = System.currentTimeMillis();
                long result = taxi2(i);
                long finish = System.currentTimeMillis();
                System.out.println("N: " + i + " , Number: " + result + ", Time: " + (finish - start)/1000.f + " seconds");
            }
        }
     
        static long taxi1(int n) {
            long current = 0;
            while (++current < Long.MAX_VALUE) {
                int solutions = 0;
                for (long i = 1; i*i + i*i <= current; i++) {
                    for (long j = i; i*i + j*j <= current; j++) {
                        if (i*i + j*j == current) {
                            solutions++;
                        }
                        if (solutions == n) {
                            return current;
                        }
                    }
                }
            }
            return Long.MAX_VALUE; // Mozno by bol slusnejsi throw...
        }
     
        static long taxi2(int n) {
            Map<Long, Integer> memory = new HashMap<>();
            for (long i = 1; i < Long.MAX_VALUE; i++) {
                for (long j = 1; j <= i; j++) {
                    long calc = i*i + j*j;
                    if (!memory.containsKey(calc)) {
                        memory.put(calc, 1);
                    } else {
                        memory.put(calc, memory.get(calc) + 1);
                    }
                    if (memory.get(calc) == n) {
                        return calc;
                    }
                }
            }
            return Long.MAX_VALUE; // Mozno by bol slusnejsi throw...
        }
    }
