package ytdlp

type Options struct {
	Format          Format
	AudioFormat     AudioFormat
	ExtractAudio    bool
	AudioQuality    string
	Output          string
	DownloadArchive string
	CookiesFile     string
	NoPlaylist      bool
	PlaylistItems   string
	WriteThumbnail  bool
	EmbedThumbnail  bool
	AddMetadata     bool
	OutputFormat    OutputFormat
	Proxy           string
	UserAgent       string
	Referer         string
	Quiet           bool
	NoWarnings      bool
	Verbose         bool
}

func (o *Options) ToArgs() []string {
	args := []string{}

	if o.Format != "" {
		args = append(args, "-f", string(o.Format))
	}
	if o.AudioFormat != "" {
		args = append(args, "--audio-format", string(o.AudioFormat))
	}
	if o.ExtractAudio {
		args = append(args, "--extract-audio")
	}
	if o.AudioQuality != "" {
		args = append(args, "--audio-quality", o.AudioQuality)
	}
	if o.Output != "" {
		args = append(args, "-o", o.Output)
	}
	if o.DownloadArchive != "" {
		args = append(args, "--download-archive", o.DownloadArchive)
	}
	if o.CookiesFile != "" {
		args = append(args, "--cookies", o.CookiesFile)
	}
	if o.NoPlaylist {
		args = append(args, "--no-playlist")
	}
	if o.PlaylistItems != "" {
		args = append(args, "--playlist-items", o.PlaylistItems)
	}
	if o.WriteThumbnail {
		args = append(args, "--write-thumbnail")
	}
	if o.EmbedThumbnail {
		args = append(args, "--embed-thumbnail")
	}
	if o.AddMetadata {
		args = append(args, "--add-metadata")
	}
	if o.OutputFormat != "" {
		args = append(args, "--merge-output-format", string(o.OutputFormat))
	}
	if o.Proxy != "" {
		args = append(args, "--proxy", o.Proxy)
	}
	if o.UserAgent != "" {
		args = append(args, "--user-agent", o.UserAgent)
	}
	if o.Referer != "" {
		args = append(args, "--referer", o.Referer)
	}
	if o.Quiet {
		args = append(args, "-q")
	}
	if o.NoWarnings {
		args = append(args, "--no-warnings")
	}
	if o.Verbose {
		args = append(args, "-v")
	}

	return args
}
