The HyperV Packer Plugin is able to create
[Hyper-V](https://www.microsoft.com/en-us/server-cloud/solutions/virtualization.aspx)
virtual machines and export them.

## Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    hyperv = {
      source  = "github.com/hashicorp/hyperv"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/hyperv
```

### Components

#### Builders

- [hyperv-iso](packer/integrations/hashicorp/hyperv/latest/components/builder/hyperv-iso.mdx) - Starts from an ISO file,
  creates a brand new Hyper-V VM, installs an OS, provisions software within
  the OS, then exports that machine to create an image. This is best for
  people who want to start from scratch.

- [hyperv-vmcx](packer/integrations/hashicorp/hyperv/latest/components/builder/hyperv-vmcx.mdx) - Clones an an existing
  virtual machine, provisions software within the OS, then exports that
  machine to create an image. This is best for people who have existing base
  images and want to customize them.
