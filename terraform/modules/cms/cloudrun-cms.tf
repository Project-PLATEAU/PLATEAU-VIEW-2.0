locals {
  reearth_cms_api_secret = [
    "REEARTH_CMS_DB",
    "REEARTH_CMS_AUTH0_CLIENTSECRET",
  ]
}

resource "google_cloud_run_service" "reearth_cms_api" {
  name                       = "reearth-cms-api"
  location                   = var.gcp_region
  autogenerate_revision_name = true
  metadata {
    annotations = {
      "run.googleapis.com/launch-stage"   = "BETA"
      "run.googleapis.com/ingress"        = "all"
      "run.googleapis.com/ingress-status" = "all"
    }
  }

  template {
    spec {
      service_account_name  = google_service_account.reearth_cms_api.email
      container_concurrency = 80
      timeout_seconds       = 1800
      containers {
        # 初回作成時にreearth/reearthを指定すると、環境変数の設定不足で立ち上がらない。
        # そのため、一時的にサンプルアプリケーションでで作成し、セットアップ完了後にgcloudでdeployを行う。
        image = "gcr.io/cloudrun/hello"
        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
        }
        ports {
          container_port = 8080
          name           = "h2c"
        }

        dynamic "env" {
          for_each = { for i in local.reearth_cms_api_secret : i => i }
          content {
            name = env.value
            value_from {
              secret_key_ref {
                name = google_secret_manager_secret.reearth_cms_api[env.value].secret_id
                key  = "latest"
              }
            }
          }
        }
        env {
          name  = "REEARTH_CMS_AUTH0_DOMAIN"
          value = "https://${var.auth0.domain}"
        }
        env {
          name  = "REEARTH_CMS_GCS_BUCKETNAME"
          value = google_storage_bucket.assets.name
        }
        env {
          name  = "REEARTH_CMS_ASSETBASEURL"
          value = "https://${local.assets_cms_domain}"
        }
        env {
          name  = "REEARTH_GCS_PUBLICATIONCACHECONTROL"
          value = "no-store"
        }
        env {
          name  = "REEARTH_CMS_ORIGINS"
          value = "https://${local.cms_domain}"
        }
        env {
          name  = "REEARTH_CMS_GRAPHQL_COMPLEXITYLIMIT"
          value = "6000"
        }
        env {
          name  = "REEARTH_CMS_HOST"
          value = "https://${local.api_cms_domain}"
        }
        env {
          name  = "REEARTH_HOST_WEB"
          value = "https://${local.cms_domain}"
        }
        env {
          name  = "REEARTH_CMS_AUTH0_AUDIENCE"
          value = "https://${local.api_cms_domain}"
        }
        env {
          name  = "REEARTH_CMS_AUTHM2M_ISS"
          value = "https://accounts.google.com"
        }
        env {
          name  = "REEARTH_CMS_AUTHM2M_AUD"
          value = "https://${local.api_cms_domain}"
        }
        env {
          name  = "REEARTH_CMS_AUTHM2M_EMAIL"
          value = google_service_account.cms_worker_m2m.email
        }
        env {
          name  = "REEARTH_CMS_TASK_GCPPROJECT"
          value = data.google_project.project.name
        }
        env {
          name  = "REEARTH_CMS_TASK_GCPREGION"
          value = "asia-northeast1"
        }
        env {
          name  = "REEARTH_CMS_TASK_QUEUENAME"
          value = "decompress"
        }
        env {
          name  = "REEARTH_CMS_TASK_SUBSCRIBERURL"
          value = "https://${local.worker_cms_domain}/api/decompress"
        }
        env {
          name  = "REEARTH_CMS_TASK_TOPIC"
          value = "cms-webhook"
        }
        env {
          name  = "GOOGLE_CLOUD_PROJECT"
          value = var.gcp_project_name
        }
        env {
          name  = "REEARTH_CMS_AUTH0_WEBCLIENTID"
          value = module.auth0.auth0_client_spa.client_id
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"         = "100"
        "run.googleapis.com/execution-environment" = "gen2"
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
  lifecycle {
    ignore_changes = [
      metadata[0].annotations,
      template[0].spec[0].containers[0].image,
      template[0].metadata[0].annotations["run.googleapis.com/client-name"],
      template[0].metadata[0].annotations["client.knative.dev/user-image"],
      template[0].metadata[0].annotations["run.googleapis.com/client-version"]
    ]
  }
  depends_on = [
    google_secret_manager_secret_version.reearth_cms_api_dummy
  ]
}

resource "google_secret_manager_secret" "reearth_cms_api" {
  for_each  = toset(local.reearth_cms_api_secret)
  secret_id = "reearth-cms-${each.value}"
  labels = {
    label = "reearth-cms"
  }
  replication {
    user_managed {
      replicas {
        location = "asia-northeast2"
      }
    }
  }
}

//MEMO: secret_managerに値が入っていないとcloudrunが起動エラーになるので、
//      あとから手動で値を入れるものに関しては先にdummyの値を入れておく
resource "google_secret_manager_secret_version" "reearth_cms_api_dummy" {
  for_each = toset([
    "REEARTH_CMS_DB",
  ])
  secret = google_secret_manager_secret.reearth_cms_api[each.value].id

  secret_data = "dummy"
  lifecycle {
    ignore_changes = [
      secret_data
    ]
  }
}

resource "google_secret_manager_secret_version" "reearth_cms_api_auth0_clientsecret" {
  secret      = google_secret_manager_secret.reearth_cms_api["REEARTH_CMS_AUTH0_CLIENTSECRET"].id
  secret_data = module.auth0.auth0_client_m2m.client_secret
}

resource "google_cloud_run_service_iam_policy" "reearth_cms_noauth" {
  location    = var.gcp_region
  project     = var.gcp_project_name
  service     = google_cloud_run_service.reearth_cms_api.name
  policy_data = data.google_iam_policy.noauth.policy_data
}
