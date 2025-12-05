package domain

type Setting struct {
	Quality           VideoQuality `json:"quality"`
	ParallelDownloads int          `json:"parallel_downloads"`
	DownloadPath      string       `json:"download_path"`
}
