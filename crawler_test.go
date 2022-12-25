package codeforcescrawler

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	crawler := NewContest(1763)
	crawler.GetTestCases("D")
	time.Sleep(time.Second * 2)
	crawler.GetTestCases("A")
	time.Sleep(time.Second * 2)
	crawler.GetTestCases("F")
}
