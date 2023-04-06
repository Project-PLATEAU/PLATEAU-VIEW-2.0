module "auth0" {
  source = "../auth0"

  spa_name          = "plateau-cms-spa"
  m2m_name          = "plateau-cms-m2m"
  login_domain      = local.cms_domain
  identifier_domain = local.api_cms_domain
  signup_name       = "signup-backend"
}