// Code generated by sysl DO NOT EDIT.
package pokeapi

import (
	"fmt"
)

// Error fulfills the error type interface for Error_
func (s Error_) Error() string {
	type plain Error_

	return fmt.Sprintf("%+v", plain(s))
}