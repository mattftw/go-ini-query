package main

import (
	"bytes"
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

func TestGetSectionProperty(t *testing.T) {

	// the ini we start with
	fromIni, _ := ini.Load([]byte(""))
	fromIni.Section("").Key("foo").SetValue("bar")
	fromIni.Section("baz").Key("foo").SetValue("foobar")

	stdout, err := testAppRawOutput(fromIni, []string{os.Args[0], "--file=-", "get", "--section=baz", "--property=foo"})
	if err != nil {
		t.Fatal(err)
	}

	expectedResult := []byte("foobar\n")
	if !bytes.Equal(stdout, expectedResult) {
		t.Fatalf("Expected %s got %s", expectedResult, stdout)
	}
}

func TestGetRootProperty(t *testing.T) {

	// the ini we start with
	fromIni, _ := ini.Load([]byte(""))
	fromIni.Section("").Key("foo").SetValue("bar")
	fromIni.Section("baz").Key("foo").SetValue("foobar")

	stdout, err := testAppRawOutput(fromIni, []string{os.Args[0], "--file=-", "get", "--property=foo"})
	if err != nil {
		t.Fatal(err)
	}

	expectedResult := []byte("bar\n")
	if !bytes.Equal(stdout, expectedResult) {
		t.Fatalf("Expected %s got %s", expectedResult, stdout)
	}
}

func TestSetRootProperty(t *testing.T) {

	inIni, _ := ini.Load([]byte(""))
	inIni.Section("").Key("foo").SetValue("bar")
	inIni.Section("baz").Key("foo").SetValue("foobar")

	resultIni, err := testAppIniOutput(inIni, []string{os.Args[0], "--file=-", "set", "--property=foo", "--value=foobaz"})
	if err != nil {
		t.Fatal(err)
	}

	expectedIni, _ := ini.Load([]byte(""))
	expectedIni.Section("").Key("foo").SetValue("foobaz")
	expectedIni.Section("baz").Key("foo").SetValue("foobar")

	expectIniMatch(t, expectedIni, resultIni)
}

func TestSetSectionProperty(t *testing.T) {

	inIni, _ := ini.Load([]byte(""))
	inIni.Section("").Key("foo").SetValue("bar")
	inIni.Section("baz").Key("foo").SetValue("foobar")

	resultIni, err := testAppIniOutput(inIni, []string{os.Args[0], "--file=-", "set", "--section=baz", "--property=foo", "--value=glengarry"})
	if err != nil {
		t.Fatal(err)
	}

	expectedIni, _ := ini.Load([]byte(""))
	expectedIni.Section("").Key("foo").SetValue("bar")
	expectedIni.Section("baz").Key("foo").SetValue("glengarry")

	expectIniMatch(t, expectedIni, resultIni)
}

func TestDeleteSectionProperty(t *testing.T) {

	inIni, _ := ini.Load([]byte(""))
	inIni.Section("").Key("foo").SetValue("bar")
	inIni.Section("baz").Key("foo").SetValue("foobar")

	resultIni, err := testAppIniOutput(inIni, []string{os.Args[0], "--file=-", "delete", "--section=baz", "--property=foo"})
	if err != nil {
		t.Fatal(err)
	}

	expectedIni, _ := ini.Load([]byte(""))
	expectedIni.Section("").Key("foo").SetValue("bar")

	expectIniMatch(t, expectedIni, resultIni)
}

func TestDeleteRootProperty(t *testing.T) {

	inIni, _ := ini.Load([]byte(""))
	inIni.Section("").Key("foo").SetValue("bar")
	inIni.Section("baz").Key("foo").SetValue("foobar")

	resultIni, err := testAppIniOutput(inIni, []string{os.Args[0], "--file=-", "delete", "--property=foo"})
	if err != nil {
		t.Fatal(err)
	}

	expectedIni, _ := ini.Load([]byte(""))
	expectedIni.Section("baz").Key("foo").SetValue("foobar")

	expectIniMatch(t, expectedIni, resultIni)
}
