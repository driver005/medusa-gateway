package database

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
)

func (h *Handler) GetCurrency(ctx context.Context, code string, config core.Filter, model models.Currency) (*models.Currency, error) {
	var m models.Currency

	if err := h.Query(ctx, config, model).Where("code = ?", code).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetCurrencys(ctx context.Context, filters models.Filter, config core.Filter, model models.Currency) ([]models.Currency, *int64, error) {
	var m = make([]models.Currency, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.Currency{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.Currency{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateCurrency(ctx context.Context, m *models.Currency) (*models.Currency, error) {
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateCurrency(ctx context.Context, m *models.Currency) (*models.Currency, error) {
	o, err := h.GetCurrency(ctx, m.Code, core.Filter{}, models.Currency{})
	if err != nil {
		return nil, helper.ParseError(err)
	}

	m.Code = o.Code
	if err := h.r.Manager(ctx).Model(&o).Updates(m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) DeleteCurrency(ctx context.Context, code string) error {
	if err := h.r.Manager(ctx).Where("code = ?", code).Delete(&models.Currency{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
