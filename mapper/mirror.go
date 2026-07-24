package mapper

import (
	"registry-sync/config"
	"registry-sync/model"
)

func ConvertMirrors(
	items []config.MirrorConfig,
) []model.Mirror {

	if len(items) == 0 {
		return nil
	}

	mirrors := make(
		[]model.Mirror,
		0,
		len(items),
	)

	for _, item := range items {

		mirrors = append(
			mirrors,
			model.Mirror{
				Name: item.Name,
				URL:  item.URL,
				Type: model.MirrorType(item.Type),
				Auth: ConvertAuth(item.Auth),
			},
		)
	}

	return mirrors
}
