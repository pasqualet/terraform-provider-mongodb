---
layout: "mongodb"
page_title: "MongoDB: user"
sidebar_current: "docs-mongodb-user"
description: |-
  Creates a user.
---

# mongodb_user

Creates a user.

## Example Usage

```
resource "mongodb_user" "user1" {
   username = "user"
   database = "test"
   password = "changeme"
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) The name of the new user.

* `password` - (Required) The user's password.

* `database` - (Required) The name of the network.

* `roles` - an array of valid MongoDB roles.

## Attributes Reference

The following attributes are exported:

* `id` - The id of the user.
* `username` - See Argument Reference above.
* `database` - See Argument Reference above.
* `password` - See Argument Reference above.
* `roles` - See Argument Reference above..
