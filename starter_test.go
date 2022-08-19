package main

import (
	"testing"
)

func TestNoDuplicates(t *testing.T) {
	if x := roadkill("=====h==yyyy==eee==n====a==="); x != "hyena" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "hyena")
	}
	if x := roadkill("======pe====nnnnnn=======================n=n=ng====u==iiii=iii==nn========================n="); x != "penguin" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "penguin")
	}
	if x := roadkill("===bb==bb==b==bb=eee==aa===rr==rr==="); x != "bear" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "bear")
	}
	if x := roadkill("===bbrrr==bb==b==bb=eee==aa===rr==rr==="); x != "??" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "??")
	}
}

func TestReverseNoDuplicates(t *testing.T) {
	if x := roadkill("=====a==nnnn==eee==y====h==="); x != "hyena" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "hyena")
	}
	if x := roadkill("n=======n====i=========i=============iiu=======u=====u======g===gg====nn=====ep=========p======pp"); x != "penguin" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "penguin")
	}
	if x := roadkill("=====r=rrr=rra=====eee======bb====b======="); x != "bear" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "bear")
	}
	if x := roadkill("===bbrrr==bb==b==brb=eee==aa===rr==rr==="); x != "??" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "??")
	}
}

func TestDuplicates(t *testing.T) {
	if x := roadkill("=====w==aaaa==llllll=ll=ll=lll==a==b====y==="); x != "wallaby" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "wallaby")
	}
	if x := roadkill("===y==b==a==lll==ll==a====w==="); x != "wallaby" {
		t.Fatalf("roadkill() got %s, wanted %s", x, "wallaby")
	}
}
