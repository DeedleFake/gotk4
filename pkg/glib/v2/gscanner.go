// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib.h>
import "C"

// ErrorType: the possible errors, used in the @v_error field of Value, when the
// token is a G_TOKEN_ERROR.
type ErrorType int

const (
	// Unknown error
	ErrorTypeUnknown ErrorType = iota
	// UnexpEOF: unexpected end of file
	ErrorTypeUnexpEOF
	// UnexpEOFInString: unterminated string constant
	ErrorTypeUnexpEOFInString
	// UnexpEOFInComment: unterminated comment
	ErrorTypeUnexpEOFInComment
	// NonDigitInConst: non-digit character in a number
	ErrorTypeNonDigitInConst
	// DigitRadix: digit beyond radix in a number
	ErrorTypeDigitRadix
	// FloatRadix: non-decimal floating point number
	ErrorTypeFloatRadix
	// FloatMalformed: malformed floating point number
	ErrorTypeFloatMalformed
)

// TokenType: the possible types of token returned from each
// g_scanner_get_next_token() call.
type TokenType int

const (
	// EOF: the end of the file
	TokenTypeEOF TokenType = 0
	// LeftParen: '(' character
	TokenTypeLeftParen TokenType = 40
	// RightParen: ')' character
	TokenTypeRightParen TokenType = 41
	// LeftCurly: '{' character
	TokenTypeLeftCurly TokenType = 123
	// RightCurly: '}' character
	TokenTypeRightCurly TokenType = 125
	// LeftBrace: '[' character
	TokenTypeLeftBrace TokenType = 91
	// RightBrace: ']' character
	TokenTypeRightBrace TokenType = 93
	// EqualSign: '=' character
	TokenTypeEqualSign TokenType = 61
	// Comma: ',' character
	TokenTypeComma TokenType = 44
	// None: not a token
	TokenTypeNone TokenType = 256
	// Error occurred
	TokenTypeError TokenType = 257
	// Char: character
	TokenTypeChar TokenType = 258
	// Binary integer
	TokenTypeBinary TokenType = 259
	// Octal integer
	TokenTypeOctal TokenType = 260
	// Int: integer
	TokenTypeInt TokenType = 261
	// Hex integer
	TokenTypeHex TokenType = 262
	// Float: floating point number
	TokenTypeFloat TokenType = 263
	// String: string
	TokenTypeString TokenType = 264
	// Symbol: symbol
	TokenTypeSymbol TokenType = 265
	// Identifier: identifier
	TokenTypeIdentifier TokenType = 266
	// IdentifierNull: null identifier
	TokenTypeIdentifierNull TokenType = 267
	// CommentSingle: one line comment
	TokenTypeCommentSingle TokenType = 268
	// CommentMulti: multi line comment
	TokenTypeCommentMulti TokenType = 269
)

// Scanner: the data structure representing a lexical scanner.
//
// You should set @input_name after creating the scanner, since it is used by
// the default message handler when displaying warnings and errors. If you are
// scanning a file, the filename would be a good choice.
//
// The @user_data and @max_parse_errors fields are not used. If you need to
// associate extra data with the scanner you can place them here.
//
// If you want to use your own message handler you can set the @msg_handler
// field. The type of the message handler function is declared by MsgFunc.
type Scanner struct {
	native C.GScanner
}

// WrapScanner wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScanner(ptr unsafe.Pointer) *Scanner {
	return (*Scanner)(ptr)
}

// Native returns the underlying C source pointer.
func (s *Scanner) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// CurLine returns the current line in the input stream (counting from 1). This
// is the line of the last token parsed via g_scanner_get_next_token().
func (scanner *Scanner) CurLine() uint {
	var _arg0 *C.GScanner // out
	var _cret C.guint     // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.g_scanner_cur_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CurPosition returns the current position in the current line (counting from
// 0). This is the position of the last token parsed via
// g_scanner_get_next_token().
func (scanner *Scanner) CurPosition() uint {
	var _arg0 *C.GScanner // out
	var _cret C.guint     // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.g_scanner_cur_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CurToken gets the current token type. This is simply the @token field in the
// #GScanner structure.
func (scanner *Scanner) CurToken() TokenType {
	var _arg0 *C.GScanner  // out
	var _cret C.GTokenType // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.g_scanner_cur_token(_arg0)

	var _tokenType TokenType // out

	_tokenType = (TokenType)(_cret)

	return _tokenType
}

// Destroy frees all memory used by the #GScanner.
func (scanner *Scanner) Destroy() {
	var _arg0 *C.GScanner // out

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	C.g_scanner_destroy(_arg0)
}

// EOF returns true if the scanner has reached the end of the file or text
// buffer.
func (scanner *Scanner) EOF() bool {
	var _arg0 *C.GScanner // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.g_scanner_eof(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NextToken parses the next token just like g_scanner_peek_next_token() and
// also removes it from the input stream. The token data is placed in the
// @token, @value, @line, and @position fields of the #GScanner structure.
func (scanner *Scanner) NextToken() TokenType {
	var _arg0 *C.GScanner  // out
	var _cret C.GTokenType // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.g_scanner_get_next_token(_arg0)

	var _tokenType TokenType // out

	_tokenType = (TokenType)(_cret)

	return _tokenType
}

// InputFile prepares to scan a file.
func (scanner *Scanner) InputFile(inputFd int) {
	var _arg0 *C.GScanner // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = C.gint(inputFd)

	C.g_scanner_input_file(_arg0, _arg1)
}

// InputText prepares to scan a text buffer.
func (scanner *Scanner) InputText(text string, textLen uint) {
	var _arg0 *C.GScanner // out
	var _arg1 *C.gchar    // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(textLen)

	C.g_scanner_input_text(_arg0, _arg1, _arg2)
}

// LookupSymbol looks up a symbol in the current scope and return its value. If
// the symbol is not bound in the current scope, nil is returned.
func (scanner *Scanner) LookupSymbol(symbol string) interface{} {
	var _arg0 *C.GScanner // out
	var _arg1 *C.gchar    // out
	var _cret C.gpointer  // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_scanner_lookup_symbol(_arg0, _arg1)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

// PeekNextToken parses the next token, without removing it from the input
// stream. The token data is placed in the @next_token, @next_value, @next_line,
// and @next_position fields of the #GScanner structure.
//
// Note that, while the token is not removed from the input stream (i.e. the
// next call to g_scanner_get_next_token() will return the same token), it will
// not be reevaluated. This can lead to surprising results when changing scope
// or the scanner configuration after peeking the next token. Getting the next
// token after switching the scope or configuration will return whatever was
// peeked before, regardless of any symbols that may have been added or removed
// in the new scope.
func (scanner *Scanner) PeekNextToken() TokenType {
	var _arg0 *C.GScanner  // out
	var _cret C.GTokenType // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.g_scanner_peek_next_token(_arg0)

	var _tokenType TokenType // out

	_tokenType = (TokenType)(_cret)

	return _tokenType
}

// ScopeAddSymbol adds a symbol to the given scope.
func (scanner *Scanner) ScopeAddSymbol(scopeId uint, symbol string, value interface{}) {
	var _arg0 *C.GScanner // out
	var _arg1 C.guint     // out
	var _arg2 *C.gchar    // out
	var _arg3 C.gpointer  // out

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = C.guint(scopeId)
	_arg2 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (C.gpointer)(box.Assign(value))

	C.g_scanner_scope_add_symbol(_arg0, _arg1, _arg2, _arg3)
}

// ScopeLookupSymbol looks up a symbol in a scope and return its value. If the
// symbol is not bound in the scope, nil is returned.
func (scanner *Scanner) ScopeLookupSymbol(scopeId uint, symbol string) interface{} {
	var _arg0 *C.GScanner // out
	var _arg1 C.guint     // out
	var _arg2 *C.gchar    // out
	var _cret C.gpointer  // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = C.guint(scopeId)
	_arg2 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_scanner_scope_lookup_symbol(_arg0, _arg1, _arg2)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

// ScopeRemoveSymbol removes a symbol from a scope.
func (scanner *Scanner) ScopeRemoveSymbol(scopeId uint, symbol string) {
	var _arg0 *C.GScanner // out
	var _arg1 C.guint     // out
	var _arg2 *C.gchar    // out

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = C.guint(scopeId)
	_arg2 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_scanner_scope_remove_symbol(_arg0, _arg1, _arg2)
}

// SetScope sets the current scope.
func (scanner *Scanner) SetScope(scopeId uint) uint {
	var _arg0 *C.GScanner // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg1 = C.guint(scopeId)

	_cret = C.g_scanner_set_scope(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SyncFileOffset rewinds the filedescriptor to the current buffer position and
// blows the file read ahead buffer. This is useful for third party uses of the
// scanners filedescriptor, which hooks onto the current scanning position.
func (scanner *Scanner) SyncFileOffset() {
	var _arg0 *C.GScanner // out

	_arg0 = (*C.GScanner)(unsafe.Pointer(scanner))

	C.g_scanner_sync_file_offset(_arg0)
}

// ScannerConfig specifies the #GScanner parser configuration. Most settings can
// be changed during the parsing phase and will affect the lexical parsing of
// the next unpeeked token.
type ScannerConfig struct {
	native C.GScannerConfig
}

// WrapScannerConfig wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScannerConfig(ptr unsafe.Pointer) *ScannerConfig {
	return (*ScannerConfig)(ptr)
}

// Native returns the underlying C source pointer.
func (s *ScannerConfig) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
