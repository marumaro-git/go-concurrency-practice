package service

import (
	"slices"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	t.Run("single job processing", func(t *testing.T) {
		jobs := make(chan int, 1)
		results := make(chan int, 1)
		defer func() {
			close(jobs)
			close(results)
		}()

		go worker(1, jobs, results)

		jobs <- 5

		select {
		case res := <-results:
			if res != 10 {
				t.Errorf("Expected result 10, got %d", res)
			}
		case <-time.After(2 * time.Second):
			t.Error("Test timed out waiting for result")
		}
	})

	t.Run("multiple jobs processing", func(t *testing.T) {
		jobs := make(chan int, 3)
		results := make(chan int, 3)
		defer func() {
			close(jobs)
			close(results)
		}()

		go worker(1, jobs, results)

		jobs <- 1
		jobs <- 2
		jobs <- 3

		expected := []int{2, 4, 6}
		for i := 0; i < 3; i++ {
			select {
			case res := <-results:
				if slices.Contains(expected, res) {
				} else {
					t.Errorf("Unexpected result %d", res)
				}
			case <-time.After(2 * time.Second):
				t.Error("Test timed out waiting for results")
			}
		}
	})
}

func TestRun(t *testing.T) {
	t.Run("normal execution", func(t *testing.T) {
		done := make(chan bool)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Run panicked: %v", r)
				}
				done <- true
			}()
			SimpleWorkerRun()
		}()

		select {
		case <-done:
			// 正常完了
		case <-time.After(10 * time.Second):
			t.Fatal("Run did not complete within 10 seconds")
		}
	})

	t.Run("concurrency effectiveness", func(t *testing.T) {
		startTime := time.Now()
		SimpleWorkerRun()
		elapsed := time.Since(startTime)

		if elapsed > 3*time.Second {
			t.Errorf("Expected execution time < 3s with 3 workers, got %v", elapsed)
		}

		if elapsed < 1500*time.Millisecond {
			t.Errorf("Execution time too fast, might not be processing correctly: %v", elapsed)
		}
	})
}

func BenchmarkWorkerPool(b *testing.B) {
	b.Run("default setup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			const numJobs = 5
			jobs := make(chan int, numJobs)
			results := make(chan int, numJobs)

			for w := 1; w <= 3; w++ {
				go worker(w, jobs, results)
			}

			for j := 1; j <= numJobs; j++ {
				jobs <- j
			}
			close(jobs)

			for a := 1; a <= numJobs; a++ {
				<-results
			}
		}
	})
}
