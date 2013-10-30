window.Goobbue = {}

$ ->
  Goobbue.CE = CE.defines('world-canvas').
    extend(Input).
    extend(Tiled).
    ready -> Goobbue.CE.Scene.call "WorldScene"
  Goobbue.settings =
    mapFile: '/api/map.json'
    characterSize: 24
  Goobbue.key =
    W: 87
    D: 68
    S: 83
    A: 65
  Goobbue.objects =
    startPosition: 1
    goalPosition: 2
    block: 3

  Goobbue.CE.Scene.new
    name: "WorldScene"
    materials:
      images:
        mapchip: '/img/mapchip.png'
        brave: '/img/brave.png'
    gotoNextStage: ->
      Goobbue.CE.Input.clearKeys([Goobbue.key.W, Goobbue.key.A, Goobbue.key.S, Goobbue.key.D])
      Goobbue.CE.Scene.call "WorldScene"
      return
    ready: (stage) ->
      arrangeMap = =>
        deferred = $.Deferred()
        
        map = @createElement()
        @map = Goobbue.CE.Tiled.new()
        @map.ready =>
          stage.append map
          
          mapWidth = @map.getWidthPixel() / @map.getTileWidth()
          collisionLayer = @map.getDataLayers()[0]
          
          startIndex = collisionLayer.indexOf(Goobbue.objects.startPosition)
          @braveInitialPositionX = startIndex % mapWidth
          @braveInitialPositionY = Math.floor(startIndex / mapWidth)
          
          goalIndex = collisionLayer.indexOf(Goobbue.objects.goalPosition)
          @goalPositionX = goalIndex % mapWidth
          @goalPositionY = Math.floor(goalIndex / mapWidth)
          
          deferred.resolve()
        @map.load(@, map, Goobbue.settings.mapFile)
        
        deferred.promise()
      
      arrangeBrave = =>
        deferred = $.Deferred()
        tileSize = @map.getTileWidth()
        margin = (tileSize - Goobbue.settings.characterSize) / 2
        
        @brave = @createElement()
        @brave.moveX = (x) -> @setPosition(@positionX + x, @positionY)
        @brave.moveY = (y) -> @setPosition(@positionX, @positionY + y)
        @brave.setPosition = (x, y) =>
          movable = true
          
          if @map
            tileID = @map.getTileInMap(x, y)
            movable = false if @map.getDataLayers()[0][tileID] == Goobbue.objects.block
          
          if x == @goalPositionX && y == @goalPositionY
            @gotoNextStage()
          
          if movable
            @brave.positionX = x
            @brave.positionY = y
          
          @brave.x = @brave.positionX * tileSize
          @brave.y = @brave.positionY * tileSize
          
          stage.refresh()
        
        positionX = @braveInitialPositionX
        positionY = @braveInitialPositionY
        
        @brave.setPosition positionX, positionY
        @brave.drawImage 'brave', margin, margin
        stage.append @brave
        
        deferred.resolve()
      
      arrangeMap().done => arrangeBrave().done
      
      Goobbue.CE.Input.keyDown Goobbue.key.W, (e) => @brave.moveY(-1)
      Goobbue.CE.Input.keyDown Goobbue.key.D, (e) => @brave.moveX(1)
      Goobbue.CE.Input.keyDown Goobbue.key.S, (e) => @brave.moveY(1)
      Goobbue.CE.Input.keyDown Goobbue.key.A, (e) => @brave.moveX(-1)
      
      return