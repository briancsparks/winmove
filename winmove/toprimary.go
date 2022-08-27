package winmove

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/briancsparks/winmove/activedevelopment/debug"
  "github.com/briancsparks/winmove/activedevelopment/grumpy"
  "github.com/briancsparks/winmove/activedevelopment/verbose/vv"

  "github.com/gonutz/w32/v2"
)

func ToPrimary2() {
  monitors := Monitors()
  primaryMonitor := monitors[0]
  grumpy.PanicIf(primaryMonitor == 0, "No primary monitor")

  if len(monitors) <= 1 {
    grumpy.Because("No secondary monitor - nothing to do")
    return
  }

  minfo := GetMonitorInfo(primaryMonitor)
  workArea := minfo.RcWork
  workArea = ShrinkBy(workArea, 0.10)

  desktop := w32.GetDesktopWindow()

  w32.EnumChildWindows(desktop, func (hwnd w32.HWND) bool {
    hmonitor := w32.MonitorFromWindow(hwnd, 0)

    if hmonitor != primaryMonitor {
      isAppWindow   := isAppWindow(hwnd)

      debug.Printf("\nWindow %8x: %5t  (%v)\n", hwnd, isAppWindow, w32.GetWindowText(hwnd))
      if !isAppWindow {
        vv.Printf("  class: %s\n", className(hwnd))
        vv.Printf("  popup: %t\n", isPopup(hwnd))
      }

      // Move to primary monitor
      if isAppWindow {
        debug.Printf("SetWindowPosW(%x, %v)\n", hwnd, workArea)

        //success := SetWindowPosW(hwnd, workArea)
        //grumpy.Unlessf(success, "SetWindowPosW(%x, %v) fail\n", hwnd, workArea)

        //// Do just one
        //break
      }
    }

    return true
  })

}


func ToPrimary() {
  var lmpi w32.MONITORINFO

  // First scan
  var primaryMonitor, secondaryMonitor w32.HMONITOR
  var onWrongMonitor []w32.HWND

  desktop := w32.GetDesktopWindow()

  w32.EnumChildWindows(desktop, func (hwnd w32.HWND) bool {
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
  if secondaryMonitor == 0 {
    grumpy.Because("No secondary monitor - nothing to do")
    return
  }
  if len(onWrongMonitor) == 0 {
    grumpy.Because("No windows found on secondary monitor")
    return
  }

  minfo := GetMonitorInfo(primaryMonitor)
  workArea := minfo.RcWork
  workArea = ShrinkBy(workArea, 0.10)

  fmt.Printf("work area: %+v\n", minfo.RcWork)

  for i, hwnd := range onWrongMonitor {

    isAppWindow   := isAppWindow(hwnd)

    debug.Printf("\n%02d: Window %8x: %5t  (%v)\n", i, hwnd, isAppWindow, w32.GetWindowText(hwnd))
    if !isAppWindow {
     vv.Printf("  class: %s\n", className(hwnd))
     vv.Printf("  popup: %t\n", isPopup(hwnd))
    }

    // Move to primary monitor
    if isAppWindow {
      debug.Printf("SetWindowPosW(%x, %v)\n", hwnd, workArea)

      //success := SetWindowPosW(hwnd, workArea)
      //grumpy.Unlessf(success, "SetWindowPosW(%x, %v) fail\n", hwnd, workArea)

      //// Do just one
      //break
    }
  }

  return
}



