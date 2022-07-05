package filter

// FilterOption A filter for filtering the colors
// populating the color bank when calling dmc.NewColorBankWithFilter
type FilterOption map[string]interface{}

// Sepia returns a filter option for filtering to a sepia color scheme
var Sepia = func(colors ...string) FilterOption {
	// TODO: populate map with Sepia colors
	sepiaMap := map[string]interface{}{}

	return sepiaMap
}

// GreyScale returns a filter option for filtering to a greyscale color scheme
var GreyScale = func(colors ...string) FilterOption {
	// TODO: populate map with GreyScale colors
	greyscaleMap := map[string]interface{}{}

	return greyscaleMap
}

// CustomColorFilter returns a filter option for filtering to a custom color scheme
var CustomColorFilter = func(colors ...string) FilterOption {
	colorMap := map[string]interface{}{}

	for _, color := range colors {
		var dud interface{}
		colorMap[color] = dud
	}

	return colorMap
}
