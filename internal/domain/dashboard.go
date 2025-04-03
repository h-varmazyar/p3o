package domain

type DashboardInfoItem struct {
	Count       uint   `json:"count"`
	Growth      string `json:"growth"`
	GrowthTrend string `json:"growth_trend"`
}
