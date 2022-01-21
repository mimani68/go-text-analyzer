package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jdkato/prose/v2"
)

func main() {
	query := `market`
	text := `EURUSD’s uptrend may be accelerated if US consumer confidence comes in below the 131.5 estimate and Fed Chairman Jerome Powell’s economic outlook bolsters market expectations of rate cuts. Overnight index swaps are pricing in a 100 percent probability of a cut from the July meeting through year-end. However, rhetoric from the central bank has not indicated that policymakers are feeling dovish to that degree.
	However, hawkish members of the Fed are finding it increasingly difficult to justify their position in light of US growth. Since February, economic activity out of the US has been broadly underperforming relative to economists’ expectations – signaling that analysts are over estimating the economy’s strength. Inflationary pressure has also been waning alongside a deterioration in global trade due to the US-China trade war.`
	paragraphList := strings.Split(text, "\n")

	result := map[string]interface{}{}
	for _, value := range paragraphList {
		doc, err := prose.NewDocument(value)
		if err != nil {
			log.Fatal(err)
		}
		for _, sent := range doc.Sentences() {
			if strings.Contains(sent.Text, query) {
				result["selectItem"] = sent.Text
				result["selectItemIndex"] = strings.Index(text, sent.Text)
			}
		}
	}
	result["pre"] = text[:result["selectItemIndex"].(int)]
	postStartingIndex := result["selectItemIndex"].(int) + len(fmt.Sprintf("%s", result["selectItem"]))
	result["post"] = text[postStartingIndex:postStartingIndex+150] + "..."
	fmt.Printf("[pre] %s \n\n", result["pre"])
	fmt.Printf("[selectItem] %s \n\n", result["selectItem"])
	fmt.Printf("[post] %s \n\n", result["post"])
	fmt.Printf("[text] %s \n\n", text)
}
