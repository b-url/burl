#!/bin/bash

origin_head_sha=$(git rev-parse origin/HEAD)

git cliff "$origin_head_sha"..HEAD | pbcopy
echo Changelog has been copied to the clipboard.
pbpaste
