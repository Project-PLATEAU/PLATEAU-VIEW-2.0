variable "base_domain" {
  type        = string
  default     = null
  description = "reearthで利用するドメインを指定してください"
}

variable "gcp_project_name" {
  type        = string
  default     = null
  description = "GCPのプロジェクト名を指定してください"
}

variable "gcp_region" {
  type        = string
  default     = "asia-northeast1"
  description = "GCPで使用するregionを指定してください"
}

variable "service_prefix" {
  type        = string
  default     = null
  description = "特定のリソースに付与するためのprefixを指定してください"
}

variable "dns_managed_zone_name" {
  type        = string
  default     = null
  description = "CloudDNSのゾーン名を指定してください"
}

variable "auth0" {
  type = object({
    domain = string
  })
  default = {
    domain = null
  }
  description = "auth0に関する設定を指定してください"
}

variable "reearth_version" {
  type    = string
  default = "0.14.1"
}

variable "plateauview" {
  type = object({
    fme_baseurl            = string
    fme_skip_quality_check = string
    ckan_org               = string
    ckan_base_url          = string
    cms_system_project     = string
    option_to              = string
    option_from            = string
    cms_plateau_project    = string
  })
}