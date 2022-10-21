package cli

import (
	"github.com/logrusorgru/aurora"
)

type ApiConfigs = apiConfigs

// SetApiConfig configs to the specified ApiConfigs
// No dynamic managed APIs by restish required.
func SetApiConfig(c ApiConfigs) {
	configs = c
}

// SetName sets the name of the APIConfig to the specified string
// Setter for private field required.
func (a *APIConfig) SetName(name string) {
	a.name = name
}

// SetCurrentConfig sets the currentConfig to the specified one.
// Setter for private variable required.
func SetCurrentConfig(apiName string) {
	if cfg, ok := configs[apiName]; ok {
		currentConfig = cfg
	}
}

func SetTTY(b bool) {
	tty = b
}

func SetAurora(a aurora.Aurora) {
	au = a
}

func ResetRegistries() {
	authHandlers = map[string]AuthHandler{}
	contentTypes = []contentTypeEntry{}
	encodings = map[string]ContentEncoding{}
	linkParsers = []LinkParser{}
	loaders = []Loader{}
}

func GetGenericClosure() func(method string, addr string, args []string) {
	return generic
}

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

func InitConfig(appName string) {
	initConfig(appName, "")
}

func InitCache(appName string) {
	initCache(appName)
}

func MatchTemplate(url, template string) string {
	return matchTemplate(url, template)
}

func EnableVerbose() {
	enableVerbose = true
}

func IsVerbose() bool {
	return enableVerbose
}
