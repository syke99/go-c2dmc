package filter

// FilterOption A filter for filtering the colors
// populating the color bank when calling dmc.NewColorBankWithFilter
type FilterOption map[string]interface{}

// Sepia returns a filter option for filtering to a sepia color scheme
var Sepia = func(colors ...string) FilterOption {
	// TODO: populate map with Sepia colors
	sepiaMap := map[string]interface{}{}

	return sepiaMap
}()

// GreyScale returns a filter option for filtering to a greyscale color scheme
var GreyScale = func(colors ...string) FilterOption {
	// TODO: populate map with GreyScale colors
	greyscaleMap := map[string]interface{}{}

	return greyscaleMap
}()

// CustomFilter returns a filter option for filtering to a custom color scheme
var CustomFilter = func(colors ...string) FilterOption {
	var colorMap map[string]interface{}
	var dud interface{}

	for _, color := range colors {
		colorMap[color] = dud
	}

	return colorMap
}

// ExtendFilter allows you to extend previously created filters with other filters
// and newly included colors
var ExtendFilter = func(baseFilter FilterOption, additionalFilters []FilterOption, colors ...string) FilterOption {
	var dud interface{}

	// loop over all additionalFilters (any previously
	// instantiated filter)
	for _, additionalFilter := range additionalFilters {
		// loop over every color in the additionalFilter
		for color, _ := range additionalFilter {
			// if the color doesn't exist in the baseFilter
			// add it
			if _, ok := baseFilter[color]; !ok {
				baseFilter[color] = dud
			}
		}
	}

	for _, color := range colors {
		// if the color doesn't exist in the baseFilter
		// add it
		if _, ok := baseFilter[color]; !ok {
			baseFilter[color] = dud
		}
	}

	return baseFilter
}
