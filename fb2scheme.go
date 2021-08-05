package fb2

import (
	"encoding/xml"
	"fmt"
)

// FictionBookScheme scheme
type FictionBookScheme struct {
	XMLName xml.Name `xml:"FictionBook"`

	stylesheet []FictionBookStylesheet

	Description FictionBookDescription `xml:"description"`

	Body string `xml:"body"`

	Binary []FictionBookBinary `xml:"binary,omitempty"`
}

// FictionBookStylesheet struct
type FictionBookStylesheet struct {
	XMLName xml.Name `xml:"stylesheet"`

	Type string `xml:"type,attr"`

	XlinkHref string `xml:"href,attr"`

	Text string `xml:",chardata"`
}

// Element
type DescriptionDocumentInfo struct {
	XMLName xml.Name `xml:"document-info"`

	Author []AuthorType `xml:"author"`

	ProgramUsed string `xml:"program-used,omitempty"`

	Date DateType `xml:"date"`

	SrcUrl []string `xml:"src-url,omitempty"`

	SrcOcr string `xml:"src-ocr"`

	Id string `xml:"id"`

	Version string `xml:"version"`

	History AnnotationType `xml:"history"`

	Publisher []AuthorType `xml:"publisher"`
}

// Element
type DescriptionPublishInfo struct {
	XMLName xml.Name `xml:"publish-info"`

	BookName string `xml:"book-name"`

	Publisher string `xml:"publisher"`

	City string `xml:"city"`

	Year string `xml:"year"`

	Isbn TextFieldType `xml:"isbn"`

	Sequence SequenceType `xml:"sequence,omitempty"`
}

// Element
type DescriptionCustomInfo struct {
	XMLName xml.Name `xml:"custom-info"`

	InfoType string `xml:"info-type,attr"`

	XmlLang string `xml:"lang,attr,omitempty"`

	Text string `xml:",chardata"`
}

// Element
type FictionBookDescription struct {
	XMLName xml.Name `xml:"description"`

	TitleInfo TitleInfoType `xml:"title-info"`

	SrcTitleInfo *TitleInfoType `xml:"src-title-info,omitempty"`

	DocumentInfo DescriptionDocumentInfo `xml:"document-info"`

	PublishInfo DescriptionPublishInfo `xml:"publish-info"`

	CustomInfo []DescriptionCustomInfo `xml:"custom-info,omitempty"`

	Output []ShareInstructionType `xml:"output,omitempty"`
}

// Element
type FictionBookBinary struct {
	XMLName xml.Name `xml:"binary"`

	ContentType string `xml:"content-type,attr"`

	Id string `xml:"id,attr"`

	Text string `xml:",chardata"`
}

// Element
type Stanza struct {
	XMLName xml.Name `xml:"stanza"`

	XmlLang string `xml:"lang,attr,omitempty"`

	Title *TitleType `xml:"title,omitempty"`

	Subtitle *PType `xml:"subtitle,omitempty"`

	V []PType `xml:"v,omitempty"`
}

// Element
type Tr struct {
	XMLName xml.Name `xml:"tr"`

	Align AlignType `xml:"align,attr,omitempty"`

	Th []TdType `xml:"th,omitempty"`

	Td []TdType `xml:"td,omitempty"`
}

// Element
type Genre struct {
	XMLName xml.Name `xml:"genre"`

	Match int64 `xml:"match,attr,omitempty"`

	Text string `xml:",chardata"`
}

// Element
type Coverpage struct {
	XMLName xml.Name `xml:"coverpage"`

	Image *InlineImageType `xml:"image"`
}

// XSD ComplexType declarations

type BodyType string

type AuthorType struct {
	XMLName xml.Name `xml:"author"`

	FirstName string `xml:"first-name"`

	MiddleName string `xml:"middle-name"`

	LastName string `xml:"last-name"`

	Nickname string `xml:"nickname"`

	HomePage []string `xml:"home-page"`

	Email []string `xml:"email"`

	Id string `xml:"id"`
}

func (a *AuthorType) String() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

type TextFieldType struct {
	XmlLang string `xml:"lang,attr,omitempty"`

	Text string `xml:",chardata"`
}

type DateType struct {
	Value string `xml:"value,attr"`

	XmlLang string `xml:"lang,attr,omitempty"`

	Text string `xml:",chardata"`
}

type TitleType struct {
	XmlLang string `xml:"lang,attr,omitempty"`

	P []PType `xml:"p,omitempty"`

	EmptyLine []string `xml:"empty-line,omitempty"`
}

type ImageType struct {
	XlinkHref string `xml:"href,attr"`

	Alt string `xml:"alt,attr,omitempty"`

	Title string `xml:"title,attr,omitempty"`

	Id string `xml:"id,attr"`
}

type PType struct {
	Id string `xml:"id,attr,omitempty"`

	Style string `xml:"style,attr,omitempty"`

	Text string `xml:",chardata"`

	XmlLang string `xml:"lang,attr,omitempty"`

	Strong []StyleType `xml:"strong,omitempty"`

	Emphasis []StyleType `xml:"emphasis,omitempty"`

	StyleElm []NamedStyleType `xml:"style,omitempty"`

	A []LinkType `xml:"a,omitempty"`

	Strikethrough []StyleType `xml:"strikethrough,omitempty"`

	Sub []StyleType `xml:"sub,omitempty"`

	Sup []StyleType `xml:"sup,omitempty"`

	Code []StyleType `xml:"code,omitempty"`

	Image []InlineImageType `xml:"image,omitempty"`
}

type CiteType struct {
	Id string `xml:"id,attr,omitempty"`

	XmlLang string `xml:"lang,attr,omitempty"`

	TextAuthor []PType `xml:"text-author,omitempty"`

	P []PType `xml:"p,omitempty"`

	Poem []PoemType `xml:"poem,omitempty"`

	EmptyLine []string `xml:"empty-line,omitempty"`

	Subtitle []PType `xml:"subtitle,omitempty"`

	Table []TableType `xml:"table,omitempty"`

	Text string `xml:",chardata"`
}

type PoemType struct {
	Id string `xml:"id,attr,omitempty"`

	XmlLang string `xml:"lang,attr,omitempty"`

	Title *TitleType `xml:"title,omitempty"`

	Epigraph []EpigraphType `xml:"epigraph,omitempty"`

	TextAuthor []PType `xml:"text-author,omitempty"`

	Date *DateType `xml:"date,omitempty"`

	Subtitle []PType `xml:"subtitle,omitempty"`

	Stanza []Stanza `xml:"stanza,omitempty"`

	Text string `xml:",chardata"`
}

type EpigraphType struct {
	Id string `xml:"id,attr,omitempty"`

	TextAuthor []PType `xml:"text-author,omitempty"`

	P []PType `xml:"p,omitempty"`

	Poem []PoemType `xml:"poem,omitempty"`

	Cite []CiteType `xml:"cite,omitempty"`

	EmptyLine []string `xml:"empty-line,omitempty"`

	Text string `xml:",chardata"`
}

type AnnotationType struct {
	Id string `xml:"id,attr,omitempty"`

	XmlLang string `xml:"lang,attr,omitempty"`

	P []PType `xml:"p,omitempty"`

	Poem []PoemType `xml:"poem,omitempty"`

	Cite []CiteType `xml:"cite,omitempty"`

	Subtitle []PType `xml:"subtitle,omitempty"`

	Table []TableType `xml:"table,omitempty"`

	EmptyLine []string `xml:"empty-line,omitempty"`

	Text string `xml:",chardata"`
}

type SectionType string

type StyleType struct {
	XmlLang string `xml:"lang,attr,omitempty"`

	Strong []StyleType `xml:"strong,omitempty"`

	Emphasis []StyleType `xml:"emphasis,omitempty"`

	Style []NamedStyleType `xml:"style,omitempty"`

	A []LinkType `xml:"a,omitempty"`

	Strikethrough []StyleType `xml:"strikethrough,omitempty"`

	Sub []StyleType `xml:"sub,omitempty"`

	Sup []StyleType `xml:"sup,omitempty"`

	Code []StyleType `xml:"code,omitempty"`

	Image []InlineImageType `xml:"image,omitempty"`

	Text string `xml:",chardata"`
}

type NamedStyleType struct {
	XmlLang string `xml:"lang,attr,omitempty"`

	Name string `xml:"name,attr"`

	Strong []StyleType `xml:"strong,omitempty"`

	Emphasis []StyleType `xml:"emphasis,omitempty"`

	Style []NamedStyleType `xml:"style,omitempty"`

	A []LinkType `xml:"a,omitempty"`

	Strikethrough []StyleType `xml:"strikethrough,omitempty"`

	Sub []StyleType `xml:"sub,omitempty"`

	Sup []StyleType `xml:"sup,omitempty"`

	Code []StyleType `xml:"code,omitempty"`

	Image []InlineImageType `xml:"image,omitempty"`

	Text string `xml:",chardata"`
}

type LinkType struct {
	XlinkHref string `xml:"href,attr"`

	Type string `xml:"type,attr,omitempty"`

	Strong []StyleLinkType `xml:"strong,omitempty"`

	Emphasis []StyleLinkType `xml:"emphasis,omitempty"`

	Style []StyleLinkType `xml:"style,omitempty"`

	Strikethrough []StyleLinkType `xml:"strikethrough,omitempty"`

	Sub []StyleLinkType `xml:"sub,omitempty"`

	Sup []StyleLinkType `xml:"sup,omitempty"`

	Code []StyleLinkType `xml:"code,omitempty"`

	Image []InlineImageType `xml:"image,omitempty"`

	Text string `xml:",chardata"`
}

type StyleLinkType struct {
	Strong []StyleLinkType `xml:"strong,omitempty"`

	Emphasis []StyleLinkType `xml:"emphasis,omitempty"`

	Style []StyleLinkType `xml:"style,omitempty"`

	Strikethrough []StyleLinkType `xml:"strikethrough,omitempty"`

	Sub []StyleLinkType `xml:"sub,omitempty"`

	Sup []StyleLinkType `xml:"sup,omitempty"`

	Code []StyleLinkType `xml:"code,omitempty"`

	Image []InlineImageType `xml:"image,omitempty"`

	Text string `xml:",chardata"`
}

type SequenceType struct {
	Name string `xml:"name,attr,omitempty"`

	Number int64 `xml:"number,attr,omitempty"`

	XmlLang string `xml:"lang,attr,omitempty"`
}

type TableType struct {
	Style string `xml:"style,attr,omitempty"`

	Id string `xml:"id,attr,omitempty"`

	Tr []Tr `xml:",any"`
}

type TitleInfoType struct {
	Genre []Genre `xml:"genre,omitempty"`

	Author []AuthorType `xml:"author"`

	BookTitle string `xml:"book-title"`

	Annotation AnnotationType `xml:"annotation"`

	Keywords string `xml:"keywords"`

	Date DateType `xml:"date"`

	Coverpage []Coverpage `xml:"coverpage"`

	Lang string `xml:"lang"`

	SrcLang string `xml:"src-lang,omitempty"`

	Translator []AuthorType `xml:"translator,omitempty"`

	Sequence SequenceType `xml:"sequence,omitempty"`
}

type ShareInstructionType struct {
	Mode ShareModesType `xml:"mode,attr,omitempty"`

	IncludeAll DocGenerationInstructionType `xml:"include-all,attr,omitempty"`

	Price float64 `xml:"price,attr,omitempty"`

	Currency string `xml:"currency,attr,omitempty"`

	Part []PartShareInstructionType `xml:"part,omitempty"`

	OutputDocumentClass []OutPutDocumentType `xml:"output-document-class,omitempty"`
}

type PartShareInstructionType struct {
	XlinkHref string `xml:"href,attr,omitempty"`

	Include DocGenerationInstructionType `xml:"include,attr,omitempty"`
}

type OutPutDocumentType struct {
	Name string `xml:"name,attr,omitempty"`

	Create DocGenerationInstructionType `xml:"create,attr,omitempty"`

	Price float64 `xml:"price,attr,omitempty"`

	Part PartShareInstructionType `xml:",any"`
}

type TdType struct {
	Id string `xml:"id,attr,omitempty"`

	Style string `xml:"style,attr,omitempty"`

	Colspan int64 `xml:"colspan,attr,omitempty"`

	Rowspan int64 `xml:"rowspan,attr,omitempty"`

	Align AlignType `xml:"align,attr,omitempty"`

	Valign VAlignType `xml:"valign,attr,omitempty"`

	XmlLang string `xml:"lang,attr,omitempty"`

	Strong []StyleType `xml:"strong,omitempty"`

	Emphasis []StyleType `xml:"emphasis,omitempty"`

	StyleElm []NamedStyleType `xml:"style,omitempty"`

	A []LinkType `xml:"a,omitempty"`

	Strikethrough []StyleType `xml:"strikethrough,omitempty"`

	Sub []StyleType `xml:"sub,omitempty"`

	Sup []StyleType `xml:"sup,omitempty"`

	Code []StyleType `xml:"code,omitempty"`

	Image []InlineImageType `xml:"image,omitempty"`

	Text string `xml:",chardata"`
}

type InlineImageType struct {
	XMLName xml.Name `xml:"image"`

	XlinkHref string `xml:"l:href,attr"`

	Alt string `xml:"alt,attr,omitempty"`
}

// XSD SimpleType declarations

type AlignType string

const AlignTypeLeft AlignType = "left"

const AlignTypeRight AlignType = "right"

const AlignTypeCenter AlignType = "center"

type VAlignType string

const VAlignTypeTop VAlignType = "top"

const VAlignTypeMiddle VAlignType = "middle"

const VAlignTypeBottom VAlignType = "bottom"

type ShareModesType string

const ShareModesTypeFree ShareModesType = "free"

const ShareModesTypePaid ShareModesType = "paid"

type DocGenerationInstructionType string

const DocGenerationInstructionTypeRequire DocGenerationInstructionType = "require"

const DocGenerationInstructionTypeAllow DocGenerationInstructionType = "allow"

const DocGenerationInstructionTypeDeny DocGenerationInstructionType = "deny"
