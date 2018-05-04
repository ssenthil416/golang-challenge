package dataservice

import (
    "encoding/gob"
    "os"
    "errors" 
)


type Dump struct {
   
   filePath string
}


func (d *Dump) Init(filePath string)error{
        //Basic Validation
        if filePath == "" || filePath == "null" || filePath == "NULL" {
                return errors.New("FilePath is missing")
        }
   d.filePath = filePath
   return nil
}

func (d *Dump)Get(name string)(int32,error){

        object := DumbList{}
        file, err := os.Open(d.filePath)
        if err == nil {
                decoder := gob.NewDecoder(file)
                if err := decoder.Decode(&object); err != nil {
                   return int32(0), err
                }
            
        }else{
          return int32(0), err
        }
        file.Close()

        for _,dump := range object.List {
            if dump.Name == name {
               return dump.Length, nil
            }
        }
    return int32(0),nil
}

func (d *Dump)Delete(name string)(string,error){
    return "test",nil
}



