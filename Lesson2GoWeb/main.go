package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// The context package is used to manage the context of a request.
	ctx := context.Background()

	// WithValue() is a function that takes a context and a key-value pair.
	// We can send everything we want as a value
	ctx = context.WithValue(ctx, "my-Key", "my-Value")
	ctx = context.WithValue(ctx, "my-Key-int", 5)
	ctx = context.WithValue(ctx, "api-key", "super-secret!")

	// * if we try to access a non-existent key-value, we'll get a nil from ctx.Value

	viewContext(ctx)

	// * A context with timeout will live only in the specified time, then it will expire
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel() // This allow us to cancel a context and propagate it between layers

	go myProcess(ctx2)

	// Use the context until we reach the timeout
	select {
	case <-ctx2.Done():
		fmt.Println("We reached ctx timeout")
		fmt.Println(ctx.Err())
	}

}

// Print the context given as an argument.
func viewContext(ctx context.Context) {
	fmt.Printf("My value in context is %s\n", ctx.Value("my-Key"))
	fmt.Printf("My numeric value in context is %d\n", ctx.Value("my-Key-int"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ohh, time out!")
			return
		default:
			fmt.Println("Doing something useless!")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// * How to aggregate routes with chi!

/* func main() {
// create a new router
router := chi.NewRouter()
// register endpoints: method + path + handler
router.Get("/health", HandlerHealthCheck)
// - profile group
router.Route("/profile", func(r chi.Router) {
r.Get("/about", HandlerProfileAbout)
r.Get("/pictures", HandlerProfilePictures)
r.Get("/friends", HandlerProfileFriends)
})
// run server
http.ListenAndServe(":8080", router)
} */

type ResponseGetByIdEmployee struct{}
