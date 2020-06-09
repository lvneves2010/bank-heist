package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true

  eludedGuards := rand.Intn(100)
  fmt.Println(eludedGuards)
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
    return
  }

  openedVault := rand.Intn(100)
  fmt.Println(openedVault)
  if isHeistOn && openedVault >= 50 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can't be opened!!")
    return
  }

  leftSafely := rand.Intn(5)
  fmt.Println(leftSafely)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Busted getting Out!!")
      case 1:
        isHeistOn = false
        fmt.Println("Turns out vault doors don't open from the inside...")
      case 2:
        isHeistOn = false
        fmt.Println("The Cops arrived earlier!!")
      default:
        isHeistOn = true
        fmt.Println("Drive!!!!!")
    }
  }
  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("We got away with: $", amtStolen)
  }
}
