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

package terms

import (
	_c "encoding/xml"
	_e "fmt"
	_a "github.com/sanhuanshisanshao/unioffice"
	_cf "github.com/sanhuanshisanshao/unioffice/common/logger"
	_ac "github.com/sanhuanshisanshao/unioffice/schema/purl.org/dc/elements"
)

func NewElementsAndRefinementsGroup() *ElementsAndRefinementsGroup {
	_gd := &ElementsAndRefinementsGroup{}
	return _gd
}

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_cfg *ElementsAndRefinementsGroup) ValidateWithPath(path string) error {
	for _cb, _ddc := range _cfg.Choice {
		if _bfd := _ddc.ValidateWithPath(_e.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _cb)); _bfd != nil {
			return _bfd
		}
	}
	return nil
}
func (_bcg *ISO3166) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0049S\u004f\u0033\u0031\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewISO3166() *ISO3166 { _cdb := &ISO3166{}; return _cdb }

type ISO3166 struct{}

func (_dgc *MESH) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_bfb, _aga := d.Token()
		if _aga != nil {
			return _e.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073", _aga)
		}
		if _ga, _begbd := _bfb.(_c.EndElement); _begbd && _ga.Name == start.Name {
			break
		}
	}
	return nil
}
func NewDDC() *DDC { _ed := &DDC{}; return _ed }

type UDC struct{}
type Point struct{}

func (_eb *IMT) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ebg, _cef := d.Token()
		if _cef != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073", _cef)
		}
		if _dfdd, _fca := _ebg.(_c.EndElement); _fca && _dfdd.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ecfe *Period) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ebb, _fgde := d.Token()
		if _fgde != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073", _fgde)
		}
		if _abc, _bbcf := _ebb.(_c.EndElement); _bbcf && _abc.Name == start.Name {
			break
		}
	}
	return nil
}

type ElementsAndRefinementsGroupChoice struct{ Any []*_ac.Any }

func (_be *DDC) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ade, _cac := d.Token()
		if _cac != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073", _cac)
		}
		if _df, _bfc := _ade.(_c.EndElement); _bfc && _df.Name == start.Name {
			break
		}
	}
	return nil
}
func (_beg *ISO3166) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ff, _badc := d.Token()
		if _badc != nil {
			return _e.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073", _badc)
		}
		if _begb, _ffg := _ff.(_c.EndElement); _ffg && _begb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cag *ElementsAndRefinementsGroup) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _cag.Choice != nil {
		for _, _fd := range _cag.Choice {
			_fd.MarshalXML(e, _c.StartElement{})
		}
	}
	return nil
}

// Validate validates the LCC and its children
func (_ceg *LCC) Validate() error { return _ceg.ValidateWithPath("\u004c\u0043\u0043") }

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_cefa *LCSH) ValidateWithPath(path string) error { return nil }

type MESH struct{}

func (_ceff *RFC1766) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0052F\u0043\u0031\u0037\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_dcbc *UDC) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0055\u0044\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the ElementOrRefinementContainer and its children
func (_bcb *ElementOrRefinementContainer) Validate() error {
	return _bcb.ValidateWithPath("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072")
}
func NewRFC3066() *RFC3066 { _gaf := &RFC3066{}; return _gaf }
func (_dg *DDC) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0044\u0044\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_dc *ElementOrRefinementContainer) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072"
	e.EncodeToken(start)
	if _dc.Choice != nil {
		for _, _dcb := range _dc.Choice {
			_dcb.MarshalXML(e, _c.StartElement{})
		}
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewIMT() *IMT { _fac := &IMT{}; return _fac }

type Period struct{}

func NewBox() *Box { _gb := &Box{}; return _gb }

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_fbe *Period) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_eed *MESH) ValidateWithPath(path string) error { return nil }
func (_ecg *LCSH) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_eca, _dcd := d.Token()
		if _dcd != nil {
			return _e.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073", _dcd)
		}
		if _gbcg, _ccb := _eca.(_c.EndElement); _ccb && _gbcg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_fdf *IMT) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0049\u004d\u0054"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_acdg *ISO639_2) ValidateWithPath(path string) error { return nil }
func (_dfb *LCC) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ced, _aa := d.Token()
		if _aa != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073", _aa)
		}
		if _deb, _bdf := _ced.(_c.EndElement); _bdf && _deb.Name == start.Name {
			break
		}
	}
	return nil
}

type Box struct{}

func NewPoint() *Point { _gga := &Point{}; return _gga }
func (_aed *RFC3066) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0052F\u0043\u0033\u0030\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewMESH() *MESH { _eecb := &MESH{}; return _eecb }
func (_ecd *ISO639_2) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ag, _cde := d.Token()
		if _cde != nil {
			return _e.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073", _cde)
		}
		if _fgb, _efd := _ag.(_c.EndElement); _efd && _fgb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_abg *W3CDTF) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0057\u0033\u0043\u0044\u0054\u0046"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_ad *DCMIType) ValidateWithPath(path string) error { return nil }

type W3CDTF struct{}

// Validate validates the TGN and its children
func (_bcf *TGN) Validate() error { return _bcf.ValidateWithPath("\u0054\u0047\u004e") }
func (_cbg *Point) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0050\u006f\u0069n\u0074"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_cec *LCSH) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u004c\u0043\u0053\u0048"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type ElementOrRefinementContainer struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

func (_aff *RFC3066) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_gc, _bce := d.Token()
		if _bce != nil {
			return _e.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073", _bce)
		}
		if _abf, _dde := _gc.(_c.EndElement); _dde && _abf.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Point and its children
func (_fcaf *Point) Validate() error { return _fcaf.ValidateWithPath("\u0050\u006f\u0069n\u0074") }

type DCMIType struct{}

func NewISO639_2() *ISO639_2 { _cbc := &ISO639_2{}; return _cbc }
func NewPeriod() *Period     { _bebc := &Period{}; return _bebc }

// Validate validates the IMT and its children
func (_bad *IMT) Validate() error { return _bad.ValidateWithPath("\u0049\u004d\u0054") }
func NewRFC1766() *RFC1766        { _dfg := &RFC1766{}; return _dfg }

type ElementsAndRefinementsGroup struct {
	Choice []*ElementsAndRefinementsGroupChoice
}
type RFC1766 struct{}

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_aac *UDC) ValidateWithPath(path string) error { return nil }

// Validate validates the DDC and its children
func (_dd *DDC) Validate() error { return _dd.ValidateWithPath("\u0044\u0044\u0043") }
func (_ce *ElementsAndRefinementsGroup) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_faf:
	for {
		_caa, _gg := d.Token()
		if _gg != nil {
			return _gg
		}
		switch _cc := _caa.(type) {
		case _c.StartElement:
			switch _cc.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_baf := NewElementsAndRefinementsGroupChoice()
				if _bbe := d.DecodeElement(&_baf.Any, &_cc); _bbe != nil {
					return _bbe
				}
				_ce.Choice = append(_ce.Choice, _baf)
			default:
				_cf.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076", _cc.Name)
				if _ggc := d.Skip(); _ggc != nil {
					return _ggc
				}
			}
		case _c.EndElement:
			break _faf
		case _c.CharData:
		}
	}
	return nil
}

type LCC struct{}
type LCSH struct{}

// Validate validates the Period and its children
func (_dccb *Period) Validate() error {
	return _dccb.ValidateWithPath("\u0050\u0065\u0072\u0069\u006f\u0064")
}

type DDC struct{}

func NewLCC() *LCC { _ge := &LCC{}; return _ge }

// Validate validates the ISO3166 and its children
func (_afb *ISO3166) Validate() error {
	return _afb.ValidateWithPath("\u0049S\u004f\u0033\u0031\u0036\u0036")
}

type ISO639_2 struct{}

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_bc *DDC) ValidateWithPath(path string) error { return nil }
func (_fcf *LCC) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u004c\u0043\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type RFC3066 struct{}

func (_cce *TGN) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0054\u0047\u004e"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewElementOrRefinementContainer() *ElementOrRefinementContainer {
	_fg := &ElementOrRefinementContainer{}
	return _fg
}
func (_bb *DCMIType) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_db, _cg := d.Token()
		if _cg != nil {
			return _e.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073", _cg)
		}
		if _bbc, _bd := _db.(_c.EndElement); _bd && _bbc.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the W3CDTF and its children
func (_bag *W3CDTF) Validate() error {
	return _bag.ValidateWithPath("\u0057\u0033\u0043\u0044\u0054\u0046")
}

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_dbf *W3CDTF) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_beb *ISO3166) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_ffd *Point) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_ec *ElementOrRefinementContainer) ValidateWithPath(path string) error {
	for _bac, _bdg := range _ec.Choice {
		if _de := _bdg.ValidateWithPath(_e.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _bac)); _de != nil {
			return _de
		}
	}
	return nil
}
func (_bfbe *UDC) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_fda, _aef := d.Token()
		if _aef != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073", _aef)
		}
		if _acc, _bge := _fda.(_c.EndElement); _bge && _acc.Name == start.Name {
			break
		}
	}
	return nil
}
func NewLCSH() *LCSH { _cbce := &LCSH{}; return _cbce }
func (_eef *ElementsAndRefinementsGroupChoice) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _eef.Any != nil {
		_dcc := _c.StartElement{Name: _c.Name{Local: "\u0064\u0063\u003a\u0061\u006e\u0079"}}
		for _, _ae := range _eef.Any {
			e.EncodeElement(_ae, _dcc)
		}
	}
	return nil
}
func NewElementsAndRefinementsGroupChoice() *ElementsAndRefinementsGroupChoice {
	_ccg := &ElementsAndRefinementsGroupChoice{}
	return _ccg
}

// Validate validates the RFC1766 and its children
func (_ccba *RFC1766) Validate() error {
	return _ccba.ValidateWithPath("\u0052F\u0043\u0031\u0037\u0036\u0036")
}
func (_cfa *Point) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_gf, _eedg := d.Token()
		if _eedg != nil {
			return _e.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073", _eedg)
		}
		if _cgb, _gda := _gf.(_c.EndElement); _gda && _cgb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_dca *Period) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0050\u0065\u0072\u0069\u006f\u0064"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_gea *TGN) ValidateWithPath(path string) error { return nil }
func (_aag *URI) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0055\u0052\u0049"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_d *Box) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0042\u006f\u0078"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the Box and its children
func (_ef *Box) Validate() error { return _ef.ValidateWithPath("\u0042\u006f\u0078") }

// Validate validates the URI and its children
func (_fcd *URI) Validate() error { return _fcd.ValidateWithPath("\u0055\u0052\u0049") }

type TGN struct{}

// Validate validates the RFC3066 and its children
func (_agd *RFC3066) Validate() error {
	return _agd.ValidateWithPath("\u0052F\u0043\u0033\u0030\u0036\u0036")
}
func (_cab *RFC1766) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_dfe, _aaf := d.Token()
		if _aaf != nil {
			return _e.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073", _aaf)
		}
		if _fcfb, _cfgb := _dfe.(_c.EndElement); _cfgb && _fcfb.Name == start.Name {
			break
		}
	}
	return nil
}

type IMT struct{}

func (_eec *ISO639_2) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the MESH and its children
func (_cdf *MESH) Validate() error { return _cdf.ValidateWithPath("\u004d\u0045\u0053\u0048") }

// Validate validates the DCMIType and its children
func (_bf *DCMIType) Validate() error {
	return _bf.ValidateWithPath("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065")
}

type URI struct{}

func (_cfc *Box) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_ee, _b := d.Token()
		if _b != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073", _b)
		}
		if _gbc, _eg := _ee.(_c.EndElement); _eg && _gbc.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the LCSH and its children
func (_ecae *LCSH) Validate() error { return _ecae.ValidateWithPath("\u004c\u0043\u0053\u0048") }
func NewDCMIType() *DCMIType        { _cd := &DCMIType{}; return _cd }

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_bdb *ElementsAndRefinementsGroupChoice) ValidateWithPath(path string) error {
	for _cbd, _aeg := range _bdb.Any {
		if _eee := _aeg.ValidateWithPath(_e.Sprintf("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d", path, _cbd)); _eee != nil {
			return _eee
		}
	}
	return nil
}

// Validate validates the ISO639_2 and its children
func (_ab *ISO639_2) Validate() error {
	return _ab.ValidateWithPath("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032")
}
func (_fef *URI) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_dgb, _cgf := d.Token()
		if _cgf != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073", _cgf)
		}
		if _acgb, _ddf := _dgb.(_c.EndElement); _ddf && _acgb.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_gec *URI) ValidateWithPath(path string) error { return nil }

// Validate validates the ElementsAndRefinementsGroup and its children
func (_fgc *ElementsAndRefinementsGroup) Validate() error {
	return _fgc.ValidateWithPath("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070")
}
func (_f *DCMIType) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_ca *Box) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_cfgg *RFC1766) ValidateWithPath(path string) error { return nil }
func (_bg *ElementOrRefinementContainer) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_fb:
	for {
		_cad, _fgd := d.Token()
		if _fgd != nil {
			return _fgd
		}
		switch _dce := _cad.(type) {
		case _c.StartElement:
			switch _dce.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_dfd := NewElementsAndRefinementsGroupChoice()
				if _ea := d.DecodeElement(&_dfd.Any, &_dce); _ea != nil {
					return _ea
				}
				_bg.Choice = append(_bg.Choice, _dfd)
			default:
				_cf.Log.Debug("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076", _dce.Name)
				if _edd := d.Skip(); _edd != nil {
					return _edd
				}
			}
		case _c.EndElement:
			break _fb
		case _c.CharData:
		}
	}
	return nil
}
func (_ecf *ElementsAndRefinementsGroupChoice) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_bcbe:
	for {
		_ege, _af := d.Token()
		if _af != nil {
			return _af
		}
		switch _bbd := _ege.(type) {
		case _c.StartElement:
			switch _bbd.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_acd := _ac.NewAny()
				if _ggd := d.DecodeElement(_acd, &_bbd); _ggd != nil {
					return _ggd
				}
				_ecf.Any = append(_ecf.Any, _acd)
			default:
				_cf.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076", _bbd.Name)
				if _fc := d.Skip(); _fc != nil {
					return _fc
				}
			}
		case _c.EndElement:
			break _bcbe
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the UDC and its children
func (_fdc *UDC) Validate() error { return _fdc.ValidateWithPath("\u0055\u0044\u0043") }
func NewURI() *URI                { _ecac := &URI{}; return _ecac }
func (_bbef *MESH) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u004d\u0045\u0053\u0048"
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_cfe *IMT) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_edf *LCC) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_bgd *RFC3066) ValidateWithPath(path string) error { return nil }
func NewW3CDTF() *W3CDTF                                 { _fed := &W3CDTF{}; return _fed }
func NewTGN() *TGN                                       { _dceb := &TGN{}; return _dceb }

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_acg *ElementsAndRefinementsGroupChoice) Validate() error {
	return _acg.ValidateWithPath("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065")
}
func (_cfea *TGN) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_eba, _edfb := d.Token()
		if _edfb != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073", _edfb)
		}
		if _fe, _bgc := _eba.(_c.EndElement); _bgc && _fe.Name == start.Name {
			break
		}
	}
	return nil
}
func NewUDC() *UDC { _fafa := &UDC{}; return _fafa }
func (_ada *W3CDTF) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for {
		_bcfg, _eeg := d.Token()
		if _eeg != nil {
			return _e.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073", _eeg)
		}
		if _bee, _cbgg := _bcfg.(_c.EndElement); _cbgg && _bee.Name == start.Name {
			break
		}
	}
	return nil
}
func init() {
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004c\u0043\u0053\u0048", NewLCSH)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004d\u0045\u0053\u0048", NewMESH)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0044\u0044\u0043", NewDDC)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004c\u0043\u0043", NewLCC)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0055\u0044\u0043", NewUDC)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0050\u0065\u0072\u0069\u006f\u0064", NewPeriod)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0057\u0033\u0043\u0044\u0054\u0046", NewW3CDTF)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065", NewDCMIType)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049\u004d\u0054", NewIMT)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0055\u0052\u0049", NewURI)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032", NewISO639_2)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0052F\u0043\u0031\u0037\u0036\u0036", NewRFC1766)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0052F\u0043\u0033\u0030\u0036\u0036", NewRFC3066)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0050\u006f\u0069n\u0074", NewPoint)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049S\u004f\u0033\u0031\u0036\u0036", NewISO3166)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0042\u006f\u0078", NewBox)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0054\u0047\u004e", NewTGN)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072", NewElementOrRefinementContainer)
	_a.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070", NewElementsAndRefinementsGroup)
}
