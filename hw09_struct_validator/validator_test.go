package hw09structvalidator

import (
	"encoding/json"
	"fmt"
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
	user := NewUser()

	err := Validate(*user)

	fmt.Println(err)

	//require.NoError(t, err)

	//tests := []struct {
	//	in          interface{}
	//	expectedErr error
	//}{
	//	{
	//		// Place your code here.
	//		in: NewUser(),
	//	},
	//	// ...
	//	// Place your code here.
	//}
	//
	//for i, tt := range tests {
	//	t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
	//		tt := tt
	//		t.Parallel()
	//
	//		// Place your code here.
	//		err := Validate(NewUser())
	//
	//		fmt.Println(err)
	//		_ = tt
	//	})
	//}
}

func NewUser() *User {
	return &User{
		ID:     "a5ca6f3f-2c56-4d83-a484-d732b23e43fb",
		Name:   "Bob",
		Age:    16,
		Email:  "mail@mail.com",
		Role:   "admin",
		Phones: []string{"89086663", "098768992"},
		meta:   nil,
	}
}
