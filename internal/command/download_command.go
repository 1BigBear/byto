package command

import (
	"bufio"
	"byto/internal/builder"
	"byto/internal/domain"
	"byto/internal/parser"
	"fmt"
	"os/exec"
	"strconv"
)

type DownloadCommand struct {
	Builder *builder.YTDLPBuilder
}

func (c *DownloadCommand) Execute(args any) error {
	media, ok := args.(*domain.Media)
	if !ok {
		return fmt.Errorf("invalid arguments, expected *domain.Media")
	}

	c.Builder.ProgressTemplate("[byto:title] %(info.title)s [byto:downloaded_bytes] %(progress.downloaded_bytes)d [byto:total_bytes] %(progress.total_bytes)d")

	ucmd := c.Builder.Build()
	cmd := exec.Command("yt-dlp", ucmd...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	parser := parser.YTDLPDownloadParser{}
	scanner := bufio.NewScanner(stdout)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			media.AppendLog(line)

			parsedData, err := parser.Parse(line)
			if err == nil {
				downloaded, _ := strconv.ParseInt(parsedData["downloaded_bytes"], 10, 64)
				total, _ := strconv.ParseInt(parsedData["total_bytes"], 10, 64)

				percentage := 0
				if total > 0 {
					percentage = int((float64(downloaded) / float64(total)) * 100)
				}

				media.UpdateProgress(downloaded, total, percentage)
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		media.SetStatus(domain.Failed)
		return err
	}

	media.SetStatus(domain.Completed)
	return nil
}
