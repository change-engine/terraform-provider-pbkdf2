resource "random_password" "example" {}

resource "pbkdf2_key" "example" {
  password = random_password.example.result
  # Output for https://github.com/change-engine/pbkdf-subtle
  format = "v1{{printf \"%06x\" .Iterations}}{{printf \"%s\" .Key}}"
}
