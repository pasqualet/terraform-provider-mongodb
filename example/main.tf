provider "mongo" {
    url = "mongodb://localhost:27017/test"
}

resource "mongo_user" "user" {
    database = "test"
    username = "user"
    password = "pass"
    roles = ["read", "dbAdmin", "userAdmin"]
}
