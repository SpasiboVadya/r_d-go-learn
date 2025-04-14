package generics

/*
Use Interface when:
	- Behavior matters (e.g. Close())	✅
	- Want to mix different types	✅
	- You don’t need compile-time type info	✅

Use Generics when:
	- You operate on values, not behavior	✅
	- Need type safety or constraints	✅
	- Reusable data structures (slices, maps)	✅
*/

import "fmt"

type Closer interface {
	Close() error
}

type File struct{}

func (f File) Close() error {
	fmt.Println("Closed File")
	return nil
}

type Stream struct{}

func (s Stream) Close() error {
	fmt.Println("Closed Stream")
	return nil
}

/* 👎 Overengineering: Generics used unnecessarily
- Adds complexity and type instantiation at compile time
- Doesn’t gain any benefit from generics here — the code only relies on behavior (Close()), not type info
- You can’t mix types even if they all satisfy Close() — no polymorphism */

func CloseAllGeneric[T Closer](items []T) {
	for _, item := range items {
		_ = item.Close()
	}
}

/* ✅ Simpler and idiomatic
- You can pass any mix of types that implement Close()
- No unnecessary generic instantiation
- Easier to read and maintain */

func CloseAllInterface(items []Closer) {
	for _, item := range items {
		_ = item.Close()
	}
}

func ShowInMain() {
	abstractClose := []Closer{File{}, File{}, Stream{}, Stream{}}
	concreteFiles := []File{{}, {}}

	CloseAllGeneric(abstractClose)
	CloseAllGeneric(concreteFiles) // concrete types supported (but no need in that)

	CloseAllInterface(abstractClose)

	// concrete types not supported (and shouldn't be)
	// uncomment to see
	//concreteStreams := []Stream{{}, {}}
	//CloseAllInterface(concreteStreams)
}
