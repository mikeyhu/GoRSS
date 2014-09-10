module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    watch: {
      scripts: {
        files: ['src/**/*', 'bower_components/**/*','test/**/*'],
        tasks: ['karma','clean','copy:to_dist','compass:dev'],
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
          environment: 'production'
        }
      },
      dev: {
        options: {
          sassDir: 'src/sass',
          cssDir: 'dist/css'
        }
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-compass');
  grunt.loadNpmTasks('grunt-karma');

  grunt.registerTask('default', ['karma:unit','clean','copy:to_dist','compass:dev']);
};
