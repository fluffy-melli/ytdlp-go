package ffmpeg

import (
	"bytes"
	"io"
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

func ToOpusStream(opusData []byte, callback func([]byte), bufferSize int) error {
	cmd := exec.Command("ffmpeg",
		"-i", "pipe:0",
		"-c:a", "libopus",
		"-f", "opus",
		"pipe:1",
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		_, _ = stdin.Write(opusData)
	}()

	buf := make([]byte, bufferSize)
	for {
		n, err := stdout.Read(buf)
		if n > 0 {
			callback(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
