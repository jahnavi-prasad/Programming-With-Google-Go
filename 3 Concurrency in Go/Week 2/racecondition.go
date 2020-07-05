package main

import "fmt"

func main() {
	a := 0
	time := 10000
	for i := 0; i < time; i++ {
		go func() {
			a++
		}()
	}
	fmt.Printf("a: %d\n", a)
}

/*
Race Condition : A race condition or race hazard is the condition of an electronics,
software, or other system where the system's substantive behavior is dependent
on the sequence or timing of other uncontrollable events.
It becomes a bug when one or more of the possible behaviors is undesirable.

Why Race condition occurs?
It occurs when two or more process can access and change the shared data at the same time.
It occurred because there were conflicting accesses to a resource .
Critical section problem may cause race condition.
*/
