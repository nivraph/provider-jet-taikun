# Using the Testsuite

To run the testsuite, please follow the required steps:

## Setting the variables

Export the following environment variables:

* TAIKUN_EMAIL
* TAIKUN_PASSWORD
* TAIKUN_USER
* PROVIDER_NAME
* PROMETHEUS_USER
* PROMETHEUS_PASSWORD
* S3_ACCESS_KEY
* S3_SECRET_KEY
* OPENSTACK_CLOUD_USER
* OPENSTACK_CLOUD_PASSWORD

## Running the tests

Once you have exported the required variables, just run:

```
make test
```
