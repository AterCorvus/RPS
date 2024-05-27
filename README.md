# RPS
Simple multiplayer game of rock paper scissors betting

To Run Server:
    because go-sqlite is a CGO enabled package, you are required to set the environment variable CGO_ENABLED=1 and have a gcc compiler present within your path.
    
    CGO_ENABLED=1 go run main.go