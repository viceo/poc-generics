package main

import (
	"fmt"
	"poc-generics/scsi3"
)

func main() {

	inquiry := scsi3.RunCmd[scsi3.CmdInquiry]()
	fmt.Println(inquiry.InquiryProperty)

	element := scsi3.RunCmd[scsi3.CmdElementStatus]()
	// element.elementProperty  /* This is not possible.. it's a private property */
	msg := element.SpecificCmdElementStatusFunction()
	fmt.Println(msg)

	senseData := element.GetSenseData()
	fmt.Printf("LengthL %d Key %s, ASC: %s, ASCQ: %s\n",
		senseData.SenseLength,
		senseData.SenseKey,
		senseData.Asc,
		senseData.Ascq,
	)
}
