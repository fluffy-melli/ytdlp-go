package ytdlp

import (
	"fmt"
	"io"
	"os/exec"
)

func Stream(url string, option *Options, do func([]byte)) error {
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

	buf := make([]byte, 32*1024)
	for {
		n, err := stdout.Read(buf)
		if n > 0 {
			do(buf[:n])
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
