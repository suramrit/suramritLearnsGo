package main 

import ("testing";
		"net/http/httptest";
		"net/http";
		"time";
		"fmt")


func TestRace(t *testing.T){

	slowServer :=  delayedServer(2)
	fastServer := delayedServer(0)
	defer slowServer.Close()
    defer fastServer.Close()


    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got,_ := race(slowURL, fastURL)

    if got != want {
        t.Errorf("got '%s', want '%s'", got, want)
    }
}

func TestTimeout(t *testing.T) {
	slowServer := delayedServer(2)
	fastServer := delayedServer(3)
	defer slowServer.Close()
    defer fastServer.Close()

    slowURL := slowServer.URL
    fastURL := fastServer.URL
    testtime := time.Second 
  
 	_,got := customrace(slowURL, fastURL, testtime)    
 	want := httpGetTimeOut 

    if got!=want {
        t.Errorf("Error expected, got nil")
    } else {
    	fmt.Printf("PASS::%v", httpGetTimeOut.Error())
    }

}

func delayedServer(t time.Duration) *httptest.Server {
	 return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(t * time.Second)
        w.WriteHeader(http.StatusOK)
    }))
}