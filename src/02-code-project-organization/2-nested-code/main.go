package main

import (
	"errors"
	"flag"
	"log"
)

func main() {
	// Use a flag to determine which function to run.
	function := flag.String("function", "", "Specify the function to run (listing1, listing2, listing3, listing4)")
	arg1 := flag.String("arg1", "", "Argument 1 for join functions")
	arg2 := flag.String("arg2", "", "Argument 2 for join functions")
	arg3 := flag.Int("arg3", 0, "Argument 3 for join functions")

	flag.Parse()

	// Call the selected function.
	switch *function {
	case "join1":
		str, err := join1(*arg1, *arg2, *arg3)
		log.Printf("%v", str)
		handleError(err)
	case "join2":
		str, err := join2(*arg1, *arg2, *arg3)
		log.Printf("%v", str)
		handleError(err)
	default:
		log.Println("Invalid function specified. Please use -function flag with one of: listing1, listing2, listing3, listing4")
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return s1 + s2, nil
}
