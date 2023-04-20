package main

import (
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"second/handlers"
	"second/repo"
	"second/server"
)

func main() {
	//db.InitDB()
	//userRepo := repo.NewUserMockRepo()
	userRepo := repo.NewUserPgRepo()
	srv := server.NewServer(handlers.NewHandlers(userRepo))
	go srv.Run()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to stop server")
	// ReadString will block until the delimiter is entered
	reader.ReadString('\n')
}
