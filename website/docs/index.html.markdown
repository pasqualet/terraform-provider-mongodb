---
layout: "mongodb"
page_title: "Provider: MongoDB"
sidebar_current: "docs-mongodb-index"
description: |-
  The MongoDB provider is used to interact with a MongoDB instance or replicaset. The provider needs to be configured with the proper credentials before it can be used.
---

# MongoDB Provider

The MongoDB provider is used to interact with
a MongoDB instance or replicaset. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
# Configure the MongoDBProvider
provider "mongodb" {
  url = "mongodbdb://localhost:27017/test"
}

# Create a user
resource "mongodb_user" "user-test" {
  # ...
}
```

## Configuration Reference

The following arguments are supported:

* `url` - (Required) A valid MongoDB url to use.

## Testing and Development

In order to run the Acceptance Tests for development, the following environment
variables must also be set:

* `MONGODB_URL` - The MongoDB url to use.
