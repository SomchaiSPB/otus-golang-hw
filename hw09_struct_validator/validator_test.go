package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			// Place your code here.
			in: User{
				ID:     "a5ca6f3f-2c56-4d83-a484-d732b23e43fb",
				Name:   "Bob",
				Age:    1,
				Email:  "mail@mail.com",
				Role:   "admin",
				Phones: []string{"12345678901", "12345678901"},
				meta:   nil,
			},
			expectedErr: MinimumValueViolationErr,
		},
		{
			// Place your code here.
			in: User{
				ID:     "a5ca6f3f-2c56-4d83-a484-d732b23e43fb",
				Name:   "Bob",
				Age:    25,
				Email:  "1112s",
				Role:   "admin",
				Phones: []string{"12345678901", "12345678901"},
				meta:   nil,
			},
			expectedErr: RegexpViolationErr,
		},
		{
			in: App{
				Version: "123",
			},
			expectedErr: LenViolationErr,
		},
		{
			in: Response{
				Code: 205,
			},
			expectedErr: NotInRangeViolationErr,
		},
		{
			// Place your code here.
			in: User{
				ID:     "a5ca6f3f-2c56-4d83-a484-d732b23e43fb",
				Name:   "Bob",
				Age:    100,
				Email:  "mail@mail.com",
				Role:   "admin",
				Phones: []string{"12345678901", "12345678901"},
				meta:   nil,
			},
			expectedErr: MaxValueViolationErr,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)

			require.EqualError(t, tt.expectedErr, err.Error())

			_ = tt
		})
	}
}

func NewUser() *User {
	return &User{
		ID:     "a5ca6f3f-2c56-4d83-a484-d732b23e43fb",
		Name:   "Bob",
		Age:    16,
		Email:  "mail@mail.com",
		Role:   "admin",
		Phones: []string{"12345678901", "12345678901"},
		meta:   nil,
	}
}
