locals {
  action_binding = [
    {
      id   = module.reearth-api.auth0_action_singup.id,
      name = module.reearth-api.auth0_action_singup.name
    },
    {
      id   = module.reearth-cms.auth0_action_singup.id,
      name = module.reearth-cms.auth0_action_singup.name
    }
  ]

}

resource "auth0_trigger_binding" "reearth_login" {

  trigger = "post-user-registration"
  dynamic "actions" {
    for_each = local.action_binding
    content {
      id           = actions.value.id
      display_name = actions.value.name
    }
  }
}
