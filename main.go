package main

import (
	"fmt"
	"poc-generics/scsi3"
)

func main() {

	inquiry := scsi3.RunCmd[scsi3.CmdInquiry]()
	fmt.Println(inquiry.InquiryProperty)

	element := scsi3.RunCmd[scsi3.CmdElementStatus]()
	msg := element.SpecificCmdElementStatusFunction()
	fmt.Println(msg)
}
