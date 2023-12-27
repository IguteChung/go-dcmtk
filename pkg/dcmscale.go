package godcmtk

// #cgo LDFLAGS: -lgodcmimage -L../dcmtk/dcmimage/dist
// extern int dcmscale(int argc, char *argv[]);
import "C"

// DCMSCALE aaa
func DCMSCALE() {
	cStrings, free := StringArray("main", "+Sxv", "200", "+un", "123.dcm", "456.dcm")
	defer free()
	C.dcmscale(6, cStrings)
}
