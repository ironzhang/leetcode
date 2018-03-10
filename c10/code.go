package leetcode

import (
	"io"
	"strings"
)

type machine struct {
	expr  string
	index int
	eof   bool
	ch    byte
	star  bool
}

func (p *machine) ok() bool {
	return p.eof
}

func (p *machine) match(ch byte) bool {
	for {
		if p.eof {
			return true
		}
		if p.ch == '.' || p.ch == ch {
			if !p.star {
				p.next()
			}
			return true
		}
		if p.star {
			p.next()
			continue
		}
		return false
	}
}

func (p *machine) next() {
	p.ch, p.eof = p.readRune()
	if p.eof {
		return
	}
	p.star = p.readStar()
}

func (p *machine) readRune() (ch byte, eof bool) {
	if p.index >= len(p.expr) {
		eof = true
		return
	}
	ch = p.expr[p.index]
	p.index++
	return
}

func (p *machine) readStar() bool {
	if p.index >= len(p.expr) {
		return false
	}
	if p.expr[p.index] == '*' {
		p.index++
		return true
	}
	return false
}

func isMatch(s string, p string) bool {
	m := machine{}
	r := strings.NewReader(s)
	m.next()
	for {
		if m.ok() {
			return true
		}
		ch, err := r.ReadByte()
		if err == io.EOF {
			return false
		}
		if !m.match(ch) {
			return false
		}
		continue
	}
}
