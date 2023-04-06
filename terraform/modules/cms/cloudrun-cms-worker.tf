locals {
  reearth_cms_worker_secret = [
    "REEARTH_CMS_WORKER_DB",
  ]
}

resource "google_cloud_run_service" "reearth_cms_worker" {
  name                       = "reearth-cms-worker"
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
      container_concurrency = 1
      timeout_seconds       = 3600
      service_account_name  = google_service_account.reearth_cms_worker.email
      containers {
        # 初回作成時にreearth/reearthを指定すると、環境変数の設定不足で立ち上がらない。
        # そのため、一時的にサンプルアプリケーションでで作成し、セットアップ完了後にgcloudでdeployを行う。
        image = "gcr.io/cloudrun/hello"
        resources {
          limits = {
            cpu    = "8000m"
            memory = "32Gi"
          }
        }
        ports {
          container_port = 8080
          name           = "h2c"
        }

        dynamic "env" {
          for_each = { for i in local.reearth_cms_worker_secret : i => i }
          content {
            name = env.value
            value_from {
              secret_key_ref {
                name = google_secret_manager_secret.reearth_cms_worker[env.value].secret_id
                key  = "latest"
              }
            }
          }
        }
        env {
          name  = "GCS_BUCKET_NAME"
          value = google_storage_bucket.assets.name
        }
        env {
          name  = "REEARTH_CMS_WORKER_PUBSUB_TOPIC"
          value = "decompress"
        }
        env {
          name  = "REEARTH_CMS_WORKER_GCP_PROJECT"
          value = data.google_project.project.name
        }
        env {
          name  = "REEARTH_CMS_WORKER_DECOMPRESSION_NUM_WORKERS"
          value = "500"
        }
        env {
          name  = "REEARTH_CMS_WORKER_DECOMPRESSION_WORKQUEUE_DEPTH"
          value = "https://${var.auth0.domain}"
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
    google_secret_manager_secret_version.reearth_cms_worker_dummy
  ]
}

resource "google_secret_manager_secret" "reearth_cms_worker" {
  for_each  = toset(local.reearth_cms_worker_secret)
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
resource "google_secret_manager_secret_version" "reearth_cms_worker_dummy" {
  for_each = toset([
    "REEARTH_CMS_WORKER_DB",
  ])
  secret = google_secret_manager_secret.reearth_cms_worker[each.value].id

  secret_data = "dummy"
  lifecycle {
    ignore_changes = [
      secret_data
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "reearth_cmw_worker_noauth" {
  location    = var.gcp_region
  project     = var.gcp_project_name
  service     = google_cloud_run_service.reearth_cms_worker.name
  policy_data = data.google_iam_policy.noauth.policy_data
}
