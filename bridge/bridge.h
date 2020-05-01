

#ifndef BRIDGE_H
#define BRIDGE_H

#include <windows.h>
#include <winuser.h>


struct WindowPos {
    long x, y, w, h;
};
struct WindowPos GetWindowPos(char* name, char* clazz);


struct MousePos {
    long x, y;
};

struct MousePos GetMousePos();


#endif