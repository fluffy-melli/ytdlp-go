package ytdlp

import "encoding/binary"

func ByteToPCM(bytes []byte) []int16 {
	pcm := make([]int16, len(bytes)/2)
	for i := 0; i < len(pcm); i++ {
		pcm[i] = int16(binary.LittleEndian.Uint16(bytes[i*2:]))
	}
	return pcm
}
