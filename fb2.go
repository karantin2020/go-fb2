package fb2

import (
	b64 "encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	etree "github.com/rupor-github/fb2converter/etree"
)

// fb2 represents FictionBook structure
type fb2 struct {
	sync.Mutex
	data FictionBookScheme
	// srcFileName string
	body       *etree.Element
	annotation *etree.Element
}

var (
	_            FB2 = (*fb2)(nil)
	dateValueFmt     = "2006-01-02"
	dateTextFmt      = "2006"
)

func NewFB2(title string) FB2 {
	v := &fb2{
		body: etree.NewElement("body"),
	}
	v.data.Description.TitleInfo.Author = []AuthorType{}
	v.data.Description.TitleInfo.Genre = []Genre{}
	v.data.Description.TitleInfo.Translator = []AuthorType{}
	v.data.Description.DocumentInfo.Author = []AuthorType{}
	v.data.Description.DocumentInfo.Publisher = []AuthorType{}
	v.body.SetText("\n")
	bt := v.body.CreateElement("title")
	bt.SetText("\n")
	bt.CreateElement("p").SetTail("\n")
	bTitle := bt.CreateElement("p")
	bTitle.SetText(title)
	bTitle.SetTail("\n")
	bt.SetTail("\n")
	v.body.SetTail("\n")
	v.SetIdentifier(uuid.Must(uuid.NewV4()).String())
	v.data.Description.TitleInfo.BookTitle = title
	v.data.Description.PublishInfo.BookName = title
	v.data.Description.DocumentInfo.Date.Value = time.Now().Format(dateValueFmt)
	v.data.Description.DocumentInfo.Date.Text = time.Now().Format(dateTextFmt)
	v.data.Description.DocumentInfo.ProgramUsed = "go-fb2"
	v.data.Description.DocumentInfo.Version = "1.0"
	return v
}

type FB2 interface {
	AddCSS(source string, mime string)
	AddImage(source, internalFilename, mimeType string) (string, error)
	AddSection(body string, sectionTitle string) error
	Title() string
	Author() string
	Description() string
	Identifier() string
	Lang() string
	SetTitle(title string)
	SetAuthor(author AuthorType)
	SetCover(srcName string) error
	SetDescription(desc string) error
	SetIdentifier(identifier string)
	SetLang(lang string)
	WriteToFile(destFilePath string) error
	WriteToString() (string, error)
	Body() *etree.Element
	Data() *FictionBookScheme
}

func (d *fb2) AddCSS(linkHref string, mimeType string) {
	d.Lock()
	defer d.Unlock()
	d.data.stylesheet = append(d.data.stylesheet, FictionBookStylesheet{
		Type:      mimeType,
		XlinkHref: linkHref,
	})
}

func (d *fb2) AddImage(sourcePath, internalFilename, mimeType string) (string, error) {
	d.Lock()
	defer d.Unlock()
	return d.addImage(sourcePath, internalFilename, mimeType)
}

func (d *fb2) addImage(sourcePath, internalFilename, mimeType string) (string, error) {
	b, err := getMedia(sourcePath)
	if err != nil {
		return "", fmt.Errorf("addImage error: %w", err)
	}
	ext := filepath.Ext(sourcePath)
	if internalFilename == "" {
		// internalFilename = filepath.Base(sourcePath)
		internalFilename = fmt.Sprintf(`_image%d%s`, len(d.data.Binary), ext)
	}
	if filepath.Ext(internalFilename) == "" {
		internalFilename += ext
	}

	if mimeType == "" {
		ext = strings.Trim(ext, ".")
		mimeType = "image/jpg"
		if ext != "" {
			mimeType = "image/" + ext
		}
	}
	d.data.Binary = append(d.data.Binary, FictionBookBinary{
		ContentType: mimeType,
		Text:        string(b),
		Id:          internalFilename,
	})

	return internalFilename, nil
}

func (d *fb2) AddSection(body string, sectionTitle string) error {
	d.Lock()
	defer d.Unlock()
	// title := etree.NewElement("title")
	// title.CreateElement("p").SetText(sectionTitle)
	content := etree.NewDocument()
	err := content.ReadFromString(body)
	if err != nil {
		return fmt.Errorf("read section body error: %w", err)
	}
	section := d.body.CreateElement("section")
	section.SetText("\n")
	title := section.CreateElement("title")
	title.CreateElement("p").SetText(sectionTitle)
	title.SetTail("\n")
	elems := content.ChildElements()
	for i := range elems {
		section.AddChild(elems[i].Copy())
	}
	sElems := section.ChildElements()
	if len(sElems) != 0 {
		sElems[len(sElems)-1].SetTail("\n")
	}
	section.SetTail("\n")
	d.body.AddChild(section)
	return nil
}

func (d *fb2) Title() string {
	d.Lock()
	defer d.Unlock()
	return d.data.Description.TitleInfo.BookTitle
}

func (d *fb2) Author() string {
	d.Lock()
	defer d.Unlock()
	authorsString := d.getAuthor()
	return authorsString
}

func (d *fb2) getAuthor() string {
	authors := []string{}
	for _, author := range d.data.Description.TitleInfo.Author {
		authors = append(authors, author.String())
	}
	authorsString := strings.Join(authors, ", ")
	return authorsString
}

func (d *fb2) Description() string {
	d.Lock()
	defer d.Unlock()
	if d.annotation == nil {
		return ""
	}
	doc := etree.NewDocument()
	doc.SetRoot(d.annotation.Copy())
	desc, err := doc.WriteToString()
	if err != nil {
		return ""
	}
	return desc
}

func (d *fb2) Identifier() string {
	d.Lock()
	defer d.Unlock()
	return d.data.Description.DocumentInfo.Id
}

func (d *fb2) Lang() string {
	d.Lock()
	defer d.Unlock()
	return d.data.Description.TitleInfo.Lang
}

func (d *fb2) SetTitle(title string) {
	d.Lock()
	defer d.Unlock()
	d.data.Description.TitleInfo.BookTitle = title
}

func (d *fb2) SetAuthor(author AuthorType) {
	d.Lock()
	defer d.Unlock()
	d.data.Description.TitleInfo.Author = append(d.data.Description.TitleInfo.Author, author)
	bAuthor := d.body.FindElement("./title/p")
	if bAuthor != nil {
		bAuthor.SetText(d.getAuthor())
	}
}

func (d *fb2) SetCover(srcName string) error {
	d.Lock()
	defer d.Unlock()
	// fmt.Println("Start SetCover")
	coverName, err := d.addImage(srcName, "cover", "")
	if err != nil {
		return fmt.Errorf("in SetCover() addImage returned error: %w", err)
	}
	d.data.Description.TitleInfo.Coverpage = append(d.data.Description.TitleInfo.Coverpage, Coverpage{
		Image: &InlineImageType{
			XlinkHref: "#" + coverName,
			Alt:       "Cover",
		},
	})
	return nil
}

func (d *fb2) SetDescription(desc string) error {
	d.Lock()
	defer d.Unlock()
	doc := etree.NewDocument()
	desc = strings.TrimSpace(desc)
	if !strings.HasPrefix(desc, "<p>") {
		desc = fmt.Sprintf(`<p>%s</p>`, desc)
	}
	err := doc.ReadFromString(desc)
	if err != nil {
		return fmt.Errorf("error SetDescription: %w", err)
	}
	d.annotation = doc.Root().Copy()
	return nil
}

func (d *fb2) SetIdentifier(identifier string) {
	d.Lock()
	defer d.Unlock()
	d.data.Description.DocumentInfo.Id = identifier
}

func (d *fb2) SetLang(lang string) {
	d.Lock()
	defer d.Unlock()
	d.data.Description.TitleInfo.Lang = lang
}

func (d *fb2) WriteToFile(destFilePath string) error {
	d.Lock()
	defer d.Unlock()
	book, err := d.WriteToString()
	if err != nil {
		return fmt.Errorf("write to file error: %w", err)
	}
	err = ioutil.WriteFile(destFilePath, []byte(book), 0644)
	if err != nil {
		return fmt.Errorf("write to file error: %w", err)
	}
	return nil
}

func (d *fb2) WriteToString() (string, error) {
	d.Lock()
	defer d.Unlock()
	data, err := xml.MarshalIndent(d.data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("write error: %w", err)
	}
	doc := etree.NewDocument()
	proc := doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	proc.TailData = "\n"
	doc.ReadFromBytes(data)
	for i := range d.data.stylesheet {
		t := d.data.stylesheet[i]
		proc := doc.CreateProcInst("xml-stylesheet",
			fmt.Sprintf(`type="%s" href="%s"`, t.Type, t.XlinkHref))
		proc.TailData = "\n"
	}
	fb := doc.Root()
	fb.CreateAttr("xmlns:l", "http://www.w3.org/1999/xlink")
	fb.CreateAttr("xmlns", "http://www.gribuser.ru/xml/fictionbook/2.0")
	desc := doc.FindElement("//title-info/annotation")
	if desc != nil {
		desc.Child = nil
		desc.AddChild(d.annotation)
	}
	body := doc.FindElement("//body")
	if body != nil {
		*body = *(d.body.Copy())
	}
	bTitle := body.FindElements("//title/p")
	if len(bTitle) < 2 {
		return "", errors.New("invalid body title structure")
	}
	out, err := doc.WriteToString()
	if err != nil {
		return "", fmt.Errorf("write to string error: %w", err)
	}
	return out, nil
}

func (d *fb2) Body() *etree.Element {
	return d.body
}

func (d *fb2) Data() *FictionBookScheme {
	return &d.data
}

func validateFileSource(source string) error {
	// fmt.Println("validateFileSource")
	u, err := url.Parse(source)
	if err != nil {
		return err
	}

	var r io.ReadCloser
	var resp *http.Response
	// If it's a URL
	if u.Scheme == "http" || u.Scheme == "https" {
		resp, err = http.Get(source)
		if err != nil {
			return err
		}
		r = resp.Body

		// Otherwise, assume it's a local file
	} else {
		r, err = os.Open(source)
	}
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	return nil
}

func getMedia(sourcePath string) (string, error) {
	// fmt.Println("Start getMedia")
	err := validateFileSource(sourcePath)
	if err != nil {
		return "", errors.New("invalid sourcePath name")
	}

	u, err := url.Parse(sourcePath)
	if err != nil {
		return "", errors.New("error parse sourcePath as url")
	}

	var r io.ReadCloser
	var resp *http.Response
	// If it's a URL
	if u.Scheme == "http" || u.Scheme == "https" {
		resp, err = http.Get(sourcePath)
		if err != nil {
			return "", errors.New("error get url sourcePath")
		}
		r = resp.Body
		// Otherwise, assume it's a local file
	} else {
		r, err = os.Open(sourcePath)
	}
	defer r.Close()
	if err != nil {
		return "", errors.New("error getting source")
	}
	b, err := ioutil.ReadAll(r)
	uEnc := b64.StdEncoding.EncodeToString([]byte(b))
	if err != nil {
		return "", errors.New("error reading source")
	}

	return string(uEnc), nil
}
