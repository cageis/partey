package src

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type PartialsBuildCommand struct {
    aggregateFile string
    partialsDir string
    commentChars string
}

func NewPartialsBuildCommand(aggregateFile string, partialsDir string, commentChars string) PartialsBuildCommand {
    return PartialsBuildCommand{aggregateFile, partialsDir, commentChars}
}

func (p PartialsBuildCommand) GetStartFlag() string {
    return fmt.Sprintf("%s PARTIALS>>>>>", p.commentChars)
}

func (p PartialsBuildCommand) GetEndFlag() string {
    return fmt.Sprintf("%s PARTIALS<<<<<", p.commentChars)
}

func (p PartialsBuildCommand) Run() {
    path, err := filepath.Abs(p.aggregateFile)
    if err != nil {
        panic(err)
    }
    agg, err := os.ReadFile(path)
    output := string(agg)
    if err != nil {
        panic(err)
    }

    startIndex := strings.Index(output, p.GetStartFlag())
    endIndex := strings.Index(output, p.GetEndFlag())

    if startIndex != -1 && endIndex != -1 {
        before := agg[:startIndex]
        after := agg[endIndex + len(p.GetEndFlag()):]
        output = string(before) + string(after)
    }

    output += p.GetStartFlag()
    output += "\n"

    files, err := ioutil.ReadDir(p.partialsDir)

    if err != nil {
        panic(err)
    }

    // Each file: read contents into var to be written later.
    for _, file := range files {
        if file.IsDir() == true {
            continue
        }

        fileContents, _ := os.ReadFile(filepath.Join(p.partialsDir, file.Name()))
        output += string(fileContents)
        output += "\n"
    }

    output += p.GetEndFlag()
    output += "\n"

    err = ioutil.WriteFile(p.aggregateFile, []byte(output), 0644)

    if err != nil {
        panic(err)
    }
}
