package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/go-pdf/fpdf"
)

const (
	scaleY                   = 1
	scaleX                   = 3.5
	radiusRatioToHeight      = 0.2
	innerRadiusToHeightRatio = 0.5
	spacingRatioToWidth      = 0.067
	digitLimit               = 6
)

var (
	dominoesPerRow int
	numPages       int
	filename       string
	orientation    string
	pageSize       string
	debug          bool
)

func dominoHeightForWidth(width float64) float64 {
	return width * (scaleY / scaleX)
}

func main() {
	flag.IntVar(&numPages, "pages", 10, "The number of pages to generate.")
	flag.IntVar(&dominoesPerRow, "per-row", 5, "The number of dominoes per row.")
	flag.StringVar(&orientation, "orientation", "L", "Page orientation: P (portrait) or L (landscape)")
	flag.StringVar(&pageSize, "page-size", "Letter", "Page size: Letter, Legal, Tabloid, A3, A4, A5")
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s [OPTION]... [FILE]\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	filename = flag.Arg(0)
	if filename == "" || numPages <= 0 || dominoesPerRow <= 0 {
		flag.Usage()
		return
	}

	if err := runPDF(); err != nil {
		log.Fatal(err)
	}
}

func computeValidNumbers() []int {
	valid := make([]int, 0)
	for i := 63; i < 4096; i++ {
		if isValid(i) {
			valid = append(valid, i)
		}
	}
	return valid
}

func runPDF() error {
	pdf := fpdf.New(orientation, "pt", pageSize, "")

	numbers := computeValidNumbers()
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	numIdx := 0
	for i := 0; i < numPages; i++ {
		pdf.AddPage()
		pageWidth, pageHeight := pdf.GetPageSize()
		marginLeft, marginTop, _, _ := pdf.GetMargins()
		topX := marginLeft
		topY := marginTop
		overallDomWidth := float64(pageWidth-2*marginLeft) / float64(dominoesPerRow)

		domWidth := scaleX / ((spacingRatioToWidth * scaleX) + scaleX) * overallDomWidth
		domHeight := dominoHeightForWidth(domWidth)

		dbgLog("domWidth: ", domWidth)
		dbgLog("domHeight: ", domHeight)

		for {
			for domRowIndex := 0; domRowIndex < dominoesPerRow; domRowIndex++ {

				dbgLog("topX: ", topX)
				dbgLog("topY: ", topY)

				number := numbers[numIdx]
				numIdx++
				dbgLog("using", number)

				pdf.SetFillColor(0, 0, 0)
				pdf.RoundedRect(topX, topY, domWidth, domHeight, domHeight*radiusRatioToHeight, "1234", "F")

				circleRadius := domWidth / 17 / 2
				circleWidth := circleRadius * 2
				innerOffsetX := topX
				innerOffsetY := topY

				dbgLog("innerOffsetX", innerOffsetX)
				dbgLog("innerOffsetY", innerOffsetY)
				numStr := fmt.Sprintf("11%012b11", number)
				for x := 0; x < 8; x++ {
					for y := 0; y < 2; y++ {
						fX := float64(x)
						fY := float64(y)

						val := numStr[x*2+y]

						if val == '1' {

							pdf.SetFillColor(255, 255, 255)
							pdf.Circle(innerOffsetX+(fX*2+1)*circleWidth+circleRadius, innerOffsetY+(fY*2+1)*circleWidth+circleRadius, circleRadius, "F")
						}

					}
				}

				topX += overallDomWidth
			}
			topY += domHeight * 2
			topX = marginLeft
			if topY > pageHeight-marginTop*2-domHeight {
				break
			}
		}
	}
	return pdf.OutputFileAndClose(filename)
}

func isValid(num int) bool {
	str := fmt.Sprintf("%012b", num)

	cnt := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			cnt++
		}
	}

	return cnt == digitLimit && !isPalindrome(str)
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func dbgLog(s ...any) {
	if os.Getenv("DEBUG") == "1" || debug {
		log.Println(s...)
	}
}
