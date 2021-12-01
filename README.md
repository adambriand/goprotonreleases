# goprotonreleases

A simple app written in Go to display the latest Proton releases (Valve Software compatibility tool for Linux).

This app checks the repositories for Proton and Proton-ge-custom (GloriousEggroll), then displays their latest tagged release,  date of the release, and download link.

## Build and Installation

To build, download the source code and build using Go (`go build goprotonreleases.go`). If you do not have GLIBC 2.32 installed you can compile it with `CGO_ENABLED=0 go build goprotonreleases.go`.
 
To install the Linux binary directly, download the latest `.tar.gz` file from Releases and uncompress.

## Usage

This is a MVP for a simple display of current Proton versions, there are no configuration options. Run `goprotonreleases` and the results will display on stdout.

Want to stay up to date on your Proton releases? Why not add it to your `.bashrc`.

## References

[Proton](https://github.com/ValveSoftware/Proton/)

[Proton GE](https://github.com/GloriousEggroll/proton-ge-custom)

## Acknowledgments

Thank you to Wine developers, Valve Software, and GloriousEggroll for creating/customizing such great software. Without you Linux gaming would not be where it is today.

Special thank you to Tom Lebreux for the code review and suggestions, and Brad Avery for the testing.
