package godcmtk

// #cgo LDFLAGS: -lgodcmdata -L../dcmtk/dcmdata/dist
// extern int dcmodify(int argc, char *argv[]);
import "C"

// DCMODIFY aaa
func DCMODIFY() {
	cStrings, free := StringArray("main", "123.dcm")
	defer free()
	C.dcmodify(2, cStrings)
}
