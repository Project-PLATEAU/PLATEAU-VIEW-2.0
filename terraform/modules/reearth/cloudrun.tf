locals {
  reearth_api_secret = [
    "REEARTH_DB",
    "REEARTH_AUTH0_CLIENTID",
    "REEARTH_AUTH0_CLIENTSECRET",
    "REEARTH_SIGNUPSECRET",
    "REEARTH_MARKETPLACE_SECRET",
  ]
}

resource "google_cloud_run_service" "reearth_api" {
  name                       = "reearth-api"
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
      service_account_name = google_service_account.reearth_api.email
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
          for_each = { for i in local.reearth_api_secret : i => i }
          content {
            name = env.value
            value_from {
              secret_key_ref {
                name = google_secret_manager_secret.reearth_api[env.value].secret_id
                key  = "latest"
              }
            }
          }
        }
        env {
          name  = "REEARTH_AUTH0_DOMAIN"
          value = "https://${var.auth0.domain}"
        }
        env {
          name  = "REEARTH_GCS_BUCKETNAME"
          value = google_storage_bucket.static.name
        }
        env {
          name  = "REEARTH_ASSETBASEURL"
          value = "https://${local.static_reearth_domain}"
        }
        env {
          name  = "REEARTH_TRACERSAMPLE"
          value = ".0"
        }
        env {
          name  = "REEARTH_GCS_PUBLICATIONCACHECONTROL"
          value = "no-store"
        }
        env {
          name  = "REEARTH_ORIGINS"
          value = "https://${local.reearth_domain}"
        }
        env {
          name  = "REEARTH_AUTHSRV_DISABLED"
          value = "true"
        }
        env {
          name  = "REEARTH_HOST"
          value = "https://${local.api_reearth_domain}"
        }
        env {
          name  = "REEARTH_HOST_WEB"
          value = "https://${local.reearth_domain}"
        }
        env {
          name  = "REEARTH_AUTH0_AUDIENCE"
          value = "https://${local.api_reearth_domain}"
        }
        env {
          name  = "REEARTH_MARKETPLACE_ENDPOINT"
          value = "https://api.marketplace.reearth.io"
        }
        env {
          name  = "GOOGLE_CLOUD_PROJECT"
          value = var.gcp_project_name
        }
        env {
          name  = "REEARTH_PUBLISHED_HOST"
          value = "{}.${local.reearth_domain}"
        }
        env {
          name  = "REEARTH_AUTH0_WEBCLIENTID"
          value = module.auth0.auth0_client_spa.client_id
        }
        env {
          name  = "REEARTH_WEB"
          value = var.cesium_ion_access_token != "" ? "cesiumIonAccessToken:${var.cesium_ion_access_token}" : ""
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
    google_secret_manager_secret_version.reearth_api_dummy
  ]
}

resource "google_secret_manager_secret" "reearth_api" {
  for_each  = toset(local.reearth_api_secret)
  secret_id = "reearth-api-${each.value}"
  labels = {
    label = "reearth-api"
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
resource "google_secret_manager_secret_version" "reearth_api_dummy" {
  for_each = toset([
    "REEARTH_DB",
    "REEARTH_MARKETPLACE_SECRET",
  ])
  secret = google_secret_manager_secret.reearth_api[each.value].id

  secret_data = "dummy"
  lifecycle {
    ignore_changes = [
      secret_data
    ]
  }
}

resource "google_secret_manager_secret_version" "reearth_api_auth0_clientsecret" {
  secret      = google_secret_manager_secret.reearth_api["REEARTH_AUTH0_CLIENTSECRET"].id
  secret_data = module.auth0.auth0_client_m2m.client_secret
}

resource "google_secret_manager_secret_version" "reearth_api_auth0_clientid" {
  secret      = google_secret_manager_secret.reearth_api["REEARTH_AUTH0_CLIENTID"].id
  secret_data = module.auth0.auth0_client_m2m.client_id
}

resource "google_secret_manager_secret_version" "reearth_api_signupsecret" {
  secret      = google_secret_manager_secret.reearth_api["REEARTH_SIGNUPSECRET"].id
  secret_data = module.auth0.action_secret.result
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "reearth_noauth" {
  location    = var.gcp_region
  project     = var.gcp_project_name
  service     = google_cloud_run_service.reearth_api.name
  policy_data = data.google_iam_policy.noauth.policy_data
}
