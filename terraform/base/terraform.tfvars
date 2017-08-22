terragrunt = {
  include {
    path = "${find_in_parent_folders()}"
  }

  terraform {
    extra_arguments "conditional_vars" {
      commands = ["${get_terraform_commands_that_need_vars()}"]

      optional_var_files = [
        "${get_tfvars_dir()}/local.tfvars",
      ]
    }
  }
}
