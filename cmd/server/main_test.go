package main_test

import (
	"github.com/scarlettmiss/bestPal/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
*
testing suit for the application.
*/
func TestApplicationCreation(t *testing.T) {
	_, err := application.New(application.Options{})
	assert.Nil(t, err)
}
