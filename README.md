/*
There are four familiar printing function defined by their output destination.
    Print, Println, Printf   -> the screen (os.Stdout)
    Sprint, Sprintln, Sprinf -> returns a string, prints nothing
    Fprint, Fprintln, Fprintf -> io.Writer writes to any w: screen, file, or memory buffer. ex: fmt.Fprint(w, ...)
    Append, Appendln, Appendf  -> appends to a []byte
*/