output "auth0_client_spa" {
  value = auth0_client.spa
}

output "auth0_client_m2m" {
  value = auth0_client.m2m
}

output "action_secret" {
  value = random_string.action_secret
}

output "action_singup" {
  value = auth0_action.singup
}