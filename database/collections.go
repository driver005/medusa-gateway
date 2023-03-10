package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

func (h *Handler) GetCollection(ctx context.Context, config core.Filter, model models.ProductCollection) (*models.ProductCollection, error) {
	var m models.ProductCollection

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetCollections(ctx context.Context, filters models.Filter, config core.Filter, model models.ClaimOrder) ([]models.ProductCollection, *int64, error) {
	var m = make([]models.ProductCollection, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.ProductCollection{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.ProductCollection{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateCollection(ctx context.Context, m *models.ProductCollection) (*models.ProductCollection, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateCollection(ctx context.Context, m *models.ProductCollection) (*models.ProductCollection, error) {
	q := models.ProductCollection{}
	q.Id = m.Id

	o, err := h.GetCollection(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteCollection(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.ProductCollection{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) GetProductCollectionsByIds(ctx context.Context, Ids []string) ([]models.ProductCollection, error) {
	var m = make([]models.ProductCollection, 0)

	for _, p := range Ids {
		Id, err := uuid.Parse(p)
		if err != nil {
			return nil, helper.ParseError(err)
		}

		q := models.ProductCollection{}
		q.Id = Id

		r, err := h.GetCollection(ctx, core.Filter{}, q)
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
