package main 

import ("time";
		"net/http";
		"errors")

var httpGetTimeOut = errors.New("Time out while getting response!")
const tenSecondTimeOut = time.Second * 10

func race(u1,u2 string) (string,error) {
	return customrace(u1,u2,tenSecondTimeOut)

}


func customrace(u1, u2 string, t time.Duration) (string, error) {
		select {
			case <-ping(u1) :
				return u1,nil
			case <-ping(u2) :
				return u2,nil
			case <-time.After(t):
        		return "", httpGetTimeOut
		}
}

/*func urltime(u string) time.Duration {
	start := time.Now()
    http.Get(u)
    duration := time.Since(start)
    return duration
}*/

func ping(u string) (chan bool) {
	ch := make (chan bool)

	go func(){
		http.Get(u)
		ch <- true
		}()
	return ch
}