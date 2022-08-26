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

