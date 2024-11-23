# Go language bindings for Pensando APIs - PSM REST and DSC gRPC

## Updating PDS Import Statements

If you update the pds directory, run the following to update the import statements:
```
gofmt -w -r '"github.com/pensando/sw/nic/apollo/agent/gen/pds/meta/pds" -> "github.com/david-gurley/gopen/pds/meta/pds"' *.go
gofmt -w -r '"github.com/pensando/sw/nic/apollo/agent/gen/pds" -> "github.com/david-gurley/gopen/pds"' *.go
```
