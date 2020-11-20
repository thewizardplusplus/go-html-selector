package htmlselector

// TagName ...
type TagName string

// ...
const (
	UniversalTag TagName = "*"
)

// AttributeName ...
type AttributeName string

// FilterGroup ...
type FilterGroup map[TagName][]AttributeName

// OptimizedFilterGroup ...
type OptimizedFilterGroup map[TagName]OptimizedAttributeFilterGroup

// OptimizedAttributeFilterGroup ...
type OptimizedAttributeFilterGroup map[AttributeName]struct{}

// OptimizeFilters ...
func OptimizeFilters(
	filters FilterGroup,
	options ...Option,
) OptimizedFilterGroup {
	config := newOptionConfig(options)

	optimizedFilters := make(OptimizedFilterGroup)
	var universalTagAttributes OptimizedAttributeFilterGroup
	for tag, attributes := range filters {
		if config.skipEmptyTags && len(attributes) == 0 {
			continue
		}

		optimizedAttributeFilters := make(OptimizedAttributeFilterGroup)
		for _, attribute := range attributes {
			if _, ok := universalTagAttributes[attribute]; ok {
				continue
			}

			optimizedAttributeFilters[attribute] = struct{}{}
		}

		optimizedFilters[tag] = optimizedAttributeFilters
		if tag == UniversalTag {
			universalTagAttributes = optimizedAttributeFilters
		}
	}

	return optimizedFilters
}
