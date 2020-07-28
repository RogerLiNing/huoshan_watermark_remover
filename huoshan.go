package huoshan

func WatermarkRemover(url string)(string, error)  {

	videoId, err := getVideoId(url)

	videoLink, err := getVideoLink(videoId)

	downloadLink, err := getVideoDownloadLink(videoLink)

	return downloadLink, err
}

