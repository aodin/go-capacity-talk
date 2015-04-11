# Go Capacity Talk

A [wat](https://www.destroyallsoftware.com/talks/wat) moment in Go.

1. Walkthrough the `recurse` function in `capacity.go`
2. Run the `TestSimple` test with `go test -run=TestSimple`
3. Run the `TestComplex` test with `go test -run=TestComplex` - it should fail
4. ???
5. Profit

The answer lies in `TestCapacity` and http://blog.golang.org/go-slices-usage-and-internals

- aodin, 2015
