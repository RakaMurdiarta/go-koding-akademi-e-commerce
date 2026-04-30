package delivery

type AddressResponse struct {
	ID           uint   `json:"id"`
	ReceiverName string `json:"receiver_name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	City         string `json:"city"`
	IsDefault    bool   `json:"is_default"`
}

type UserAddresses []AddressResponse
