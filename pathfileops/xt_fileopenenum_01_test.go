package pathfileops

import (
	"os"
	"testing"
)

func TestFileOpenStatus_New_01(t *testing.T) {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

	fOpStatus, err := FileOpenStatus{}.New(FOpenType.WriteOnly(),
		FOpenMode.Append(), FOpenMode.Truncate())

	if err != nil {
		t.Errorf("Error returned by FileOpenStatus{}.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by FileOpenStatus{}.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

}
