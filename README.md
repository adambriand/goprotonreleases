# goprotonreleases

A simple app written in Go to display the latest Proton releases (Valve Software compatibility tool for Linux).

This app checks the repositories for Proton and Proton-ge-custom (GloriousEggroll), then displays their latest tagged release,  date of the release, and download links to stdout.

## Build and Installation

To build, download the source code and build using Go (`go build goprotonreleases.go`).

To run the Linux binary directly, get the latest .tar.gz from Releases, uncompress, and run.

## Configuration

This is a MVP for a simple display of current Proton versions, there are no configuration options currently.

## References

[Proton](https://github.com/ValveSoftware/Proton/)

[Proton GE](https://github.com/GloriousEggroll/proton-ge-custom)

## Acknowledgments

Thank you to Wine developers, Valve Software, and GloriousEggroll for creating/customizing such great software. Without you Linux gaming would not be where it is today.

Special thank you to Tom Lebreux for the code review and suggestions, and Brad Avery for the testing.
