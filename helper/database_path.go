package helper


type Path string

var DatabasePath = struct {
    DATABASE         Path
    USERS		 Path
}{
    DATABASE: "backgroundservice",
    USERS: "users",
}