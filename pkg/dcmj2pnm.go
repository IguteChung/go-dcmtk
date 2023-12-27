package godcmtk

// #cgo LDFLAGS: -lgodcmjpeg -L../dcmtk/dcmjpeg/dist
// extern int dcm2pnm(int argc, char *argv[]);
import "C"

// DCMJ2PNM aaa
func DCMJ2PNM() {
	cStrings, free := StringArray("main", "123.dcm", "123.jpg")
	defer free()
	C.dcm2pnm(3, cStrings)
}
