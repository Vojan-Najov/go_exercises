func main() {
	r := chi.NewRouter

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", HelloWorld)
		r.Get("/{AssetUrl}", GetAsset)
		r.Get("/manage/url/{path}", FetchAssetDetailsByURL)
		r.Get("/manage/id/{path}", FetchAssetDetailsByID)
	})

	// PrivateRoutes
	// Require Authentication
	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Post("/manage", CreateAsset)
	})
}
