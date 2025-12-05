package command

import (
	"byto/internal/builder"
	"os/exec"
)

type UpdateCommand struct {
	Builder *builder.YTDLPBuilder
}

func (c *UpdateCommand) Execute(args any) error {
	ucmd := c.Builder.Build()
	cmd := exec.Command("yt-dlp", ucmd...)

	return cmd.Run()
}
