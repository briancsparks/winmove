package vvvx

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/briancsparks/winmove/activedevelopment"
)

func Printf(format string, a ...any) (int, error) {
  if !shouldVvvx() {
    return 0, nil
  }

  return fmt.Printf(format, a...)
}

func Print(a ...any) (int, error) {
  if !shouldVvvx() {
    return 0, nil
  }

  return fmt.Print(a...)
}

func Println(a ...any) (int, error) {
  if !shouldVvvx() {
    return 0, nil
  }

  return fmt.Println(a...)
}

func shouldVvvx() bool {
  return (activedevelopment.ConfigVerbosity() >= 3) && activedevelopment.DoX()
}
