package main

import "os"

func ExampleMain_lowercase() {
	os.Args = []string{"spongebob", "these", "are", "claps"}
	main()
	// Output: these 👏 are 👏 claps
}

func ExampleMain_uppercase() {
	os.Args = []string{"spongebob", "what", "if", "there", "is", "a", "👏", "in", "it"}
	main()
	// Output: what 👏 if 👏 there 👏 is 👏 a 👏 in 👏 it
}
