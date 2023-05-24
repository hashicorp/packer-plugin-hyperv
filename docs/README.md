# Hyper-V Plugins

The HyperV Packer Plugin is able to create
[Hyper-V](https://www.microsoft.com/en-us/server-cloud/solutions/virtualization.aspx)
virtual machines and export them.

- [hyperv-iso](/docs/builders/hyperv-iso.mdx) - Starts from an ISO file,
  creates a brand new Hyper-V VM, installs an OS, provisions software within
  the OS, then exports that machine to create an image. This is best for
  people who want to start from scratch.

- [hyperv-vmcx](/docs/builders/hyperv-vmcx.mdx) - Clones an an existing
  virtual machine, provisions software within the OS, then exports that
  machine to create an image. This is best for people who have existing base
  images and want to customize them.

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    name = {
      version = ">= 1.1.0"
      source  = "github.com/hashicorp/name"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/hashicorp/packer-plugin-name/releases).
Once you have downloaded the latest archive corresponding to your target OS,
uncompress it to retrieve the plugin binary file corresponding to your platform.
To install the plugin, please follow the Packer documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


#### From Source

If you prefer to build the plugin from its source code, clone the GitHub
repository locally and run the command `go build` from the root
directory. Upon successful compilation, a `packer-plugin-name` plugin
binary file can be found in the root directory.
To install the compiled plugin, please follow the official Packer documentation
on [installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).
