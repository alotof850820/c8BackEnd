package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (a *AddressService) CreateAddress(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		Name:    a.Name,
		Phone:   a.Phone,
		Address: a.Address,
		UserID:  uId,
	}

	err := addressDao.CreateAddress(address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (a *AddressService) GetAddresses(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	addressDao := dao.NewAddressDao(ctx)

	addresses, err := addressDao.GetAddressesByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddresses(addresses),
	}
}

func (a *AddressService) GetAddress(ctx context.Context, aId string, uId uint) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	address, err := addressDao.GetAddressByAId(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}

func (a *AddressService) UpdateAddress(ctx context.Context, aId string, uId uint) serializer.Response {
	code := e.Success
	addressId, _ := strconv.Atoi(aId)
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uId,
		Name:    a.Name,
		Phone:   a.Phone,
		Address: a.Address,
	}
	err := addressDao.UpdateAddressByAId(uint(addressId), address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}

func (a *AddressService) DeleteAddress(ctx context.Context, aId string, uId uint) serializer.Response {
	code := e.Success
	addressId, _ := strconv.Atoi(aId)
	addressDao := dao.NewAddressDao(ctx)
	err := addressDao.DeleteAddressByAId(uint(addressId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
