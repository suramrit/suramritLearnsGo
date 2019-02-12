package main 

import ("time";
		"net/http")
func race(u1, u2 string) string{


	aDuration := urltime(u1)
	bDuration := urltime(u2)

	if aDuration > bDuration {
		return u2
	}

	return u1
}

func urltime(u string) time.Duration {
	start := time.Now()
    http.Get(u)
    duration := time.Since(start)
    return duration
}