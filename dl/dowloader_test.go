package dl

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	link := "https://hls.videocc.net/c3a4ce8fd2/2/c3a4ce8fd22811d3e0ec4533a3774062.m3u8?device=desktop&pid=1599527119892X1092333"
	output := "/Users/evrins/Downloads/out.mp4"
	downloader, err := NewTask(output, link)
	if err != nil {
		panic(err)
	}
	err = downloader.Start(25)
	if err != nil {
		panic(err)
	}
}
