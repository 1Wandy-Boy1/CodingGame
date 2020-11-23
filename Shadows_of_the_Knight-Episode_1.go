# CodingGame

package main

import ("fmt"
       "strings")

func main() {
    var W, H int
    fmt.Scan(&W, &H)
    var N int
    fmt.Scan(&N)    
    var X0, Y0 int
    fmt.Scan(&X0, &Y0)

    var X1, Y1 int

    var X2 int
    X2 = W -1

    var Y2 int
    Y2 = H - 1 
    for {
        var bombDir string
        fmt.Scan(&bombDir)
       if strings.ContainsAny(bombDir, "U") == true {
        Y2 = Y0 - 1;
      } else if strings.ContainsAny(bombDir, "D") == true {
        Y1 = Y0 + 1;
      }

      if strings.ContainsAny(bombDir, "L") == true {
        X2 = X0 - 1;
      } else if strings.ContainsAny(bombDir, "R") == true {
        X1 = X0 + 1;
      }
       X0 = X1 + (X2 - X1) / 2
       Y0 = Y1 + (Y2 - Y1) / 2
       
    fmt.Println(X0, Y0)
    }
}
