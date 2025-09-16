    // parent/embed.go
    package main

    import _ "embed"

    //go:embed hello.txt
    var EmbeddedFile string