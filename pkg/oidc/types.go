package oidc

import (
	"encoding/json"
	"strings"
	"time"

	"golang.org/x/text/language"
)

type Audience []string

func (a *Audience) UnmarshalJSON(text []byte) error {
	var i interface{}
	err := json.Unmarshal(text, &i)
	if err != nil {
		return err
	}
	switch aud := i.(type) {
	case []interface{}:
		*a = make([]string, len(aud))
		for i, audience := range aud {
			(*a)[i] = audience.(string)
		}
	case string:
		*a = []string{aud}
	}
	return nil
}

type Display string

func (d *Display) UnmarshalText(text []byte) error {
	display := Display(text)
	switch display {
	case DisplayPage, DisplayPopup, DisplayTouch, DisplayWAP:
		*d = display
	}
	return nil
}

type Gender string

type Locales []language.Tag

func (l *Locales) UnmarshalText(text []byte) error {
	locales := strings.Split(string(text), " ")
	for _, locale := range locales {
		tag, err := language.Parse(locale)
		if err == nil && !tag.IsRoot() {
			*l = append(*l, tag)
		}
	}
	return nil
}

type MaxAge *uint

func NewMaxAge(i uint) MaxAge {
	return &i
}

type SpaceDelimitedArray []string

type Prompt SpaceDelimitedArray

type ResponseType string

func (s SpaceDelimitedArray) Encode() string {
	return strings.Join(s, " ")
}

func (s *SpaceDelimitedArray) UnmarshalText(text []byte) error {
	*s = strings.Split(string(text), " ")
	return nil
}

func (s SpaceDelimitedArray) MarshalText() ([]byte, error) {
	return []byte(s.Encode()), nil
}

func (s SpaceDelimitedArray) MarshalJSON() ([]byte, error) {
	return json.Marshal((s).Encode())
}

func (s *SpaceDelimitedArray) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = strings.Split(str, " ")
	return nil
}

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	*t = Time(time.Unix(i, 0).UTC())
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*t).UTC().Unix())
}
