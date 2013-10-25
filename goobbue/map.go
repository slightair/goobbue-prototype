package goobbue

const (
    MapWidth  = 21
    MapHeight = 15
    TileWidth = 32
    TileHeight = 32
    MapChipWidth = 320
    MapChipHeight = 320
)

const (
    Wall  = 33
    Grass = 42
)

const (
    None = 0
    Block = 1
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
    
    for y := 0; y < MapHeight; y++ {
        for x := 0; x < MapWidth; x++ {
            index := y * MapWidth + x
            if y == 0 || y == MapHeight - 1 || x == 0 || x == MapWidth - 1 {
                mapData[index] = Wall
            } else {
                mapData[index] = Grass
            }
        }
    }
    
    for i, v := range mapData {
        if v == Wall {
            objectData[i] = Block
        } else {
            objectData[i] = None
        }
    }
    
    return []*Layer {
        NewLayer("object", objectData, false),
        NewLayer("map", mapData, true),
    }
}
