package z3

// #include <stdlib.h>
// #include "go-z3.h"
import "C"

// AST represents an AST value in Z3.
//
// AST memory management is automatically managed by the Context it
// is contained within. When the Context is freed, so are the AST nodes.
type AST struct {
	rawCtx C.Z3_context
	rawAST C.Z3_ast
}

// String returns a human-friendly string version of the AST.
func (a *AST) String() string {
	return C.GoString(C.Z3_ast_to_string(a.rawCtx, a.rawAST))
}

// DeclName returns the name of a declaration. The AST value must be a
// func declaration for this to work.
func (a *AST) DeclName() *Symbol {
	return &Symbol{
		rawCtx: a.rawCtx,
		rawSymbol: C.Z3_get_decl_name(
			a.rawCtx, C.Z3_to_func_decl(a.rawCtx, a.rawAST)),
	}
}

//-------------------------------------------------------------------
// Var, Literal Creation
//-------------------------------------------------------------------

// Const declares a variable. It is called "Const" since internally
// this is equivalent to create a function that always returns a constant
// value. From an initial user perspective this may be confusing but go-z3
// is following identical naming convention.
func (c *Context) Const(s *Symbol, typ *Sort) *AST {
	return &AST{
		rawCtx: c.raw,
		rawAST: C.Z3_mk_const(c.raw, s.rawSymbol, typ.rawSort),
	}
}

// Int creates an integer type that fit in a 32-bit integer.
//
// Maps: Z3_mk_int
func (c *Context) Int(v int, typ *Sort) *AST {
	return &AST{
		rawCtx: c.raw,
		rawAST: C.Z3_mk_int(c.raw, C.int(v), typ.rawSort),
	}
}

// Int64 creates an integer type that fit in a 64-bit integer.
//
// Maps: Z3_mk_int64
func (c *Context) Int64(v int64, typ *Sort) *AST {
	return &AST{
		rawCtx: c.raw,
		rawAST: C.Z3_mk_int64(c.raw, C.longlong(v), typ.rawSort),
	}
}

// True creates the value "true".
//
// Maps: Z3_mk_true
func (c *Context) True() *AST {
	return &AST{
		rawCtx: c.raw,
		rawAST: C.Z3_mk_true(c.raw),
	}
}

// False creates the value "false".
//
// Maps: Z3_mk_false
func (c *Context) False() *AST {
	return &AST{
		rawCtx: c.raw,
		rawAST: C.Z3_mk_false(c.raw),
	}
}

//-------------------------------------------------------------------
// Value Readers
//-------------------------------------------------------------------

// Int gets the integer value of this AST. The value must be able to fit
// into a 32-bit integer.
//
// Maps: Z3_get_numeral_int
func (a *AST) Int() int {
	var dst C.int
	C.Z3_get_numeral_int(a.rawCtx, a.rawAST, &dst)
	return int(dst)
}

// Int64 gets the integer value of this AST. The value must be able to fit
// into a 64-bit integer.
//
// Maps: Z3_get_numeral_int64
func (a *AST) Int64() int64 {
	var dst C.longlong
	C.Z3_get_numeral_int64(a.rawCtx, a.rawAST, &dst)
	return int64(dst)
}
