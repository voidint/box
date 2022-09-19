package fn

import "time"

// Retry Retry function call
func Retry(retries int, retryDelay time.Duration, do func() (again bool, err error)) (err error) {
	var again bool

	if retries < 0 {
		for {
			if again, err = do(); !again {
				break
			}
			time.Sleep(retryDelay)
		}

	} else {
		for i := 0; i <= retries; i++ {
			if again, err = do(); !again {
				break
			}
			time.Sleep(retryDelay)
		}
	}

	return err
}
