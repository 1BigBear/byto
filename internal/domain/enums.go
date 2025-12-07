package domain

type VideoQuality int

const (
	Quality360p VideoQuality = iota
	Quality480p
	Quality720p
	Quality1080p
	Quality1440p
	Quality2160p
)

type DownloadStatus int

const (
	Pending DownloadStatus = iota
	InProgress
	Completed
	Failed
	Paused
)
