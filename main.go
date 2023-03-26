package main

import (
	"fmt"
	"net/http"

	"library-sample/handlers"
	"library-sample/router"
)

func main() {
	mux := http.NewServeMux()
	appRouter := router.NewRouter()

	bookRoutes := handlers.BookRoutes()
	appRouter.RegisterRoutes(bookRoutes)

	userRoutes := handlers.UserRoutes()
	appRouter.RegisterRoutes(userRoutes)

	borrowRoutes := handlers.BorrowRoutes()
	appRouter.RegisterRoutes(borrowRoutes)

	appRouter.SetupRoutes(mux)

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", mux)
}
