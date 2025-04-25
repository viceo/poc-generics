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

// Part of the error handling example
type CmdWithPanicError struct{}

// This will fail if defined in RunCmd call
type AnyOtherStructNotACmd struct{}

// It may exist other concrete types that implements Run() interface
// but this will also fail because Type Assertion
type AnyOtherStructNotACmdWithRunnableInterface struct{}

// Each type has it's own parameters
type CmdElementStatus struct {
	Cmd
	elementProperty string
}

// Each type can have it's own functions
func (cmd CmdElementStatus) SpecificCmdElementStatusFunction() string {
	return cmd.elementProperty
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

// Sentinel Error
var ErrPanicCommand = errors.New("this will panic everytime... as an example")
var ErrUnkownCommand = errors.New("unknown scsi v3 command")

func NewCmdWithPanicError() CmdWithPanicError {
	panic(ErrPanicCommand)
	// return CmdWithPanicError{}
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

func (AnyOtherStructNotACmdWithRunnableInterface) Run() {
	/* Some other Run code.. but not a SCSI CMD */
}
func (CmdWithPanicError) Run() {
	/* Some other Run code.. supposed to panic everytime */
}

func RunCmd[CMD Runnable]() (cmd CMD, cmderr error) {
	// You can recover from any function... here as an example
	// the ErrUnkownCommand will be recovered and sent back as an
	// error
	defer func() {
		if r := recover(); r != nil {
			// Type assertion to convert the recovered value to an error
			err, ok := r.(error)
			if ok && errors.Is(err, ErrUnkownCommand) {
				// Handle the specific error
				cmderr = err
				return
				// Add your error handling logic here
			} else {
				// Bubble up...
				panic(err)
			}
		}
	}()

	// Create the appropriate type based on generic parameter
	switch any(cmd).(type) {
	case CmdInquiry:
		// Use type assertion to figure out what to create
		concreteStruct := any(NewCmdInquiry()).(CMD)
		cmd = concreteStruct
	case CmdElementStatus:
		// Use type assertion to figure out what to create
		concreteStruct := any(NewCmdElementStatus()).(CMD)
		cmd = concreteStruct
	case CmdWithPanicError:
		concreteStruct := any(NewCmdWithPanicError()).(CMD)
		cmd = concreteStruct
	default:
		// This error will bubble up to main to be handled
		// https://go.dev/wiki/PanicAndRecover
		panic(ErrUnkownCommand)
	}

	cmd.Run()
	return cmd, nil
}
