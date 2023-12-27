package godcmtk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDCMDUMP(t *testing.T) {
	b, _ := os.ReadFile("testdata/dump.txt")
	result, err := parseOutput(string(b))
	assert.EqualValues(t, &DumpResult{
		Files: []*DumpFile{
			&DumpFile{
				Name: "dcm/123.dcm",
				Tags: map[string]*DumpTag{
					"0002,0000": &DumpTag{
						Tag:         "0002,0000",
						VR:          "UL",
						Value:       "208",
						Description: "FileMetaInformationGroupLength",
					},
					"0002,0001": &DumpTag{
						Tag:         "0002,0001",
						VR:          "OB",
						Value:       `00\01`,
						Description: "FileMetaInformationVersion",
					},
					"0002,0003": &DumpTag{
						Tag:         "0002,0003",
						VR:          "UI",
						Value:       "[1.3.6.1.4.1.14519.5.2.1.2857.5885.216794745831170257495208176572]",
						Description: "MediaStorageSOPInstanceUID",
					},
					"0032,1064": &DumpTag{
						Tag:         "0032,1064",
						VR:          "SQ",
						Value:       "(Sequence with explicit length #=1)",
						Description: "RequestedProcedureCodeSequence",
					},
					"fffe,e000": &DumpTag{
						Tag:         "fffe,e000",
						VR:          "na",
						Value:       "(Item with explicit length #=3)",
						Description: "Item",
					},
					"7fe0,0010": &DumpTag{
						Tag:         "7fe0,0010",
						VR:          "OW",
						Value:       `0000\0000\0000\0000\0000\0000\0000\0000\0000\0000\0000\0000\0000...`,
						Description: "PixelData",
					},
				},
			},
		},
	}, result)
	assert.NoError(t, err)
}
