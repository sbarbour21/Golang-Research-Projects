package main

import "fmt"

func main() {
    const a = 42
    const b float32 = 3

    var i int = a
    var f32 float32 = b
    var f64 float64 = float64(b)

    fmt.Println(i)
    fmt.Println(f32, f64)

    const c = iota
    fmt.Println(c)

    const (
        d = 2 * 5
        e
        f = iota
        g
        h = 10 * iota
    )

    fmt.Println(d, e, f, g, h)
}
