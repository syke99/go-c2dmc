package filter

// FilterOption returns a filter for filtering the colors
// populating the color bank when calling dmc.NewColorBankWithFilter
type FilterOption func(colors ...string) map[string]interface{}

// Sepia returns a filter option for filtering to a sepia color scheme
var Sepia FilterOption = func(colors ...string) map[string]interface{} {
	// TODO: populate map with Sepia colors
	return map[string]interface{}{}
}

// GreyScale returns a filter option for filtering to a greyscale color scheme
var GreyScale FilterOption = func(colors ...string) map[string]interface{} {
	// TODO: populate map with GreyScale colors
	return map[string]interface{}{}
}

// CustomColorFilter returns a filter option for filtering to a custom color scheme
var CustomColorFilter FilterOption = func(colors ...string) map[string]interface{} {
	colorMap := map[string]interface{}{}

	for _, color := range colors {
		var dud interface{}
		colorMap[color] = dud
	}

	return colorMap
}
