package tekmetric

import (
	"context"
	"fmt"
)

// ============================================================================
// Models
// ============================================================================

// Shop represents a Tekmetric shop
type Shop struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Nickname             string  `json:"nickname"`
	Phone                string  `json:"phone"`
	Email                string  `json:"email"`
	Website              string  `json:"website"`
	Address              Address `json:"address"`
	ROCustomLabelEnabled bool    `json:"roCustomLabelEnabled"`
}

// ============================================================================
// API Methods
// ============================================================================

// GetShops returns all shops accessible by the current token
func (c *Client) GetShops(ctx context.Context) ([]Shop, error) {
	var shops []Shop
	if err := c.doRequest(ctx, "GET", "/api/v1/shops", nil, &shops); err != nil {
		return nil, err
	}
	return shops, nil
}

// GetShop returns a specific shop by ID
func (c *Client) GetShop(ctx context.Context, id int) (*Shop, error) {
	var shop Shop
	path := fmt.Sprintf("/api/v1/shops/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &shop); err != nil {
		return nil, err
	}
	return &shop, nil
}
