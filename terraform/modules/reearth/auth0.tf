module "auth0" {
  source = "../auth0"

  spa_name          = "plateau-reearth-spa"
  m2m_name          = "plateau-reearth-m2m"
  login_domain      = local.reearth_domain
  identifier_domain = local.api_reearth_domain
  signup_name       = "signup-reearth"
}