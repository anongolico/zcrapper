package helpers

import (
	"strings"

	"github.com/anongolico/ib/schemas"
)


// ScanFormats returns a map of string slices, containing
// the hashes of the files from a given Post
func ScanFormats(post *schemas.Post) (map[string][]string, int) {
	// declare variables used through the function
	formats := make(map[string][]string, 0)
	var format string

	commentsWithAttachment := 0

	for _, v := range post.Comentarios {
		hash := v.Media.Url

		if hash == "" || strings.Contains(hash, "//") {
			continue
		}

		commentsWithAttachment++
		_, format, _ = strings.Cut(hash, ".")
		formats[format] = append(formats[format], hash)
	}

	return formats, commentsWithAttachment
}
