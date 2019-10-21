package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// toto je ta globalna premenna

var a int64



// toto nemeniť, len použiť

func input(prompt string) {

	var i string

	fmt.Println(prompt)

	fmt.Scan(&i)
	a, _ = strconv.ParseInt(i, 10, 64)

}
//01
func akratbKodujRek1(){
	if a>0{
		a--
		akratbKodujRek1()
		a*=5
	}else{
		input("zadaj druhe cele cislo:")
		if a<0{
			a*=-1
			akratbKodujRek2()
			a*=-1
		}else {
			akratbKodujRek2()
		}

	}
}
func akratbKodujRek2(){
	if a>0{
		a--
		akratbKodujRek2()
		a*=3
	}else{
		a=1
	}

}

func akratb(){
	input("zadaj prve cele cislo:")
	if a<0{
		a*=-1
		akratbKodujRek1()
		a*=-1
	}else{
		akratbKodujRek1()
	}
	urobnasobenie()
	if a<0{
		a*=-1
	ziskajCislo()
	a*=-1}else {ziskajCislo()}
	fmt.Print( "cislo je: ", a)
}

func urobnasobenie() {
	b:=a
	for; a%5==0;{
		a/=5
		b/=5
		urobnasobenieRek()
		}
}

func urobnasobenieRek() {
	if a%3==0 &&a!=0{
		a/=3
		urobnasobenieRek()
		a*= 2*3

	}
}
func ziskajCislo(){
	if a%2==0&&a!=0{
		a/=2
		ziskajCislo()
		a++
	}else{
		a=0
	}
}
//02
func sumN(){
	input("zadaj cislo:")
	sumnInt1()
	fmt.Print("sucet je: ", a)
}
func sumnInt1() {
	if a > 0 {
		a--
		sumnInt1()
		a++
	} else {
		if a < 0 {
			a++
			sumnInt1()
			a--
		} else {
			input("zadaj cislo:")
			if a!=0{
				sumnInt1()
			}
		}

	}

}

/*
03. úloha je v screenshotoch a v cvičenie.txt, ospravedlnujem sa 1. screenshot vnikol v čase,
som ešte nepokladal za nutné písať aj do ...txt, ten som prepísal až neskôr, preto tam môžu byť rozdiely.......
 */


// 4


/*
func comb(a int64,b int64)*big.Int{
	var n *big.Int
	ch:=make(chan *big.Int)
	ch2:=make(chan *big.Int)
	ch3:=make(chan *big.Int)
	ch4:=make(chan *big.Int)
	go func() {ch<-faktorial(a)}()
	go func() {ch2<-faktorial(b)}()
	go func() {ch3<-faktorial(a-b)}()
	go func() {n =<-ch2
	ch4<-new(big.Int).Mul(<-ch3,n)}()
	return new(big.Int).Div(<-ch, n)
}
func faktorial(n int64)*big.Int{
	return new(big.Int).MulRange(1, n)
}

func main() {
	//akratb()//funguje len pre <(-7*7, (7*7)>, ak som niekde zabudol b:= a;b++,
	// tak to b som potrebovl iba na debugovanie, lebo narychlo som nevedel nastaviť, aby to ukazovalo globálne prem.
	//sumN()
	start:=time.Now()
	v:=comb(50000, 25000)
	fmt.Print("time=%v",time.Since(start))
	str := v.String()
	print("   prefix: ", str[:10],"           ma " , v.BitLen()," cifier           ")
 	fmt.Printf("\"formating=%v\n",time.Since(start))
/*
	start=time.Now()
	v=comb(100000, 50000)
	fmt.Print("time=%v",time.Since(start))
	str = v.String()
	print("   prefix: ", str[:10],"           ma " , v.BitLen()," cifier           ")
	fmt.Printf("formating=%v\n",time.Since(start))
	start=time.Now()
	v=comb(1000000, 500000)
	fmt.Print("time=%v",time.Since(start))
	str = v.String()
	print("   prefix: ", str[:10],"           ma " , v.BitLen()," cifier           ")
	fmt.Printf("\tformating=%v\n",time.Since(start))

	start=time.Now()
	v=comb(10000000, 5000000)
	fmt.Print("\ttime=%v",time.Since(start))
	str = v.String()
	print("   prefix: ", str[:10],"           ma " , v.BitLen()," cifier           ")
	fmt.Printf("\tformating=%v\n",time.Since(start))
		comb(50000, 25000) ma 15050 cifier, time=19ms
		comb(100000, 50000) ma 30101 cifier, time=88ms
		comb(1000000, 500000) ma 301027 cifier, time=7sec
		comb(10000000, 5000000)
}*/
// https://github.com/paradigmy/Kod/blob/master/PR03/comb.go, neviem, aká bola pointa, ale snažil som sa to upraviť......
func combPara(n int64, k int64) *big.Int {
	citatelch := make(chan *big.Int)
	go Factorial(n, citatelch)
	menovatel1ch := make(chan *big.Int)
	go Factorial(k, menovatel1ch)
	delenie1ch := make(chan *big.Int)
	go func() {
		citatel := <-citatelch
		menovatel1 := <-menovatel1ch
		citatel.Div(citatel, menovatel1)
		delenie1ch <- citatel
	}()
	delenie2ch := make(chan *big.Int)
	go func() {
		fn_k := make(chan  *big.Int )
		go func() {f :=big.NewInt(1)
			fn_k<-f.Mul(f, big.NewInt(n-k))}()
		menovatel := <-delenie1ch
		menovatel.Div(menovatel, <-fn_k)
		delenie2ch <- menovatel	}()
	return <-delenie2ch
}
func comb1Para(n int64, k int64) *big.Int {
	citatelch := make(chan *big.Int)
	go func() {		citatelch <- big.NewInt(1).MulRange(n-k+1, n)	}()
	menovatelch := make(chan *big.Int)
	go func() {		menovatelch <- big.NewInt(1).MulRange(1, k)}()
	citatel := <-citatelch
	citatel.Div(citatel, <-menovatelch)
	return citatel
}
func Factorial(n int64, ch chan *big.Int) {
	f :=big.NewInt(1)
	if n>1{ch <- f.Mul(f, big.NewInt(n))} else {ch<-f}
}
func main() {
	var n int64
	n = 50000
	start := time.Now()
	l:=comb1Para(n, n/2).String()
	fmt.Printf("comb1Para(%d, %d) ma %d cifier,, prvych 10 cisel:%10s,   time=%v\n", n, n/2, len(l),l[:10], time.Since(start))
	n = 100000
	start = time.Now()
	l = comb1Para(n, n/2).String()
	fmt.Printf("comb1Para(%d, %d) ma %d cifier,, prvych 10 cisel:%10s,   time=%v\n", n, n/2, len(l),l[:10], time.Since(start))
	n = 1000000
	start = time.Now()
	l = comb1Para(n, n/2).String()
	fmt.Printf("comb1Para(%d, %d) ma %d cifier,, prvych 10 cisel:%10s,   time=%v\n", n, n/2, len(l),l[:10], time.Since(start))
	n = 10000000
	start = time.Now()
	l = comb1Para(n, n/2).String()
	fmt.Printf("comb1Para(%d, %d) ma %d cifier,, prvych 10 cisel:%10s,   time=%v\n", n, n/2, len(l),l[:10], time.Since(start))
	//n = 10000000
	//start = time.Now()
	//fmt.Printf("comb1Para(%d, %d) ma %d cifier, time=%v\n", n, n/2, len(comb1Para(n, n/2).String()), time.Since(start))
}
