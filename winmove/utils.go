package winmove

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "fmt"
)

func u() {
  fmt.Printf("\n")
}


func inSet(s map[string]struct{}, key string) bool {
  _, present := s[key]
  return present
}

func scaleInt(n int, x float64) int {
  return int(float64(n) * 1.25)
}

func scaleInt32(n int32, x float64) int32 {
  return int32(float64(n) * 1.25)
}

