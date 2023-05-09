output "policy_name" {
  description = "Unique name of the policy"
  value       = powerplatform_data_loss_prevention_policy.my_policy.name
}

output "policy_display_name" {
  description = "Display name of the policy"
  value       = powerplatform_data_loss_prevention_policy.my_policy.display_name
}