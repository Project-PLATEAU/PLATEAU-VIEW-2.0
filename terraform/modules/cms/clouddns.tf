data "google_dns_managed_zone" "cms" {
  name = var.dns_managed_zone_name
}

resource "google_dns_record_set" "api" {
  name = "${local.api_cms_domain}."
  type = "CNAME"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.cms.name
  rrdatas      = ["${local.cms_domain}."]
}

resource "google_dns_record_set" "plateauview_api" {
  name = "${local.api_domain}."
  type = "CNAME"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.cms.name
  rrdatas      = ["${local.cms_domain}."]
}

resource "google_dns_record_set" "assets" {
  name = "${local.assets_cms_domain}."
  type = "CNAME"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.cms.name
  rrdatas      = ["${local.cms_domain}."]
}

resource "google_dns_record_set" "app" {
  name = "${local.cms_domain}."
  type = "A"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.cms.name
  rrdatas      = [google_compute_global_address.cms_lb.address]
}

resource "google_dns_record_set" "worker" {
  name = "${local.worker_cms_domain}."
  type = "CNAME"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.cms.name
  rrdatas      = ["${local.cms_domain}."]
}