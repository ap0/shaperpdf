# shaperpdf

Simple program to create pages of unique shaper domino patterns for the [Shaper Origin](https://www.shapertools.com/en-us/origin).  Maybe there are other repos out there that do this; I didn't look -- I thought it'd be fun to reverse engineer it.

Using a pair of calipers and a little trial and error it appears the domino patterns must have ten white dots (bits) each, they must have two on the outsides of the shape, and the patterns can't be "palindromic" (i.e., they can't be read the same upside down as right side up).

They appear to scan fine with the origin.  You could probably use shipping tape to stick them down, or do something more permanent if using for a permanent jig.

Running:

```go
go run . -pages 3 -per-row 4 output.pdf
```