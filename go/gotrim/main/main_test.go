package main

import (
	"strings"
	"testing"

	"github.com/paulopraxedes/estudos/go/gotrim"
)

const text = " \n\nFoo Bar  Foo  \n\n  "

func TestLeftTrim(t *testing.T) {
	textOk := "\n\nFoo Bar  Foo  \n\n  "

	if len(text)-1 != len(gotrim.LeftTrim(text)) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.LeftTrim(text)))
	} else {
		if strings.Compare(textOk, gotrim.LeftTrim(text)) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}

func TestRightTrim(t *testing.T) {
	textOk := " \n\nFoo Bar  Foo  \n\n"

	if len(text)-2 != len(gotrim.RightTrim(text)) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.RightTrim(text)))
	} else {
		if strings.Compare(textOk, gotrim.RightTrim(text)) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}

func TestTopTrim(t *testing.T) {
	textOk := "Foo Bar  Foo  \n\n  "

	if len(text)-3 != len(gotrim.TopTrim(gotrim.LeftTrim(text))) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.TopTrim(gotrim.LeftTrim(text))))
	} else {
		if strings.Compare(textOk, gotrim.TopTrim(gotrim.LeftTrim(text))) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}

func TestBottomTrim(t *testing.T) {
	textOk := " \n\nFoo Bar  Foo  "

	if len(text)-4 != len(gotrim.BottomTrim(gotrim.RightTrim(text))) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.BottomTrim(gotrim.RightTrim(text))))
	} else {
		if strings.Compare(textOk, gotrim.BottomTrim(gotrim.RightTrim(text))) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}

func TestVerticalTrim(t *testing.T) {
	textOk := "Foo Bar  Foo  "

	if len(text)-7 != len(gotrim.BottomTrim(gotrim.RightTrim(gotrim.TopTrim(gotrim.LeftTrim(text))))) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.BottomTrim(gotrim.RightTrim(gotrim.TopTrim(gotrim.LeftTrim(text))))))
	} else {
		if strings.Compare(textOk, gotrim.BottomTrim(gotrim.RightTrim(gotrim.TopTrim(gotrim.LeftTrim(text))))) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}

func TestHorizontalTrim(t *testing.T) {
	textOk := "\n\nFoo Bar  Foo  \n\n"

	if len(text)-3 != len(gotrim.RightTrim(gotrim.LeftTrim(text))) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.RightTrim(gotrim.LeftTrim(text))))
	} else {
		if strings.Compare(textOk, gotrim.RightTrim(gotrim.LeftTrim(text))) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}
