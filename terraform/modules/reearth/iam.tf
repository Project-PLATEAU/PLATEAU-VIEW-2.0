resource "google_service_account" "reearth_api" {
  account_id   = "reearth-api"
  display_name = "Service Account for reearth api"
}

resource "google_project_iam_member" "reearth_api" {
  role    = google_project_iam_custom_role.reearth_api.id
  project = data.google_project.project.project_id

  member = "serviceAccount:${google_service_account.reearth_api.email}"
}

resource "google_project_iam_custom_role" "reearth_api" {
  role_id     = "reearthapi"
  project     = data.google_project.project.project_id
  title       = "reearth-api"
  description = "iam role for reearth-api"
  stage       = "GA"
  permissions = [
    "cloudprofiler.profiles.create",
    "cloudprofiler.profiles.update",
    "pubsub.topics.publish",
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",
    "storage.objects.update",
    "secretmanager.versions.access",
  ]
}