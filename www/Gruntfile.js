module.exports = function (grunt) {
        grunt.initConfig({
                "steal-build": {
                        "bundle": {
                                options: {
                                        system: {
                                                config: "bower.json!bower"
                                        }
                                }
                        }
                }
        });

        grunt.loadNpmTasks('steal-tools');
        grunt.registerTask('build', ['steal-build']);
};
