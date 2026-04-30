package common

type UserRole string

const (
	RoleBuyer  UserRole = "buyer"
	RoleSeller UserRole = "seller"
	RoleAdmin  UserRole = "admin"
)

type Provider string

const (
	ProviderLocal  Provider = "local"
	ProviderGoogle Provider = "google"
	ProviderGithub Provider = "github"
)

type DiscountType string

const (
	DiscountFixed      DiscountType = "fixed"
	DiscountPercentage DiscountType = "percentage"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentPaid      PaymentStatus = "paid"
	PaymentExpired   PaymentStatus = "expired"
	PaymentCancelled PaymentStatus = "cancelled"
)

type OrderStatus string

const (
	OrderWaitingPayment OrderStatus = "waiting_payment"
	OrderProcessed      OrderStatus = "processed"
	OrderShipped        OrderStatus = "shipped"
	OrderDelivered      OrderStatus = "delivered"
	OrderCancelled      OrderStatus = "cancelled"
	OrderCompleted      OrderStatus = "completed"
)
