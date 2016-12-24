package listener

import (
  "sync"
)

type GroupListener struct {
  Users []*UserListener
}

var instance *GroupListener
var once sync.Once

func GetInstance() *GroupListener {
  once.Do(func(){
    instance = &GroupListener{make([]*UserListener, 0)}
  })
  return instance
}
