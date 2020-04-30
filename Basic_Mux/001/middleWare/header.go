package middleWare

import "net/http"

func setHeaders (w *http.ResponseWriter) {
	(*w).Header().Set("Mike", "Winning")
}
