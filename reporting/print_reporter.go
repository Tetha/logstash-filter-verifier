package reporting

import (
	"fmt"
	"strings"
	"os"
)

// The print reporter should overall output the same as the code does at
// the moment.
//
// The prefix mostly exists so I can distinguish the existing prints 
// and the prints from the  print reporter.
type PrintReporter struct {
	Prefix string
};

func (pr PrintReporter) BeforeTest(testDescription string) {
	fmt.Printf("%sRunning tests in %s...\n", pr.Prefix, testDescription)
}

func (pr PrintReporter) TestError(err error) {
	userError("%sTestcase failed, continuing with the rest: %s", pr.Prefix, err)
}

func (pr PrintReporter)	StartCompare(testNum int, testCount int, testSet string, testDescription string) {
	fmt.Printf("%sComparing message %d of %d from %s%s...\n", pr.Prefix, testNum, testCount, testSet, testDescription)
}

func (pr PrintReporter) TestSuccess() {
	fmt.Printf("%sTest Successful\n", pr.Prefix)
}

// userError prints an error message to stderr.
// TODO: do something with this copy-pasted method.
func userError(format string, a ...interface{}) {
	if strings.HasSuffix(format, "\n") {
		fmt.Fprintf(os.Stderr, format, a...)
	} else {
		fmt.Fprintf(os.Stderr, format+"\n", a...)
	}
}