
resource "google_compute_target_http_proxy" "reearth" {
  name       = "reearth-common-http-targetproxy"
  proxy_bind = "false"
  url_map    = google_compute_url_map.reearth.id
}



resource "google_compute_target_https_proxy" "reearth" {
  name            = "reearth-common-https-targetproxy"
  url_map         = google_compute_url_map.reearth.id
  certificate_map = "//certificatemanager.googleapis.com/${google_certificate_manager_certificate_map.reearth.id}"
}

resource "google_certificate_manager_certificate_map" "reearth" {
  name = "reearth-cert-map"
}

resource "google_certificate_manager_certificate_map_entry" "reearth_primary" {
  name        = "reearth-cert-map-primary"
  description = "reearth wildcard"
  map         = google_certificate_manager_certificate_map.reearth.name

  certificates = [google_certificate_manager_certificate.reearth_wildcard.id]

  matcher = "PRIMARY"
}

resource "google_certificate_manager_certificate" "reearth_wildcard" {
  name        = "reearth-wildcard"
  description = "reearth wildcard cert"
  scope       = "DEFAULT"
  managed {
    domains = [
      google_certificate_manager_dns_authorization.reearth_wildcard.domain,
      "*.${google_certificate_manager_dns_authorization.reearth_wildcard.domain}"
    ]
    dns_authorizations = [
      google_certificate_manager_dns_authorization.reearth_wildcard.id,
    ]
  }
}

resource "google_certificate_manager_dns_authorization" "reearth_wildcard" {
  name        = "reearth-wildcard-dns-auth"
  description = "reearth wildcard dns auth"
  domain      = local.reearth_domain
}

resource "google_compute_global_address" "reearth_lb" {
  name = "reearth-common-lb"
}

resource "google_compute_global_forwarding_rule" "reearth_https" {
  name       = "reearth-common-https"
  target     = google_compute_target_https_proxy.reearth.self_link
  port_range = "443"
  ip_address = google_compute_global_address.reearth_lb.address

  depends_on = [google_compute_url_map.reearth]
}


resource "google_compute_global_forwarding_rule" "reearth_http" {
  name       = "reearth-common-http-redirect"
  target     = google_compute_target_http_proxy.reearth.self_link
  port_range = "80"
  ip_address = google_compute_global_address.reearth_lb.address

  depends_on = [google_compute_url_map.reearth]
}

resource "google_compute_url_map" "reearth_redirect" {
  name = "reearth-https-redirect"
  default_url_redirect {
    https_redirect         = "true"
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    strip_query            = "false"
  }

  description = "HTTP to HTTPS redirect forwarding rule"
}

resource "google_compute_url_map" "reearth" {
  name        = "reearth-common-urlmap"
  description = "reearth common urlmap"

  default_service = google_compute_backend_service.reearth_api.self_link

  host_rule {
    hosts = [
      local.api_reearth_domain,
      "*.${local.reearth_domain}",
      local.reearth_domain,
    ]
    path_matcher = "path-matcher-3"
  }

  path_matcher {
    default_service = google_compute_backend_service.reearth_api.self_link
    name            = "path-matcher-3"
  }
}

resource "google_compute_backend_bucket" "static_backend" {
  name        = "reearth-static-backend"
  bucket_name = google_storage_bucket.static.name
  enable_cdn  = true
  cdn_policy {
    signed_url_cache_max_age_sec = 7200
  }
}


resource "google_compute_region_network_endpoint_group" "reearth_api" {
  name                  = "reearth-api-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "asia-northeast1"
  cloud_run {
    service = google_cloud_run_service.reearth_api.name
  }
}

resource "google_compute_backend_service" "reearth_api" {
  affinity_cookie_ttl_sec = "0"
  enable_cdn              = true
  cdn_policy {
    signed_url_cache_max_age_sec = 7200
  }
  backend {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "0"
    group                        = google_compute_region_network_endpoint_group.reearth_api.id
    max_connections              = "0"
    max_connections_per_endpoint = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_endpoint        = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  description                     = "reearth-api-neg"
  load_balancing_scheme           = "EXTERNAL"

  log_config {
    enable      = "true"
    sample_rate = "1"
  }

  name             = "reearth-api-backend"
  port_name        = "http"
  protocol         = "HTTPS"
  session_affinity = "NONE"
  timeout_sec      = "30"
}