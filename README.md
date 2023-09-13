# shaperpdf

Simple program to create pages of unique domino patterns for the [Shaper Origin](https://www.shapertools.com/en-us/origin) rather than purchasing Shaper Tape.  Maybe there are other repos out there that do this; I didn't look -- I thought it'd be fun to reverse engineer it.

Using a pair of calipers and a little trial and error it appears the domino patterns must have ten white dots (bits) each, they must have two on the outsides of the shape, and the patterns can't be "palindromic" (i.e., they can't be read the same upside down as right side up).  My best guess is that the dominoes have to be sized identically to the tape/plate/workstation.  These look extremely close to 1/2" and 1 11/16" height and width respectively.

They seem to scan fine with the Origin.  I have had the best luck printing these on a self-adhering sheet of paper and then trimming them using a paper cutter and applying them to the work surface.  Sticking them down with e.g. shipping tape doesn't work very well.

Note that typical laser-printed self-adhering sheets are not particularly durable, especially when you're rubbing on them with the Origin and having sawdust around.  They black ink tends to get worn off fairly quickly, which can cause problems for reading while cutting with these.  The actual Shaper Tape seems to be much more durable.  If you do go this route, go for a lot of density so there's more dominoes for the Origin to reference.  For cuts where absolute accuracy matters, stick with the tape.

I use this mostly for doing cutouts of things that don't have to be super precise (e.g. cutting out wooden holiday decorations). I frequently have to blow the sheets off with compressed air during cutting.  I've gotten some errors during cutting with them that seem to be related to the dominos getting obscured by debris.  I have never had this problem with the real shaper tape.

Cost-wise, this approach has worked out to be about 20% of the cost of Shaper Tape.  However, if you're using expensive materials, that extra 80% may be a small price to pay for knowing that the tape will always work.

Running:

```go
go run . -pages 5 output.pdf
```
