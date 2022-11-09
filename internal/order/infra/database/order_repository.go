package database

import (
	"devfullcycle/gointensivo/internal/order/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	err := r.DB.Create(&order)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

//func (r *OrderRepository) Save(order *entity.Order) error {
//	stmt, err := r.DB.Create("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	orer := entity.Order{}
	err := r.DB.Find(orer).Scan(&total)
	if err.Error != nil {
		return 0, err.Error
	}
	return total, nil
}
