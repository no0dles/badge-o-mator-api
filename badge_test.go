package main

import (
	"testing"
	"bytes"
	"fmt"
	"bufio"
	"io/ioutil"
)

func TestEmptyBadge(t *testing.T) {
	_, err := CreateBanner("", "", "sm", "1")
	if err != nil {
		t.Fail()
	}
}

func TestInvalidColor(t *testing.T) {
	_, err := CreateBanner("", "", "sm", "0")
	if err == nil {
		t.Fail()
	}
}

func TestInvalidSize(t *testing.T) {
	_, err := CreateBanner("", "", "sm", "0")
	if err == nil {
		t.Fail()
	}
}

func TestCreateSvg(t *testing.T) {
	banner, err := CreateBanner("test", "test", "sm", "1")
	if err != nil {
		t.Fail()
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	generateBanner(writer, banner)
}