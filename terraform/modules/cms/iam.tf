resource "google_service_account" "reearth_cms_api" {
  account_id   = "reearth-cms-api"
  display_name = "Service Account for reearth cms api"
}

resource "google_project_iam_member" "reearth_cms_api" {
  role    = google_project_iam_custom_role.reearth_cms_api.id
  project = data.google_project.project.project_id

  member = "serviceAccount:${google_service_account.reearth_cms_api.email}"
}

resource "google_project_iam_custom_role" "reearth_cms_api" {
  role_id     = "reearthcmsapi"
  project     = data.google_project.project.project_id
  title       = "reearth-cms-api"
  description = "iam role for reearth-cms-api"
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
    "cloudtasks.tasks.create",
  ]
}



resource "google_service_account" "reearth_cms_worker" {
  account_id   = "reearth-cms-worker"
  display_name = "Service Account for reearth cms worker"
}

resource "google_project_iam_member" "reearth_cms_worker" {
  role    = google_project_iam_custom_role.reearth_cms_worker.id
  project = data.google_project.project.project_id

  member = "serviceAccount:${google_service_account.reearth_cms_worker.email}"
}

resource "google_project_iam_custom_role" "reearth_cms_worker" {
  role_id     = "reearthcmsworker"
  project     = data.google_project.project.project_id
  title       = "reearth-cms-worker"
  description = "iam role for reearth-cms-worker"
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

resource "google_service_account" "plateauview_api" {
  account_id   = "plateauview-api"
  display_name = "Service Account for plateauview api"
}

resource "google_project_iam_member" "plateauview_api" {
  role    = google_project_iam_custom_role.plateauview_api.id
  project = data.google_project.project.project_id

  member = "serviceAccount:${google_service_account.plateauview_api.email}"
}

resource "google_project_iam_custom_role" "plateauview_api" {
  role_id     = "plateauviewapi"
  project     = data.google_project.project.project_id
  title       = "plateauview-api"
  description = "iam role for plateauview api"
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


resource "google_service_account" "cms_worker_m2m" {
  account_id   = "cms-worker-m2m"
  display_name = "Service Account for cms worker m2m"
}

resource "google_project_iam_member" "cms_worker_m2m" {
  role    = google_project_iam_custom_role.cms_worker_m2m.id
  project = data.google_project.project.project_id

  member = "serviceAccount:${google_service_account.cms_worker_m2m.email}"
}

resource "google_project_iam_custom_role" "cms_worker_m2m" {
  role_id     = "cmsworkerm2m"
  project     = data.google_project.project.project_id
  title       = "cmsworkerm2m"
  description = "iam role for cmsworkerm2m"
  stage       = "GA"
  permissions = [
    "iam.serviceAccounts.actAs",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.getAccessToken",
    "iam.serviceAccounts.getOpenIdToken",
    "iam.serviceAccounts.implicitDelegation",
    "iam.serviceAccounts.list",
    "iam.serviceAccounts.signBlob",
    "iam.serviceAccounts.signJwt",
    "resourcemanager.projects.get",
    "run.jobs.run",
  ]
}