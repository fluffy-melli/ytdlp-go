package ffmpeg

import (
	"bytes"
	"os/exec"
)

func ToOpus(opusData []byte, callback func([]byte)) error {
	cmd := exec.Command("ffmpeg",
		"-i", "pipe:0",
		"-c:a", "libopus",
		"-f", "opus",
		"pipe:1",
	)

	cmd.Stdin = bytes.NewReader(opusData)

	var buffer bytes.Buffer
	cmd.Stdout = &buffer

	var errorBuffer bytes.Buffer
	cmd.Stderr = &errorBuffer

	err := cmd.Run()
	if err != nil {
		return err
	}

	callback(buffer.Bytes())

	return nil
}
