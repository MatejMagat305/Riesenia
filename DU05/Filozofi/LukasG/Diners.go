package philosophers

var DELAY = 10

func Dinner(N int, fixed bool) {
	simulation := newSimulation(N, fixed)
	simulation.runSimulation()
}
