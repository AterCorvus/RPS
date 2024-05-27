package main

import (
	controllers "github.com/AterCorvus/RPS/Server/src/Controllers"
	api "github.com/AterCorvus/RPS/Server/src/api"
)

func main() {
	// Connect to SQLite
	controllers.InitSQLiteCon()

	// Start the server
	api.StartServer()
	/*
		userService := services.NewUserService()
		err := userService.RegisterUser("test", "test")
		if err != nil {
			fmt.Println(err)
		}
		result := userService.LoginUser("test", "test")
		println(result)
		result = userService.LogoutUser("test", "test")
		println(result)*/

	// Create a new RPC server
	/*server := rpc.NewServer()
	server.Register(con)

	// Listen for incoming connections
	listener, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := listener.Accept()
		go server.ServeConn(conn)
	}*/
}
