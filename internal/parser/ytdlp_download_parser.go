package parser

import (
	"errors"
	"regexp"
)

type YTDLPDownloadParser struct{}

// [byto:title] %(info.title)s [byto:downloaded_bytes] %(progress.downloaded_bytes)d [byto:total_bytes] %(progress.total_bytes)d
var logRegex = regexp.MustCompile(`\[byto:title\]\s+(.*?)\s+\[byto:downloaded_bytes\]\s+(\d+)\s+\[byto:total_bytes\]\s+(\d+)`)

func (p YTDLPDownloadParser) Parse(input string) (map[string]string, error) {
	matches := logRegex.FindStringSubmatch(input)
	if len(matches) < 4 {
		return nil, errors.New("failed to parse log line: format mismatch")
	}

	result := make(map[string]string)
	result["title"] = matches[1]
	result["downloaded_bytes"] = matches[2]
	result["total_bytes"] = matches[3]

	return result, nil
}
