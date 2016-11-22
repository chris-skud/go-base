package main

import "fmt"

// Hook func
type Hook func(msg string)

// SharedLibOfSomeSort is a thing that holds a ref to the hook func
type SharedLibOfSomeSort struct {
	Hook
}

// DoSomething demonstrates some shared package behavior
func (s *SharedLibOfSomeSort) DoSomething() {
	// do some stuff.

	// now call injected hook to do whatever the caller wanted.
	s.Hook("hey the thing worked")
}

func main() {
	funky := func(m string) {

		// now i can do something based-on or what the shared lib passes in
		if m == "fail" {
			panic("something bad happened")
		}

		// or decorate a string for logging
		m = fmt.Sprintf("The stuff i want to add to the log: %s", m)
		fmt.Println(m)

	}
	lib := &SharedLibOfSomeSort{Hook: funky}
	lib.DoSomething()
}
