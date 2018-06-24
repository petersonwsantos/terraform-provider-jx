provider "jx" {
  name        = "my-jx-cluster"
  endpoint    = "https://127.0.0.1"
  certificate = "some generated cert..."
}

resource "jx_install" "teste" {
  name             = "teste"
  admin_password   = "1q2w3e4a"
  jx_provider      = "kubernetes"
  git_provider_url = "https://github.com"
  git_owner        = "opstricks"
  git_user         = "opstricks"
  git_token        = "f1cfa8bb880640e4e347ffed18ae7d88cb3de07db"
}

resource "jx_team" "team1" {
  name = "team1"
}

resource "jx_team" "team2" {
  name = "team2"
}

resource "jx_environment" "staging" {
  name               = "staging"
  promotion_strategy = "auto"
  order              = "100"
  namespace          = "jx-staging"
}

resource "jx_environment" "production" {
  name               = "production"
  promotion_strategy = "manual"
  order              = "200"
  namespace          = "jx-production"
}
