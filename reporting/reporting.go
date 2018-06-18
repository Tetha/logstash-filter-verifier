package reporting;

type Reporter interface {
	BeforeTest(testDescription string);

	TestError(e error)
	TestSuccess()

	StartCompare(testNum int, testCount int, testSet string, testDescription string);
}