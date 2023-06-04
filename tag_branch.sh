#!/bin/sh

git tag --delete release
git push https://github.com/young-yang03/cloud-mta-build-tool.git --delete release

git tag release
git push -u origin release

