#!/usr/bin/env bash
cd ../
find . -name ".idea" -exec rm -rf {} \;
find . -name "*.iml" -exec rm -rf {} \;
find . -name "target" -exec rm -rf {} \;
find . -name ".DS_Store" -exec rm -rf {} \;
find . -name ".project" -exec rm -rf {} \;
find . -name "node_modules" -exec rm -rf {} \;