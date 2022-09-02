package activedevelopment

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// -------------------------------------------------------------------------------------------------------------------

type Config struct {
  ActiveD         bool
  ShowDebug       bool
  LogApis         bool
  Verbosity       int
  DoX             bool

  FeatureFlags    map[string]struct{}
  FeatureTags     map[string]string
  FeatureTagsEx   map[string]interface{}
}

// -------------------------------------------------------------------------------------------------------------------

var theConfig *Config

func init() {
  theConfig = &Config{}

  theConfig.ActiveD     = true

  // Release
  theConfig.ShowDebug = false
  theConfig.Verbosity = 0
  theConfig.DoX       = false
  theConfig.LogApis   = false

  if !theConfig.ActiveD {
    // Debug / ActiveDevelopment
    theConfig.ShowDebug   = true
    theConfig.Verbosity   = 1
    theConfig.DoX         = true
    theConfig.LogApis     = true
  }
}

// -------------------------------------------------------------------------------------------------------------------

func ConfigIsActiveD() bool {
  return theConfig.ActiveD
}

// -------------------------------------------------------------------------------------------------------------------

func ActiveD() bool {
  return ConfigIsActiveD()
}

// -------------------------------------------------------------------------------------------------------------------

func IsProd() bool {
  return !ActiveD()
}

// -------------------------------------------------------------------------------------------------------------------

func ConfigVerbosity() int {
  return theConfig.Verbosity
}

// -------------------------------------------------------------------------------------------------------------------

func Verbosity() int {
  return theConfig.Verbosity
}

// -------------------------------------------------------------------------------------------------------------------

func DoX() bool {
  return theConfig.DoX
}

// -------------------------------------------------------------------------------------------------------------------

func ConfigIf(n int) bool {
  if n == 0 {
    return false
  }
  return true
}

// -------------------------------------------------------------------------------------------------------------------

func FeatureFlag(flag string) bool {
  _, present := theConfig.FeatureTags[flag]
  return present
}

// -------------------------------------------------------------------------------------------------------------------

func OnFeatureFlag(flag string, fn func() error) error {
  if FeatureFlag(flag) {
    return fn()
  }
  return nil
}

