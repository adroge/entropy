package entropy

// SetEntropyMock allows replacing the call variable with the
// mock code found in the mock_entropy directory.
//
//	func (test *EntropyTestSuite) TestMock() {
//		controller := gomock.NewController(test.T())
//		entropyMock := mock_entropy.NewMockMethods(controller)
//		entropy.SetEntropyMock(entropyMock)
//		entropyMock.EXPECT().Calculate(gomock.Any()).Return(entropy.EntropyResult{}, entropy.ErrUnexpectedRune)
//		_, err := entropy.Calculate("some string")
//		require.NotNil(test.T(), err)
//		assert.True(test.T(), errors.Is(err, entropy.ErrUnexpectedRune))
//	}
func SetEntropyMock(mock Methods) {
	call = mock
}
