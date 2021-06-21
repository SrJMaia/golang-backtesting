// Package fibonnaci generate fibonacci numbers
package fibonacci

// Generate
func Fibo(x int) {
  fmt.Println("Fibonacci Sequence")
  prev, pos := 0, 1
  for ; x != 0; x-- {
    pos += prev
    prev = pos - prev
    }
  fmt.Println(prev)
}
