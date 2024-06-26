provider "google" {
  project = var.gcp_project_id
  region  = var.cloudrun_location
}

terraform {
  backend "gcs" {
    bucket = "bucket-capstone-425808"
    prefix = "terraform/state"
  }
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_v2_service.production.location
  project  = var.gcp_project_id
  service  = google_cloud_run_v2_service.production.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

resource "google_project_service" "service_usage" {
  project = var.gcp_project_id
  service = "serviceusage.googleapis.com"
}

resource "google_project_service" "cloud_run_api" {
  service = "run.googleapis.com"
}

resource "google_storage_bucket_iam_policy" "editor" {
  bucket      = google_storage_bucket.bucket.name
  policy_data = data.google_iam_policy.viewer.policy_data
}


resource "google_cloud_run_v2_service" "production" {
  name         = var.cloudrun_name
  location     = var.cloudrun_location
  launch_stage = "BETA"
  ingress      = "INGRESS_TRAFFIC_ALL"


  template {
    containers {
      image = format("%s@%s", var.image, data.docker_registry_image.get_metadata.sha256_digest)
      ports {
        container_port = 8080
      }

      volume_mounts {
        name       = "bucket"
        mount_path = "/models"
      }

      resources {
        limits = {
          cpu    = "1000m"
          memory = "2Gi"
        }

        cpu_idle = true

      }
    }

    volumes {
      name = "bucket"
      gcs {
        bucket    = google_storage_bucket.bucket.name
        read_only = false
      }
    }
  }

  depends_on = [
    google_project_service.cloud_run_api
  ]
}

resource "google_storage_bucket" "bucket" {
  project       = var.gcp_project_id
  name          = var.bucket_name
  location      = var.bucket_location
  storage_class = "STANDARD"

}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

data "google_iam_policy" "viewer" {
  binding {
    role = "roles/storage.objectViewer"
    members = [
      "allUsers",
    ]
  }
}

data "google_client_config" "default" {}

provider "docker" {
  registry_auth {
    address  = "asia-southeast2-docker.pkg.dev"
    username = "oauth2accesstoken"
    password = data.google_client_config.default.access_token
  }
}

data "docker_registry_image" "get_metadata" {
  name = format("%s:%s", var.image, "latest")
}
