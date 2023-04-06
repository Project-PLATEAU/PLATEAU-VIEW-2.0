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

variable "cesium_ion_access_token" {
  type    = string
  default = ""
}