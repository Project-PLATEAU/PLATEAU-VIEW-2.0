terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.50"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 4.50"
    }
    random = {
      source = "hashicorp/random"
    }
    auth0 = {
      source  = "auth0/auth0"
      version = "0.43.0"
    }
  }
  required_version = ">= v1.3.7"

  backend "gcs" {
    bucket = "plateau-test2-terraform-tfstate"
  }
}
