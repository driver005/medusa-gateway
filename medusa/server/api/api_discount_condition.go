/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteDiscountsDiscountConditionsCondition - Delete a Condition
func DeleteDiscountsDiscountConditionsCondition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetDiscountsDiscountConditionsCondition - Get a Condition
func GetDiscountsDiscountConditionsCondition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostDiscountsDiscountConditions - Create a Condition
func PostDiscountsDiscountConditions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
