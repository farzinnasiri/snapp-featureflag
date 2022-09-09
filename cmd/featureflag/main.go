package main

import "snapp-featureflag/internal/app/featureflag"

func main() {
	app, err := featureflag.CreateApp()
	if err != nil {
		panic(err)
	}

	app.Run()

}
