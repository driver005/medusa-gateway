package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

func (h *Handler) GetProductType(ctx context.Context, config core.Filter, model models.ProductType) (*models.ProductType, error) {
	var m models.ProductType

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetProductTypes(ctx context.Context, filters models.Filter, config core.Filter, model models.ProductType) ([]models.ProductType, *int64, error) {
	var m = make([]models.ProductType, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ProductType{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ProductType{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateProductType(ctx context.Context, m *models.ProductType) (*models.ProductType, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateProductType(ctx context.Context, m *models.ProductType) (*models.ProductType, error) {
	q := models.ProductType{}
	q.Id = m.Id

	o, err := h.GetProductType(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteProductType(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ProductType{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetProductTypesByIds(ctx context.Context, ProductIds []string) ([]models.ProductType, error) {
	var m = make([]models.ProductType, 0)

	for _, p := range ProductIds {
		Id, err := uuid.Parse(p)
		if err != nil {
			return nil, helper.ParseError(err)
		}

		q := models.ProductType{}
		q.Id = Id

		r, err := h.GetProductType(ctx, core.Filter{}, q)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
