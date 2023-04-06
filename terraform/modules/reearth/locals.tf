locals {
  reearth_domain        = "reearth.${var.base_domain}"
  api_reearth_domain    = "api.${local.reearth_domain}"
  static_reearth_domain = "static.${local.reearth_domain}"
}