package main

var NoopSignal = struct{}{}

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

func NoopPromise(length int) (func(), func(error), func() error) {
	errChan := make(chan error)
	signalChan := make(chan struct{}, length)
	resolve, reject, status := StatusPromise(signalChan, errChan)
	pass := func() {
		resolve(NoopSignal)
	}
	cry := func(err error) {
		reject(err)
	}
	wait := func() error {
		for i := 0; i < length; i++ {
			_, err := status()
			if err != nil {
				return err
			}
		}
		return nil
	}
	return pass, cry, wait
}
