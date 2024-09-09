package featureflags

var flags = map[string]bool{
	"new_product_page": false,
	"advanced_search":  true,
}

func IsEnabled(flag string) bool {
	if enabled, exists := flags[flag]; exists {
		return enabled
	}
	return false
}
