package choices

import (
	"strconv"

	"github.com/three-plus-three/forms"
	"github.com/three-plus-three/modules/permissions"
)

func Role(db *permissions.DB, hasEmpty bool) []forms.InputChoice {
	choices := []forms.InputChoice{}
	var roles []permissions.Role
	err := db.Roles().Where().All(&roles)
	if err != nil {
		panic(err)
	}
	if hasEmpty {
		choices = append(choices, forms.InputChoice{Value: "", Label: "(无)"})
	}

	for _, role := range roles {
		choices = append(choices, forms.InputChoice{Value: strconv.FormatInt(role.ID, 10) + ":" + role.Name + ":" + role.Description, Label: role.Name})
	}

	return choices
}

func Remove(db *permissions.DB, hasEmpty bool, selectedRoles []permissions.Role) []forms.InputChoice {
	choices := []forms.InputChoice{}
	var roles []permissions.Role
	err := db.Roles().Where().All(&roles)
	if err != nil {
		panic(err)
	}
	if hasEmpty {
		choices = append(choices, forms.InputChoice{Value: "", Label: "(无)"})
	}

	for _, role := range roles {
		if len(selectedRoles) != 0 {
			for _, selectedRole := range selectedRoles {
				if selectedRole.ID != role.ID {
					choices = append(choices, forms.InputChoice{Value: strconv.FormatInt(role.ID, 10) + ":" + role.Name + ":" + role.Description, Label: role.Name})
				}
			}
		} else {
			choices = append(choices, forms.InputChoice{Value: strconv.FormatInt(role.ID, 10) + ":" + role.Name + ":" + role.Description, Label: role.Name})
		}
	}
	return choices
}
