package main 

import ("testing";
		"net/http/httptest";
		"net/http";
		"time")


func TestRace(t *testing.T){

	slowServer := delayedServer(100)
	fastServer := delayedServer(0)


    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got := race(slowURL, fastURL)

    if got != want {
        t.Errorf("got '%s', want '%s'", got, want)
    }

    slowServer.Close()
    fastServer.Close()
}

func delayedServer(t time.Duration) *httptest.Server {
	 return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(t * time.Millisecond)
        w.WriteHeader(http.StatusOK)
    }))
}