module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    watch: {
      scripts: {
        files: ['src/**/*', 'bower_components/**/*'],
        tasks: ['clean','copy:to_dist'],
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
            src: '**/*',
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
        }
  });

  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-jasmine');
  grunt.loadNpmTasks('grunt-karma');

  grunt.registerTask('default', ['clean','copy:to_dist']);
};
