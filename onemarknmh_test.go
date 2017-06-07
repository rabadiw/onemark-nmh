package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os/exec"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	fmt.Println("Testing Hello")

	msg := "{\"Text\":\"Hi\"}"
	expected := "{\"Text\":\"Hello!\"}"

	actual := getReply(msg)

	if !bytes.Equal([]byte(actual), []byte(expected)) {
		t.Errorf("Expected for %q is %q but got %q", msg, expected, actual)
		t.Fail()
	}
}

func TestApiUrl(t *testing.T) {
	fmt.Println("Testing TestApiUrl")

	msg := "{\"Text\":\"--APIURL\"}"
	expected := "{\"Text\":\"http://localhost:3010/api/\"}"

	actual := getReply(msg)

	if !bytes.Equal([]byte(actual), []byte(expected)) {
		t.Errorf("Expected for %q is %q (%v) but got %q (%v)", msg, expected, len(expected), actual, len(actual))
		t.Fail()
	}
}

func TestEnvParam(t *testing.T) {
	fmt.Println("Testing TestEnvParam")

	// results can be checked in log file
	_, err := exec.Command("go", "run", "onemarknmh.go", "-env", "../.env").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	_, err = exec.Command("go", "run", "onemarknmh.go", "-env", ".env").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func getReply(msg string) string {
	cmd := exec.Command("go", "run", "onemarknmh.go")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	defer stdin.Close()
	defer stdout.Close()

	// Write the msg length then the msg
	length := make([]byte, 4)
	binary.LittleEndian.PutUint32(length, uint32(len(msg)))
	io.Copy(stdin, bytes.NewBuffer(length))
	fmt.Fprint(stdin, msg)

	time.Sleep(1 * time.Second)

	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	buf.Next(4)
	s := buf.String()

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

	return s
}
