# Go Versionizer

This is a simple tool to generate incremental version numbers for projects. 

It is designed to be used in a CI/CD pipeline to automatically generate version numbers for each build. The code is written in Go and is designed to be used as a command line tool, it requires git to be installed on the system.

## Usage
To return the next version number for a project, run the following command in the root of the project:
```shell
go-versionizer --bump [major|minor|patch] [optional path to git repository, defaults to current directory]
```

The tool will return the next version number based on the latest git tag. If no tags are found, the tool will return `0.0.1`.

## Example
```shell
go-versionizer --bump patch /path/to/git/repo
```

## Example which creates a new tag in the repository
```shell
go-versionizer --bump patch --create /path/to/git/repo
```