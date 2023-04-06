data "google_dns_managed_zone" "reearth" {
  name = var.dns_managed_zone_name
}

resource "google_dns_record_set" "api" {
  name = "api.reearth.${data.google_dns_managed_zone.reearth.dns_name}"
  type = "A"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.reearth.name
  rrdatas      = [google_compute_global_address.reearth_lb.address]
}

resource "google_dns_record_set" "static" {
  name = "static.reearth.${data.google_dns_managed_zone.reearth.dns_name}"
  type = "A"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.reearth.name
  rrdatas      = [google_compute_global_address.reearth_lb.address]
}

resource "google_dns_record_set" "app" {
  name = "reearth.${data.google_dns_managed_zone.reearth.dns_name}"
  type = "A"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.reearth.name
  rrdatas      = [google_compute_global_address.reearth_lb.address]
}

resource "google_dns_record_set" "published" {
  name = "*.reearth.${data.google_dns_managed_zone.reearth.dns_name}"
  type = "A"
  ttl  = 60

  managed_zone = data.google_dns_managed_zone.reearth.name
  rrdatas      = [google_compute_global_address.reearth_lb.address]
}


resource "google_dns_record_set" "reearth_cert_authorization" {
  name = google_certificate_manager_dns_authorization.reearth_wildcard.dns_resource_record.0.name
  type = google_certificate_manager_dns_authorization.reearth_wildcard.dns_resource_record.0.type
  ttl  = 10

  managed_zone = data.google_dns_managed_zone.reearth.name
  rrdatas = [
    google_certificate_manager_dns_authorization.reearth_wildcard.dns_resource_record.0.data,
    # google_certificate_manager_dns_authorization.reearth.dns_resource_record.0.data
  ]

}