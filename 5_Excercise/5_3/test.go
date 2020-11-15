func TestPhilosophers(t *testing.T) {

	var COUNT = 5

	// start table for 5 philosophers
	table := NewTable(COUNT)

	// create 5 philosophers and run parallel 
	for i := 0; i < COUNT; i++ {
		philosopher := Philosopher(i, table)
		go philosopher.run()
	}
	go table.run()

	// simulate 10 milliseconds --> check output
	time.Sleep(10 * time.Millisecond)
}