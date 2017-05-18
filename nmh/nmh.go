package nmh

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"os"
	"strings"
)

// Message type model for in/out communication
type Message struct {
	Text string
}

// Receive msg from Stdin.
func Receive() {
	//for {
	reader := bufio.NewReader(os.Stdin)
	length := make([]byte, 4)
	reader.Read(length)
	lengthNum := readMessageLength(length)
	content := make([]byte, lengthNum)
	reader.Read(content)
	echoMessage(content)
	//}
}

// Send message on Stdout
func Send(msg string) {
	byteMsg := encodeMessage(msg)
	var msgBuf bytes.Buffer
	writeMessageLength(byteMsg)
	msgBuf.Write(byteMsg)
	msgBuf.WriteTo(os.Stdout)
}

func echoMessage(msg []byte) {
	content := strings.ToLower(decodeMessage(msg))
	if content == "hi" {
		Send("Hello!")
	} else if content == "--apiurl" {
		v := GetEnvValue("ONEMARK_API_URL")
		Send(v)
	} else {
		Send(content)
	}
}

func encodeMessage(msg string) []byte {
	return dataToBytes(Message{Text: msg})
}

func decodeMessage(msg []byte) string {
	var msgJSON Message
	json.Unmarshal(msg, &msgJSON)
	return msgJSON.Text
}

func dataToBytes(msg Message) []byte {
	byteMsg, _ := json.Marshal(msg)
	return byteMsg
}

func writeMessageLength(msg []byte) {
	binary.Write(os.Stdout, binary.LittleEndian, uint32(len(msg)))
}

func readMessageLength(msg []byte) int {
	var length uint32
	buf := bytes.NewBuffer(msg)
	binary.Read(buf, binary.LittleEndian, &length)
	return int(length)
}
