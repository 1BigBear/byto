package domain

type Setting struct {
	Quality           VideoQuality `json:"quality"`
	ParallelDownloads int          `json:"parallel_downloads"`
	DownloadPath      string       `json:"download_path"`
}

func NewSetting() *Setting {
	return &Setting{
		Quality:           Quality1080p,
		ParallelDownloads: 3,
		DownloadPath:      "./downloads",
	}
}

func (s *Setting) Update(quality VideoQuality, parallelDownloads int, downloadPath string) {
	s.Quality = quality
	s.ParallelDownloads = parallelDownloads
	s.DownloadPath = downloadPath
}
