package entropy_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/adroge/entropy"
	"github.com/adroge/entropy/mock_entropy"
)

type EntropyTestSuite struct {
	suite.Suite
}

type EntropyMockTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(EntropyTestSuite))
	suite.Run(t, new(EntropyMockTestSuite))
}

func (test *EntropyTestSuite) SetupTest() {
	alphabets := []string{
		`abcdefghijklmnopqrstuvwxyz`,
		`ABCDEFGHIJKLMNOPQRSTUVWXYZ`,
		`1234567890`,
		`!@#$%^&*`,
		`()[]{}<>`,
		`~-_=+|;:',./?\ "` + "`",
	}
	err := entropy.Alphabets(alphabets)
	require.Nil(test.T(), err)
	err = entropy.Bounds(30.0, 40.0, 60.0, 127.0)
	require.Nil(test.T(), err)
	err = entropy.Descriptions(entropy.DescriptionTags{
		Invalid:    "invalid",
		VeryWeak:   "very weak",
		Weak:       "weak",
		Reasonable: "reasonable",
		Strong:     "strong",
		VeryStrong: "very strong"})
	require.Nil(test.T(), err)
}

func (test *EntropyTestSuite) TestCalculateEmptyString() {
	result, err := entropy.Calculate("")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), "very weak", result.String())
	assert.Equal(test.T(), entropy.VeryWeak, result.Evaluation)
}

func (test *EntropyTestSuite) TestVeryWeak() {
	result, err := entropy.Calculate("abcd")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 18.801758872564367, result.Bits)
	assert.Equal(test.T(), "very weak", result.String())
	assert.Equal(test.T(), entropy.VeryWeak, result.Evaluation)
}

func (test *EntropyTestSuite) TestWeak() {
	result, err := entropy.Calculate("abcdefgh")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 37.603517745128734, result.Bits)
	assert.Equal(test.T(), "weak", result.String())
	assert.Equal(test.T(), entropy.Weak, result.Evaluation)
}

func (test *EntropyTestSuite) TestReasonable() {
	result, err := entropy.Calculate("abcdEFGH")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 45.603517745128734, result.Bits)
	assert.Equal(test.T(), "reasonable", result.String())
	assert.Equal(test.T(), entropy.Reasonable, result.Evaluation)
}

func (test *EntropyTestSuite) TestReasonableWithNumber() {
	result, err := entropy.Calculate("abcdE1GH")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 47.633570483095, result.Bits)
	assert.Equal(test.T(), "reasonable", result.String())
	assert.Equal(test.T(), entropy.Reasonable, result.Evaluation)
}

func (test *EntropyTestSuite) TestReasonableWithNumberAndSymbol() {
	result, err := entropy.Calculate("ab$dE1GH")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 49.03426413555973, result.Bits)
	assert.Equal(test.T(), "reasonable", result.String())
	assert.Equal(test.T(), entropy.Reasonable, result.Evaluation)
}

func (test *EntropyTestSuite) TestStrong() {
	result, err := entropy.Calculate("Bc8$5yjvK>8Y")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 75.42482662634698, result.Bits)
	assert.Equal(test.T(), "strong", result.String())
	assert.Equal(test.T(), entropy.Strong, result.Evaluation)
}

func (test *EntropyTestSuite) TestVeryStrong() {
	result, err := entropy.Calculate("J90;1]6rtpZ4ny;.EZ:wW")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 135.3018134128233, result.Bits)
	assert.Equal(test.T(), "very strong", result.String())
	assert.Equal(test.T(), entropy.VeryStrong, result.Evaluation)
}

func (test *EntropyTestSuite) TestRuneNotInAlphabet() {
	_, err := entropy.Calculate("zufälliges Passwort")
	require.NotNil(test.T(), err)
	assert.True(test.T(), errors.Is(err, entropy.ErrUnexpectedRune))
	assert.Equal(test.T(), "unexpected rune in input: ä", err.Error())
}

func (test *EntropyTestSuite) TestRuneInLatinAlphabet() {
	alphabets := []string{
		`äabcdefghijklmnopqrstuvwxyz`,
		`ABCDEFGHIJKLMNOPQRSTUVWXYZ`,
		`~-_=+|;:',./? \"` + "`",
	}
	err := entropy.Alphabets(alphabets)
	assert.Nil(test.T(), err)
	result, err := entropy.Calculate("zufälliges Passwort")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), "strong", result.String())
	assert.Equal(test.T(), entropy.Strong, result.Evaluation)
}

func (test *EntropyTestSuite) TestAlphabetsEmpty() {
	err := entropy.Alphabets([]string{})
	require.NotNil(test.T(), err)
	assert.True(test.T(), errors.Is(err, entropy.ErrInvalidAlphabet))
}

func (test *EntropyTestSuite) TestChangeBounds() {
	err := entropy.Bounds(5, 10, 15, 20)
	assert.Nil(test.T(), err)
	result, err := entropy.Calculate("monkey")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 28.202638308846552, result.Bits)
	assert.Equal(test.T(), "very strong", result.String())
	assert.Equal(test.T(), entropy.VeryStrong, result.Evaluation)
}

func (test *EntropyTestSuite) TestChangeBoundsBadInput() {
	err := entropy.Bounds(50, 10, 15, 20)
	assert.NotNil(test.T(), err)
	assert.True(test.T(), errors.Is(err, entropy.ErrInvalidEntropy))
}

func (test *EntropyTestSuite) TestChangeDescriptionInvalid() {
	err := entropy.Descriptions(entropy.DescriptionTags{"", "", "", "", "", ""})
	assert.NotNil(test.T(), err)
	assert.True(test.T(), errors.Is(err, entropy.ErrInvalidDescription))
}

func (test *EntropyTestSuite) TestChangeDescriptionVeryWeak() {
	err := entropy.Descriptions(entropy.DescriptionTags{"invalid", "pitiful", "laughable", "passable", "impressive", "ludicrous"})
	assert.Nil(test.T(), err)
	result, err := entropy.Calculate("monkey")
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), "pitiful", result.String())
	assert.Equal(test.T(), entropy.VeryWeak, result.Evaluation)
}

func (test *EntropyTestSuite) TestEntropyBounds() {
	bounds := entropy.EntropyBounds()
	assert.EqualValues(test.T(), []float64{30.0, 40.0, 60.0, 127.0}, bounds)
}

func (test *EntropyMockTestSuite) TestMock() {
	controller := gomock.NewController(test.T())
	entropyMock := mock_entropy.NewMockFunction(controller)
	entropy.SetMock(entropyMock)
	entropyMock.EXPECT().Calculate(gomock.Any()).Return(entropy.EntropyResult{}, entropy.ErrUnexpectedRune)
	_, err := entropy.Calculate("password")
	require.NotNil(test.T(), err)
	assert.True(test.T(), errors.Is(err, entropy.ErrUnexpectedRune))
}
