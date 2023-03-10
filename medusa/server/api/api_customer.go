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

// GetCustomers - List Customers
func GetCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetCustomersCustomer - Get a Customer
func GetCustomersCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostCustomers - Create a Customer
func PostCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostCustomersCustomer - Update a Customer
func PostCustomersCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
