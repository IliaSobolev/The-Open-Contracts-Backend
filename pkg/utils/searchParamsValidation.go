package utils

import "TOC/pkg/domain"

func LangValidation(lang string) bool {
	found := false
	for _, language := range domain.Languages {
		if language == lang {
			found = true
			break
		}
	}
	return found
}
