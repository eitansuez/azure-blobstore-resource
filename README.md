# azure-blobstore-resource
A concourse resource to interact with the azure blob service.

> NOTE: The resource has been moved from the `czero` dockerhub account to the `pcfabr` dockerhub
> account. If your pipeline is currently using the resource from `czero` it should be switched
> to `pcfabr`.

## Source Configuration

* `storage_account_name`: *Required.* The storage account name on Azure.

* `storage_account_key`: *Required.* The storage account access key for the storage account on Azure.

* `container`: *Required.* The name of the container in the storage account.

* `base_url`: *Optional.* The storage endpoint to use for the resource. Defaults to the
  Azure Public Cloud (core.windows.net).

### Filenames

One of the two options must be specified:

* `regexp`: *Optional.* The pattern to match filenames against. At least one capture group
  must be specified, with parentheses to extract the version. If multiple capture groups are
  provided the first group is used by default, but if a group is named `version` that will
  be extracted as the version. Semantic versions and numbers are supported for versioning.

* `versioned_file`: *Optional.* The file name of the blob to be managed by the resource.
  The resource only pulls the latest snapshot. If the blob doesn't have a snapshot, the
  resource will not find the blob. A new snapshot must also be created when a blob is
  updated for the resource to successfully check new versions.

## Behavior

### `check`: Extract snapshot versions from the container.

Checks for new snapshot versions for a file. If a blob exists without a snapshot
the file will not be found.

### `in`: Fetch a blob from the container.

Places the blob file in the destination.

#### Parameters

* `unpack`: *Optional.* If true, the blob will be unpacked before running the task. Supports
  tar, zip, gzip files.

### `out`: Upload a blob to the container.

Uploads a file to the container. If `regexp` is specified, the new file will be uploaded
to the directory that the regex searches in. If `versioned_file` is specified, the
new file will be uploaded as a new snapshot of that file.

#### Parameters

* `file`: *Required.* Path to the file to upload, provided by an output of a task. If multiple
  files are matched by the glob, an error is raised. The file that matches the glob will be
  uploaded into the directory specified by the regexp. Only supports bash glob expansion, not
  regex.

## Example Configuration

An example pipeline exists in the `example` directory.

### Resource

When using Azure snapshots:

```yaml
resource_types:
- name: azure-blobstore
  type: docker-image
  source:
    repository: pcfabr/azure-blobstore-resource

resources:
  - name: terraform-state
    type: azure-blobstore
    source:
      storage_account_name: {{storage_account_name}}
      storage_account_key: {{storage_account_key}}
      container: {{container}}
      versioned_file: terraform.tfstate
```

or with regexp:

```yaml
resource_types:
- name: azure-blobstore
  type: docker-image
  source:
    repository: pcfabr/azure-blobstore-resource

resources:
  - name: my-release
    type: azure-blobstore
    source:
      storage_account_name: {{storage_account_name}}
      storage_account_key: {{storage_account_key}}
      container: {{container}}
      regexp: release-(.*).tgz
```

### Plan

```yaml
- get: terraform-state
```

```yaml
- put: terraform-state
  params:
    file: terraform-state/terraform.tfstate
```

and with regexp:

```yaml
- put: my-release
  params:
    file: my-release/release-*.tgz
```
