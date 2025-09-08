package model

import "time"

type Type string

const (
TypeOpening Type = "opening"
TypeEnding  Type = "ending"
TypeOST     Type = "ost"
)

type Anime struct {
ID        int64     
Title     string    
CreatedAt time.Time 
}

type Singer struct {
ID        int64     
Name      string    
CreatedAt time.Time 
}

type Opening struct {
ID          int64     
AnimeID     int64     
SingerID    int64     
Type        Type      
Title       string    
OrderNumber int       
CreatedAt   time.Time 
}
