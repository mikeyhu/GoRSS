module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    watch: {
      scripts: {
        files: ['src/**/*'],
        tasks: ['clean','copy'],
        options: {
          spawn: false,
        },
      },
    },

    clean: ["dist/"],

    copy: {
      files: {
        cwd: 'src/',
        src: '**/*',
        dest: 'dist/',
        expand: true
      }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-watch');


  grunt.registerTask('default', ['clean','copy']);
};
