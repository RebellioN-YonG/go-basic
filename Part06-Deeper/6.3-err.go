package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// write a file
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Errors are values.")

	if err != nil {
		// once writing , we should close the file
		// to avoid leaking resources
		f.Close()
		return err
	}
	f.Close()
	return err

}

// the defer keyword is used to call a func
// at the end of the current function
func deferproverb(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	// using defer to simplify the code
	// everytime the function returns, the f.close() will be called
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err

}

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.w, s)
}

func safeproverb(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()

	sw := safeWriter{w: f}

	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")

	return sw.err

}

const rows, cols = 9, 9

type grid [rows][cols]int8

func (g *grid) Set(row, col int, digit int8) error {
	if !inBounds(row, col) {
		return ErrBounds
	}

	if !validDigit(digit) {
		return ErrDigit
	}

	g[row][col] = digit
	return nil
}

func inBounds(row, col int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if col < 0 || col >= cols {
		return false
	}

	return true
}

// sudoku2
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// custom errors types
// error type is a in-builtinterface
type error interface {
	Error() string
}

type SudoError []error

func (se SudoError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *grid) Set2(row, col int, digit int8) error {

	var errs SudoError

	if !inBounds(row, col) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][col] = digit
	return nil
}

func main() {

	// when func returns an error
	// the last value of the func is the error
	files, err := os.ReadDir(`.`)

	// after calling a func, we should checking errors immediatly
	// if there's no error, err value will return nil
	// else, other values returned by the func is not guaranteed
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	err2 := deferproverb(`proverbs2.txt`)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

	// handle errors gracefully
	err3 := safeproverb(`proverbs3.txt`)
	if err3 != nil {
		fmt.Println(err3)
		os.Exit(1)
	}

	var g grid
	err4 := g.Set(0, 0, 15)

	if err4 != nil {
		switch err4 {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de parametres hors limites")
		default:
			fmt.Println(err)
		}
		// os.Exit(1)
	}

	// err5.(SudoError) is a type assertion
	// errs's type is SudoError
	// If it is, ok will be true,
	// and errs will be a SudokuError, giving access to the
	err5 := g.Set2(10, 0, 15)
	if err5 != nil {
		if errs, ok := err5.(SudoError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		// os.Exit(1)
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("i forgot my towel")

}
