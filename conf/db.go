package conf

import (
    "encoding/json"
    "io/ioutil"
)

type ConfDB struct {
    Host    string `json:"host"`
    Port    int    `json:"port"`
    DbName  string `json:"db-name"`
    Charset string `json:"charset"`
    User    string `json:"user"`
    Pass    string `json:"pass"`
}

func ReadConfDB() (*ConfDB, error) {

    const conffile = "/Users/yudai/GoProjects3/conf/db.json"

    conf := new(ConfDB)

    cValue, err := ioutil.ReadFile(conffile)
    if err != nil {
        return conf, err
    }

    err = json.Unmarshal([]byte(cValue), conf)
    if err != nil {
        return conf, err
    }
    return conf, nil
}

