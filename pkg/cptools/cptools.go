package cptools

import (
  "strconv"
  "fmt"
  "os"
  "io"
  "strings"
  "errors"
  )

func Reverse(input string) (result string) {
  for _, c := range input {
    result = string(c) + result
  }
  return result
}

func Inspect(input string, digits bool) (count int, kind string) {
  if !digits {
    return len(input), "char"
  }
  return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
  for _, c := range input {
    _, err := strconv.Atoi(string(c))
    if err == nil {
      count++
    }
  }
  return count
}

func contains(s []string, t string) (res bool) {
  for _, x := range s {
    if x == t {
      return true
    }
  }
  return false
}

func Copy(src, dst string) error {
  in, err := os.Open(src)
  if err != nil {
    return err
  }
  defer in.Close()

  out, err := os.Create(dst)
  if err != nil {
    return err
  }
  defer out.Close()

  _, err = io.Copy(out, in)
  if err != nil {
    return err
  }
  return out.Close()
}


func NewFile(input string, lang string, templateDir string, noTemplate bool) (result string, error bool) {
  // Check if template dir is available
  files, err := os.ReadDir(templateDir)
  if err != nil {
    // fmt.Println("Err: ", err)
    result = fmt.Sprintf("Cannot access '%s' directory\n", templateDir)
    error = true
    return
  }
  // Check if there's available languages
  langs := []string{}
  for _, file := range files {
    langs = append(langs, strings.Split(file.Name(), ".")[0])
  }
  if !contains(langs, lang) {
    result = fmt.Sprintf("There is no template in '%s' language\n", lang)
    error = true
    return
  }
  // Check if there is already a file
  filename := fmt.Sprintf("%s.%s", input, lang)
  templateName := fmt.Sprintf("%s/%s.txt", templateDir, lang)
  if _, err := os.Stat(filename); err == nil {
    result = fmt.Sprintf("The file '%s' is already exists", filename)
    error = true
    return
  } else if !errors.Is(err, os.ErrNotExist) {
    result = err.Error()
    error = true
    return
  } 
  
  if err := Copy(templateName, filename); err != nil {
    fmt.Println("Error in copying template")
    error = true
    return
  }

  result = fmt.Sprintf("Create %s.%s successfully", input, lang)
  error = false
  return
}
