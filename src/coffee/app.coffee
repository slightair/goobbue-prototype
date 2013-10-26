window.Goobbue = {}

$ ->
  Goobbue.CE = CE.defines('world-canvas').
    extend(Input).
    extend(Tiled).
    ready -> Goobbue.CE.Scene.call "WorldScene"
  Goobbue.settings =
    mapFile: '/api/map.json'
    characterSize: 24
    tileSize: 32
    initialPositionX: 1
    initialPositionY: 1

  Goobbue.CE.Scene.new
    name: "WorldScene"
    materials:
      images:
        mapchip: '/img/mapchip.png'
        brave: '/img/brave.png'
    ready: (stage) ->
      arrangeMap = =>
        deferred = $.Deferred()
        
        map = @createElement()
        @map = Goobbue.CE.Tiled.new()
        @map.ready ->
          stage.append map
          deferred.resolve()
        @map.load(@, map, Goobbue.settings.mapFile)
        
        deferred.promise()
      
      arrangeBrave = =>
        deferred = $.Deferred()
        margin = (Goobbue.settings.tileSize - Goobbue.settings.characterSize) / 2
        
        @brave = @createElement()
        @brave.moveX = (x) -> @setPosition(@positionX + x, @positionY)
        @brave.moveY = (y) -> @setPosition(@positionX, @positionY + y)
        @brave.setPosition = (x, y) =>
          movable = true
          
          if @map
            tileID = @map.getTileInMap(x, y)
            movable = false if @map.getDataLayers()[0][tileID] > 0
          
          if movable
            @brave.positionX = x
            @brave.positionY = y
          
          @brave.x = @brave.positionX * Goobbue.settings.tileSize
          @brave.y = @brave.positionY * Goobbue.settings.tileSize
          
          stage.refresh()
        
        positionX = Goobbue.settings.initialPositionX
        positionY = Goobbue.settings.initialPositionY
        
        @brave.setPosition positionX, positionY
        @brave.drawImage 'brave', margin, margin
        stage.append @brave
        
        deferred.resolve()
      
      arrangeMap().done => arrangeBrave().done
      
      Goobbue.CE.Input.keyDown 87, (e) => @brave.moveY(-1)
      Goobbue.CE.Input.keyDown 68, (e) => @brave.moveX(1)
      Goobbue.CE.Input.keyDown 83, (e) => @brave.moveY(1)
      Goobbue.CE.Input.keyDown 65, (e) => @brave.moveX(-1)
      
      return