package goobbue

import "math/rand"

const (
    MapWidth  = 21
    MapHeight = 15
    TileWidth = 32
    TileHeight = 32
    MapChipWidth = 320
    MapChipHeight = 320
)

const (
    Void  = 0
    Wall  = 33
    Grass = 42
)

const (
    None  = iota
    Start = iota
    Goal  = iota
    Block = iota
)

type Map struct {
    Width int                       `json:"width"`
    Height int                      `json:"height"`
    Layers []*Layer                 `json:"layers"`
    Orientation string              `json:"orientation"`
    TileWidth int                   `json:"tilewidth"`
    TileHeight int                  `json:"tileheight"`
    TileSets []*TileSet             `json:"tilesets"`
    Properties map[string]*Property `json:"properties"`
    Version int                     `json:"version"`
}

type Layer struct {
    Name string                     `json:"name"`
    Type string                     `json:"type"`
    X int                           `json:"x"`
    Y int                           `json:"y"`
    Width int                       `json:"width"`
    Height int                      `json:"height"`
    Data []int                      `json:"data"`
    Opacity int                     `json:"opacity"`
    Visible bool                    `json:"visible"`
}

type TileSet struct {
    FirstGID int                    `json:"firstgid"`
    Image string                    `json:"image"`
    ImageWidth int                  `json:"imagewidth"`
    ImageHeight int                 `json:"imageheight"`
    Margin int                      `json:"margin"`
    Name string                     `json:"name"`
    Properties map[string]*Property `json:"properties"`
    Spacing int                     `json:"spacing"`
    TileWidth int                   `json:"tilewidth"`
    TileHeight int                  `json:"tileheight"`
}

type Property string

func NewMap() *Map {
    return &Map{
        MapWidth, MapHeight,
        createLayers(),
        "orthogonal",
        TileWidth, TileHeight,
        []*TileSet { NewTileSet("mapchip") },
        map[string]*Property {},
        1,
    }
}

func NewLayer(name string, data []int, visible bool) *Layer {
    return &Layer{
        name,
        "tilelayer",
        0, 0,
        MapWidth, MapHeight,
        data,
        1,
        visible,
    }
}

func NewTileSet(name string) *TileSet {
    return &TileSet{
        1,
        "mapchip.png",
        MapChipWidth, MapChipHeight,
        0,
        name,
        map[string]*Property {},
        0,
        TileWidth, TileHeight,
    }
}

func createLayers() []*Layer {
    mapData := make([]int, MapWidth * MapHeight)
    objectData := make([]int, MapWidth * MapHeight)
    
    createMaze(mapData)
    
    for i, v := range mapData {
        if v == Wall {
            objectData[i] = Block
        } else {
            objectData[i] = None
        }
    }
    
    blindAlleys := searchBlindAlleys(mapData)
    
    for i := len(blindAlleys); i > 1; i-- {
        a := i - 1
        b := rand.Intn(i)
        blindAlleys[a], blindAlleys[b] = blindAlleys[b], blindAlleys[a]
    }
    
    startIndex := blindAlleys[0]
    goalIndex := blindAlleys[len(blindAlleys)-1]
    
    objectData[startIndex] = Start
    objectData[goalIndex] = Goal
    
    return []*Layer {
        NewLayer("object", objectData, false),
        NewLayer("map", mapData, true),
    }
}

func createMaze(mapData []int) {
    for y := 1; y < MapHeight - 1; y++ {
        for x := 1; x < MapWidth - 1; x++ {
            mapData[mapIndex(x, y)] = Wall
        }
    }
    
    startX := rand.Intn((MapWidth - 2) / 2) * 2 + 2
    startY := rand.Intn((MapHeight - 2) / 2) * 2 + 2
    mapData[mapIndex(startX, startY)] = Void
    
    buildRoad(mapData, startX, startY)
    
    for i, o := range mapData {
        if o == Void {
            mapData[i] = Grass
        }
    }
}

func searchBlindAlleys(mapData []int) []int {
    blindAlleys := []int{}
    for y := 2; y < MapHeight - 2; y++ {
        for x := 2; x < MapWidth - 2; x++ {
            index := mapIndex(x, y)
            
            if mapData[index] == Wall {
                continue
            }
            
            count := 0
            if mapData[mapIndex(x - 1, y)] == Wall { count++ }
            if mapData[mapIndex(x + 1, y)] == Wall { count++ }
            if mapData[mapIndex(x, y - 1)] == Wall { count++ }
            if mapData[mapIndex(x, y + 1)] == Wall { count++ }
            
            if count == 3 {
                blindAlleys = append(blindAlleys, index)
            }
        }
    }
    return blindAlleys
}

func mapIndex(x, y int) int {
    return y * MapWidth + x
}

type buildRoadFunc func(mapData []int, startX, startY int)

func buildRoad(mapData []int, startX, startY int) {
    functions := []buildRoadFunc {
        buidRoadToTop,
        buidRoadToRight,
        buidRoadToLeft,
        buidRoadToBottom,
    }
    
    for i := len(functions); i > 1; i-- {
        a := i - 1
        b := rand.Intn(i)
        functions[a], functions[b] = functions[b], functions[a]
    }
    
    for _, f := range functions {
        f(mapData, startX, startY)
    }
}

func buidRoadToTop(mapData []int, startX, startY int) {
    target := mapIndex(startX, startY - 2)
    if mapData[target] != Void {
        mapData[mapIndex(startX, startY - 1)] = Void
        mapData[target] = Void
        buildRoad(mapData, startX, startY - 2)
    }
}

func buidRoadToRight(mapData []int, startX, startY int) {
    target := mapIndex(startX + 2, startY)
    if mapData[target] != Void {
        mapData[mapIndex(startX + 1, startY)] = Void
        mapData[target] = Void
        buildRoad(mapData, startX + 2, startY)
    }
}

func buidRoadToBottom(mapData []int, startX, startY int) {
    target := mapIndex(startX, startY + 2)
    if mapData[target] != Void {
        mapData[mapIndex(startX, startY + 1)] = Void
        mapData[target] = Void
        buildRoad(mapData, startX, startY + 2)
    }
}

func buidRoadToLeft(mapData []int, startX, startY int) {
    target := mapIndex(startX - 2, startY)
    if mapData[target] != Void {
        mapData[mapIndex(startX - 1, startY)] = Void
        mapData[target] = Void
        buildRoad(mapData, startX - 2, startY)
    }
}
