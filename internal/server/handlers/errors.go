package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/internal/server/views"
)

func resolveDuplicateKey(c *gin.Context, err repositories.ErrDuplicateKey) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error": err.Error(),
	})
}

func resolveNotFound(c *gin.Context, err repositories.ErrNotFound) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": err.Error(),
	})
}

func resolveInvalidValue(c *gin.Context, err views.ErrInvalidValue) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error":        err.Error(),
		"field":        err.FieldName,
		"valid_values": err.ValidValues,
	})
}

func ResolveError(c *gin.Context, err error) {
	var errDuplicateKey repositories.ErrDuplicateKey
	var errNotFound repositories.ErrNotFound
	var errInvalidValue views.ErrInvalidValue

	if errors.As(err, &errDuplicateKey) {
		resolveDuplicateKey(c, errDuplicateKey)
	} else if errors.As(err, &errNotFound) {
		resolveNotFound(c, errNotFound)
	} else if errors.As(err, &errInvalidValue) {
		resolveInvalidValue(c, errInvalidValue)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error processing request"})
	}
}
