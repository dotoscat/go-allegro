// Package physfs provides support for Allegro's PhysicsFS addon.
package physfs

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_physfs.h>
import "C"

// This function sets both the ALLEGRO_FILE_INTERFACE and ALLEGRO_FS_INTERFACE
// for the calling thread.
func UseFileInterface() {
	C.al_set_physfs_file_interface()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_physfs_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}
