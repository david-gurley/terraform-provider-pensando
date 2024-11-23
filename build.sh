#!/usr/bin/env bash

go build -o terraform-provider-pensando
mv terraform-provider-pensando ~/.terraform.d/plugins/darwin_amd64/
