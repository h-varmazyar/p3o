package link

func (c *Controller) Counts(ctx *gin.Context) {
	if totalLinkCount, err := c.linkModel.TotalCounts(ctx); err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
	} else {
		utils.JsonHttpResponse(ctx, totalLinkCount, nil, true)
	}
}