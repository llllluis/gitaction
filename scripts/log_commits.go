package main
 
import (

  "fmt"

  "os"

  "os/exec"

  "path/filepath"

  "time"

)
 
func main() {

  cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an : %s")

  out, err := cmd.Output()

  if err != nil {

    fmt.Printf("Error ejecutandogit log %v\n",err)

    os.Exit(1)

  }
 
  //Creació de la carpeta log

  logDir := filepath.Join("..", "log")

  if _, err := os.Stat(logDir); os.IsNotExist(err) {

    //Definim permisos d'escriptura

    err = os.Mkdir(logDir, 0755)

    if err != nil {

      fmt.Printf("Error creando directorio %s: %v\n",logDir,err);

      os.Exit(1)

    }

  }
 
  //Generar nom de l'arxiu

  currentTime := time.Now().Format("2006-01-02_15:04:05") //Mascara de YYYY-mm-dd HH-ii-ss

  logFile := filepath.Join(logDir,fmt.Sprintf("commits_%s.txt",currentTime))
 
  //Escriure l'arxiu

  err := fmt.Sprintf("S'han escrit els ultims 3 commits del repositori:\n%s",string(out))

  if err != nil {

    fmt.Printf("S'ha produit un error creant en %s %v\n",logFile,err)

    os.Exit(1)

  }
 
  fmt.Printf("S'ha creat l'arxiu de log en %s\n",logFile)

}
                 
                    
