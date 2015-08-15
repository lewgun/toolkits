package gists

//camlistore localdisk.go
func rateLimiter(){
	rateLimiter := make(chan bool, maxParallelStats)
	for _, ref := range blobs {
		go func(ref *blobref.BlobRef) {
			rateLimiter <- true  //流控, good
			errCh <- statSend(ref, waitSeconds > 0)
			<-rateLimiter
		}
	}
}(ref)
