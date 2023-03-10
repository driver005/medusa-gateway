package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

func (h *Handler) GetReturnReason(ctx context.Context, config core.Filter, model models.ReturnReason) (*models.ReturnReason, error) {
	var m models.ReturnReason

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetReturnReasons(ctx context.Context, filters models.Filter, config core.Filter, model models.ReturnReason) ([]models.ReturnReason, *int64, error) {
	var m = make([]models.ReturnReason, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ReturnReason{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ReturnReason{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateReturnReason(ctx context.Context, m *models.ReturnReason) (*models.ReturnReason, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateReturnReason(ctx context.Context, m *models.ReturnReason) (*models.ReturnReason, error) {
	q := models.ReturnReason{}
	q.Id = m.Id

	o, err := h.GetReturnReason(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteReturnReason(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ReturnReason{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}
