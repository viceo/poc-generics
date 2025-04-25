package test1

import "errors"

type Runnable interface {
	Run(cdb []byte)
}

type CmdElementStatus struct {
	elementProperty string
}

func (cmd CmdElementStatus) Run(cdb []byte) {
	// Run Actual SCSI Command
}

type CmdInquiry struct {
	InquiryProperty string
}

func (cmd CmdInquiry) Run(cdb []byte) {
	// Run Actual SCSI Command
}

func (x CmdElementStatus) SpecificCmdElementStatusFunction() string {
	return x.elementProperty
}

func RunCmd[CMD Runnable]() CMD {

	// Create the appropriate type based on generic parameter
	var cmd CMD
	var cdb []byte
	switch any(cmd).(type) {
	case CmdInquiry:
		// Initialize with SgIo and the specific property
		cdb = []byte{0x12}
		concreteStruct := any(CmdInquiry{
			InquiryProperty: "default inquiry value", // Set appropriate default or parameter
		}).(CMD)
		concreteStruct.Run(cdb)
		return concreteStruct
	case CmdElementStatus:
		// Initialize with SgIo and the specific property
		cdb = []byte{0x8B}
		concreteStruct := any(CmdElementStatus{
			elementProperty: "default element value", // Set appropriate default or parameter
		}).(CMD)
		concreteStruct.Run(cdb)
		return concreteStruct
	default:
		panic(errors.New("unknown Command"))
	}
}
