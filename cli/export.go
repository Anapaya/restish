package cli

import (
	"github.com/logrusorgru/aurora"
)

type ApiConfigs = apiConfigs

// SetApiConfig sets the configs.
func SetApiConfig(c ApiConfigs) {
	configs = c
}

// SetName sets the name of the APIConfig.
func (a *APIConfig) SetName(name string) {
	a.name = name
}

// SetCurrentConfig sets the currentConfig.
func SetCurrentConfig(apiName string) {
	if cfg, ok := configs[apiName]; ok {
		currentConfig = cfg
	}
}

// SetTTY sets the tty.
func SetTTY(b bool) {
	tty = b
}

// SetAurora sets the aurora.
func SetAurora(a aurora.Aurora) {
	au = a
}

// ResetRegistries resets the registries used for internal bookkeeping.
func ResetRegistries() {
	authHandlers = map[string]AuthHandler{}
	contentTypes = []contentTypeEntry{}
	encodings = map[string]ContentEncoding{}
	linkParsers = []LinkParser{}
	loaders = []Loader{}
}

// GetGenericClosure exposes the internal generic function.
func GetGenericClosure() func(method string, addr string, args []string) {
	return generic
}

// GetEditClosure exposes the internal edit function.
func GetEditClosure() func(
	addr string,
	args []string,
	interactive,
	noPrompt bool,
	exitFunc func(int),
	editMarshal func(interface{}) ([]byte, error),
	editUnmarshal func([]byte, interface{}) error, ext string,
) {
	return edit
}

// InitConfig calls the internal initConfig function.
func InitConfig(appName string) {
	initConfig(appName, "")
}

// InitCache calls the internal initCache function.
func InitCache(appName string) {
	initCache(appName)
}

// MatchTemplate calls the internal matchTemplate function.
func MatchTemplate(url, template string) string {
	return matchTemplate(url, template)
}

// EnableVerbose sets the enableVerbose boolean to true.
func EnableVerbose() {
	enableVerbose = true
}

// IsVerbose returns the state of the internal enableVerbose boolean.
func IsVerbose() bool {
	return enableVerbose
}
