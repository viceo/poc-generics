package scsi3

import (
	"errors"
	"fmt"
)

// The base Command type and the interface required
// for each command
type Cmd struct {
	cdb         []byte
	senseBuffer []byte
}

type SenseData struct {
	SenseLength uint8  `json:"senseLength"`
	SenseKey    string `json:"senseKey"`
	Asc         string `json:"asc"`
	Ascq        string `json:"ascq"`
}

// Return the Sense Buffer
func (cmd Cmd) GetSenseData() SenseData {
	// Made up values... just as an example
	// taken from senseBuffer

	/* senseLength := cmd.senseBuffer[0] */
	/* senseKey    := fmt.Sprintf("%02x", cmd.senseBuffer[1]) */
	return SenseData{
		SenseLength: 10,
		SenseKey:    "70",
		Asc:         "00",
		Ascq:        "01",
	}
}

type Runnable interface{ Run() }

// Each type has it's own parameters
type CmdElementStatus struct {
	Cmd
	elementProperty string
}
type CmdInquiry struct {
	Cmd
	InquiryProperty string
}

// Structs can't have constructors...
// The golang way... New* methods
func NewCmdElementStatus() CmdElementStatus {
	return CmdElementStatus{
		elementProperty: "This is an element",
		Cmd: Cmd{
			cdb: []byte{0x0B},
		},
	}
}

func NewCmdInquiry() CmdInquiry {
	return CmdInquiry{
		InquiryProperty: "This is an Inquiry",
		Cmd: Cmd{
			cdb: []byte{0x12},
		},
	}
}

// Commands implements Runnable
func (cmd CmdElementStatus) Run() {
	/* Run Actual SCSI Command*/
	fmt.Printf("Running CDB %x\n", cmd.cdb)
}
func (cmd CmdInquiry) Run() {
	/* Run Actual SCSI Command */
	fmt.Printf("Running CDB %x\n", cmd.cdb)
}

// Each type can have it's own functions
func (cmd CmdElementStatus) SpecificCmdElementStatusFunction() string {
	return cmd.elementProperty
}

// Sentinel Error
var ErrUnkownCommand = errors.New("unknown scsi v3 command")

func RunCmd[CMD Runnable]() CMD {

	// Create the appropriate type based on generic parameter
	var cmd CMD
	switch any(cmd).(type) {
	case CmdInquiry:

		// Use type assertion to figure out what to create
		concreteStruct := any(NewCmdInquiry()).(CMD)
		concreteStruct.Run()
		return concreteStruct
	case CmdElementStatus:

		// Use type assertion to figure out what to create
		concreteStruct := any(NewCmdElementStatus()).(CMD)
		concreteStruct.Run()
		return concreteStruct
	default:
		panic(ErrUnkownCommand)
	}
}
