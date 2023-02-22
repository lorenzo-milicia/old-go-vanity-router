package main

import (
	"fmt"
	"github.com/syumai/workers"
	"net/http"
	"strings"
)

func main() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, req *http.Request) {
		domainPath := req.URL.Path
		registryPath := registryPath(domainPath)

		w.Write([]byte(html(domainPath, registryPath)))
	}
	workers.Serve(handler)
}

func registryPath(path string) string {
	if path == "/" {
		return ""
	}
	return strings.Replace(path, "/libs", "/go-libs", 1)
}

func html(domainPath string, registryPath string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta name="go-import" content="go.lorenzomilicia.dev%v git https://github.com/lorenzo-milicia%v">
		</head>
		<body>
			Test...
		</body>
	</html>
`, domainPath, registryPath)
}
