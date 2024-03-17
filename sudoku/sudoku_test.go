package sudoku

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SudokuSuite struct {
	suite.Suite
}

func (suite *SudokuSuite) SetupTest() {
}

func (suite *SudokuSuite) TestIsSafe() {
	b := &Board{
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	suite.True(isSafe(b, 0, 0, 2))
	suite.False(isSafe(b, 0, 0, 1))
	suite.False(isSafe(b, 0, 0, 3))
}

func (suite *SudokuSuite) TestFindEmpty() {
	b := &Board{
		{7, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	r, c, ok := findEmpty(b)
	suite.True(ok)
	suite.Equal(0, r)
	suite.Equal(2, c)

	b = &Board{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	r, c, ok = findEmpty(b)
	suite.False(ok)
	suite.Equal(0, r)
	suite.Equal(0, c)
}

func TestSudokuSuite(t *testing.T) {
	suite.Run(t, new(SudokuSuite))
}
