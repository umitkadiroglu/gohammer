package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/tubuarge/GoHammer/util"
)

type TestResults struct {
	TestStartTimestamp   time.Time
	TestEndTimestamp     time.Time
	OverallExecutionTime time.Duration

	TotalTxCount int
}

type LogClient struct {
	LogFile    *os.File
	TestResult *TestResults
}

func NewLogClient(filename string) (*LogClient, error) {
	file, err := CreateLogFile(filename)
	if err != nil {
		return nil, err
	}
	return &LogClient{
		LogFile:    file,
		TestResult: &TestResults{},
	}, nil
}

func CreateLogFile(filename string) (*os.File, error) {
	ts := util.GetFormattedTimestampNow()
	fullFilename := ts + " " + filename
	file, err := os.Create(fullFilename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (l *LogClient) WriteFile(data []byte) error {
	_, err := l.LogFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (l *LogClient) WriteTestResults() error {
	strData := fmt.Sprintf(`
	Test Started At: %v\n
	Test Ended At: %v\n
	Total Test Execution Time: %v\n
	Total Transaction Count: %d\n`,
		l.TestResult.TestStartTimestamp,
		l.TestResult.TestEndTimestamp,
		l.TestResult.OverallExecutionTime,
		l.TestResult.TotalTxCount)

	err := l.WriteFile([]byte(strData))
	if err != nil {
		return err
	}
	return nil
}

func (l *LogClient) CloseFile() error {
	err := l.LogFile.Close()
	if err != nil {
		return err
	}
	return nil
}
