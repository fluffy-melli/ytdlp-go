package ytdlp

import (
	"fmt"
	"io"
	"os/exec"
)

func Stream(url string, option *Options, callback func([]byte), bufferSize int) error {
	var args []string
	if option != nil {
		args = append(option.ToArgs(), "-o", "-", url)
	} else {
		args = []string{"-o", "-", url}
	}

	cmd := exec.Command("yt-dlp", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start yt-dlp: %w", err)
	}

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
			return fmt.Errorf("stream read error: %w", err)
		}
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("yt-dlp error: %w", err)
	}

	return nil
}

func Run(url string, option *Options, callback func([]byte)) error {
	var buffer []byte

	err := Stream(url, option, func(chunk []byte) {
		buffer = append(buffer, chunk...)
	}, 4096)

	if err != nil {
		return err
	}

	callback(buffer)

	return nil
}
