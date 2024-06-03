package serializer

import "gin_mall_tmp/model"

type Address struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
}

func BuildAddress(address *model.Address) Address {
	return Address{
		Id:        address.ID,
		UserId:    address.UserID,
		Name:      address.Name,
		Phone:     address.Phone,
		Address:   address.Address,
		CreatedAt: address.CreatedAt.Unix(),
	}
}

func BuildAddresses(addresses []*model.Address) (addressList []Address) {
	for _, address := range addresses {
		addressList = append(addressList, BuildAddress(address))
	}
	return
}
