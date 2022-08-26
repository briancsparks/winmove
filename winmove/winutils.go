package winmove

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/winmove/activedevelopment/grumpy"
  "github.com/gonutz/w32/v2"
  "syscall"
)

var topLevelPopupClassNames map[string]struct{} = map[string]struct{}{
  "X410_XAppWin" : struct{}{},
}


var (
  user32      = syscall.NewLazyDLL("user32.dll")
  getParent   = user32.NewProc("GetParent")
)

func isPopup(hwnd w32.HWND) bool {
  isPopup := false

  var style, styleEx uint32
  _=styleEx

  info, success := w32.GetWindowInfo(w32.HWND(hwnd))
  grumpy.Unlessf(success, "  GetWindowInfo(%x) fail\n", hwnd)
  if success {
    style, styleEx = info.DwStyle, info.DwExStyle
    isPopup = isPopup || (style & w32.WS_POPUP) == w32.WS_POPUP
    isPopup = isPopup || (style & w32.WS_POPUPWINDOW) == w32.WS_POPUPWINDOW
  }

  return isPopup
}

func isAppWindow(hwnd w32.HWND) bool {
  if GetParent(hwnd) == 0 {
    if !isPopup(hwnd) {
      return true
    }

    className := className(hwnd)
    if inSet(topLevelPopupClassNames, className) {
      return true
    }
    return false
  }

  return false
}

func className(hwnd w32.HWND) string {
  //className := ""
  name, success := w32.GetClassName(hwnd)
  grumpy.Unlessf(success, "  GetClassName(%x) fail\n", hwnd)
  if success {
    return name
  }
  return ""
}


func GetParent(of w32.HWND) w32.HWND {
  ret, _, _ := getParent.Call(uintptr(of))
  return w32.HWND(ret)
}


