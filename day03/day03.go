package day03

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(data []byte) {
	lexer := newLexer(data)
	parser := newParser(lexer)
	sum := 0

	for parser.currentToken.tokenType != EOF {
		if parser.currentToken.tokenType == MUL {
			sum += parser.parseMulExpression()
		}
		parser.nextToken()
	}

	fmt.Printf("part1: %d\n", sum)
}

func part2(data []byte) {
	lexer := newLexer(data)
	parser := newParser(lexer)
	sum := 0
	enabled := true

	for parser.currentToken.tokenType != EOF {
		if parser.currentToken.tokenType == MUL && enabled {
			sum += parser.parseMulExpression()
		}

		if parser.currentToken.tokenType == DO && parser.parseDoExpression() {
			enabled = true
		}

		if parser.currentToken.tokenType == DONT && parser.parseDontExpression() {
			enabled = false
		}

		parser.nextToken()
	}

	fmt.Printf("part1: %d\n", sum)
}

func Run() {
	data, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(data)
	part2(data)
}

type lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
}

func newLexer(data []byte) *lexer {
	l := &lexer{input: data}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *lexer) readIdentifier() string {
	position := l.position
	for isValidIdentifierValue(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

type tokenType = string

type token struct {
	literal   string
	tokenType tokenType
}

const (
	EOF     tokenType = "EOF"
	LPAREN  tokenType = "LPAREN"
	RPAREN  tokenType = "RPAREN"
	COMMA   tokenType = "COMMA"
	MUL     tokenType = "MUL"
	INT     tokenType = "INT"
	ILLEGAL tokenType = "ILLEGAL"
	DO      tokenType = "DO"
	DONT    tokenType = "DON'T"
)

var keywords = map[string]tokenType{
	"do":    DO,
	"don't": DONT,
	"mul":   MUL,
}

func (l *lexer) nextToken() token {
	var t token

	switch l.ch {
	case '(':
		t = token{tokenType: LPAREN, literal: "("}
	case ')':
		t = token{tokenType: RPAREN, literal: ")"}
	case ',':
		t = token{tokenType: COMMA, literal: ","}
	case 0:
		t = token{tokenType: EOF, literal: "EOF"}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			if tokenType, ok := keywords[ident]; ok {
				t = token{tokenType: tokenType, literal: ident}
			} else {
				t = token{tokenType: ILLEGAL, literal: ident}
			}
			return t
		} else if isDigit(l.ch) {
			num := l.readNumber()
			t = token{tokenType: INT, literal: num}
			return t
		} else {
			t = token{tokenType: ILLEGAL, literal: string(l.ch)}
		}
	}

	l.readChar()
	return t
}

type parser struct {
	lexer        *lexer
	currentToken token
	peekToken    token
}

func newParser(lexer *lexer) *parser {
	p := &parser{lexer: lexer}
	return p
}

func (p *parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.nextToken()
}

func (p *parser) parseMulExpression() int {
	result := 1

	if p.peekToken.tokenType != LPAREN {
		return 0
	}
	p.nextToken()

	for {
		if p.peekToken.tokenType != INT {
			return 0
		}

		p.nextToken()
		value, err := strconv.Atoi(p.currentToken.literal)
		if err != nil {
			return 0
		}

		result *= value

		if p.peekToken.tokenType != COMMA && p.peekToken.tokenType != RPAREN {
			return 0
		}

		if p.peekToken.tokenType == RPAREN {
			p.nextToken()
			return result
		} else {
			p.nextToken()
		}
	}
}

func (p *parser) parseDoExpression() bool {
	if p.peekToken.tokenType != LPAREN {
		return false
	}
	p.nextToken()

	if p.peekToken.tokenType != RPAREN {
		return false
	}

	p.nextToken()
	return true
}

func (p *parser) parseDontExpression() bool {
	if p.peekToken.tokenType != LPAREN {
		return false
	}
	p.nextToken()

	if p.peekToken.tokenType != RPAREN {
		return false
	}

	p.nextToken()
	return true
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isValidIdentifierValue(ch byte) bool {
	return isLetter(ch) || ch == '\''
}
