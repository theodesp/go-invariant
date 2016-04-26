package main

import (
        "fmt"        
		"github.com/namsral/flag"
)

const (
	development = iota
	production 
)

type InvariantError struct {
	What string
}

func (e InvariantError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

func invariant(condition bool, format string, args ...interface{}) error {
	var env int
	
	flag.IntVar(&env, "environment", development, "current environment (development or production)")
    flag.Parse()
	
	if env != production {
	    if format == "" {
	      return InvariantError{"invariant requires an error message argument"};
	    }
 	}
	return nil
}

func main() {
	if err := invariant(1==0, "qweqew"); err != nil {
		fmt.Println(err)
	}
}

