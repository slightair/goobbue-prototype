module.exports = (grunt) ->
  grunt.initConfig
    pkg: grunt.file.readJSON 'package.json'
    watch:
      files: ['src/coffee/**/*.coffee']
      tasks: ['build']
    coffee:
      compile:
        options:
          join: true
          sourceMap: false
        files:
          'src/js/goobbue.js': [
            'src/coffee/app.coffee'
          ]
    uglify:
      compile:
        files: [
          expand: true
          cwd: 'src/js/'
          src: ['**/*.js']
          dest: 'public/js/'
          ext: '-min.js'
        ]
    bower:
      install:
        options:
          targetDir: './public/lib'
          layout: 'byType'
          install: true
          verbose: true
          cleanTargetDir: true
          cleanBowerDir: false
  grunt.loadNpmTasks 'grunt-contrib-coffee'
  grunt.loadNpmTasks 'grunt-contrib-watch'
  grunt.loadNpmTasks 'grunt-contrib-uglify'
  grunt.loadNpmTasks 'grunt-bower-task'
  grunt.registerTask 'build', ['coffee', 'uglify']
  grunt.registerTask 'default', ['watch']
  return
