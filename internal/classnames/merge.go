package classnames

import twmerge "github.com/Oudwins/tailwind-merge-go"

// Merge resolves conflicting Tailwind utility classes while preserving order.
func Merge(classes ...string) string {
	return twmerge.Merge(classes...)
}
