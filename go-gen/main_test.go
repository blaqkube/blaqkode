package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (s *Suite) SetupTest() {
	s.VariableThatShouldStartAtFive = 5
}

func (s *Suite) TestExample() {
	assert.Equal(s.T(), 5, s.VariableThatShouldStartAtFive)
	s.Equal(5, s.VariableThatShouldStartAtFive)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
