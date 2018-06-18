package reporting;

// The Quiet Reporter discards every event.
// 
// Eventually this could be used as a null object so we don't
// have to worry about nil's all over the damn place, especially
// if the user requests a quiet mode with a minimum output.
//
// Currently I've added this because I don't want to worry about
// concurrency in the parallel tests.
type QuietReporter struct {};

func (qr QuietReporter) BeforeTest(testDescription string) {}
func (qr QuietReporter) TestError(err error) {}
func (qr QuietReporter) TestSuccess() {}
func (qr QuietReporter)	StartCompare(testNum int, testCount int, testSet string, testDescription string) {}