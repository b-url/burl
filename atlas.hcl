env "local" {
  src = "file://schemas/schema.hcl"
  url = "postgres://develop:develop_secret@database:5432/develop?sslmode=disable"
  dev = "postgres://develop:develop_secret@postgres:5432/develop?sslmode=disable"
}
