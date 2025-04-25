package main

import (
	"errors"
	"fmt"
	"poc-generics/scsi3"
)

func main() {
	// All Error will be caught here
	defer errorTrap()

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

	// Impossible to define as this type doesn't implement Runnable
	// failed := scsi3.RunCmd[scsi3.AnyOtherStructNotACmd]()

	_ = scsi3.RunCmd[scsi3.AnyOtherStructNotACmdWithRunnableInterface]()

}

func errorTrap() {
	if r := recover(); r != nil {
		// Type assertion to convert the recovered value to an error
		err, ok := r.(error)
		if ok && errors.Is(err, scsi3.ErrUnkownCommand) {
			// Handle the specific error
			fmt.Println("Caught unknown command error")
			// Add your error handling logic here
		} else {
			// Either not an error type or not the specific error we're looking for
			fmt.Printf("Recovered from: %v\n", r)
		}
	}
}
