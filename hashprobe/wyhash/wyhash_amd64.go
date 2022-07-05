//go:build !purego

package wyhash

//go:noescape
func MultiHash32(hashes []uintptr, values []uint32, seed uintptr)

//go:noescape
func MultiHash64(hashes []uintptr, values []uint64, seed uintptr)
