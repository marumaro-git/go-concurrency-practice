package service

import "sync"

func OnceSample() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
			// NOTE: Do関数は一度だけ実行されるため2回目以降は実行されない。
			once.Do(increment)
		}()
	}

	wg.Wait()
	println("Count:", count)
}
