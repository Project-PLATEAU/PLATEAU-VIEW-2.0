resource "auth0_client" "spa" {
  name = var.spa_name

  app_type = "spa"

  initiate_login_uri = "https://${var.login_domain}"

  token_endpoint_auth_method = "none"

  allowed_logout_urls = [
    "https://${var.login_domain}"
  ]
  allowed_origins = [
    "https://${var.login_domain}"
  ]
  callbacks = [
    "https://${var.login_domain}"
  ]
  web_origins = [
    "https://${var.login_domain}"
  ]
}

resource "auth0_client" "m2m" {
  name = var.m2m_name

  app_type           = "non_interactive"
  initiate_login_uri = "https://${var.login_domain}"

  allowed_logout_urls = [
    "https://${var.login_domain}"
  ]
  allowed_origins = [
    "https://${var.login_domain}"
  ]
  callbacks = [
    "https://${var.login_domain}"
  ]
  web_origins = [
    "https://${var.login_domain}"
  ]
}

resource "auth0_resource_server" "signup" {
  name       = var.spa_name
  identifier = "https://${var.identifier_domain}"
}

resource "auth0_action" "singup" {
  name    = var.signup_name
  runtime = "node16"
  deploy  = true
  code    = file("${path.module}/source/signup-reearth.js")

  supported_triggers {
    id      = "post-user-registration"
    version = "v2"
  }

  dependencies {
    name    = "axios"
    version = "1.3.2"
  }

  secrets {
    name  = "secret"
    value = random_string.action_secret.result
  }

  secrets {
    name  = "api"
    value = "https://${var.identifier_domain}/api/signup"
  }
}

resource "random_string" "action_secret" {
  length  = 32
  special = false
}
