# go_effect_sort
Golang slice sorting library that takes user functions for comparison and element movement

## Installation
Run the following command in your project directory:
```golang
go get github.com/gabe-lee/go_effect_sort@latest
```
## Usage Example
```golang
import esort "github.com/gabe-lee/go_effect_sort"

func main() {
    var mySlice = []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}
    fmt.Printf("%s\n", mySlice)
    esort.InsertionSortSwap(mySlice, esort.GreaterThanOrd, esort.SwapNoSideEffect)
    fmt.Printf("%s\n", mySlice)
}
```
Output:
```
Hello World!
 !HWdellloor
```
## Supported Algorithms
  - Insertion Sort (Remove compare val, move all larger vals up, re-insert compare val
  - Insertion Sort (swapping each element like a Bubble Sort)
