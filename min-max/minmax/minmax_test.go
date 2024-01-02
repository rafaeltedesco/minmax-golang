package minmax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SuiteMinMax struct {
	suite.Suite
	arr []int32
}

func (s *SuiteMinMax) SetupTest() {
	s.arr = []int32{1, 2, 3, 4, 5}
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(SuiteMinMax))
}

func (s *SuiteMinMax) TestGetHighest() {
	var (
		expectedPos int   = 4
		expectEl    int32 = 5
	)
	el, pos := getHighest(s.arr)

	assert.Equal(s.T(), expectEl, el)
	assert.Equal(s.T(), expectedPos, pos)
}

func (s *SuiteMinMax) TestGetLowest() {
	var (
		expectedPos int   = 0
		expectedEl  int32 = 1
	)
	el, pos := getLowest(s.arr)
	assert.Equal(s.T(), expectedEl, el)
	assert.Equal(s.T(), expectedPos, pos)
}

func (s *SuiteMinMax) TestGetMaxUsingEvaluator() {
	var (
		expectedPos int   = 4
		expectedEl  int32 = 5
	)
	el, pos := GetEl(s.arr, maxEvaluator)
	assert.Equal(s.T(), expectedEl, el)
	assert.Equal(s.T(), expectedPos, pos)
}

func (s *SuiteMinMax) TestGetMinUsingEvaluator() {
	var (
		expectedPos int   = 0
		expectEl    int32 = 1
	)
	el, pos := GetEl(s.arr, minEvaluator)

	assert.Equal(s.T(), expectEl, el)
	assert.Equal(s.T(), expectedPos, pos)
}

func (s *SuiteMinMax) TestGetNLowestElementsFromArray() {
	var (
		expectedArr = []int32{1, 2, 3, 4}
	)

	arr := GetNLowestElementsFromArray(4, s.arr)

	assert.Equal(s.T(), expectedArr, arr)
}

func (s *SuiteMinMax) TestSortAscending() {
	var (
		expectedArr = s.arr
	)
	unsortedArr := []int32{5, 4, 3, 2, 1}
	arr := Sort(unsortedArr, ascending)
	assert.Equal(s.T(), expectedArr, arr)
}

func (s *SuiteMinMax) TestGetNHighestElementsFromArray() {
	var (
		expectedArr = []int32{2, 3, 4, 5}
	)

	arr := GetNHighestElementsFromArray(4, s.arr)

	assert.Equal(s.T(), expectedArr, arr)
}

func (s *SuiteMinMax) TestSumMinNElements() {
	var (
		expected int32 = 10
	)
	response, err := SumMinNElements(4, s.arr)
	s.Nil(err)
	assert.Equal(s.T(), expected, response)
}

func (s *SuiteMinMax) TestSumMaxNElements() {
	var (
		expected int32 = 14
	)
	response, err := SumMaxNElements(4, s.arr)
	s.Nil(err)
	assert.Equal(s.T(), expected, response)
}

func (s *SuiteMinMax) TestCannotAskForMoreElementsThanIsInSlice() {
	_, err := SumMaxNElements(10, s.arr)
	s.NotNil(err)
	assert.Equal(s.T(), "invalid length", err.Error())
	_, err = SumMinNElements(5, s.arr)
	s.NotNil(err)
	assert.Equal(s.T(), "invalid length", err.Error())
}

func (s *SuiteMinMax) TestNewMinMaxCase() {
	var (
		inputs = []int32{256741038, 623958417, 467905213, 714532089, 938071625}
		output = []int64{2063136757, 2744467344}
	)
	result, err := SumMaxNElements(len(inputs)-1, inputs)
	s.Nil(err)
	s.Equal(output[1], result)

	result, err = SumMinNElements(len(inputs)-1, inputs)
	s.Nil(err)
	s.Equal(output[0], result)
}
