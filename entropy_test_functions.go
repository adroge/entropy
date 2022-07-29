package entropy

// SetMock allows replacing the call variable with the
// mock code found in the mock_entropy directory.
//
//	func (test *EntropyTestSuite) TestMock() {
//		controller := gomock.NewController(test.T())
//		entropyMock := mock_entropy.NewMockFunction(controller)
//		entropy.SetMock(entropyMock)
//		entropyMock.EXPECT().Calculate(gomock.Any()).Return(entropy.EntropyResult{}, entropy.ErrUnexpectedRune)
//		_, err := entropy.Calculate("password")
//		require.NotNil(test.T(), err)
//		assert.True(test.T(), errors.Is(err, entropy.ErrUnexpectedRune))
//	}
func SetMock(mock Function) {
	call = mock
}
