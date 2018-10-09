package invariant

import (
	"fmt"
	"flag"
)

const (
	development = "development"
	production  = "production"
)

type InvariantError struct {
	What string
	Name string
}

func (e InvariantError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

func (e InvariantError) String() string {
	return fmt.Sprintf("%v", e.What)
}

func Invariant(condition bool, format string, args ...interface{}) error {
	var env string

	flag.StringVar(&env, "env", development, "current environment (development or production)")
	flag.Parse()

	if env != production {
		if format == "" {
			return InvariantError{"invariant requires an error message argument", "Invariant Violation"}
		}
	}

	if !condition {
		var error InvariantError

		if env != development {
			error = InvariantError{"invariant exception in production environment. Please use development flag to see the full error message", "Invariant Violation"}
		} else {
			error = InvariantError{fmt.Sprintf(format, args...), "Invariant Violation"}
		}

		return error
	}
	return nil
}
