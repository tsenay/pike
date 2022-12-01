resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.ApiManagement/service/read",
      "Microsoft.ApiManagement/service/write",
      "Microsoft.ApiManagement/service/delete",
      "Microsoft.ApiManagement/service/operationresults/read",
      "Microsoft.ApiManagement/service/policies/read",
      "Microsoft.ApiManagement/service/portalsettings/read",
      "Microsoft.ApiManagement/service/tenant/listSecrets/action"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
