#!/bin/bash

source config.sh

RESOURCE="organization"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=55

export CP_ORG_NAME="test-or-$TAIKUN_USER-$RANDOM"
echo -n "$CP_ORG_NAME" > ref
export CP_ORG_DESCRIPTION="test-$TAIKUN_USER-desc"
export CP_ORG_DISCOUNT=42
export CP_ORG_CITY="test-organization-city"
export CP_ORG_BEMAIL="$TAIKUN_USER-organization@empty.com"
export CP_ORG_EMAIL="$TAIKUN_USER-organization@empty2.com"
export CP_ORG_PHONE=+441989
