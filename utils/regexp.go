package utils

import "regexp"

var (
	RegexpPaginationContainerID = regexp.MustCompile(`^(\d+)/(\d+)$`)
)
