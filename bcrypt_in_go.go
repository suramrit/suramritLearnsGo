package main

import("fmt"
 		"golang.org/x/crypto/bcrypt" // -- go get golang.org/x/crypto/bcrypt
)

func main() {
	p := "123123" //bad password 
	p_encrypt , err := bcrypt.GenerateFromPassword([]byte(gp), 5)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(gp)
	fmt.Printf("%s\n", p_encrypt) // good hash 
}