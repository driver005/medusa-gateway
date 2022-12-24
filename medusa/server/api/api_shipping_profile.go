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

// DeleteShippingProfilesProfile - Delete a Shipping Profile
func DeleteShippingProfilesProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetShippingProfiles - List Shipping Profiles
func GetShippingProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetShippingProfilesProfile - Get a Shipping Profile
func GetShippingProfilesProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostShippingProfiles - Create a Shipping Profile
func PostShippingProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostShippingProfilesProfile - Update a Shipping Profile
func PostShippingProfilesProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}