package hats

func RunSimulation(N int) {
	simulation := newSimulation(N)
	simulation.printConfiguration()
	simulation.run()
	simulation.check()
}
