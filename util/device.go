package util

import "github.com/gin-gonic/gin"

// GetIPAddress gets the IP address of the request
func GetIPAddress(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

// GetDeviceInfo gets the device information (user agent) from the request
func GetDeviceInfo(c *gin.Context) string {
	device := c.GetHeader("User-Agent")
	return device
}
