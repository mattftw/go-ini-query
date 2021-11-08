package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

func expectIniMatch(t *testing.T, expectedIni *ini.File, resultIni *ini.File) {

	expectedBytes, err := convertIniToBytes(expectedIni)
	if err != nil {
		t.Fatal(err)
	}

	resultBytes, err := convertIniToBytes(resultIni)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expectedBytes, resultBytes) {
		t.Fatalf("expected\n---\n%s\n---\n got \n---\n%s\n---\n", expectedBytes, resultBytes)
	}
}

func convertIniToBytes(ini *ini.File) ([]byte, error) {
	var out bytes.Buffer
	writer := bufio.NewWriter(&out)
	if _, err := ini.WriteTo(writer); err != nil {
		return nil, err
	}

	if err := writer.Flush(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func testAppRawOutput(inCfg *ini.File, args []string) (outBytes []byte, err error) {

	// construct stdin
	inBytes, err := convertIniToBytes(inCfg)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert ini to bytes")
	}
	app.Reader = bytes.NewBuffer(inBytes)

	// fudge stdout
	var stdoutBuffer bytes.Buffer
	appWriter := bufio.NewWriter(&stdoutBuffer)
	app.Writer = appWriter

	// run the app
	err = app.Run(args)
	if err != nil {
		return nil, err
	}

	// flush stdout writer to ensure everything is in the buffer
	if err := appWriter.Flush(); err != nil {
		return nil, errors.Wrapf(err, "unable to flush stdout writer")
	}

	// return stdout to the caller
	return stdoutBuffer.Bytes(), nil
}

func testAppIniOutput(inCfg *ini.File, args []string) (outCfg *ini.File, err error) {

	outBytes, err := testAppRawOutput(inCfg, args)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get app output")
	}

	outCfg, err = ini.Load(outBytes)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to convert output into ini")
	}

	return outCfg, nil
}
