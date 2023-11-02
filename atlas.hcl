data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./internal/migrate"
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  url = "postgres://hostal_app_user:hostal_app_pass@localhost:5433/hostal_app_db?search_path=public&sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}