package grumpy

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "fmt"
  "log"
)

func Becausef(format string, a ...any) {
  fmt.Printf(format, a...)
}

func Because(a ...any) {
  fmt.Println(a...)
}


func Dief(format string, a ...any) {
  log.Panicf(format, a...)
}

func Die(a ...any) {
  log.Panic(a...)
}


func Panicf(format string, a ...any) {
  log.Panicf(format, a...)
}

func Panic(a ...any) {
  log.Panic(a...)
}


func Fatalf(format string, a ...any) {
  log.Fatalf(format, a...)
}

func Fatal(a ...any) {
  log.Fatal(a...)
}


