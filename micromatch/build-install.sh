#!/bin/sh

# clean env
rm -rf ./micromatch-wrapper-linux ./micromatch-wrapper-macos ./micromatch-wrapper-win.exe
rm -rf node_modules

# install pkg
npm install pkg
npx pkg --version

# build micromatch wrapper
npm install
npx pkg ./

# install and test micromatch wrapper
cp ./micromatch-wrapper-win.exe $GOPATH/bin/micromatch-wrapper.exe
micromatch-wrapper.exe -h

# clean and copy micromatch wrapper to target path for release to GitHub
rm -rf ../micromatch-wrapper-Linux ../micromatch-wrapper-Darwin ../micromatch-wrapper-Windows
mv ./micromatch-wrapper-linux ../micromatch-wrapper-Linux
mv ./micromatch-wrapper-macos ../micromatch-wrapper-Darwin
mv ./micromatch-wrapper-win.exe ../micromatch-wrapper-Windows