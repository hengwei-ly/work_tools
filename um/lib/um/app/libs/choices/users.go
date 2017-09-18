package choices

import (
	"strconv"

	"github.com/three-plus-three/forms"
	"github.com/three-plus-three/modules/permissions"
)

func Users(db *permissions.DB, hasEmpty bool) []forms.InputChoice {
	choices := []forms.InputChoice{}
	var users []permissions.User
	err := db.Users().Where().All(&users)
	if err != nil {
		panic(err)
	}

	if hasEmpty {
		choices = append(choices, forms.InputChoice{Value: "", Label: "(无)"})
	}

	for _, user := range users {
		choices = append(choices, forms.InputChoice{Value: strconv.FormatInt(user.ID, 10), Label: user.Name})
	}

	return choices
}
