package main

func StatusPromise[T any](ch chan T, err chan error) (func(T), func(error), func() (T, error)) {
	resolve := func(result T) {
		ch <- result
	}
	reject := func(errValue error) {
		err <- errValue
	}

	status := func() (T, error) {
		select {
		case value := <-ch:
			return value, nil
		case errValue := <-err:
			var zero T
			return zero, errValue
		}
	}

	return resolve, reject, status
}
