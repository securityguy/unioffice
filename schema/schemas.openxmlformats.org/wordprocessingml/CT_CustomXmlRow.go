// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type CT_CustomXmlRow struct {
	// Custom XML Element Namespace
	UriAttr *string
	// Custom XML Element Name
	ElementAttr string
	// Custom XML Element Properties
	CustomXmlPr          *CT_CustomXmlPr
	EG_ContentRowContent []*EG_ContentRowContent
}

func NewCT_CustomXmlRow() *CT_CustomXmlRow {
	ret := &CT_CustomXmlRow{}
	return ret
}

func (m *CT_CustomXmlRow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:element"},
		Value: fmt.Sprintf("%v", m.ElementAttr)})
	e.EncodeToken(start)
	if m.CustomXmlPr != nil {
		secustomXmlPr := xml.StartElement{Name: xml.Name{Local: "w:customXmlPr"}}
		e.EncodeElement(m.CustomXmlPr, secustomXmlPr)
	}
	if m.EG_ContentRowContent != nil {
		for _, c := range m.EG_ContentRowContent {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomXmlRow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
		}
		if attr.Name.Local == "element" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ElementAttr = parsed
		}
	}
lCT_CustomXmlRow:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "customXmlPr":
				m.CustomXmlPr = NewCT_CustomXmlPr()
				if err := d.DecodeElement(m.CustomXmlPr, &el); err != nil {
					return err
				}
			case "tr":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmp := NewCT_Row()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentrowcontent.Tr = append(tmpcontentrowcontent.Tr, tmp)
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
			case "customXml":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmpcontentrowcontent.CustomXml = NewCT_CustomXmlRow()
				if err := d.DecodeElement(tmpcontentrowcontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
			case "sdt":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmpcontentrowcontent.Sdt = NewCT_SdtRow()
				if err := d.DecodeElement(tmpcontentrowcontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
			case "proofErr":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "ins":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "del":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveFrom":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveTo":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
			case "bookmarkStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "bookmarkEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeStart":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeEnd":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "oMathPara":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case "oMath":
				tmpcontentrowcontent := NewEG_ContentRowContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_ContentRowContent = append(m.EG_ContentRowContent, tmpcontentrowcontent)
				tmpcontentrowcontent.EG_RunLevelElts = append(tmpcontentrowcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				log.Printf("skipping unsupported element on CT_CustomXmlRow %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomXmlRow
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomXmlRow and its children
func (m *CT_CustomXmlRow) Validate() error {
	return m.ValidateWithPath("CT_CustomXmlRow")
}

// ValidateWithPath validates the CT_CustomXmlRow and its children, prefixing error messages with path
func (m *CT_CustomXmlRow) ValidateWithPath(path string) error {
	if m.CustomXmlPr != nil {
		if err := m.CustomXmlPr.ValidateWithPath(path + "/CustomXmlPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_ContentRowContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ContentRowContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
