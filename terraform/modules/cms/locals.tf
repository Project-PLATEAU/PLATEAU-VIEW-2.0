locals {
  cms_domain        = "cms.${var.base_domain}"
  api_domain        = "api.${var.base_domain}"
  api_cms_domain    = "api.${local.cms_domain}"
  assets_cms_domain = "assets.${local.cms_domain}"
  worker_cms_domain = "worker.${local.cms_domain}"
}