package main
import (
	"fmt"
	"github.com/snwfdhmp/errlog"
	"errors"
	
)
func main() {
    fmt.Println("Program start")

    wrapingFunc() //call to our important function

    fmt.Println("Program end")
}

func wrapingFunc() {
    someBigFunction() // call some func 
}

func someBigFunction() {
    someDumbFunction() // just random calls
    someSmallFunction() // just random calls
    someDumbFunction() // just random calls

    // Here it can fail, so instead of `if err  != nil` we use `errlog.Debug(err)`
    if err := someNastyFunction(); errlog.Debug(err) {
        return
    }

    someSmallFunction() // just random calls
    someDumbFunction() // just random calls
}

func someSmallFunction() {
    _ = fmt.Sprintf("I do things !")
}

func someNastyFunction() error {
    return errors.New("I'm failing for some reason") // simulate an error
}

func someDumbFunction() bool {
    return false // just random things
}