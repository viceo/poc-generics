package test1

type SgIo struct {
	cdb []byte
}

type IDescriptable interface {
	GetCdb() []byte
}

func (x SgIo) GetCdb() []byte {
	return x.cdb
}

type CmdElementStatus struct {
	SgIo
	elementProperty string
}

func (x CmdElementStatus) SpecificCmdElementStatusFunction() string {
	return x.elementProperty
}

type CmdInquiry struct {
	SgIo
	InquiryProperty string
}

func RunCmd[CMD IDescriptable]() CMD {

	// Create the appropriate type based on generic parameter
	var cmd CMD
	var cdb []byte
	switch any(cmd).(type) {
	case CmdInquiry:
		// Initialize with SgIo and the specific property
		cdb = []byte{0x12}
		return any(CmdInquiry{
			SgIo:            SgIo{cdb: cdb},
			InquiryProperty: "default inquiry value", // Set appropriate default or parameter
		}).(CMD)
	case CmdElementStatus:
		// Initialize with SgIo and the specific property
		cdb = []byte{0x8B}
		return any(CmdElementStatus{
			SgIo:            SgIo{cdb: cdb},
			elementProperty: "default element value", // Set appropriate default or parameter
		}).(CMD)
	}

	// Create and return a zero value if no match
	var zeroValue CMD
	return zeroValue
}
