
resource "google_compute_target_http_proxy" "cms" {
  name       = "${var.service_prefix}-http-targetproxy"
  proxy_bind = "false"
  url_map    = google_compute_url_map.cms.id
}

resource "google_compute_target_https_proxy" "cms" {
  name             = "${var.service_prefix}-common-https-targetproxy"
  url_map          = google_compute_url_map.cms.id
  ssl_certificates = [google_compute_managed_ssl_certificate.common.id]
}

resource "google_compute_managed_ssl_certificate" "common" {
  name = "${var.service_prefix}-common-cert"

  managed {
    domains = [
      local.cms_domain,
      local.api_domain,
      local.api_cms_domain,
      local.assets_cms_domain,
      local.worker_cms_domain
    ]
  }
}

resource "google_compute_global_address" "cms_lb" {
  name = "${var.service_prefix}-lb"
}

resource "google_compute_global_forwarding_rule" "cms_https" {
  name       = "${var.service_prefix}-https"
  target     = google_compute_target_https_proxy.cms.self_link
  port_range = "443"
  ip_address = google_compute_global_address.cms_lb.address

  depends_on = [google_compute_url_map.cms]
}

resource "google_compute_global_forwarding_rule" "reearth_http" {
  name       = "${var.service_prefix}-http-redirect"
  target     = google_compute_target_http_proxy.cms.self_link
  port_range = "80"
  ip_address = google_compute_global_address.cms_lb.address

  depends_on = [google_compute_url_map.cms]
}

resource "google_compute_url_map" "cms_redirect" {
  name = "${var.service_prefix}-https-redirect"
  default_url_redirect {
    https_redirect         = "true"
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    strip_query            = "false"
  }

  description = "HTTP to HTTPS redirect forwarding rule"
}

resource "google_compute_url_map" "cms" {
  name        = "cms-common-urlmap"
  description = "cms common urlmap"

  default_service = google_compute_backend_service.cms_api.self_link

  host_rule {
    hosts = [
      local.cms_domain,
      local.api_cms_domain,
    ]
    path_matcher = "path-matcher-1"
  }

  path_matcher {
    default_service = google_compute_backend_service.cms_api.self_link
    name            = "path-matcher-1"
  }

  host_rule {
    hosts = [
      local.assets_cms_domain,
    ]
    path_matcher = "path-matcher-2"
  }

  path_matcher {
    default_service = google_compute_backend_bucket.assets_backend.self_link
    name            = "path-matcher-2"
  }

  host_rule {
    hosts = [
      local.api_domain,
    ]
    path_matcher = "path-matcher-3"
  }

  path_matcher {
    default_service = google_compute_backend_service.plateauview_api.self_link
    name            = "path-matcher-3"
  }

  host_rule {
    hosts = [
      local.worker_cms_domain,
    ]
    path_matcher = "path-matcher-4"
  }

  path_matcher {
    default_service = google_compute_backend_service.cms_worker.self_link
    name            = "path-matcher-4"
  }

}

resource "google_compute_backend_bucket" "assets_backend" {
  name        = "${var.service_prefix}-assets-backend"
  bucket_name = google_storage_bucket.assets.name
  enable_cdn  = true
  cdn_policy {
    signed_url_cache_max_age_sec = 7200
  }
}


resource "google_compute_region_network_endpoint_group" "cms_api" {
  name                  = "${var.service_prefix}-cms-api-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "asia-northeast1"
  cloud_run {
    service = google_cloud_run_service.reearth_cms_api.name
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_compute_backend_service" "cms_api" {
  affinity_cookie_ttl_sec = "0"

  enable_cdn = true
  cdn_policy {
    signed_url_cache_max_age_sec = 7200
  }

  backend {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "0"
    group                        = google_compute_region_network_endpoint_group.cms_api.id
    max_connections              = "0"
    max_connections_per_endpoint = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_endpoint        = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  description                     = "cms-api-neg"
  load_balancing_scheme           = "EXTERNAL"

  log_config {
    enable      = "true"
    sample_rate = "1"
  }

  name             = "cms-api-backend"
  port_name        = "http"
  protocol         = "HTTPS"
  session_affinity = "NONE"
  timeout_sec      = "30"

  lifecycle {
    create_before_destroy = true
  }

}


resource "google_compute_region_network_endpoint_group" "plateauview_api" {
  name                  = "${var.service_prefix}-plateauview-api-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "asia-northeast1"
  cloud_run {
    service = google_cloud_run_service.plateauview_api.name
  }
}

resource "google_compute_backend_service" "plateauview_api" {
  affinity_cookie_ttl_sec = "0"
  enable_cdn              = true
  cdn_policy {
    signed_url_cache_max_age_sec = 7200
  }

  backend {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "0"
    group                        = google_compute_region_network_endpoint_group.plateauview_api.id
    max_connections              = "0"
    max_connections_per_endpoint = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_endpoint        = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  description                     = "plateauview-api-neg"
  load_balancing_scheme           = "EXTERNAL"

  log_config {
    enable      = "true"
    sample_rate = "1"
  }

  name             = "plateauview-api-backend"
  port_name        = "http"
  protocol         = "HTTPS"
  session_affinity = "NONE"
  timeout_sec      = "30"
}

resource "google_compute_region_network_endpoint_group" "cms_worker" {
  name                  = "${var.service_prefix}-cms-worker-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "asia-northeast1"
  cloud_run {
    service = google_cloud_run_service.reearth_cms_worker.name
  }
}

resource "google_compute_backend_service" "cms_worker" {
  affinity_cookie_ttl_sec = "0"
  backend {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "0"
    group                        = google_compute_region_network_endpoint_group.cms_worker.id
    max_connections              = "0"
    max_connections_per_endpoint = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_endpoint        = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  description                     = "cms-worker-neg"
  enable_cdn                      = "false"
  load_balancing_scheme           = "EXTERNAL"

  log_config {
    enable      = "true"
    sample_rate = "1"
  }

  name             = "cms-worker-backend"
  port_name        = "http"
  protocol         = "HTTPS"
  session_affinity = "NONE"
  timeout_sec      = "30"
}