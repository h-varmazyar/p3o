package links

type Key struct {
	Key string `uri:"key" binding:"required"`
}

type Link struct {
	Key string `uri:"key" binding:"required"`
	URL string `uri:"url" binding:"required"`
}
