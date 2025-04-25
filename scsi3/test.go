package scsi3

import (
	"errors"
	"fmt"
)

// The base Command type and the interface required
// for each command
type Runnable interface{ Run() }
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
	cmd.senseBuffer = []byte{0x70, 0x00, 0x01, 0xF0}
	return SenseData{
		SenseLength: cmd.senseBuffer[3],
		SenseKey:    fmt.Sprintf("%02x", cmd.senseBuffer[0]),
		Asc:         fmt.Sprintf("%02x", cmd.senseBuffer[1]),
		Ascq:        fmt.Sprintf("%02x", cmd.senseBuffer[2]),
	}
}

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

// This will fail if defined in RunCmd call
type AnyOtherStructNotACmd struct{}

// It may exist other concrete types that implements Run() interface
// but this will also fail because Type Assertion
type AnyOtherStructNotACmdWithRunnableInterface struct{}

func (AnyOtherStructNotACmdWithRunnableInterface) Run() {
	/* Some other Run code.. but not a SCSI CMD */
}

func RunCmd[CMD Runnable]() CMD {

	// Create the appropriate type based on generic parameter
	var cmd CMD
	switch any(cmd).(type) {
	case CmdInquiry:

		// Use type assertion to figure out what to create
		concreteStruct := any(NewCmdInquiry()).(CMD)
		cmd = concreteStruct
	case CmdElementStatus:

		// Use type assertion to figure out what to create
		concreteStruct := any(NewCmdElementStatus()).(CMD)
		cmd = concreteStruct
	default:
		panic(ErrUnkownCommand)
	}
	cmd.Run()
	return cmd
}
