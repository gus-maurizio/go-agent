package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "github.com/sirupsen/logrus"
    // "github.com/logmatic/logmatic-go"
)

type conf struct {
    Logfile  string  `yaml:"logfile"`
    Time     int64   `yaml:"time"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("agent.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c conf

    // use JSONFormatter
    log.SetFormatter(&logmatic.JSONFormatter{})
    // log an event as usual with logrus
    log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1 }).Info("My first ssl event from golang")

    c.getConf()

    fmt.Println(c)
}
