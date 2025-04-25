package main

import (
	"fmt"
	"poc-generics/scsi3"
)

func main() {
	// All Error will be caught here
	defer errorTrap()

	inquiry, err := scsi3.RunCmd[scsi3.CmdInquiry]()
	if err != nil {
		fmt.Printf("returned error: %s", err)
	}
	fmt.Println(inquiry.InquiryProperty)

	element, err := scsi3.RunCmd[scsi3.CmdElementStatus]()
	if err != nil {
		fmt.Printf("returned error: %s\n", err)
	}

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

	_, err = scsi3.RunCmd[scsi3.AnyOtherStructNotACmdWithRunnableInterface]()
	if err != nil {
		fmt.Printf("returned error: %s\n", err)
	}

	_, err = scsi3.RunCmd[scsi3.CmdWithPanicError]()
	if err != nil {
		fmt.Printf("returned error: %s\n", err)
	}
}

func errorTrap() {
	if r := recover(); r != nil {
		// Type assertion to convert the recovered value to an error
		err, ok := r.(error)
		if ok {
			fmt.Printf("recovered in main error trap(%t): %s\n", ok, err)
		} else {
			fmt.Printf("recovered in main error trap(%t): %s\n", ok, err)
		}
	}
}
