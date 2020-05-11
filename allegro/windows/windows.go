// Package windows provides support for Allegro's Windows-specific functions.
package windows

// #include <allegro5/allegro_windows.h>
import "C"
import (
	"github.com/dradtke/go-allegro/allegro"
)

// Returns the handle to the window that the passed display is using.
func WindowHandle(d *allegro.Display) uintptr {
	hwnd := C.al_get_win_window_handle((*C.ALLEGRO_DISPLAY)(d))
	return hwnd
}
