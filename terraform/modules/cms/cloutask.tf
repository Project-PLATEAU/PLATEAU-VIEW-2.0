resource "google_cloud_tasks_queue" "cms_decompress" {
  name     = "decompress"
  location = var.gcp_region

  rate_limits {
    max_concurrent_dispatches = 1000
    max_dispatches_per_second = 500
  }

  retry_config {
    max_attempts  = 50
    max_backoff   = "3600s"
    max_doublings = 16
    min_backoff   = "10s"
  }

  stackdriver_logging_config {
    sampling_ratio = 1.0
  }
}

resource "google_pubsub_topic" "cms_webhook" {
  name = "cms-webhook"
}

resource "google_pubsub_subscription" "cms_webhook" {
  name  = "cms-webhook"
  topic = google_pubsub_topic.cms_webhook.name

  push_config {
    push_endpoint = "https://${local.worker_cms_domain}/api/webhook"
    oidc_token {
      service_account_email = google_service_account.cms_worker_m2m.email
    }
  }

  retry_policy {
    maximum_backoff = "600s"
    minimum_backoff = "10s"
  }

  expiration_policy {
    ttl = ""
  }
}

resource "google_pubsub_topic" "cms_decompress" {
  name = "decompress"
}

#TODO: subscriptionの設定を見直す
resource "google_pubsub_subscription" "cms_notify" {
  name  = "notify-cms"
  topic = google_pubsub_topic.cms_decompress.name

  push_config {
    push_endpoint = "https://${local.api_cms_domain}/api/notify"
    oidc_token {
      service_account_email = google_service_account.cms_worker_m2m.email
      audience              = "https://${local.api_cms_domain}"
    }
  }

  message_retention_duration = "604800s"
  retry_policy {
    maximum_backoff = "600s"
    minimum_backoff = "10s"
  }

  expiration_policy {
    ttl = ""
  }
}
