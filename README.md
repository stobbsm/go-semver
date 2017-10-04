# go-semver
Semantic versioning support in Go

## Install
``` go get github.com/stobbsm/go-semver ```

## Usage
To import use:
``` import version "github.com/stobbsm/go-semver" ```

### Set version
Use the Set method to set the version number.
Set takes 5 arguments:
1. Major uint
2. Minor uint
3. Patch uint
4. Release string
5. Build string

If your version is 1.2.3-alpha+build1
``` version.Set(1, 2, 3, "alpha", "build1") ```
Or, to omit release and build versioning,
``` version.Set(1, 2, 3, "", "") ```

### Get version
The Get method returns the version as a string. For the examples above, they would be:
``` "1.2.3-alpha+build1" ``` and ``` "1.2.3" ``` respectively
