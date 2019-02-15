//Working on https://gophercises.com/exercises/ ---- check out for small projects on golang! 

//Disclaimer: The code was developed while following execises at https://gophercises.com/exercises/ and were solely developed from an educational perspective. 


package readhtml

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

//"net/httptest")

var exHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page
  <span> which is totally awesome!</span>
  </a>
</body>
</html>`

func TestConvert(t *testing.T) {
	t.Run("Test Convert", func(t *testing.T) {
		var ret map[string]interface{}
		e := json.Unmarshal(Convert("sample.html"), &ret)
		if e != nil {
			fmt.Println(e.Error())
		}
		got := ret["href"]
		want := "/wikipedia"
		if got != want {
			t.Errorf("Wanted %v, got %v\n", want, got)
		}
	})
}

func TestParse(t *testing.T) {

	t.Run("Test io.reader parse", func(t *testing.T) {
		fmt.Println("Testing new parse")

		page := strings.NewReader(exHtml)

		docs, err := Parse(page)

		if err != nil {
			t.Errorf("Incorrect Parsing! ")
		}

		fmt.Printf("Got::%+v", docs)
	})

}
