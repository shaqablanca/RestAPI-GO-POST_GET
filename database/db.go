package database

import "orness/models"

var Db map[uint32]models.Note = make(map[uint32]models.Note)      // Internal Memory by using Mapping DataBase for Notes
var RDb map[string][]models.Note = make(map[string][]models.Note) // Relational DataBase
