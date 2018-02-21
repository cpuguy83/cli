---
title: "config ls"
description: "The config ls command description and usage"
keywords: ["config, ls"]
---

<!-- This file is maintained within the docker/cli GitHub
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# config ls

```Markdown
Usage:	docker config ls [OPTIONS]

List configs

Aliases:
  ls, list

Options:
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print configs using a Go template
      --help            Print usage
  -q, --quiet           Only display IDs
```

## Description

Run this command on a manager node to list the configs in the swarm.

For detailed information about using configs, refer to [manage sensitive data with Docker configs](https://docs.docker.com/engine/swarm/configs/).

## Examples

```bash
$ docker config ls

ID                          NAME                        CREATED             UPDATED
6697bflskwj1998km1gnnjr38   q5s5570vtvnimefos1fyeo2u2   6 weeks ago         6 weeks ago
9u9hk4br2ej0wgngkga6rp4hq   my_config                   5 weeks ago         5 weeks ago
mem02h8n73mybpgqjf0kfi1n0   test_config                 3 seconds ago       3 seconds ago
```

### Filtering

The filtering flag (`-f` or `--filter`) format is a `key=value` pair. If there is more
than one filter, then pass multiple flags (e.g., `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* [id](config_ls.md#id) (config's ID)
* [label](config_ls.md#label) (`label=<key>` or `label=<key>=<value>`)
* [name](config_ls.md#name) (config's name)

#### id

The `id` filter matches all or prefix of a config's id.

```bash
$ docker config ls -f "id=6697bflskwj1998km1gnnjr38"

ID                          NAME                        CREATED             UPDATED
6697bflskwj1998km1gnnjr38   q5s5570vtvnimefos1fyeo2u2   6 weeks ago         6 weeks ago
```

#### label

The `label` filter matches configs based on the presence of a `label` alone or
a `label` and a value.

The following filter matches all configs with a `project` label regardless of
its value:

```bash
$ docker config ls --filter label=project

ID                          NAME                        CREATED             UPDATED
mem02h8n73mybpgqjf0kfi1n0   test_config                 About an hour ago   About an hour ago
```

The following filter matches only services with the `project` label with the
`project-a` value.

```bash
$ docker service ls --filter label=project=test

ID                          NAME                        CREATED             UPDATED
mem02h8n73mybpgqjf0kfi1n0   test_config                 About an hour ago   About an hour ago
```

#### name

The `name` filter matches on all or prefix of a config's name.

The following filter matches config with a name containing a prefix of `test`.

```bash
$ docker config ls --filter name=test_config

ID                          NAME                        CREATED             UPDATED
mem02h8n73mybpgqjf0kfi1n0   test_config                 About an hour ago   About an hour ago
```

### Format the output

The formatting option (`--format`) pretty prints configs output
using a Go template.

Valid placeholders for the Go template are listed below:

| Placeholder  | Description                                                                          |
| ------------ | ------------------------------------------------------------------------------------ |
| `.ID`        | Config ID                                                                            |
| `.Name`      | Config name                                                                          |
| `.CreatedAt` | Time when the config was created                                                     |
| `.UpdatedAt` | Time when the config was updated                                                     |
| `.Labels`    | All labels assigned to the config                                                    |
| `.Label`     | Value of a specific label for this config. For example `{{.Label "config.ssh.key"}}` |

When using the `--format` option, the `config ls` command will either
output the data exactly as the template declares or, when using the
`table` directive, will include column headers as well.

The following example uses a template without headers and outputs the
`ID` and `Name` entries separated by a colon for all images:

```bash
$ docker config ls --format "{{.ID}}: {{.Name}}"

77af4d6b9913: config-1
b6fa739cedf5: config-2
78a85c484f71: config-3
```

To list all configs with their name and created date in a table format you
can use:

```bash
$ docker config ls --format "table {{.ID}}\t{{.Name}}\t{{.CreatedAt}}"

ID                  NAME                      CREATED
77af4d6b9913        config-1                  5 minutes ago
b6fa739cedf5        config-2                  3 hours ago
78a85c484f71        config-3                  10 days ago
```

## Related commands

* [config create](config_create.md)
* [config inspect](config_inspect.md)
* [config rm](config_rm.md)
