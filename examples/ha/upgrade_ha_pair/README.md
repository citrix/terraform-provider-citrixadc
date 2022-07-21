## Upgrade HA pair

This example shows how to upgrade an HA pair in an idempotent manner.

It relies on the `nsversion` data source to fetch the target ADC current version
and compares it with the desired version set by the user.

If the versions do not match then the count variable present on all resources
will get a value of 1 and the resources will be created.

If the versions match then the count variable will take a value of 0 which
will cause most resources to be deleted.

The delete operation for the resources present is a noop so no
changes will be applied to the target ADCs.

After that any subsequent application of the configuration will result in
an empty plan.
