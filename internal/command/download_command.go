package command

import (
	"byto/internal/builder"
	"byto/internal/domain"
	"os/exec"
)

type DownloadCommand struct {
	Builder *builder.YTDLPBuilder
}

func (c *DownloadCommand) Execute(args any) error {
	ucmd := c.Builder.Build()
	cmd := exec.Command("yt-dlp", ucmd...)
	media, ok := args.(*domain.Media)
	if ok {
		//parsing logic
	}
	return cmd.Run()
}
