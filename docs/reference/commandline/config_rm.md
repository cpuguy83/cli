---
title: "config rm"
description: "The config rm command description and usage"
keywords: ["config, rm"]
---

<!-- This file is maintained within the docker/cli GitHub
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# config rm

```Markdown
Usage:	docker config rm CONFIG [CONFIG...]

Remove one or more configs

Aliases:
  rm, remove

Options:
      --help   Print usage
```

## Description

Removes the specified configs from the swarm. This command has to be run
targeting a manager node.

For detailed information about using configs, refer to [manage sensitive data with Docker configs](https://docs.docker.com/engine/swarm/configs/).

## Examples

This example removes a config:

```bash
$ docker config rm config.json
sapth4csdo5b6wz2p5uimh5xg
```

> **Warning**: Unlike `docker rm`, this command does not ask for confirmation
> before removing a config.


## Related commands

* [config create](config_create.md)
* [config inspect](config_inspect.md)
* [config ls](config_ls.md)
