package api

import "net/http"

func ListenAndServe(listenAddress, appLang string) *http.Server {

	routes := http.NewServeMux()
	routes.HandleFunc("api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api/v1"))
	})

	srv := &http.Server{
		Addr:    listenAddress,
		Handler: routes,
	}

	go func() {
		panic(srv.ListenAndServe())
	}()

	return srv
}
