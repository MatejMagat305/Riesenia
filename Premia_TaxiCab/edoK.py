for n = 2 is smallest number = 50
Seconds: 0.027939199999999997
for n = 3 is smallest number = 325
Seconds: 0.008985799999999988
for n = 4 is smallest number = 1105
Seconds: 0.00939509999999999
for n = 5 is smallest number = 5525
Seconds: 0.009948300000000021
for n = 6 is smallest number = 5525
Seconds: 0.0086503
for n = 7 is smallest number = 27625
Seconds: 0.020514900000000003
for n = 8 is smallest number = 27625
Seconds: 0.020082499999999948
for n = 9 is smallest number = 71825
Seconds: 0.04956559999999999
for n = 10 is smallest number = 138125
Seconds: 0.10288699999999995
for n = 11 is smallest number = 160225
Seconds: 0.1175062
for n = 12 is smallest number = 160225
Seconds: 0.11559070000000005
for n = 13 is smallest number = 801125
Seconds: 0.7043029
for n = 14 is smallest number = 801125
Seconds: 0.6713866
for n = 15 is smallest number = 801125
Seconds: 0.6761963
for n = 16 is smallest number = 801125
Seconds: 0.6736508999999997



# Eduard Krivanek
import time
class TaxiNumber:
	def __init__(self):
		self.numbers = {}
		self._end = False
 
	def findNumber(self, n):
		self.numbers = {}
		result =  self.__calculate(n)
		return 'for n = ' + str(n) + ' is smallest number = ' + str(result[0])
                #return 'for n = ' + str(n) + ' is smallest number = ' + str(result[0])

	def __calculate(self, n):
		number = 1
		while(True):
			# ak number = 10, tak sa robi 1 * 1 + 10 * 10 , 2 * 2 + 10 * 10 ...
			for i in range(number + 1):
				if i == 0:
					continue
				calc = i**2 + number**2
				if calc not in self.numbers:
					self.numbers[calc] = []
				self.numbers[calc].append([i, number])
				#print(self.numbers)
				if len(self.numbers[calc]) == n:
					return [calc , self.numbers[calc]]
			number += 1
 
# neviem , hadam je to ok
for i in range(15) :
        t = TaxiNumber()
        t1 = time.perf_counter()
        print(t.findNumber(i+2))
        t2 = time.perf_counter()
        print('Seconds:', t2 - t1)
 
