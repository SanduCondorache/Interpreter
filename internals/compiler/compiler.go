package compiler

import (
	"github.com/SanduCondorache/Interpreter/ast"
	"github.com/SanduCondorache/Interpreter/internals/code"
	"github.com/SanduCondorache/Interpreter/internals/object"
)

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)

			if err != nil {
				return err
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)

		if err != nil {
			return err
		}
	case *ast.InfixExpression:
		err := c.Compile(node.Left)

		if err != nil {
			return err
		}

		err = c.Compile(node.Right)

		if err != nil {
			return nil
		}

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))
	}

	return nil
}

func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)

	return len(c.constants) - 1
}

func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)

	return pos
}

func (c *Compiler) addInstruction(ins []byte) int {
	powNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)

	return powNewInstruction
}
