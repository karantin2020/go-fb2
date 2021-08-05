package fb2test

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	etree "github.com/rupor-github/fb2converter/etree"
)

func readFile(t *testing.T) []byte {
	t.Helper()
	f, err := os.ReadFile("./test1.fb2")
	if err != nil {
		t.Errorf("go-fb2: read test file error: %s\n", err)
	}
	return f
}

func TestDocument(t *testing.T) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile("./test1.fb2")
	if err != nil {
		t.Errorf("go-fb2: read file error: %s\n", err)
	}
	el := doc.FindElement("./FictionBook/body/section/title")
	if el != nil {
		doc1 := etree.NewDocument()
		doc1.AddChild(el.Copy())
		// fmt.Printf("%#v\n", *(el))
		doc1.WriteTo(os.Stdout)
		fmt.Println("")
	} else {
		fmt.Println("no match element")
	}
}

func TestAddChild(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="style.xsl"?>
<People>
  <!--These are all known people-->
  <Person name="Jon O&apos;Reilly"/>
  <Person name="Sally">
  </Person>
</People>`
	addData := `
    <bookstore>
      <book>
        <title>Great Expectations</title>
        <author>Charles Dickens</author>
      </book>
      <book>
        <title>Ulysses</title>
        <author>James Joyce</author>
      </book>
    </bookstore>`
	type PersonType struct {
		XMLName xml.Name `xml:"Person"`

		Name string `xml:"name,attr"`
	}
	type People struct {
		Person []PersonType `xml:"Person,omitempty"`
	}
	v := People{}
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		t.Errorf("go-fb2: unmarshal file error: %s\n", err)
	}
	// fmt.Printf("%+v\n", v)
	_ = addData
	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Errorf("go-fb2: unmarshal error: %s\n", err)
	}
	// fmt.Println(string(output))
	doc := etree.NewDocument()
	err = doc.ReadFromBytes(output)
	if err != nil {
		t.Errorf("go-fb2: read doc error: %s\n", err)
	}
	// doc.Indent(2)
	// _, err = doc.WriteTo(os.Stdout)
	// if err != nil {
	// 	t.Errorf("go-fb2: write error: %s\n", err)
	// }
	// fmt.Println("")
	el := doc.FindElement("//Person[@name='Sally']")
	// fmt.Printf("%#v\n", *el)
	doc1 := etree.NewDocument()
	err = doc1.ReadFromString(addData)
	if err != nil {
		t.Errorf("go-fb2: read doc error: %s\n", err)
	}
	// doc1.Indent(2)
	if el != nil {
		el.Child = nil
		el.SetText(`
    `)
		el.AddChild(doc1.FindElement("./bookstore").Copy().SetTail("\n  "))
	}
	// doc.Indent(2)
	s, err := doc.WriteToString()
	if err != nil {
		t.Errorf("go-fb2: write error: %s\n", err)
	}
	fmt.Println(s)
}

func TestPrintRoot(t *testing.T) {
	v := `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="style.xsl"?>
<People>
  <!--These are all known people-->
  <Person name="Jon O&apos;Reilly"/>
  <Person name="Sally"/>
</People>`
	doc := etree.NewDocument()
	doc.ReadFromString(v)
	out, _ := doc.WriteToString()
	fmt.Println(out)
	doc1 := etree.NewDocument()
	doc1.SetRoot(doc.Root().Copy())
	out, _ = doc1.WriteToString()
	fmt.Println(out)

}
