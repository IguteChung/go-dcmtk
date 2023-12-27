package main

import godcmtk "go-dcmtk/pkg"

func main() {
	godcmtk.DCMDUMP("dcm/123.dcm", nil)
	// godcmtk.DCMODIFY()
	// godcmtk.DCMSCALE()
	// godcmtk.DCMJ2PNM()
}
