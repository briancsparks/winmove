package grumpy

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

func Iff(b bool, format string, a ...any) {
  if b {
    Becausef(format, a...)
  }
}

func If(b bool, a ...any) {
  if b {
    Because(a...)
  }
}


func Unlessf(b bool, format string, a ...any) {
  if !b {
    Becausef(format, a...)
  }
}

func Unless(b bool, a ...any) {
  if !b {
    Because(a...)
  }
}



func DieIff(b bool, format string, a ...any) {
  if b {
    Dief(format, a...)
  }
}

func DieIf(b bool, a ...any) {
  if b {
    Die(a...)
  }
}


func DieUnlessf(b bool, format string, a ...any) {
  if !b {
    Dief(format, a...)
  }
}

func DieUnless(b bool, a ...any) {
  if !b {
    Die(a...)
  }
}



func PanicIff(b bool, format string, a ...any) {
  if b {
    Panicf(format, a...)
  }
}

func PanicIf(b bool, a ...any) {
  if b {
    Panic(a...)
  }
}


func PanicUnlessf(b bool, format string, a ...any) {
  if !b {
    Panicf(format, a...)
  }
}

func PanicUnless(b bool, a ...any) {
  if !b {
    Panic(a...)
  }
}



func FatalIff(b bool, format string, a ...any) {
  if b {
    Fatalf(format, a...)
  }
}

func FatalIf(b bool, a ...any) {
  if b {
    Fatal(a...)
  }
}


func FatalUnlessf(b bool, format string, a ...any) {
  if !b {
    Fatalf(format, a...)
  }
}

func FatalUnless(b bool, a ...any) {
  if !b {
    Fatal(a...)
  }
}



