# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Hyper-V"
  description = "The Hyper-V plugin can be used with HashiCorp Packer to create custom images from Hyper-V."
  identifier = "packer/hashicorp/hyperv"
  component {
    type = "builder"
    name = "Hyper-V VMCX"
    slug = "vmcx"
  }
  component {
    type = "builder"
    name = "Hyper-V ISO"
    slug = "iso"
  }
}
