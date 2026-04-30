package delivery

type AddressRequest struct {
	ReceiverName string `json:"receiver_name" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Address      string `json:"address" validate:"required"`
	City         string `json:"city" validate:"required"`
	Province     string `json:"province" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	IsDefault    bool   `json:"is_default"`
}
