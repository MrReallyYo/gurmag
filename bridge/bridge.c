#include "bridge.h"

struct WindowPos GetWindowPos(char *name, char *clazz) {
  HWND hWnd = FindWindowExA(0, 0, clazz, name);
  struct WindowPos pos;
  if (0 != hWnd) {
    RECT rect;
    GetWindowRect(hWnd, &rect);
    pos.x = rect.left;
    pos.y = rect.top;
    pos.w = rect.right - rect.left;
    pos.h = rect.bottom - rect.top;
  } else {
    pos.x, pos.y, pos.w, pos.h = -1;
  }
  return pos;
}

struct MousePos GetMousePos() {
  POINT mouse;
  boolean bla = GetCursorPos(&mouse);
  struct MousePos pos;
  pos.x = mouse.x;
  pos.y = mouse.y;
  return pos;
}