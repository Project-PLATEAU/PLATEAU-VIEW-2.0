locals {
  plateauview_api_secret = [
    "REEARTH_PLATEAUVIEW_SECRET",
    "REEARTH_PLATEAUVIEW_CMS_WEBHOOK_SECRET",
    "REEARTH_PLATEAUVIEW_CMS_TOKEN",
    "REEARTH_PLATEAUVIEW_FME_TOKEN",
    "REEARTH_PLATEAUVIEW_CKAN_TOKEN",
    "REEARTH_PLATEAUVIEW_SENDGRID_APIKEY",
  ]
  plateauview_api_ramdom = [
    "REEARTH_PLATEAUVIEW_CMS_WEBHOOK_SECRET",
    "REEARTH_PLATEAUVIEW_SECRET",
    "REEARTH_PLATEAUVIEW_SIDEBAR_TOKEN",
    "REEARTH_PLATEAUVIEW_SDK_TOKEN",
  ]
}

resource "google_cloud_run_service" "plateauview_api" {
  name                       = "plateauview-api"
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
      service_account_name = google_service_account.plateauview_api.email
      timeout_seconds      = 3600
      containers {
        # 初回作成時にreearth/reearthを指定すると、環境変数の設定不足で立ち上がらない。
        # そのため、一時的にサンプルアプリケーションでで作成し、セットアップ完了後にgcloudでdeployを行う。
        image = "gcr.io/cloudrun/hello"
        resources {
          limits = {
            cpu    = "1000m"
            memory = "1Gi"
          }
        }
        ports {
          container_port = 8080
          name           = "h2c"
        }

        dynamic "env" {
          for_each = { for i in local.plateauview_api_secret : i => i }
          content {
            name = env.value
            value_from {
              secret_key_ref {
                name = google_secret_manager_secret.plateauview_api[env.value].secret_id
                key  = "latest"
              }
            }
          }
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_FME_BASEURL"
          value = var.plateauview.fme_baseurl
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_CMS_BASEURL"
          value = "https://${local.api_cms_domain}"
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_CKAN_BASEURL"
          value = var.plateauview.ckan_base_url
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_CKAN_ORG"
          value = var.plateauview.ckan_org
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_CMS_SYSTEMPROJECT"
          value = var.plateauview.cms_system_project
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_OPINION_TO"
          value = var.plateauview.option_to
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_OPINION_FROM"
          value = var.plateauview.option_from
        }
        env {
          name  = "REEARTH_PLATEAUVIEW_CMS_PLATEAUPROJECT"
          value = var.plateauview.cms_plateau_project
        }
        env {
          name  = "GOOGLE_CLOUD_PROJECT"
          value = var.gcp_project_name
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
    google_secret_manager_secret_version.plateauview_api_dummy
  ]
}

resource "random_string" "plateauview_env" {
  length  = 32
  special = false
}


resource "google_secret_manager_secret" "plateauview_api" {
  for_each  = toset(local.plateauview_api_secret)
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
resource "google_secret_manager_secret_version" "plateauview_api_dummy" {
  for_each = toset([
    "REEARTH_PLATEAUVIEW_CMS_TOKEN",
    "REEARTH_PLATEAUVIEW_FME_TOKEN",
    "REEARTH_PLATEAUVIEW_CKAN_TOKEN",
    "REEARTH_PLATEAUVIEW_SENDGRID_APIKEY",
  ])
  secret = google_secret_manager_secret.plateauview_api[each.value].id

  secret_data = "dummy"
  lifecycle {
    ignore_changes = [
      secret_data
    ]
  }
}
resource "random_string" "plateauview_api_env" {
  for_each = toset(local.plateauview_api_ramdom)
  length   = 32
  special  = false
}
resource "google_secret_manager_secret_version" "plateauview_api_secret" {
  secret      = google_secret_manager_secret.plateauview_api["REEARTH_PLATEAUVIEW_SECRET"].id
  secret_data = random_string.plateauview_api_env["REEARTH_PLATEAUVIEW_SECRET"].result
}

resource "google_secret_manager_secret_version" "plateauview_api_webhook_secret" {
  secret      = google_secret_manager_secret.plateauview_api["REEARTH_PLATEAUVIEW_CMS_WEBHOOK_SECRET"].id
  secret_data = random_string.plateauview_api_env["REEARTH_PLATEAUVIEW_CMS_WEBHOOK_SECRET"].result
}

resource "google_cloud_run_service_iam_policy" "plateauview_noauth" {
  location    = var.gcp_region
  project     = var.gcp_project_name
  service     = google_cloud_run_service.plateauview_api.name
  policy_data = data.google_iam_policy.noauth.policy_data
}
