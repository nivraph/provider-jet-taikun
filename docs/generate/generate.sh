#!/usr/bin/env bash

set -e

PROJECT_PWD="$PWD"
CRDS_DIR=$PROJECT_PWD"/package/crds"
CRDS_LIST=`ls $CRDS_DIR`

pip install mdutils

python docs/generate/generate.py