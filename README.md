# GoArkitect

<p align="center">
<img src="docs/assets/goarkitect.logo.jpg" alt="crkitect" title="goarkitect" />
</p>

This project gives developers the ability to describe and check many architectural constraints of a project using a composable set of rules described in one or multiple yaml files.

## Example usage

```sh
# validate the default config file (.goarkitect.yaml) and outputs the result in json
goarkitect validate --output=json

# validate the custom .ark.yaml config file
goarkitect validate .ark.yaml

# validate the custom .ark.yaml config file and all the config files found in the .ark/ folder
goarkitect validate .ark.yaml .ark/

# verify that the current folder follows the rules specified in the default config file (.goarkitect.yaml)
goarkitect verify

# verify that the current folder follows the rules specified in the .ark/ folder and outputs the result in json
goarkitect verify .ark/ --output=json
```
