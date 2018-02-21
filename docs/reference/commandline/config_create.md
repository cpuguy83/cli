---
title: "config create"
description: "The config create command description and usage"
keywords: ["config, create"]
---

<!-- This file is maintained within the docker/cli GitHub
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# config create

```Markdown
Usage:	docker config create [OPTIONS] SECRET file|-

Create a config from a file or STDIN as content

Options:
      --help                    Print usage
  -l, --label list              Config labels (default [])
      --template-driver string  Driver to use to parse templated configs
```

## Description

Creates a config using standard input or from a file for the config content. You must run this command on a manager node. 

For detailed information about using configs, refer to [manage sensitive data with Docker configs](https://docs.docker.com/engine/swarm/configs/).

## Examples

### Create a config

```bash
$ echo <config> | docker config create my_config -

onakdyv307se2tl7nl20anokv

$ docker config ls

ID                          NAME                CREATED             UPDATED
onakdyv307se2tl7nl20anokv   my_config           6 seconds ago       6 seconds ago
```

### Create a config with a file

```bash
$ docker config create my_config ./config.json

dg426haahpi5ezmkkj5kyl3sn

$ docker config ls

ID                          NAME                CREATED             UPDATED
dg426haahpi5ezmkkj5kyl3sn   my_config           7 seconds ago       7 seconds ago
```

### Create a config with labels

```bash
$ docker config create --label env=dev \
                       --label rev=20170324 \
                       my_config ./config.json

eo7jnzguqgtpdah3cm5srfb97
```

```none
$ docker config inspect my_config

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

### Create a templated config

You can create a config using templates which reference other configs or secrets.
A `template-driver` flag is provided to control this.

Currently only the `golang` driver is supported, which has two keywords:

- `config` - references a config object to be included inline, e.g. `config "foo"`
- `secret` - references a secret object to be included inline, e.g. `secret "bar"`

You can mix configs and secrets in a template, along with any other data you want
to include.

Example template: `Hello {{ config "name" }}, your passwod is {{ secret "password"}}`

More complex examples can be created by taking advantage of
[Go's templating engine](https://golang.org/pkg/text/template/)

Usage:

```bash
$ docker config create my_config1 ./config1.json
eo7jnzguqgtpdah3cm5srfb97
$ docker config create my_config1 ./config1.json
dg426haahpi5ezmkkj5kyl3sn
$ echo '{{ config "data1"}}{{ config "data2"}} | docker config create my_templated_config -
onakdyv307se2tl7nl20anokv
```

Notice that the configs referenced template are not the names of the real config,
instead these are evaluated at runtime. Here is an example of how to use this
templated config:

```bash
$ docker service create \
  --config source=my_config1,target=data1 \
  --config source=my_config_2,target=data2 \
	--config source=my_templated_config,target=combined \
	busybox top
xc4zutdyh63abg17osqgntc1n
```

Each of the configs, or secrets, referenced in the template must be used in the
service, and the targets must match what is defined in the template rather than
the real name.

## Related commands

* [config inspect](config_inspect.md)
* [config ls](config_ls.md)
* [config rm](config_rm.md)
