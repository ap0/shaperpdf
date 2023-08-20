# shaperpdf

Simple program to create pages of unique shaper domino patterns for the [Shaper Origin](https://www.shapertools.com/en-us/origin).  Maybe there are other repos out there that do this; I didn't look -- I thought it'd be fun to reverse engineer it.

Using a pair of calipers and a little trial and error it appears the domino patterns must have ten white dots (bits) each, they must have two on the outsides of the shape, and the patterns can't be "palindromic" (i.e., they can't be read the same upside down as right side up).  My best guess is that the dominoes have to be sized identically to the tape/plate/workstation.  These look extremely close to 1/2" and 1 11/16" height and width respectively.

They appear to scan fine with the Origin.  I have had the best luck printing these on a self-adhering sheet of paper and then trimming them using a paper cutter and applying them.  Sticking them down with e.g. shipping tape doesn't work very well.

Running:

```go
go run . -pages 5 output.pdf
```