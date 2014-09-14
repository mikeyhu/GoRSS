module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    watch: {
      scripts: {
        files: ['src/**/*', 'bower_components/**/*', 'test/**/*'],
        tasks: ['karma', 'clean', 'copy:to_dist', 'compass:dev'],
        options: {
          spawn: false,
        },
      },
    },

    clean: ["dist/"],

    copy: {
      to_dist: {
        files: [
          {
            cwd: 'src/',
            src: '**/*.html',
            dest: 'dist/',
            expand: true },
          {
            cwd: 'src/',
            src: '**/*.js',
            dest: 'dist/',
            expand: true },
          {
            cwd: 'bower_components/',
            src: '**/*',
            dest: 'dist/',
            expand: true
          }
        ]
      }
    },

    karma: {
      unit: {
        configFile: 'karma.conf.js'
      }
    },

    compass: {
      dist: {
        options: {
          sassDir: 'src/sass',
          cssDir: 'dist/css',
          environment: 'production',
          require: 'bootstrap-sass'
        }
      },
      dev: {
        options: {
          sassDir: 'src/sass',
          cssDir: 'dist/css',
          outputStyle: 'expanded',
          require: 'bootstrap-sass'
        }
      }
    },

    jslint: {
      client: {
        src: ['src/js/*.js'],
        directives: {
          node: true,
          browser: true,
          todo: true,
          indent: 2,
          predef: ['app','angular']
        },
        options: {
          log: 'out/server-lint.log',
          jslintXml: 'out/server-jslint.xml',
          errorsOnly: true,
          failOnError: false
        }
      },
      test: {
        src: ['test/js/*.js'],
        directives: {
          node: true,
          todo: true,
          indent: 2,
          predef: ['describe','it','inject','expect','beforeEach','module']
        },
        options: {
          log: 'out/server-lint.log',
          jslintXml: 'out/server-jslint.xml',
          errorsOnly: true,
          failOnError: false
        }
      }
    }

  });

  grunt.loadNpmTasks('grunt-jslint');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-compass');
  grunt.loadNpmTasks('grunt-karma');

  grunt.registerTask('default', ['karma:unit', 'clean', 'copy:to_dist', 'compass:dev']);
  grunt.registerTask('production', ['jslint', 'karma:unit', 'clean', 'copy:to_dist', 'compass:dist']);
};
