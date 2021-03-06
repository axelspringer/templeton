[![Build Status](https://travis-ci.org/axelspringer/templeton.svg?branch=master)](https://travis-ci.org/axelspringer/templeton)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Templeton

Templeton is a tiny tool to run commands in [Docker](https://docker.io) with a preconfigured environment from the [System Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html). It does fetches the values from a configured parameter store path and configures the runtime environment of the command to execute.

## Getting Started

The most simple case is, that you want to execute a command with all the available parameters from the store. Folders are converted to `_` in the environment variable.

> it is recommended to use `TEMPLETON_PATH` to configure the path to fetch
> all parameters are prefixed with `SSM` when not otherwise specified

```
templeton --path /example run echo "$SSM_DEV_PARAMETER
```

This would suggest that there is a key in `/example/dev/parameter`.

## Config

You can use `--help` to print out all the available options.

### TEMPLETON_PATH (--path)

Configures the path in the Parameter Store from which to fetch the configs.

## License
[Apache-2.0](/LICENSE)
