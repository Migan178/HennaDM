package utils

import (
	"fmt"
	"strings"
)

const (
	PaginationEmbedPrev    = "#mf-pages/prev$"
	PaginationEmbedPages   = "#mf-pages/pages$"
	PaginationEmbedNext    = "#mf-pages/next$"
	PaginationEmbedModal   = "#mf-pages/modal$"
	PaginationEmbedSetPage = "#mf-pages/modal/set$"
)

func MakePaginationEmbedPrev(id string) string {
	return fmt.Sprintf("%s%s", PaginationEmbedPrev, id)
}

func MakePaginationEmbedPages(id string) string {
	return fmt.Sprintf("%s%s", PaginationEmbedPages, id)
}

func MakePaginationEmbedNext(id string) string {
	return fmt.Sprintf("%s%s", PaginationEmbedNext, id)
}

func MakePaginationEmbedModal(id string) string {
	return fmt.Sprintf("%s%s", PaginationEmbedModal, id)
}

func MakePaginationEmbedSetPage(id string) string {
	return fmt.Sprintf("%s%s", PaginationEmbedSetPage, id)
}

func GetPaginationEmbedID(customID string) string {
	switch {
	case strings.HasPrefix(customID, PaginationEmbedPrev):
		return customID[len(PaginationEmbedPrev):]
	case strings.HasPrefix(customID, PaginationEmbedPages):
		return customID[len(PaginationEmbedPages):]
	case strings.HasPrefix(customID, PaginationEmbedNext):
		return customID[len(PaginationEmbedNext):]
	case strings.HasPrefix(customID, PaginationEmbedModal):
		return customID[len(PaginationEmbedModal):]
	case strings.HasPrefix(customID, PaginationEmbedSetPage):
		return customID[len(PaginationEmbedSetPage):]
	default:
		return customID
	}
}

func GetPaginationEmbedUserID(id string) string {
	return RegexpPaginationContainerID.FindAllStringSubmatch(id, 1)[0][1]
}
