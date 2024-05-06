#!/bin/bash
set -e

# Get all of the tags
git fetch --all --tags

# Get highest tag number
VERSION="$(git tag | sort -V | tail -1)"

# Get number parts and increase last one by 1
VNUM1=$(echo "$VERSION" | cut -d"." -f1)
VNUM2=$(echo "$VERSION" | cut -d"." -f2)
VNUM3=$(echo "$VERSION" | cut -d"." -f3)
VNUM1=`echo $VNUM1 | sed 's/v//'`

VNUM3=$((VNUM3+1))

# Create new tag
NEW_TAG="v$VNUM1.$VNUM2.$VNUM3"

echo "$NEW_TAG"