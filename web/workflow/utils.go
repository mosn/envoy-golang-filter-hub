package main

import "fmt"

func BuildTagName(pluginName string, version string) string {
	return fmt.Sprintf("%s|v%s", pluginName, version)
}
