package fb2

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	etree "github.com/rupor-github/fb2converter/etree"
)

func Test_fb2_AddSection(t *testing.T) {
	type fields struct {
		FB2
	}
	type args struct {
		body         string
		sectionTitle string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				FB2: NewFB2("Test1"),
			},
			args: args{
				sectionTitle: "Test1Section",
				body: `<p>Chapter text.</p>
<p><strong>Strong text.</strong></p>
<subtitle>* * *</subtitle>
<p>Published on: <a l:href="https://g.ve/test">https://g.ve/test</a></p>`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.FB2.AddSection(tt.args.body, tt.args.sectionTitle); (err != nil) != tt.wantErr {
				t.Errorf("fb2.AddSection() error = %v, wantErr %v", err, tt.wantErr)
			}
			doc := etree.NewDocument()
			doc.SetRoot(tt.fields.FB2.Body())
			// doc.Indent(2)
			got, err := doc.WriteToString()
			if err != nil {
				t.Errorf("NewFB2() got err: %v", err)
			}
			fmt.Println(got)
		})
	}
}

func TestNewFB2(t *testing.T) {
	type args struct {
		title string
	}
	type want struct {
		title string
		body  string
	}
	tests := []struct {
		name string
		args args
		want
	}{
		// Test data
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewFB2(tt.args.title)
			doc := etree.NewDocument()
			doc.SetRoot(v.Body())
			got, err := doc.WriteToString()
			if err != nil {
				t.Errorf("NewFB2() got err: %v", err)
			}
			if tt.want.body != got {
				t.Errorf("NewFB2() body want: %s, got err: %s", tt.want.body, got)
			}
			if tt.want.title != v.Data().Description.TitleInfo.BookTitle {
				t.Errorf("NewFB2() title want: %s, got err: %s", tt.want.title,
					v.Data().Description.TitleInfo.BookTitle)
			}
			fmt.Println(got)
		})
	}
}

func Test_fb2_AddImage(t *testing.T) {
	type fields struct {
		b FB2
	}
	type args struct {
		sourcePath       string
		internalFilename string
		mimeType         string
	}
	b, err := os.ReadFile("./testdata/avatar.base64")
	if err != nil {
		t.Errorf("fb2.AddImage() read test base64 file error %v", err)
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantName string
		wantBin  string
		wantErr  bool
	}{
		{
			name: "Test1",
			fields: fields{
				b: NewFB2("Test1Title"),
			},
			args: args{
				sourcePath:       "./testdata/avatar.jpeg",
				internalFilename: "",
				mimeType:         "",
			},
			wantName: "_image0.jpeg",
			wantBin:  string(b),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.b
			got, err := d.AddImage(tt.args.sourcePath, tt.args.internalFilename, tt.args.mimeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("fb2.AddImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantName {
				t.Errorf("fb2.AddImage() = %v, wantName %v", got, tt.wantName)
			}
			if len(d.Data().Binary) != 1 {
				t.Errorf("fb2.AddImage() bin length is: %d", len(d.Data().Binary))
			}
			if d.Data().Binary[0].Text != tt.wantBin {
				t.Errorf("fb2.AddImage() wrong binary content")
				fmt.Printf("d.Data().Binary[0].Text: %v\n", d.Data().Binary[0].Text)
			}
		})
	}
}

func Test_fb2_WriteToString(t *testing.T) {
	type fields struct {
		b FB2
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Test1",
			fields: fields{
				b: NewFB2("Test1Title"),
			},
			want:    `<?xml version="1.0" encoding="UTF-8"?>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.b
			if d == nil {
				t.Errorf("fb2.WriteToString() nil interface FB2")
			}
			sectionTitle := "Test1Section"
			body := `<p>Chapter text.</p>
<p><strong>Strong text.</strong></p>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Stulti autem malorum memoria torquentur, sapientes bona praeterita grata recordatione renovata delectant. Hanc quoque iucunditatem, si vis, transfer in animum; Duo Reges: constructio interrete. Quae qui non vident, nihil umquam magnum ac cognitione dignum amaverunt. Sed erat aequius Triarium aliquid de dissensione nostra iudicare. Igitur neque stultorum quisquam beatus neque sapientium non beatus. Negat enim summo bono afferre incrementum diem. Quo plebiscito decreta a senatu est consuli quaestio Cn. Iubet igitur nos Pythius Apollo noscere nosmet ipsos. Explanetur igitur. </p>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Stulti autem malorum memoria torquentur, sapientes bona praeterita grata recordatione renovata delectant. Hanc quoque iucunditatem, si vis, transfer in animum; Duo Reges: constructio interrete. Quae qui non vident, nihil umquam magnum ac cognitione dignum amaverunt. Sed erat aequius Triarium aliquid de dissensione nostra iudicare. Igitur neque stultorum quisquam beatus neque sapientium non beatus. Negat enim summo bono afferre incrementum diem. Quo plebiscito decreta a senatu est consuli quaestio Cn. Iubet igitur nos Pythius Apollo noscere nosmet ipsos. Explanetur igitur. </p>
<subtitle>* * *</subtitle>
<p>Published on: <a l:href="https://g.ve/test">https://g.ve/test</a></p>`
			// iName, err := d.AddImage("./testdata/AirPlane_400x600.jpg", "cover.jpg", "")
			// if err != nil {
			// 	t.Errorf("fb2.WriteToString() AddImage error = %v", err)

			// }
			err := d.SetCover("./testdata/AirPlane_400x600.jpg")
			if err != nil {
				t.Errorf("fb2.WriteToString() SetCover error = %v", err)
			}
			d.SetAuthor(AuthorType{
				FirstName: "TestFirstName",
				LastName:  "TestLastName",
			})
			d.SetDescription(`
			Первородная сущность Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut alios omittam, hunc appello, quem ille unum secutus est. Ut placet, inquit, etsi enim illud erat aptius, aequum cuique concedere. Quae quo sunt excelsiores, eo dant clariora indicia naturae. Sin tantum modo ad indicia veteris memoriae cognoscenda, curiosorum. Duo Reges: constructio interrete. Claudii libidini, qui tum erat summo ne imperio, dederetur. 
Quorum sine causa fieri nihil putandum est. Ita multa dicunt, quae vix intellegam. Nunc reliqua videamus, nisi aut ad haec, Cato, dicere aliquid vis aut nos iam longiores sumus. Non enim solum Torquatus dixit quid sentiret, sed etiam cur. Iam id ipsum absurdum, maximum malum neglegi. Sapientem locupletat ipsa natura, cuius divitias Epicurus parabiles esse docuit. Ait enim se, si uratur, Quam hoc suave! dicturum. Scaevola tribunus plebis ferret ad plebem vellentne de ea re quaeri. 
Teneo, inquit, finem illi videri nihil dolere. Tu enim ista lenius, hic Stoicorum more nos vexat. Cur deinde Metrodori liberos commendas? Quae autem natura suae primae institutionis oblita est? 
Apud ceteros autem philosophos, qui quaesivit aliquid, tacet; Satis est tibi in te, satis in legibus, satis in mediocribus amicitiis praesidii. Egone non intellego, quid sit don Graece, Latine voluptas? Sed quae tandem ista ratio est? Non est igitur voluptas bonum. Nummus in Croesi divitiis obscuratur, pars est tamen divitiarum. Et harum quidem rerum facilis est et expedita distinctio. 

`)
			d.AddSection(body, sectionTitle+"1")
			d.AddSection(body, sectionTitle+"2")
			got, err := d.WriteToString()
			if err != nil {
				t.Errorf("fb2.WriteToString() error = %v", err)
			}
			// if got != tt.want {
			// 	t.Errorf("fb2.WriteToString() = %v, want %v", got, tt.want)
			// }
			fmt.Println(got)
			err = ioutil.WriteFile("test_out.fb2", []byte(got), 0644)
			if err != nil {
				t.Errorf("open file error: %v", err)
			}
			// f, err := os.Create("test_out.fb2")
			// if err != nil {
			// 	t.Errorf("open file error: %v", err)
			// }
			// fmt.Fprintln(f, got)
			// f.Close()
		})
	}
}

func Test_fb2_Description(t *testing.T) {
	type fields struct {
		b FB2
	}
	testB := NewFB2("Test1Title")
	testB.SetDescription(`<p>Hello, <p>Hello, World</p>World</p>`)
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test1 Positive",
			fields: fields{
				testB,
			},
			want: `<p>Hello, <p>Hello, World</p>World</p>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.b
			// d.Description()
			if got := d.Description(); got != tt.want {
				t.Errorf("fb2.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fb2_SetDescription(t *testing.T) {
	type fields struct {
		b FB2
	}
	type args struct {
		desc string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test1 positive",
			fields: fields{
				NewFB2("Test1Title"),
			},
			args: args{
				`<p>Hello, <p>Hello, World</p>World</p>`,
			},
			wantErr: false,
		},
		{
			name: "Test1 positive with invalid input xml",
			fields: fields{
				NewFB2("Test1Title"),
			},
			args: args{
				`Negative Hello, <p>Hello, World</p>World</p>`,
			},
			wantErr: false,
		},
		{
			name: "Test1 negative with invalid input xml",
			fields: fields{
				NewFB2("Test1Title"),
			},
			args: args{
				`Negative Hello, Hello, World</p>World</p>`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.b
			if err := d.SetDescription(tt.args.desc); (err != nil) != tt.wantErr {
				fmt.Println(d.Description())
				t.Errorf("fb2.SetDescription() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("error: %v\n", err)
				fmt.Printf("description: %v\n", d.Description())
			}
		})
	}
}
