package util

import (
	"fmt"

	"gorm.io/gorm"
)

type Teacher struct{}

// gorm的hook改寫會自動調用
// 事務開始前
func (t *Teacher) BeforeSave(tx *gorm.DB) error { // 進行一些數據檢查或初始化。
	fmt.Println("hook BeforeSave")
	return nil
}

func (t *Teacher) AfterSave(tx *gorm.DB) error { // 進行一些清理工作或後續處理。
	fmt.Println("hook AfterSave")
	return nil
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) error { // 設置默認值或檢查數據。
	fmt.Println("hook BeforeCreate")
	return nil
}

func (t *Teacher) AfterCreate(tx *gorm.DB) error { // 處理創建完成後的操作。
	fmt.Println("hook AfterCreate")
	return nil
}

func (t *Teacher) BeforeUpdate(tx *gorm.DB) error { // 檢查或修改數據。
	fmt.Println("hook BeforeUpdate")
	return nil
}

func (t *Teacher) AfterUpdate(tx *gorm.DB) error { // 處理更新完成後的操作。
	fmt.Println("hook AfterUpdate")
	return nil
}

func (t *Teacher) BeforeDelete(tx *gorm.DB) error { // 檢查是否允許刪除。
	fmt.Println("hook BeforeDelete")
	return nil
}

func (t *Teacher) AfterDelete(tx *gorm.DB) error { // 清理相關數據或進行後續處理。
	fmt.Println("hook AfterDelete")
	return nil
}

func (t *Teacher) AfterFind(tx *gorm.DB) error { // 處理查詢結果。
	fmt.Println("hook AfterFind")
	return nil
}
