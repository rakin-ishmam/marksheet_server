package user_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/user"
)

func TestValid(t *testing.T) {
	tt := []struct {
		name string
		val  user.Name
		res  bool
	}{
		{
			name: "min len",
			val:  "r",
			res:  false,
		},
		{
			name: "invalid char",
			val:  "_testtest",
			res:  false,
		},
		{
			name: "valid",
			val:  "usertest",
			res:  true,
		},
	}

	for _, tcase := range tt {
		if res := tcase.val.Valid(); res != tcase.res {
			t.Run(tcase.name, func(t *testing.T) {
				t.Fatalf("%s expected %v but got %v", tcase.name, tcase.res, res)
			})
		}

	}
}
