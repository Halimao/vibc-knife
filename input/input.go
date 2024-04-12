package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	isatty "github.com/mattn/go-isatty"
)

// GetConfirmation will request user give the confirmation from stdin.
// "y", "Y", "yes", "YES", and "Yes" all count as confirmations.
// If the input is not recognized, it returns false and a nil error.
func GetConfirmation(prompt string, r *bufio.Reader, w io.Writer) (bool, error) {
	if inputIsTty() {
		_, _ = fmt.Fprintf(w, "%s [y/N]: ", prompt)
	}

	response, err := readLineFromBuf(r)
	if err != nil {
		return false, err
	}

	response = strings.TrimSpace(response)
	if len(response) == 0 {
		return false, nil
	}

	response = strings.ToLower(response)
	if response[0] == 'y' {
		return true, nil
	}

	return false, nil
}

// GetString simply returns the trimmed string output of a given reader.
func GetString(prompt string, buf *bufio.Reader) (string, error) {
	if inputIsTty() && prompt != "" {
		fmt.Fprintf(os.Stderr, "> %s\n", prompt)
	}

	out, err := readLineFromBuf(buf)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out), nil
}

// inputIsTty returns true iff we have an interactive prompt,
// where we can disable echo and request to repeat the password.
// If false, we can optimize for piped input from another command
func inputIsTty() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
}

// readLineFromBuf reads one line from reader.
// Subsequent calls reuse the same buffer, so we don't lose
// any input when reading a password twice (to verify)
func readLineFromBuf(buf *bufio.Reader) (string, error) {
	pass, err := buf.ReadString('\n')

	switch {
	case errors.Is(err, io.EOF):
		// If by any chance the error is EOF, but we were actually able to read
		// something from the reader then don't return the EOF error.
		// If we didn't read anything from the reader and got the EOF error, then
		// it's safe to return EOF back to the caller.
		if len(pass) > 0 {
			// exit the switch statement
			break
		}
		return "", err

	case err != nil:
		return "", err
	}

	return strings.TrimSpace(pass), nil
}
