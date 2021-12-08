# BigFibonacciNumbers

## Compiling and running
1. Clone the repository.
```
git clone https://github.com/fabiancdng/BigFibonacciNumbers
```

2. `cd` into the folder.
```
cd BigFibonacciNumbers/
```

3. Compilation using the [Go compiler](https://golang.org).
```
go build -o BigFibonacciNumbers ./cmd/bigfibonaccinumbers/main.go
```

4. All set! Run the program using:
```
./BigFibonacciNumbers
```

## Installing the package
> For example for use in other projects
```
go get -u github.com/fabiancdng/BigFibonacciNumbers
```


## Example implementation
```
package main

import (
	"fmt"

	"github.com/fabiancdng/BigFibonacciNumbers/pkg/fibcalculator"
)

func main() {
	result := fibcalculator.FibCalculate(3)
	fmt.Println(result)
}
```

**Copyright &copy; 2021 Fabian Reinders (fabiancdng)**