package api

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type customError struct{}

func (customError) Code() int {
	return 404
}

func (customError) Error() string {
	return "whatever"
}

func TestErrorIs(t *testing.T) {
	type testErrors struct {
		err         error
		expectedRes bool
		desc        string
	}
	err := HttpError{code: 404, message: "NotFound"}

	wrappedErr := errors.Wrap(err, "wrapped error")
	tests := []testErrors{
		{
			err:         err,
			expectedRes: true,
			desc:        "SameError",
		},
		{
			err:         errors.New("random error"),
			expectedRes: false,
			desc:        "UnequalError",
		},
		{
			err:         HttpError{code: 404, message: "NotFound"},
			expectedRes: true,
			desc:        "NewError",
		},

		{
			err:         customError{},
			expectedRes: true,
			desc:        "CustomError",
		},
	}

	for _, test := range tests {
		res := errors.Is(err, test.err)
		assert.Equal(t, res, test.expectedRes, test.desc)

		res = errors.Is(wrappedErr, test.err)
		assert.Equal(t, res, test.expectedRes, "Wrapped_"+test.desc)
	}
}
