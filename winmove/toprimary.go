package winmove

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/winmove/activedevelopment/debug"
  "github.com/briancsparks/winmove/activedevelopment/grumpy"
  "github.com/briancsparks/winmove/activedevelopment/verbose/vv"

  "github.com/gonutz/w32/v2"
)



func ToPrimary() {
  var lmpi w32.MONITORINFO

  // First scan
  var primaryMonitor, secondaryMonitor w32.HMONITOR
  var onWrongMonitor []w32.HWND

  desktop := w32.GetDesktopWindow()

  w32.EnumChildWindows(desktop, func (hwnd w32.HWND) bool {
    winfo, success := w32.GetWindowInfo(hwnd)
    _= winfo
    grumpy.PanicIf(!success, "GetWindowInfo fail")

    hmon := w32.MonitorFromWindow(hwnd, 0)
    w32.GetMonitorInfo(hmon, &lmpi)

    if (lmpi.DwFlags & w32.MONITORINFOF_PRIMARY) == w32.MONITORINFOF_PRIMARY {
      primaryMonitor = hmon
    } else {
      secondaryMonitor = hmon
      if GetParent(hwnd) == 0 {
        onWrongMonitor = append(onWrongMonitor, hwnd)
      }
    }

    return true
  })
  grumpy.PanicIf(primaryMonitor == 0, "No primary monitor")
  //if primaryMonitor == 0 {
  //  log.Panic("No primary monitor")
  //}
  if secondaryMonitor == 0 {
    grumpy.Because("No secondary monitor - nothing to do")
    return
  }
  if len(onWrongMonitor) == 0 {
    grumpy.Because("No windows found on secondary monitor")
    return
  }

  for _, hwnd := range onWrongMonitor {
    //vvvx.Printf("Window %v:  (%v)\n", hwnd, w32.GetWindowText(hwnd))
    //vvvx.Printf("  Text: %v\n", w32.GetWindowText(hwnd))
    //vvvx.Printf("  Rect: %v\n", w32.GetWindowRect(hwnd))

    //owner := w32.GetWindow(hwnd, w32.GW_OWNER)
    //_= owner
    //vvvx.Printf("  owner: %v\n", owner)
    //vvvx.Printf("  desktop: %v\n", desktop)
    //if owner == desktop {
    //  vvvx.Println("  ------------------------------- top")
    //}

    className := ""
    name, success := w32.GetClassName(hwnd)
    grumpy.Unlessf(success, "  GetClassName(%x) fail\n", hwnd)
    if success {
      className = name
    }

    //myHWND := MyHWND(hwnd)
    //isPopup := myHWND.isPopup()

    isPopup := isPopup(hwnd)
/*    isPopup := false
    var style, styleEx uint32
    _=styleEx

    info, success := w32.GetWindowInfo(hwnd)
    grumpy.Unlessf(success, "  GetWindowInfo(%x) fail\n", hwnd)
    if success {
      style, styleEx = info.DwStyle, info.DwExStyle
      isPopup = isPopup || (style & w32.WS_POPUP) == w32.WS_POPUP
      isPopup = isPopup || (style & w32.WS_POPUPWINDOW) == w32.WS_POPUPWINDOW
    }
*/
    isAppWindow := isAppWindow(hwnd)
/*    isAppWindow := false
    _= isAppWindow

    parent := GetParent(hwnd)
    if parent == 0 {
      if !isPopup {
        isAppWindow = true
      } else {
        isAppWindow = false
        if inSet(topLevelPopupClassNames, className) {
         isAppWindow = true
        }
      }
    }
*/
    debug.Printf("Window %8x: %5t  (%v)\n", hwnd, isAppWindow, w32.GetWindowText(hwnd))
    if !isAppWindow {
     vv.Printf("  class: %s\n", className)
     vv.Printf("  popup: %t\n", isPopup)
    }

    //if isAppWindow {
    //  debug.Printf("Window %x:  (%v) popup? %t\n", hwnd, w32.GetWindowText(hwnd), isPopup)
    //}
  }

  return
}
