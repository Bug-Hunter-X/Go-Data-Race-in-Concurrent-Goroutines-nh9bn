```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock()
                        count += i
                        mu.Unlock()
                }(i)
        }
        wg.Wait()
        fmt.Println("Count:", count) // Output is consistent
}
```