package ddlog_go

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	mockNow, _ := time.Parse("2006-01-02 15:04:06", "2017-09-01 12:00:00")
	setNow(mockNow)
	code := m.Run()
	os.Exit(code)
}

func setNow(t time.Time) {
	timeNowFunc = func() time.Time { return t }
}

func TestINFO(t *testing.T) {
	var buf bytes.Buffer
	ddl := NewddLogger("INFOTest", &buf)
	ddl.Attr("test", "info").INFO("1")
	expected := "INFOTest 967809600 1 loglevel=INFO test=info\n"
	if buf.String() != expected {
		t.Errorf("INFOTest failed. Returned value is %v", buf.String())
	}
}

func TestDEBUG(t *testing.T) {
	var buf bytes.Buffer
	ddl := NewddLogger("DEBUGTest", &buf)
	ddl.Attr("test", "debug").Attr("hoge", "fuga").DEBUG("2")
	expected := "DEBUGTest 967809600 2 loglevel=DEBUG test=debug hoge=fuga\n"
	if buf.String() != expected {
		t.Errorf("DEBUGTest failed. Returned value is %v", buf.String())
	}
}

func TestWARN(t *testing.T) {
	var buf bytes.Buffer
	ddl := NewddLogger("WARNTest", &buf)
	ddl.Attr("test", "warn").Attr("foo", "bar").WARN("3")
	expected := "WARNTest 967809600 3 loglevel=WARN test=warn foo=bar\n"
	if buf.String() != expected {
		t.Errorf("WARNTest failed. Returned value is %v", buf.String())
	}
}

func TestERROR(t *testing.T) {
	var buf bytes.Buffer
	ddl := NewddLogger("ERRORTest", &buf)
	ddl.Attr("test", "error").Attr("bar", "buz").ERROR("4")
	expected := "ERRORTest 967809600 4 loglevel=ERROR test=error bar=buz\n"
	if buf.String() != expected {
		t.Errorf("DEBUGTest failed. Returned value is %v", buf.String())
	}
}

func TestFATAL(t *testing.T) {
	var buf bytes.Buffer
	ddl := NewddLogger("FATALTest", &buf)
	ddl.FATAL("5")
	expected := "FATALTest 967809600 5 loglevel=FATAL\n"
	if buf.String() != expected {
		t.Errorf("DEBUGTest failed. Returned value is %v", buf.String())
	}
}
