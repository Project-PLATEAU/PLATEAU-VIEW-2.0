provider "google" {
  project = var.gcp_project_name
  region  = var.gcp_region
}

provider "google-beta" {
  project = var.gcp_project_name
  region  = var.gcp_region
}

provider "auth0" {
  domain    = var.auth0_provider.domain
  client_id = var.auth0_provider.client_id
}