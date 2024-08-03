package loadsimulation

func SimulateUsers(users int, testCase *testcase.TestCase) {
	for i := 0; i < users; i++ {
		go func() {
			// Simulate user
		}()
	}
}
