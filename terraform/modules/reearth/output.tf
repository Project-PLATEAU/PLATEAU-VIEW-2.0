
output "auth0_action_singup" {
  value = module.auth0.action_singup
}

output "plateauview_reearth_url" {
  value = local.reearth_domain
}

output "plateauview_reearth_api_url" {
  value = local.api_reearth_domain
}
