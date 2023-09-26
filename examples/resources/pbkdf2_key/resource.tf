resource "random_password" "example" {}

resource "pbkdf2_key" "example" {
  password = random_password.example.result
}
