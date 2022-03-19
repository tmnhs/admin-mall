package mall

type RouterGroup struct {
	UserRouter
	CategoryRouter
	IdleItemRouter
	PurchaseItemRouter
	OrderRouter
	AdRouter
	FeedbackRouter
	ReportRouter
	OrderHistoryRouter
}
