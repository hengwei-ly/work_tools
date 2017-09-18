package choices

import (
	"github.com/three-plus-three/forms"
	"github.com/three-plus-three/modules/permissions"
)

func Tags(hasEmpty bool) ([]forms.InputChoice, error) {
	choices := []forms.InputChoice{}
	tags, err := permissions.GetPermissionTags()
	if err != nil {
		return nil, err
	}

	if hasEmpty {
		choices = append(choices, forms.InputChoice{Value: " ", Label: "(空)"})
	}

	for _, tag := range tags {
		choices = append(choices, forms.InputChoice{Value: tag.ID, Label: tag.Name})
	}
	return choices, nil
}
