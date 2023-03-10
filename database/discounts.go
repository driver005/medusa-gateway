package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
)

func (h *Handler) GetDiscount(ctx context.Context, config types.FilterableDiscountProps, model models.Discount) (*models.Discount, error) {
	var m models.Discount

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetDiscounts(ctx context.Context, filters models.Filter, config types.FilterableDiscountProps, model models.Discount) ([]models.Discount, *int64, error) {
	var m = make([]models.Discount, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.Discount{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Discount{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) GetDiscountByCode(ctx context.Context, code string) (*models.Discount, error) {
	var m models.Discount

	if err := h.r.Manager(ctx).Where("code = ?", code).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) CreateDiscount(ctx context.Context, m *models.Discount) (*models.Discount, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateDiscount(ctx context.Context, m *models.Discount) (*models.Discount, error) {
	q := models.Discount{}
	q.Id = m.Id

	o, err := h.GetDiscount(ctx, types.FilterableDiscountProps{}, q)
	if err != nil {
		return nil, helper.ParseError(err)
	}

	m.Id = o.Id
	m.UpdatedAt = time.Now().UTC().Round(time.Second)
	if err := h.r.Manager(ctx).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteDiscount(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.Discount{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) DeleteDynamicCode(ctx context.Context, id uuid.UUID, code string) error {
	if err := h.r.Manager(ctx).Where("id = ? AND code = ?", id, code).Delete(&models.Discount{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
