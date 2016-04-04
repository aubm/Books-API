#!/bin/bash

# http://ricostacruz.com/cheatsheets/travis-gh-pages.html

set -o errexit

# clear and re-create the public directory
rm -rf docs;
mkdir docs;

git config --global user.email "travis@nomail.com"
git config --global user.name "Travis CI"

# run our compile script, discussed above
postmanerator \
    -collection=postman/collection.json \
    -output=docs/index.html \
    -theme=default \
    -ignored-request-headers="dbscripts" \
    -ignored-response-headers="Content-Length,Date"

# go to the public directory and create a *new* Git repo
cd docs
git init

# The first and only commit to this new Git repo contains all the
# files present with the commit message "Deploy to GitHub Pages".
git add .
git commit -m "Deploy to GitHub Pages"

# Force push from the current repo's master branch to the remote
# repo's gh-pages branch. (All previous history on the gh-pages branch
# will be lost, since we are overwriting it.) We redirect any output to
# /dev/null to hide any sensitive credential data that might otherwise be exposed.
git push --quiet --force "https://${GITHUB_TOKEN}@github.com/${GITHUB_REPO}.git" master:gh-pages > /dev/null 2>&1
