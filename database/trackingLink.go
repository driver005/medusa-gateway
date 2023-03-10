package database

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

func (h *Handler) GetTrackingLink(ctx context.Context, config core.Filter, model models.TrackingLink) (*models.TrackingLink, error) {
	var m models.TrackingLink

	if err := h.Query(ctx, config, model).First(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return &m, nil
}

func (h *Handler) GetTrackingLinks(ctx context.Context, filters models.Filter, config core.Filter, model models.TrackingLink) ([]models.TrackingLink, *int64, error) {
	var m = make([]models.TrackingLink, 0)

	if err := h.Query(ctx, config, model).Scopes(Paginate(filters.Offset, filters.Size)).Order("id").Find(&m).Error; err != nil {
		return nil, nil, helper.ParseError(err)
	}

	n := h.Query(ctx, config, model).Find(&models.TrackingLink{})
	if n.Error != nil {
		return nil, nil, helper.ParseError(n.Error)
	}

	if m == nil {
		m = []models.TrackingLink{}
	}

	return m, &n.RowsAffected, nil
}

func (h *Handler) CreateTrackingLink(ctx context.Context, m *models.TrackingLink) (*models.TrackingLink, error) {
	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt
	if err := h.r.Manager(ctx).Create(&m).Error; err != nil {
		return nil, helper.ParseError(err)
	}

	return m, nil
}

func (h *Handler) UpdateTrackingLink(ctx context.Context, m *models.TrackingLink) (*models.TrackingLink, error) {
	q := models.TrackingLink{}
	q.Id = m.Id

	o, err := h.GetTrackingLink(ctx, core.Filter{}, q)
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

func (h *Handler) DeleteTrackingLink(ctx context.Context, id uuid.UUID) error {
	if err := h.r.Manager(ctx).Where("id = ?", id).Delete(&models.TrackingLink{}).Error; err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (h *Handler) CreateTrackingLinks(ctx context.Context, TrackingNumbers []string) ([]models.TrackingLink, error) {
	var m = make([]models.TrackingLink, 0)

	for _, tns := range TrackingNumbers {
		r, err := h.CreateTrackingLink(ctx, &models.TrackingLink{TrackingNumber: tns})
		if err != nil {
			return nil, helper.ParseError(err)
		}
		m = append(m, *r)
	}

	return m, nil
}
