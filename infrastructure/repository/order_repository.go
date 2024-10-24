package repository

import (
	"DDD-OrderingSystem/Domain/Model"
	"gorm.io/gorm"
)

// OrderRepository 订单仓库接口
type OrderRepository interface {
	Save(ctx context.Context, order *Model.Order) error
	FindById(ctx context.Context, id uint64) (*Model.Order, error)
	FindAll(ctx context.Context) ([]Model.Order, error)
}

// OrderRepositoryImpl 订单仓库实现
type OrderRepositoryImpl struct {
	db *gorm.DB
}

// NewOrderRepository 创建订单仓库实例
func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db: db,
	}
}

// Save 保存订单
func (r *OrderRepositoryImpl) Save(ctx context.Context, order *Model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

// FindById 通过ID查找订单
func (r *OrderRepositoryImpl) FindById(ctx context.Context, id uint64) (*Model.Order, error) {
	var order Model.Order
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// FindAll 查找所有订单
func (r *OrderRepositoryImpl) FindAll(ctx context.Context) ([]Model.Order, error) {
	var orders []Model.Order
	if err := r.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
