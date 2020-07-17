package htmlselector

// TagName ...
type TagName string

// AttributeName ...
type AttributeName string

// FilterGroup ...
type FilterGroup map[TagName][]AttributeName

// OptimizedFilterGroup ...
type OptimizedFilterGroup map[TagName]OptimizedAttributeFilterGroup

// OptimizedAttributeFilterGroup ...
type OptimizedAttributeFilterGroup map[AttributeName]struct{}

// OptimizeFilters ...
func OptimizeFilters(filters FilterGroup) OptimizedFilterGroup {
	optimizedFilters := make(OptimizedFilterGroup)
	for tag, attributes := range filters {
		optimizedAttributeFilters := make(OptimizedAttributeFilterGroup)
		for _, attribute := range attributes {
			optimizedAttributeFilters[attribute] = struct{}{}
		}

		optimizedFilters[tag] = optimizedAttributeFilters
	}

	return optimizedFilters
}
