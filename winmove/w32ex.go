package winmove

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/winmove/activedevelopment/grumpy"
  "github.com/gonutz/w32/v2"
  "syscall"
)


// -------------------------------------------------------------------------------------------------------------------

var (
  user32      = syscall.NewLazyDLL("user32.dll")
  getParent   = user32.NewProc("GetParent")
)

// -------------------------------------------------------------------------------------------------------------------

func EnumDisplayMonitors(hdc w32.HDC, clip *w32.RECT, fnEnum func(hmonitor w32.HMONITOR, hdc w32.HDC, lprect *w32.RECT, lparam w32.LPARAM) bool, dwData uintptr) bool {
  f := syscall.NewCallback(func(hmonitor w32.HMONITOR, hdc w32.HDC, lprect *w32.RECT, lparam w32.LPARAM) bool {
    if fnEnum(hmonitor, hdc, lprect, lparam) {
      return true
    }
    return false
  })
  return w32.EnumDisplayMonitors(hdc, clip, f, dwData)
}

// -------------------------------------------------------------------------------------------------------------------

func GetMonitorInfo(hmonitor w32.HMONITOR) w32.MONITORINFO {
  var lmpi w32.MONITORINFO

  success := w32.GetMonitorInfo(hmonitor, &lmpi)
  grumpy.Unlessf(success, "  GetMonitorInfo(%x) fail\n", hmonitor)

  return lmpi
}

// -------------------------------------------------------------------------------------------------------------------

func GetParent(of w32.HWND) w32.HWND {
  ret, _, _ := getParent.Call(uintptr(of))
  return w32.HWND(ret)
}

