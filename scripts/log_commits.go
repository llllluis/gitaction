package main

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "time"
)

func main() {
  cmd := exec.Commnad("git","log","-n","3","--pretty=format:%h -%an, : %s"
  out, err := cmd.Output()
  if err != nil {
      fmt.Printf("Error ejecutando git log %v\n", err)
      os.Exit(1)
    }

  //creacio carpeta log
  logDir := filepath.Join("..","log")
  if _, err := os.Stat(logDir); os.IsNotExist(err) {
    //definim permisos d'escriptura
    err := os.Mkdir(logDir, 0755)
    if err != nil {
      fmt.Printf("Error creando directorio %s: %v\n", logDir, err)
      os.Exit(1)
    }
  }

  //generar nom arxiu
  currentTime := time.Now().Format("2006-01-02_15:04:05") //Mascara de YYYY-mm-DD HH-ii-ss
  logfile := filepath.Join(logDir,fmt.Sprintf("Commits_%s.txt",currentTime)

  //escriure arxiu
  err = fmt.Sprintf("S'han escrit els últim 3 commits del repositori:\n%s",string(out))
  if err != nil {
    fmt.Printf("S'ha produit un errer creant en %s %v\n", logfile,err)
    os.Exit(1)
  }

  fmt.Printf("S'ha creat l'arxiu log en %s\n",logfile)
}
                      
                      
