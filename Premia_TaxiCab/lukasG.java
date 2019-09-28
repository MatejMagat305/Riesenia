/*
i=1 = 2
Seconds:0.0
i=2 = 50
Seconds:0.0
i=3 = 325
Seconds:0.002
i=4 = 1105
Seconds:0.003
i=5 = 5525
Seconds:0.009
i=6 = 5525
Seconds:0.01
i=7 = 27625
Seconds:0.163
i=8 = 27625
Seconds:0.159
i=9 = 71825
Seconds:1.051
i=10 = 138125
Seconds:3.726
i=11 = 160225
Seconds:5.003
i=12 = 160225
Seconds:4.997
i=13 = 801125
Seconds:93.639
i=14 = 801125
Seconds:92.141
 */


public class TaxiCab {
        // Na zaciatok jednoducha O(Long.MAX_VALUE^3) implementacia s jemne optimalizovanymi cyklami.
        // Ala "Som zvedavy, co dostanem za minimalnu namahu".
        // Sranda konci niekde okolo n = 12.
     
        public static void main(String[] args) {
        	for (int i=1; i < 15; i++) {
        		long start = System.currentTimeMillis();
        		System.out.println("i=" + i + " = " + taxi(i));
        		long end = System.currentTimeMillis();
        		System.out.println("Seconds:" + (float)(end-start)/1000);
        		
        	}
        }
     
        static long taxi(int n) {
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
    }
    
