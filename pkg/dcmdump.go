package godcmtk

// #cgo LDFLAGS: -lgodcmdata -L../dcmtk/dcmdata/dist
// extern int dcmdump(int argc, char *argv[], char **output);
import "C"
import "fmt"
import "strings"
import "regexp"

// ArgDCMDUMP defines the options for DCMDUMP.
type ArgDCMDUMP struct {
	// input options:
	// input file format:
	// read file format or data set (default)
	ReadFile bool `arg:"+f"`
	// read file format only
	ReadFileOnly bool `arg:"+fo"`
	// read data set without file meta information
	ReadDataset bool `arg:"-f"`

	// input transfer syntax:
	// use TS recognition (default)
	RedXferAuto bool `arg:"-t="`
	// ignore TS specified in the file meta header
	ReadXferDetect bool `arg:"-td"`
	// read with explicit VR little endian TS
	ReadXferLittle bool `arg:"-te"`
	// read with explicit VR big endian TS
	ReadXferBig bool `arg:"-tb"`
	// read with implicit VR little endian TS
	ReadXferImplicit bool `arg:"-ti"`

	// input files:
	// scan directories for input files (dcmfile-in)
	ScanDirectories bool `arg:"+sd"`
	// [p]attern: string (only with --scan-directories)
	// pattern for filename matching (wildcards)
	ScanPattern string `arg:"+sp"`
	// do not recurse within directories (default)
	NoRecurse bool `arg:"-r"`
	// recurse within specified directories
	Recurse bool `arg:"+r"`

	// long tag values:
	// load very long tag values (default)
	LoadAll bool `arg:"+M"`
	// do not load very long values (e.g. pixel data)
	LoadShort bool `arg:"-M"`
	// [k]bytes: integer (4..4194302, default: 4)
	// set threshold for long values to k kbytes
	MaxReadLength int `arg:"+R"`

	// parsing of file meta information:
	// use file meta information group length (default)
	UseMetaLength bool `arg:"+ml"`
	// ignore file meta information group length
	IgnoreMetaLength bool `arg:"-ml"`

	// parsing of odd-length attributes:
	// accept odd length attributes (default)
	AcceptOddLength bool `arg:"+ao"`
	// assume real length is one byte larger
	AssumeEvenLength bool `arg:"+ae"`

	// handling of explicit VR:
	// use explicit VR from dataset (default)
	UseExplicitVR bool `arg:"+ev"`
	// ignore explicit VR (prefer data dictionary)
	IgnoreExplicitVR bool `arg:"-ev"`

	// handling of non-standard VR:
	// treat non-standard VR as unknown (default)
	TreatAsUnknown bool `arg:"+vr"`
	// try to read with implicit VR little endian TS
	AssumeImplicit bool `arg:"-vr"`

	// handling of undefined length UN elements:
	// read undefined len UN as implicit VR (default)
	EnableCP246 bool `arg:"+ui"`
	// read undefined len UN as explicit VR
	DisableCP246 bool `arg:"-ui"`

	// handling of defined length UN elements:
	// retain elements as UN (default)
	RetainUN bool `arg:"-uc"`
	// convert to real VR if known
	ConvertUN bool `arg:"+uc"`

	// handling of private max-length elements (implicit VR):
	// read as defined in dictionary (default)
	MaxLengthDict bool `arg:"-sq"`
	// read as sequence with undefined length
	MaxLengthSeq bool `arg:"+sq"`

	// handling of wrong delimitation items:
	// use delimitation items from dataset (default)
	UseDelimItems bool `arg:"-rd"`
	// replace wrong sequence/item delimitation items
	ReplaceWrongDelim bool `arg:"+rd"`

	// handling of illegal undefined length OB/OW elements:
	// reject dataset with illegal element (default)
	IllegalObowRej bool `arg:"-oi"`
	// convert undefined length OB/OW element to SQ
	IllegalObowConv bool `arg:"+oi"`

	// handling of VOI LUT Sequence with OW VR and explicit length:
	// reject dataset with illegal VOI LUT (default)
	IllegalVoiRej bool `arg:"-vi"`
	// convert illegal VOI LUT to SQ
	IllegalVoiConv bool `arg:"+vi"`

	// handling of explicit length pixel data for encaps. transfer syntaxes:
	// abort on explicit length pixel data (default)
	AbortExplPixdata bool `arg:"-pe"`
	// use explicit length pixel data
	UseExplPixdata bool `arg:"+pe"`

	// general handling of parser errors:
	// try to recover from parse errors
	IgnoreParseErrors bool `arg:"+Ep"`
	// handle parse errors and stop parsing (default)
	HandleParseErrors bool `arg:"-Ep"`

	// other parsing options:
	// [t]ag: "gggg,eeee" or dictionary name
	// stop parsing after element specified by t
	StopAfterElem bool `arg:"+st"`
	// [t]ag: "gggg,eeee" or dictionary name
	// stop parsing before element specified by t
	StopBeforeElem bool `arg:"+sb"`

	// automatic data correction:
	// enable automatic data correction (default)
	EnableCorrection bool `arg:"+dc"`
	// disable automatic data correction
	DisableCorrection bool `arg:"-dc"`

	// bitstream format of deflated input:
	// expect deflated bitstream (default)
	BitstreamDeflated bool `arg:"+bd"`
	// expect deflated zlib bitstream
	BitstreamZlib bool `arg:"+bz"`

	// processing options:
	// specific character set:
	// convert all element values that are affected
	// by Specific Character Set (0008,0005) to UTF-8
	ConvertToUTF8 bool `arg:"+U8"`
}

// DumpResult defines the output for DCMDUMP.
type DumpResult struct {
	Files []*DumpFile
}

// DumpFile defines the output for each dcm file for DCMDUMP.
type DumpFile struct {
	Name string
	Tags map[string]*DumpTag
}

// DumpTag defines the output for each tag in dcm file for DCMDUMP.
type DumpTag struct {
	Tag         string
	VR          string
	Description string
	Value       string
}

// DCMDUMP dumps DICOM file and data set
func DCMDUMP(dcmfileIn string, args *ArgDCMDUMP) error {
	argStrs, err := MarshalArgs(args)
	if err != nil {
		return fmt.Errorf("failed to parse args %+v for dcmdump: %v", args, err)
	}

	argStrs = append([]string{"dcmdump"}, argStrs...)
	argStrs = append(argStrs, dcmfileIn)
	cStrings, free := StringArray(argStrs...)
	defer free()

	out, free := EmptyString()
	defer free()

	C.dcmdump(C.int(len(argStrs)), cStrings, &out)
	fmt.Println("orz", C.GoString(out))

	return nil
}

func parseOutput(output string) (*DumpResult, error) {
	p := `^# dcmdump \(\d+\/\d+\): (.+)`
	pattern := `\((\w{4},\w{4})\)\s(\w{2})\s(.*)\s+#\s+\d+,\s\d+\s(\w+)`
	r := regexp.MustCompile(p)
	regex := regexp.MustCompile(pattern)
	results := &DumpResult{}
	var f *DumpFile
	for _, line := range strings.Split(output, "\n") {
		ms := r.FindStringSubmatch(line)
		if len(ms) == 2 {
			// file name detected.
			f = &DumpFile{
				Name: ms[1],
				Tags: map[string]*DumpTag{},
			}
			results.Files = append(results.Files, f)
			continue
		}

		matches := regex.FindStringSubmatch(line)
		if len(matches) == 5 && f != nil {
			// tag detected.
			tag := &DumpTag{
				Tag:         matches[1],
				VR:          matches[2],
				Value:       strings.TrimSpace(matches[3]),
				Description: strings.TrimSpace(matches[4]),
			}

			f.Tags[tag.Tag] = tag
		}
	}

	return results, nil
}
