package huoshan

func WatermarkRemover(url string)(string, error)  {

	videoId, err := getVideoId(url)
	
	if len(videoId) == 0 {
		return "", nil
	}
	
	videoLink, err := getVideoLink(videoId)

	downloadLink, err := getVideoDownloadLink(videoLink)

	return downloadLink, err
}

