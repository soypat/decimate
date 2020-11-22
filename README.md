# decimate

Need to graph a curve obtained through a simulation and the amount of points makes pgfplots run out of memory? Or do you simply want to reduce the size of your vector graphics?
---
Decimate will reduce the amount of points on your curve drastically while maintaining visual fidelity.

### Decimate is run from command line
run `decimate -h` for help

## Example

![Lots of data points](_assets/bigbig.png)
Above is graphed data from a 134MB file (8.4 million data entries). Following image is downsampled data using decimate (0.1 tolerance). Resulting files have a collective size of less than a megabyte.

![Less data points but identical to above](_assets/smolbig.png)

Data has been reduced over twohundredfold.

## Installation

You can download the latest release from https://github.com/soypat/decimate/releases.

If you prefer to build from source you'll need to install Go. Once installed run

```console
go build .
``` 

in the directory and a binary should be generated shortly.

