package main

import (
	"github.com/davecgh/go-spew/spew"
	"fmt"
)
type Project struct {
	Id      int64  `json:"project_id"`
	Title   string `json:"title"`
	Name    string `json:"name"`
	Data    string `json:"data"`
	Commits string `json:"commits"`
}
func main() {



	o:=Project{1,"title","name","data","commits"}
	spew.Dump(o)
	fmt.Print("pppppppppppppppp")
}
