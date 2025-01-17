//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package vmldrawing

import (
	_ddc "encoding/xml"
	_d "fmt"
	_a "github.com/sanhuanshisanshao/unioffice"
	_f "github.com/sanhuanshisanshao/unioffice/schema/soo/ofc/sharedTypes"
	_c "github.com/sanhuanshisanshao/unioffice/schema/urn/schemas_microsoft_com/office/excel"
	_ac "github.com/sanhuanshisanshao/unioffice/schema/urn/schemas_microsoft_com/vml"
	_dg "strconv"
	_dd "strings"
)

// SetHeight set height of shape.
func (_bed *ShapeStyle) SetHeight(height int64) { _bed._ff = height }

// IsItalic returns true if text is italic.
func (_abeb *TextpathStyle) IsItalic() bool { return _abeb._cd }

// Height return height of shape.
func (_edg *ShapeStyle) Height() int64 { return _edg._ff }

// Width return width of shape.
func (_ef *ShapeStyle) Width() int64 { return _ef._cc }

// SetItalic sets text to italic.
func (_dbb *TextpathStyle) SetItalic(italic bool) { _dbb._cd = italic }

// CreateFormula creates F element for typeFormulas.
func CreateFormula(s string) *_ac.CT_F { _fcg := _ac.NewCT_F(); _fcg.EqnAttr = &s; return _fcg }

type Container struct {
	Layout    *_ac.OfcShapelayout
	ShapeType *_ac.Shapetype
	Shape     []*_ac.Shape
}

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing() *Container {
	_b := NewContainer()
	_b.Layout = _ac.NewOfcShapelayout()
	_b.Layout.ExtAttr = _ac.ST_ExtEdit
	_b.Layout.Idmap = _ac.NewOfcCT_IdMap()
	_b.Layout.Idmap.DataAttr = _a.String("\u0031")
	_b.Layout.Idmap.ExtAttr = _ac.ST_ExtEdit
	_b.ShapeType = _ac.NewShapetype()
	_b.ShapeType.IdAttr = _a.String("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_b.ShapeType.CoordsizeAttr = _a.String("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030")
	_b.ShapeType.SptAttr = _a.Float32(202)
	_b.ShapeType.PathAttr = _a.String("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065")
	_ce := _ac.NewEG_ShapeElements()
	_b.ShapeType.EG_ShapeElements = append(_b.ShapeType.EG_ShapeElements, _ce)
	_ce.Path = _ac.NewPath()
	_ce.Path.GradientshapeokAttr = _f.ST_TrueFalseT
	_ce.Path.ConnecttypeAttr = _ac.OfcST_ConnectTypeRect
	return _b
}

// IsBold returns true if text is bold.
func (_fa *TextpathStyle) IsBold() bool { return _fa._db }

// SetBold sets text to bold.
func (_dgd *TextpathStyle) SetBold(bold bool) { _dgd._db = bold }
func NewContainer() *Container                { return &Container{} }

// NewTextpathStyle accept value of string style attribute of element v:textpath and format it to generate TextpathStyle.
func NewTextpathStyle(style string) TextpathStyle {
	_gb := TextpathStyle{_gd: "\u0022C\u0061\u006c\u0069\u0062\u0072\u0069\"", _gdg: 44, _db: false, _cd: false}
	_dae := _dd.Split(style, "\u003b")
	for _, _fg := range _dae {
		_gag := _dd.Split(_fg, "\u003a")
		if len(_gag) != 2 {
			continue
		}
		switch _gag[0] {
		case "f\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069\u006c\u0079":
			_gb._gd = _gag[1]
			break
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_gb._gdg, _ = _dg.ParseInt(_dd.ReplaceAll(_gag[1], "\u0070\u0074", ""), 10, 64)
			break
		case "f\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074":
			_gb._db = _gag[1] == "\u0062\u006f\u006c\u0064"
			break
		case "\u0066\u006f\u006e\u0074\u002d\u0073\u0074\u0079\u006c\u0065":
			_gb._cd = _gag[1] == "\u0069\u0074\u0061\u006c\u0069\u0063"
			break
		}
	}
	return _gb
}

// TextpathStyle is style attribute of element v:textpath.
type TextpathStyle struct {
	_gd  string
	_gdg int64
	_db  bool
	_cd  bool
}

// ToString formatting ShapeStyle to string.
func (_ca *ShapeStyle) String() string {
	_ddf := ""
	_ddf += _d.Sprintf("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0025\u0073\u003b", _ca._eec)
	_ddf += _d.Sprintf("\u006da\u0072g\u0069\u006e\u002d\u006c\u0065\u0066\u0074\u003a\u0025\u0064\u003b", _ca._fb)
	_ddf += _d.Sprintf("\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006fp\u003a\u0025\u0064\u003b", _ca._cbb)
	_ddf += _d.Sprintf("w\u0069\u0064\u0074\u0068\u003a\u0025\u0064\u0070\u0074\u003b", _ca._cc)
	_ddf += _d.Sprintf("\u0068\u0065\u0069g\u0068\u0074\u003a\u0025\u0064\u0070\u0074\u003b", _ca._ff)
	_ddf += _d.Sprintf("z\u002d\u0069\u006e\u0064\u0065\u0078\u003a\u0025\u0064\u003b", _ca._fe)
	_ddf += _d.Sprintf("m\u0073\u006f\u002d\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069z\u006f\u006e\u0074a\u006c:\u0025\u0073\u003b", _ca._be)
	_ddf += _d.Sprintf("\u006d\u0073o-\u0070\u006f\u0073i\u0074\u0069\u006f\u006e-ho\u0072iz\u006f\u006e\u0074\u0061\u006c\u002d\u0072el\u0061\u0074\u0069\u0076\u0065\u003a\u0025s\u003b", _ca._eb)
	_ddf += _d.Sprintf("\u006ds\u006f\u002d\u0070\u006fs\u0069\u0074\u0069\u006f\u006e-\u0076e\u0072t\u0069\u0063\u0061\u006c\u003a\u0025\u0073;", _ca._fcc)
	_ddf += _d.Sprintf("\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076e\u0072t\u0069c\u0061l\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065\u003a\u0025\u0073\u003b", _ca._baa)
	return _ddf
}
func (_fc *Container) MarshalXML(e *_ddc.Encoder, start _ddc.StartElement) error {
	start.Attr = append(start.Attr, _ddc.Attr{Name: _ddc.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0076"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"})
	start.Attr = append(start.Attr, _ddc.Attr{Name: _ddc.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006f"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"})
	start.Attr = append(start.Attr, _ddc.Attr{Name: _ddc.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Name.Local = "\u0078\u006d\u006c"
	e.EncodeToken(start)
	if _fc.Layout != nil {
		_dgb := _ddc.StartElement{Name: _ddc.Name{Local: "\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}}
		e.EncodeElement(_fc.Layout, _dgb)
	}
	if _fc.ShapeType != nil {
		_ad := _ddc.StartElement{Name: _ddc.Name{Local: "v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}}
		e.EncodeElement(_fc.ShapeType, _ad)
	}
	for _, _ae := range _fc.Shape {
		_cbe := _ddc.StartElement{Name: _ddc.Name{Local: "\u0076:\u0073\u0068\u0061\u0070\u0065"}}
		e.EncodeElement(_ae, _cbe)
	}
	return e.EncodeToken(_ddc.EndElement{Name: start.Name})
}

// SetFontFamily sets text's fontFamily.
func (_cfe *TextpathStyle) SetFontFamily(fontFamily string) { _cfe._gd = fontFamily }

// FontSize returns fontSize of the text.
func (_cg *TextpathStyle) FontSize() int64 { return _cg._gdg }

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape(col, row int64) *_ac.Shape {
	_acb := _ac.NewShape()
	_acb.IdAttr = _a.String(_d.Sprintf("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064", col, row))
	_acb.TypeAttr = _a.String("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_acb.StyleAttr = _a.String("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e")
	_acb.FillcolorAttr = _a.String("\u0023f\u0062\u0066\u0036\u0064\u0036")
	_acb.StrokecolorAttr = _a.String("\u0023e\u0064\u0065\u0061\u0061\u0031")
	_dgc := _ac.NewEG_ShapeElements()
	_dgc.Fill = _ac.NewFill()
	_dgc.Fill.Color2Attr = _a.String("\u0023f\u0062\u0066\u0065\u0038\u0032")
	_dgc.Fill.AngleAttr = _a.Float64(-180)
	_dgc.Fill.TypeAttr = _ac.ST_FillTypeGradient
	_dgc.Fill.Fill = _ac.NewOfcFill()
	_dgc.Fill.Fill.ExtAttr = _ac.ST_ExtView
	_dgc.Fill.Fill.TypeAttr = _ac.OfcST_FillTypeGradientUnscaled
	_acb.EG_ShapeElements = append(_acb.EG_ShapeElements, _dgc)
	_bc := _ac.NewEG_ShapeElements()
	_bc.Shadow = _ac.NewShadow()
	_bc.Shadow.OnAttr = _f.ST_TrueFalseT
	_bc.Shadow.ObscuredAttr = _f.ST_TrueFalseT
	_acb.EG_ShapeElements = append(_acb.EG_ShapeElements, _bc)
	_cb := _ac.NewEG_ShapeElements()
	_cb.Path = _ac.NewPath()
	_cb.Path.ConnecttypeAttr = _ac.OfcST_ConnectTypeNone
	_acb.EG_ShapeElements = append(_acb.EG_ShapeElements, _cb)
	_aca := _ac.NewEG_ShapeElements()
	_aca.Textbox = _ac.NewTextbox()
	_aca.Textbox.StyleAttr = _a.String("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f")
	_acb.EG_ShapeElements = append(_acb.EG_ShapeElements, _aca)
	_g := _ac.NewEG_ShapeElements()
	_g.ClientData = _c.NewClientData()
	_g.ClientData.ObjectTypeAttr = _c.ST_ObjectTypeNote
	_g.ClientData.MoveWithCells = _f.ST_TrueFalseBlankT
	_g.ClientData.SizeWithCells = _f.ST_TrueFalseBlankT
	_g.ClientData.Anchor = _a.String("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033")
	_g.ClientData.AutoFill = _f.ST_TrueFalseBlankFalse
	_g.ClientData.Row = _a.Int64(row)
	_g.ClientData.Column = _a.Int64(col)
	_acb.EG_ShapeElements = append(_acb.EG_ShapeElements, _g)
	return _acb
}
func (_edd *Container) UnmarshalXML(d *_ddc.Decoder, start _ddc.StartElement) error {
	_edd.Shape = nil
_bg:
	for {
		_ba, _cbed := d.Token()
		if _cbed != nil {
			return _cbed
		}
		switch _bcf := _ba.(type) {
		case _ddc.StartElement:
			switch _bcf.Name.Local {
			case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":
				_edd.Layout = _ac.NewOfcShapelayout()
				if _ec := d.DecodeElement(_edd.Layout, &_bcf); _ec != nil {
					return _ec
				}
			case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":
				_edd.ShapeType = _ac.NewShapetype()
				if _ee := d.DecodeElement(_edd.ShapeType, &_bcf); _ee != nil {
					return _ee
				}
			case "\u0073\u0068\u0061p\u0065":
				_bad := _ac.NewShape()
				if _ab := d.DecodeElement(_bad, &_bcf); _ab != nil {
					return _ab
				}
				_edd.Shape = append(_edd.Shape, _bad)
			}
		case _ddc.EndElement:
			break _bg
		}
	}
	return nil
}

// NewShapeStyle accept value of string style attribute in v:shape and format it to generate ShapeStyle.
func NewShapeStyle(style string) ShapeStyle {
	_beg := ShapeStyle{_cc: 0, _ff: 0}
	_ga := _dd.Split(style, "\u003b")
	for _, _da := range _ga {
		_abe := _dd.Split(_da, "\u003a")
		if len(_abe) != 2 {
			continue
		}
		switch _abe[0] {
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_beg._eec = _abe[1]
			break
		case "m\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074":
			_beg._fb, _ = _dg.ParseInt(_abe[1], 10, 64)
			break
		case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070":
			_beg._cbb, _ = _dg.ParseInt(_abe[1], 10, 64)
			break
		case "\u0077\u0069\u0064t\u0068":
			_beg._cc, _ = _dg.ParseInt(_dd.ReplaceAll(_abe[1], "\u0070\u0074", ""), 10, 64)
			break
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_beg._ff, _ = _dg.ParseInt(_dd.ReplaceAll(_abe[1], "\u0070\u0074", ""), 10, 64)
			break
		case "\u007a-\u0069\u006e\u0064\u0065\u0078":
			_beg._fe, _ = _dg.ParseInt(_abe[1], 10, 64)
			break
		case "\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c":
			_beg._be = _abe[1]
			break
		case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006fr\u0069z\u006f\u006e\u0074\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":
			_beg._eb = _abe[1]
			break
		case "m\u0073\u006f\u002d\u0070os\u0069t\u0069\u006f\u006e\u002d\u0076e\u0072\u0074\u0069\u0063\u0061\u006c":
			_beg._fcc = _abe[1]
			break
		case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069o\u006e\u002d\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":
			_beg._baa = _abe[1]
			break
		}
	}
	return _beg
}

// SetWidth set width of shape.
func (_gf *ShapeStyle) SetWidth(width int64) { _gf._cc = width }

// ToString generate string of TextpathStyle.
func (_ebe *TextpathStyle) String() string {
	_cgc := ""
	_cgc += _d.Sprintf("\u0066o\u006et\u002d\u0066\u0061\u006d\u0069\u006c\u0079\u003a\u0025\u0073\u003b", _ebe._gd)
	_cgc += _d.Sprintf("\u0066o\u006et\u002d\u0073\u0069\u007a\u0065\u003a\u0025\u0064\u0070\u0074\u003b", _ebe._gdg)
	if _ebe._cd {
		_cgc += _d.Sprintf("\u0066o\u006et\u002d\u0073\u0074\u0079\u006ce\u003a\u0069t\u0061\u006c\u0069\u0063\u003b")
	}
	if _ebe._db {
		_cgc += _d.Sprintf("\u0066\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074\u003ab\u006f\u006c\u0064\u003b")
	}
	return _cgc
}

// FontFamily returns fontFamily of the text.
func (_cf *TextpathStyle) FontFamily() string { return _cf._gd }

// ShapeStyle is style attribute of v:shape element.
type ShapeStyle struct {
	_eec string
	_fb  int64
	_cbb int64
	_cc  int64
	_ff  int64
	_fe  int64
	_be  string
	_eb  string
	_fcc string
	_baa string
}

// SetFontSize sets text's fontSize.
func (_ea *TextpathStyle) SetFontSize(fontSize int64) { _ea._gdg = fontSize }
