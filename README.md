# Bootkube

Bootkube provides an entire kubernetes control plane in a single binary, and includes commands to render and consume assets for bootstrapping a self-hosted kubernetes cluster.

## Usage

Bootkube has two modes of operation. 

### Render assets

First, you can use bootkube to render out all of the assets (including kubernetes object manifests, TLS assets and kubeconfig) that you need to run a self-hosted kubernetes cluster. This feature is still experimental and changing rapidly.

To use this feature, run:

```
bootkube render <options>
```

You can customize the generated manifests by passing flags to the command. For more information on the supported commands, run `bootkube help render`.

### Start bootkube

To start bootkube use the `start` subcommand:

```
bootkube start <options>
```

Bootkube expects a directory containing the manifests to be provided as a command line flag, as well as other TLS assets (all of which can be taken from the `render` command). To see the available flags, run `bootkube help start`.

## Build

First, clone the repo into the proper location in your $GOPATH:

```
git clone git@github.com:coreos/bootkube.git && cd bootkube
```

Then, to build:

```
make all
```

And optionally, to install into $GOPATH/bin:

```
make install
```

## License

bootkube is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
