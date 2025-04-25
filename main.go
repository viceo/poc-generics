package main

import (
	"fmt"
	"poc-generics/test1"
)

func main() {

	inquiry := test1.RunCmd[test1.CmdInquiry]()
	fmt.Println(inquiry.InquiryProperty)

	element := test1.RunCmd[test1.CmdElementStatus]()
	msg := element.SpecificCmdElementStatusFunction()
	fmt.Println(msg)
}
