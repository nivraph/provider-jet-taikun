#!/bin/bash

source config.sh

RESOURCE="access-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_ACCESS_NAME="ac-pr-$TAIKUN_USER-$RANDOM"
echo -n "$CP_ACCESS_NAME" > ref
export CP_ACCESS_ORG="test-$TAIKUN_USER-org-attach"
