package tarantool

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	_ "github.com/tarantool/go-tarantool/uuid"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type relation struct {
	RelationUID string  `json:"relation_uid"`
	ArticleUIDA string  `json:"article_uid_a"`
	ArticleUIDB string  `json:"article_uid_b"`
	Coefficient float64 `json:"coefficient"`
}
type article struct {
	ArticleUID string `json:"article_uid"`
	Title      string `json:"title"`
}
type articleKeywords struct {
	ArticleUID string   `json:"article_uid"`
	Keywords   []string `json:"keywords"`
}
type keywords struct {
	Uid        string `json:"uid"`
	ArticleUid string `json:"article_uid"`
	Keyword    string `json:"keywords"`
}

// Search структура описывающая альбом

func (a *article) arrayLen() int { return 2 }

// EncodeMsgpack кодирует Msgpack для передачи в  tarantool
func (a *article) EncodeMsgpack(enc *msgpack.Encoder) error {
	e := encoder{e: enc}
	e.EncodeArrayLen(a.arrayLen())
	e.EncodeString(a.ArticleUID)
	e.EncodeString(a.Title)
	return e.err
}

// DecodeMsgpack декодирует Msgpack полученный из  tarantool
func (a *article) DecodeMsgpack(dec *msgpack.Decoder) error {
	d := decoder{d: dec}
	if l := d.DecodeArrayLen(); l != a.arrayLen() {
		return fmt.Errorf("unexpected array len: %d, want %d", l, a.arrayLen())
	}
	a.ArticleUID = d.DecodeString()
	a.Title = d.DecodeString()
	return nil
}

func (r *relation) arrayLen() int { return 2 }

func (r *relation) EncodeMsgpack(enc *msgpack.Encoder) error {
	e := encoder{e: enc}
	e.EncodeArrayLen(r.arrayLen())
	e.EncodeString(r.RelationUID)
	e.EncodeString(r.ArticleUIDA)
	e.EncodeString(r.ArticleUIDB)
	e.EncodeFloat64(r.Coefficient)
	return e.err
}

// DecodeMsgpack декодирует Msgpack полученный из  tarantool
func (r *relation) DecodeMsgpack(dec *msgpack.Decoder) error {
	d := decoder{d: dec}
	if l := d.DecodeArrayLen(); l != r.arrayLen() {
		return fmt.Errorf("unexpected array len: %d, want %d", l, r.arrayLen())
	}
	r.RelationUID = d.DecodeString()
	r.ArticleUIDA = d.DecodeString()
	r.ArticleUIDB = d.DecodeString()
	r.Coefficient = d.DecodeFloat64()
	return nil
}

func (a *articleKeywords) arrayLen() int { return 2 }

func (a *articleKeywords) EncodeMsgpack(enc *msgpack.Encoder) error {
	e := encoder{e: enc}
	e.EncodeArrayLen(a.arrayLen())
	e.EncodeString(a.ArticleUID)
	e.EncodeValue(reflect.ValueOf(a.Keywords))

	return e.err
}

// DecodeMsgpack декодирует Msgpack полученный из  tarantool
func (a *articleKeywords) DecodeMsgpack(dec *msgpack.Decoder) error {
	d := decoder{d: dec}
	if l := d.DecodeArrayLen(); l != a.arrayLen() {
		return fmt.Errorf("unexpected array len: %d, want %d", l, a.arrayLen())
	}
	a.ArticleUID = d.DecodeString()
	return nil
}

func (a *keywords) arrayLen() int { return 3 }

func (a *keywords) EncodeMsgpack(enc *msgpack.Encoder) error {
	e := encoder{e: enc}
	e.EncodeString(a.Uid)
	e.EncodeString(a.ArticleUid)
	e.EncodeString(a.Keyword)

	return e.err
}

// DecodeMsgpack декодирует Msgpack полученный из  tarantool
func (a *keywords) DecodeMsgpack(dec *msgpack.Decoder) error {
	d := decoder{d: dec}
	if l := d.DecodeArrayLen(); l != a.arrayLen() {
		return fmt.Errorf("unexpected array len: %d, want %d", l, a.arrayLen())
	}
	a.Uid = d.DecodeString()
	a.ArticleUid = d.DecodeString()
	a.Keyword = d.DecodeString()
	return nil
}

type encoder struct {
	e   *msgpack.Encoder
	err error
}

// EncodeArrayLen Кодирует длинну сообщения
func (e *encoder) EncodeArrayLen(l int) {
	if e.err == nil {
		e.err = e.e.EncodeArrayLen(l)
	}
}

// EncodeString кодирует строку для передачи в tarantool
func (e *encoder) EncodeString(v string) {
	if e.err == nil {
		e.err = e.e.EncodeString(v)
	}
}

// EncodeValue кодирует произвольное значение для передачи в tarantool
func (e *encoder) EncodeValue(v reflect.Value) {
	if e.err == nil {
		e.err = e.e.EncodeValue(v)
	}
}

func (e *encoder) EncodeFloat64(v float64) {
	if e.err == nil {
		e.err = e.e.EncodeFloat64(v)
	}
}

type decoder struct {
	d   *msgpack.Decoder
	err error
}

// DecodeArrayLen декокодирует длинну массива полученную из tarantool
func (d *decoder) DecodeArrayLen() (l int) {
	if d.err == nil {
		l, d.err = d.d.DecodeArrayLen()
	}
	return
}

// DecodeString декокодирует строку полученную из tarantool
func (d *decoder) DecodeString() (s string) {
	if d.err == nil {
		s, d.err = d.d.DecodeString()
	}
	return
}

// decodeStringArray декокодирует []string полученный из tarantool
func (d *decoder) decodeStringArray() (s []string) {
	if d.err == nil {
		v, err := d.d.DecodeSlice()
		d.err = err

		if d.err == nil {
			for _, i := range v {
				s = append(s, i.(string))
			}
		}
	}
	return
}

func (d *decoder) decodeUUID() (u uuid.UUID) {
	if d.err == nil {
		var v interface{}
		v, d.err = d.d.DecodeInterface()
		if d.err == nil {
			u = v.(uuid.UUID)
		}
	}
	return
}

func (d *decoder) DecodeFloat64() (v float64) {
	if d.err == nil {
		var u float64
		u, d.err = d.d.DecodeFloat64()
		if d.err == nil {
			v = u
		}
	}
	return
}

type SearchTest struct {
	SearchRequest string `json:"searchRequest"`
	TagIds        []int  `json:"tagIds"`
	FileTypeId    int    `json:"fileTypeId"`
	RecentTypeId  int    `json:"recentTypeId"`
}

func (a *SearchTest) arrayLen() int { return 4 }

// EncodeMsgpack кодирует Msgpack для передачи в  tarantool
func (a *SearchTest) EncodeMsgpack(enc *msgpack.Encoder) error {
	e := encoder{e: enc}
	e.EncodeArrayLen(a.arrayLen())
	e.EncodeValue(reflect.ValueOf(a.SearchRequest))
	e.EncodeValue(reflect.ValueOf(a.TagIds))
	e.EncodeValue(reflect.ValueOf(a.FileTypeId))
	e.EncodeValue(reflect.ValueOf(a.RecentTypeId))
	return e.err
}
