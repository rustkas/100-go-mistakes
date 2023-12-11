package main

import (
	"flag"
	"log"
	"net/http"
)

var tracing bool

func main() {
	// Use a flag to determine which function to run.
	function := flag.String("function", "", "Specify the function to run (listing1, listing2, listing3, listing4)")
	flag.Parse()

	// Call the selected function.
	switch *function {
	case "listing1":
		err := listing1()
		handleError(err)
	case "listing2":
		err := listing2()
		handleError(err)
	case "listing3":
		err := listing3()
		handleError(err)
	case "listing4":
		err := listing4()
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

func listing1() error {
	var client *http.Client
	tracing = true
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		_ = client
	}

	_ = client
	return nil
}

func listing2() error {
	var client *http.Client
	tracing = true
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c
		log.Println(client)
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}

func listing3() error {
	var client *http.Client
	var err error
	tracing = true
	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

func listing4() error {
	var client *http.Client
	var err error
	tracing = true
	if tracing {
		client, err = createClientWithTracing()
		log.Println(client)
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	_ = client
	return nil
}

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
