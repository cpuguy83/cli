---
title: "config inspect"
description: "The config inspect command description and usage"
keywords: ["config, inspect"]
---

<!-- This file is maintained within the docker/cli GitHub
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# config inspect

```Markdown
Usage:  docker config inspect [OPTIONS] CONFIG [CONFIG...]

Display detailed information on one or more configs

Options:
  -f, --format string   Format the output using the given Go template
      --help            Print usage
```

## Description

Inspects the specified config. This command has to be run targeting a manager
node.

By default, this renders all results in a JSON array. If a format is specified,
the given template will be executed for each result.

Go's [text/template](http://golang.org/pkg/text/template/) package
describes all the details of the format.

For detailed information about using configs, refer to [manage sensitive data with Docker configs](https://docs.docker.com/engine/swarm/configs/).

## Examples

### Inspect a config by name or ID

You can inspect a config, either by its *name*, or *ID*

For example, given the following config:

```bash
$ docker config ls

ID                          NAME                CREATED             UPDATED
eo7jnzguqgtpdah3cm5srfb97   my_config           3 minutes ago       3 minutes ago
```

```none
$ docker config inspect config.json

[
    {
        "ID": "eo7jnzguqgtpdah3cm5srfb97",
        "Version": {
            "Index": 17
        },
        "CreatedAt": "2017-03-24T08:15:09.735271783Z",
        "UpdatedAt": "2017-03-24T08:15:09.735271783Z",
        "Spec": {
            "Name": "my_config",
            "Labels": {
                "env": "dev",
                "rev": "20170324"
            }
        }
    }
]
```

### Formatting

You can use the --format option to obtain specific information about a
config. The following example command outputs the creation time of the
config.

```bash
$ docker config inspect --format='{{.CreatedAt}}' eo7jnzguqgtpdah3cm5srfb97

2017-03-24 08:15:09.735271783 +0000 UTC
```


## Related commands

* [config create](config_create.md)
* [config ls](config_ls.md)
* [config rm](config_rm.md)
