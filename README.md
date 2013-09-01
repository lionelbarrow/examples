### Examples

A lightweight behavioral testing library for Go. 

Here's how it works: Write regular tests using the Go standard library, making your assertions with the [quiz library](https://github.com/benmills/quiz). If you ever feel like you're writing tests that should be grouped together, use a `When` or `Describe` function to group them, like this:

```go
import (                                                           
  e "github.com/lionelbarrow/examples"                             
  "testing"                                                        
)                                                                  
                                                                   
func TestMyClass(t *testing.T) {                                   
  e.When("writing software", t,                                    
    e.It("helps to have tests", func(ex *e.Example) {              
      ex.Expect(1).ToEqual(1)                                      
    }),                                                            
                                                                   
    e.It("helps to have correct tests", func(ex *e.Example) {      
      ex.Expect(1).ToEqual(2)                                      
    }),                                                            
  )                                                                
}
```

Run the tests with `go test` as usual. With the code above, this will output:

```
--- FAIL: TestMyClass (0.00 seconds)                                         
examples.go:32: When writing software it helps to have correct tests:
                Expected 2 to equal 1.                                       
                /Users/lionel/fake_examples/fake_test.go:15                
```

As far as the `go test` command is concerned, everything within the `TestMyClass` function is the same test. The Go standard library does a number of useful things (such as automatic parallelization) to tests run through the `go test` command, so it makes sense to keep your test groups small.
