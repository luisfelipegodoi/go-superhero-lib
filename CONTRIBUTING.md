## Start

This repository works as an external lib for Go projects. Its function is to externalize useful features for developments.

## Commit format

[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

We suggest that messages follow the pattern of _conventional commit_.

More details:
- [Conventional Commit Specifications](https://www.conventionalcommits.org/)
- [Rules of @commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional).

## Pull requests

Regardless of the contribution to be made (without source code and / or documents), operationally speaking, we have 2 ways to make * merge requests *: locally, via the command line, using Git in conjunction with Github, or online , directly editing the files in Github and requesting a merge request later, all graphically.

## Localhost edition of code

* Fork * or * merge * the project. Create a new branch from the master. Make the changes you want. Submit an MR to the master. Wait for approval and merge.

### Keep the *branch* up to date with the root repository

- If you don't already have your *fork* clone locally, crete it with:
`git clone https://github.com/seuusuario/go-superhero-lib.git`

- Added a remote for your root repository:
`git remote add reporaiz https://github.com/go-superhero-lib.git` (*rootrepo* nickname for the api root repository. You can use anything name*)

- Update your local repo from the remote root repository
`git fetch reporaiz`

- Go to your branch:
`git checkout my-branch`

- Update your branch with the with the updates from master of root repository
`git pull --rebase reporaiz master`

- Update your remote *branch*
`git push origin mybranch`

- In case of any conflict when doing the `push`, you can the use the command:
`git push origin --force-with-lease`.