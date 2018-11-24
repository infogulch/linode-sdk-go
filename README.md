# linode-sdk-go (unreleased)

A Go SDK for the Linode API ( https://developers.linode.com/api/ ) generated with openapi-generator via Docker.

Some simple scripts (hopefully folded into a CI system in the future):
* `./download.sh`: downloads the current Linode API document to overwrite `./openapi.yaml`.
* `./validate.sh`: validates `./openapi.yaml` using openapi-generator-cli validation.
* `./generate.sh`: generates Go code based on `./openapi.yaml` into `./lsdk` directory.

