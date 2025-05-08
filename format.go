package ytdlp

type Format string

const (
	FormatBestVideo          Format = "bestvideo"
	FormatBestAudio          Format = "bestaudio"
	FormatBest               Format = "best"
	FormatWorst              Format = "worst"
	FormatAudioOnly          Format = "audioonly"
	FormatVideoOnly          Format = "videoonly"
	FormatAll                Format = "all"
	FormatAllAudio           Format = "all-audio"
	FormatAllVideo           Format = "all-video"
	FormatAllVideoOnly       Format = "all-video-only"
	FormatAllAudioBest       Format = "all-audio-best"
	FormatAllVideoWorst      Format = "all-video-worst"
	FormatAllAudioBestVideo  Format = "all-audio-bestvideo"
	FormatBestVideoBestAudio Format = "bestvideo+bestaudio"
)

type OutputFormat string

const (
	OutputFormatMP4  OutputFormat = "mp4"
	OutputFormatMKV  OutputFormat = "mkv"
	OutputFormatWEBM OutputFormat = "webm"
	OutputFormatAVI  OutputFormat = "avi"
	OutputFormatMOV  OutputFormat = "mov"
	OutputFormatFLV  OutputFormat = "flv"
	OutputFormatMPEG OutputFormat = "mpeg"
	OutputFormat3GP  OutputFormat = "3gp"
	OutputFormatOGG  OutputFormat = "ogg"
	OutputFormatWMV  OutputFormat = "wmv"
	OutputFormatTS   OutputFormat = "ts"
	OutputFormatM4V  OutputFormat = "m4v"
	OutputFormatF4V  OutputFormat = "f4v"
	OutputFormatMPG  OutputFormat = "mpg"
	OutputFormatSWF  OutputFormat = "swf"
)

type AudioFormat string

const (
	AudioFormatMP3     AudioFormat = "mp3"
	AudioFormatM4A     AudioFormat = "m4a"
	AudioFormatAAC     AudioFormat = "aac"
	AudioFormatWAV     AudioFormat = "wav"
	AudioFormatFLAC    AudioFormat = "flac"
	AudioFormatOGG     AudioFormat = "ogg"
	AudioFormatOPUS    AudioFormat = "opus"
	AudioFormatWAVPACK AudioFormat = "wavpack"
	AudioFormatAIFF    AudioFormat = "aiff"
	AudioFormatAU      AudioFormat = "au"
	AudioFormatAC3     AudioFormat = "ac3"
	AudioFormatDTS     AudioFormat = "dts"
	AudioFormatAMR     AudioFormat = "amr"
	AudioFormatPCM     AudioFormat = "pcm"
	AudioFormatALAC    AudioFormat = "alac"
	AudioFormatDSD     AudioFormat = "dsd"
	AudioFormatDSF     AudioFormat = "dsf"
	AudioFormatDFF     AudioFormat = "dff"
	AudioFormatMKA     AudioFormat = "mka"
	AudioFormatM4B     AudioFormat = "m4b"
	AudioFormatM4P     AudioFormat = "m4p"
	AudioFormatM4R     AudioFormat = "m4r"
	AudioFormatM4V     AudioFormat = "m4v"
)
