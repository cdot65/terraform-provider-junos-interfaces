---
page_title: "Provider: Junos IPsec"
subcategory: ""
description: |-
  Terraform provider for interacting with Juniper's NETCONF API for provisioning IPsec configurations on SRX firewalls.
---

# Juniper IPsec Provider

-> Visit the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) Learn tutorials for an interactive getting started experience.

The `terraform-provider-junos-ipsec` provider is used to interact with a Juniper SRX firewall. This provider is meant to serve as an educational tool to show users how:
1. use providers to [create, read, update and delete (CRUD) resources](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers) using Terraform.
1. create a custom Terraform provider.

To learn how to re-create the `terraform-provider-junos-ipsec` provider, refer to the [Official Documentation](https://cdot65.github.io/juniper-terraform-srx/) for a walkthrough.

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "ipsec" {
  host     = "enterprise-fw1"
  port     = 22
  sshkey   = ""
  username = "terraform"
  password = "juniper123"
}
```

## Schema

### Optional

- **username** (String, Optional) Username to authenticate to SRX firewall
- **password** (String, Optional) Password to authenticate to SRX firewall
- **host** (String, Optional) SRX firewall address
- **port** (String, Optional) SRX firewall port number
- **sshkey** (String, Optional) SRX firewall SSH key
