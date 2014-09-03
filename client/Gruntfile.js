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
            expand: true
          },
          {
            cwd: 'bower_components/',
            src: '**/*',
            dest: 'dist/',
            expand: true
          }
        ]
      }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-watch');


  grunt.registerTask('default', ['clean','copy:to_dist']);
};
