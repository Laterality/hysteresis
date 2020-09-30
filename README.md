# Hysteresis - DB migration tool using flyway

Hysteresis is DB migration tool based on [flyway](https://flywaydb.org/documentation/configfiles) internally.

## Prerequisites

Tool uses [container-wrapped flyway](https://hub.docker.com/r/flyway/flyway). So there is no need to download such as JDK or flyway jar.

You only need running docker daemon to use this tool.

## Usage

```bash
# Show migration information
./hyst info

# Perform migration
./hyst migrate
```

## Configuration

The tool just mounts flyway configuration file on `./conf/${profile}.conf` into flyway container.

Also mounts `.sql` files on `./sql` directory.

So, structure of working directory will look like following:

```
./
|- hyst # built binary
|- conf
|  |- local.conf # Connection information for local database
|  `- dev.conf   # Connection information for some other environment database
`- sql
   |- V1__Initial_tables.sql
   `- V2__Add_some_index.sql
```

The profile is `local` by default. So if you want to perform migration on you local DB, just run:

```bash
./flyway migrate
```

When if you perform migration with `develop` configuration:

```bash
./flyway -p develop migrate
```

Then, Command above uses `./conf/develop.conf` file for running flyway.

## Build

```bash
make build-linux # Builds linux binary
make build-mac   # Builds mac os binary
```

Then you can run binary named `hyst`.
