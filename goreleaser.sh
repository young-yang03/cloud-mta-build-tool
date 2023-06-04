#!/bin/sh

git tag --delete v.1.2.229
git push https://github.com/young-yang03/cloud-mta-build-tool.git --delete v.1.2.229

git tag v.1.2.229
git push https://github.com/young-yang03/cloud-mta-build-tool.git v.1.2.229