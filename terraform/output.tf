output "plateauview_cms_webhook_url" {
  value = "https://api.${var.base_domain}/webhook"
}

output "plateauview_cms_webhook_secret" {
  value = module.reearth-cms.plateauview_cms_webhook_secret
}

output "plateauview_sdk_token" {
  value = module.reearth-cms.plateauview_sdk_token
}

output "plateauview_sidebar_token" {
  value = module.reearth-cms.plateauview_sidebar_token
}

output "plateauview_cms_url" {
  value = module.reearth-cms.plateauview_cms_url
}

output "plateauview_reearth_url" {
  value = module.reearth-api.plateauview_reearth_url
}

output "plateauview_sidecar_url" {
  value = module.reearth-cms.plateauview_api_url
}
