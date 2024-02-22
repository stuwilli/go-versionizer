# Go Versionizer

This is a simple tool to generate incremental version numbers for projects. 

It is designed to be used in a CI/CD pipeline to automatically generate version numbers for each build. The code is written in Go and is designed to be used as a command line tool, it requires git to be installed on the system.

The tool will return the next version number based on the latest git tag. If no tags are found, the tool will return `0.0.1`.

## Usage

To return the next version number for a project, run the following command in the root of the project:
```shell
go-versionizer --bump [major|minor|patch] --create --push [optional path to git repository, defaults to current directory]
```

### Command Flags/Arguments
The optional `--bump` flag is used to specify the type of version bump to be performed. The options are `major`, `minor` or `patch`. If no option is provided, the tool will default to `patch`.

The optional `--create` flag can be used to create a new tag in the repository with the new version number. Default is `false`.

The optional `--push` flag can be used to push the new tag to the remote repository. Default is `false`.

The final parameter is the path to the git repository. If no path is provided, the tool will default to the current directory.

## Example basic usage
```shell
versionizer --bump patch /path/to/git/repo
```

## Example returning new version number only (no push or tag creation) in the current directory
```shell
versionizer --bump patch
```

## Example which creates a new tag in the repository
```shell
versionizer --bump patch --create /path/to/git/repo
```

## Example which creates a new tag in the repository and pushes it to the remote
```shell
versionizer --bump patch --create --push /path/to/git/repo
```
