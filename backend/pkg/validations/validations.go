package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValidName(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	regex := `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
	return regexp.MustCompile(regex).MatchString(value)
}

/*
Validation for path that must start with /
no trailing slash
segments contain only letters, digits, underscores, or dashes
no empty segments (i.e., no //)
*/

func IsValidPath(fl validator.FieldLevel) bool {
	pathRegex := regexp.MustCompile(`^\/[a-zA-Z0-9\-\/_]*$`)
	return pathRegex.MatchString(fl.Field().String())
}

/*
Validatrion for validate a host name
RFC 1123-compliant regex (without lookaheads)
(?i)           → case-insensitive
^...$          → anchors the pattern to the entire string
([a-z0-9]      → label must start with alphanumeric
([a-z0-9\-]{0,61}[a-z0-9])? → middle of label: alphanum or hyphen, ends with alphanum
\.)*           → allow multiple labels ending in a dot
(...)+         → final label without a trailing dot
*/
func IsValidHostName(fl validator.FieldLevel) bool {
	hostRegex := regexp.MustCompile(`(?i)^([a-z0-9]([a-z0-9\-]{0,61}[a-z0-9])?\.)*([a-z0-9]([a-z0-9\-]{0,61}[a-z0-9])?)$`)
	return hostRegex.MatchString(fl.Field().String())
}

// Regular expression to match:
// - http:// or https://
// - domain name (with optional subdomains), or IPv4/IPv6
// - optional port number
// - optional path
func IsValidBackendUrl(fl validator.FieldLevel) bool {
	var backendURLRegex = `^https?://([a-zA-Z0-9\-._~%]+|\[[a-fA-F0-9:]+\])(:\d+)?(/[\S]*)?$`
	re := regexp.MustCompile(backendURLRegex)
	return re.MatchString(fl.Field().String())
}
