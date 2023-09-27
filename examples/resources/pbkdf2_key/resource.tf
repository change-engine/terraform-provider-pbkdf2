resource "random_password" "example" {}

resource "pbkdf2_key" "example" {
  password = random_password.example.result
  # Output for https://github.com/change-engine/pbkdf-subtle
  format = "v01{{bin 3 .Iterations}}{{printf \"%s\" .Key}}"
}
