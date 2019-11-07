//pouzivam kostru kod z premie, kde sa problem riesil vseobecne

// klobuky
package main

import (
	"fmt"
	"time"
)

var farby []bool = []bool{true, false}
var agenti []*Hrac           // pole vsetych agentov tu budu traja
var finito = make(chan bool) // kanal sluzi len na ukoncenie

type Hrac struct { // agent definovany ako objekt
	id     int       // id agenta
	klobuk bool      // farba jeho klobuka
	ch     chan bool // kanal, na ktorom pocuva, hovorit sa mozu len farby, bool
}

func (agent Hrac) toString() string {
	return fmt.Sprintf("agent %v ma klobuk %v ", agent.id, agent.klobuk)
}

// aby sme nepodvadzali, tak toto je jedina funkcia, ktorou agenti pristupuju k polu klobukov
func (agent Hrac) vidim() []bool {
	k := []bool{}
	for i := 0; i < len(agenti); i++ {
		if agent.id == i {
			continue
		}
		k = append(k, agenti[i].klobuk)
	}
	return k
}

// a este tato, ta nam overi, ci sa mylime
func (agent Hrac) maPravdu(b bool) bool {
	return agent.klobuk == b
}

//-------------------------------------------------------------------------------
// zivot agenta
func (agent Hrac) run() {
	go func() { // agent je gorutina
		coVidim := agent.vidim()
		biele := 0
		for _, i := range coVidim { //zratam biele co vidim
			if i {
				biele++
			}
		}
		x := true
		switch biele {
		case 0: // ak vidim 2 cierne viem ze mam biely
			x = true
		case 1: //ak vidim 1 cierny tak cakam 2s
			time.Sleep(time.Duration(2) * time.Second)
			select {
			case <-agent.ch: //ak niekto nieco povedal tak musel vidiet 2 cierne teda mam aj ja cierny
				//ak nepovedal tak mam biely
				x = false
			default:
				x = true
			}
		case 2: //ak vidim 0 ciernych tak cakam 4s
			time.Sleep(time.Duration(4) * time.Second)
			select {
			case <-agent.ch: //nikto nemohol vidiet 2 cierne ale mohli vidiet 1 cierny a teda nieco
				//povedali a viem, ze mam cierny inak biely
				x = false
			default:
				x = true
			}
		}
		time.Sleep(time.Duration(1) * time.Second) //toto tu je aby sa zabezpecilo, ze sa kazdy skontroluje
		//v danom kroku kym im nieco poslem

		for _, a := range agenti {
			a.ch <- x
		}

		verdict := "MYLI SA"
		if agent.maPravdu(x) {
			verdict = "NE" + verdict
		}
		fmt.Printf("agent %v tvrdi, ze ma farbu %v \t a %v  \tlebo vidi %v \n", agent.id, x, verdict, agent.vidim())
		finito <- true
	}()
}

//----------------------------------------------------------------------------------------------------
func klobuky_test() {
	//su iba 3 moznosti tak sa testuju vsetky 3
	f := 0
	fmt.Println("2 Cierne:")
	agenti = []*Hrac{&Hrac{0, false, make(chan bool, 3)}, &Hrac{1, false, make(chan bool, 3)}, &Hrac{2, true, make(chan bool, 3)}}
	for i := 0; i < len(agenti); i++ {
		agenti[i].run()
	}
	f = 0
	for f < 3 {
		<-finito
		f++
	}
	/////////////////////////////
	fmt.Println("1 cierny:")
	agenti = []*Hrac{&Hrac{0, false, make(chan bool, 3)}, &Hrac{1, true, make(chan bool, 3)}, &Hrac{2, true, make(chan bool, 3)}}
	for i := 0; i < len(agenti); i++ {
		go agenti[i].run()
	}
	f = 0
	for f < 3 {
		<-finito
		f++
	}
	/////////////////////////////////////////////////
	fmt.Println("0 ciernych:")
	agenti = []*Hrac{&Hrac{0, true, make(chan bool, 3)}, &Hrac{1, true, make(chan bool, 3)}, &Hrac{2, true, make(chan bool, 3)}}
	for i := 0; i < len(agenti); i++ {
		go agenti[i].run()
	}
	f = 0
	for f < 3 {
		<-finito
		f++
	}
}
