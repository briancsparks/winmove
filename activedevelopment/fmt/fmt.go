package afmt

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "fmt"
)

func Printf(format string, a ...any) (n int, err error) {
  //return 0, nil
  return fmt.Printf(format, a...)
}

func Print(a ...any) (n int, err error) {
  return fmt.Print(a...)
}

func Println(a ...any) (n int, err error) {
  return fmt.Println(a...)
}


