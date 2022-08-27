package winmove

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/winmove/activedevelopment/grumpy"
  "github.com/gonutz/w32/v2"
  "image"
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

func Monitors() []w32.HMONITOR {

  r2 := []w32.HMONITOR{0}
  EnumDisplayMonitors(0, nil, func(hmonitor w32.HMONITOR, hdc w32.HDC, lprect *w32.RECT, lparam w32.LPARAM) bool {
    if isPrimaryMonitor(hmonitor) {
      r2[0] = hmonitor
    } else {
      r2 = append(r2, hmonitor)
    }

    return true
  }, 0)

  return r2
}

func isPrimaryMonitor(hmonitor w32.HMONITOR) bool {
  var lmpi w32.MONITORINFO
  success := w32.GetMonitorInfo(hmonitor, &lmpi)
  grumpy.Unlessf(success, "  GetMonitorInfo(%x) fail\n", hmonitor)
  return success && ((lmpi.DwFlags & w32.MONITORINFOF_PRIMARY) == w32.MONITORINFOF_PRIMARY)
}

// --------------------------------------------------------------------------------------------------------------------

func SetWindowPos(hwnd w32.HWND, x, y, dx, dy int) bool {
  if x == 0 || y == 0 || dx == 0 || dy == 0 {
    grumpy.Becausef("NOT setting windows pos for %x: Rect: x: %v, y: %v, dx: %v, dy: %v\n", hwnd, x, y, dx, dy)
    return false
  }

  //vvvx.Printf("setting windows pos for %v (0x%x): Rect: x: %v, y: %v, dx: %v, dy: %v\n", hwnd, hwnd, x, y, dx, dy)
  return w32.SetWindowPos(hwnd, w32.HWND_TOP, x, y, dx, dy, w32.SWP_NOZORDER)
}

// --------------------------------------------------------------------------------------------------------------------

func SetWindowPosW(hwnd w32.HWND, rect w32.RECT) bool {
  //vvvx.Printf("SetWindowPosW(%x): rect: %+v\n", hwnd, rect)
  return SetWindowPos(hwnd, int(rect.Left), int(rect.Top), Width(rect), Height(rect))
}


// ====================================================================================================================

// --------------------------------------------------------------------------------------------------------------------

func WinRECT(ir image.Rectangle) w32.RECT {
 rect := w32.RECT{Left: int32(ir.Min.X), Top: int32(ir.Min.Y), Right: int32(ir.Max.X), Bottom: int32(ir.Max.Y)}
 return rect
}

// --------------------------------------------------------------------------------------------------------------------

func ImageRectangle(wr w32.RECT) image.Rectangle {
 rect := image.Rect(int(wr.Left), int(wr.Top), int(wr.Right), int(wr.Bottom))
 return rect
}

// --------------------------------------------------------------------------------------------------------------------

func Width(r w32.RECT) int {
  return int(r.Right - r.Left)
}

func Height(r w32.RECT) int {
  return int(r.Bottom - r.Top)
}

func Width32(r w32.RECT) int32 {
  return r.Right - r.Left
}

func Height32(r w32.RECT) int32 {
  return r.Bottom - r.Top
}

func ShrinkBy(r w32.RECT, delta float64) w32.RECT {
  w, h := Width32(r), Height32(r)
  dx, dy := int32(float64(w) * delta), int32(float64(h) * delta)

  return Shrink(r, dx, dy)
}

func Shrink(r w32.RECT, dx, dy int32) w32.RECT {
  dxOn2, dyOn2 := dx/2, dy/2

  return w32.RECT{
    Left:   r.Left + dxOn2,
    Top:    r.Top + dyOn2,
    Right:  r.Right - dxOn2,
    Bottom: r.Bottom - dyOn2,
  }
}
