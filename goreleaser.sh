#!/bin/sh

git tag --delete v1.2.228
git push https://github.com/young-yang03/cloud-mta-build-tool.git --delete v1.2.228

git tag v1.2.228
git push https://github.com/young-yang03/cloud-mta-build-tool.git v1.2.228