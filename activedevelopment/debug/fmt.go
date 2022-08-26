package debug

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/briancsparks/winmove/activedevelopment"
)

func Printf(format string, a ...any) {
  if !shouldDebug() {
    return
  }

  fmt.Printf(format, a...)
}

func Print(a ...any) {
  if !shouldDebug() {
    return
  }

  fmt.Print(a...)
}

func Println(a ...any) {
  if !shouldDebug() {
    return
  }

  fmt.Println(a...)
}



func shouldDebug() bool {
  return activedevelopment.ConfigVerbosity() >= 1
}


