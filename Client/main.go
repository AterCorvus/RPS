package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	authentification "github.com/AterCorvus/RPS/CLient/src/Authentification"
	ui "github.com/AterCorvus/RPS/CLient/src/UI"
)

func logout() {
    fmt.Println("Logout")
    authentification.Logout()
}

func main() {
    // Logout before exit
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigs
        logout()
        os.Exit(0)
    }()

    ui.LoginOrRegisterUser()
}