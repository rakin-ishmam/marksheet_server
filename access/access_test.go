package access_test

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/access"
	"github.com/rakin-ishmam/marksheet_server/user"
)

func TestNewAccessor(t *testing.T) {
	tt := []struct {
		name  string
		owner user.Name
		users string

		expErr error
	}{
		{
			"valid1",
			user.GlobalUser(),
			fmt.Sprintf("%v", access.NewUserRight(user.TestUser(), access.SuperRights())),
			nil,
		},

		{
			"valid2",
			user.GlobalUser(),
			fmt.Sprintf(
				"%v u%v",
				access.NewUserRight(user.TestUser(), access.SuperRights()),
				access.NewUserRight(user.TestUser(), access.SuperRights()),
			),
			nil,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			accessor, err := access.NewAccessor(v.owner, v.users)
			test := testformat.NewTest(
				fmt.Sprintf("err %v", v.name),
				testformat.ConvVF(v.expErr),
				testformat.ConvVF(err),
			)
			if terr := test.Test(); terr != nil {
				t.Fatal(terr.Error())
				return
			}

			test = testformat.NewTest(
				fmt.Sprintf("value %v", v.name),
				testformat.ConvVF(v.users),
				testformat.ConvVF(accessor.String()),
			)

			if terr := test.Test(); terr != nil {
				t.Fatal(terr.Error())
			}

		})
	}
}
