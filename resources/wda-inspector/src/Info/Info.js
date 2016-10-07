var Mustache = require('mustache');
var tpl = {
    info: require('./tpl/info.html')
};

var Info = function($el) {


    var _render = function($content) {
        $el
            .empty()
            .append($content);
    };

    this.update = function(info) {
        var $content = $(Mustache.render(tpl.info, {info: info}));
        _render($content);
    };

};

module.exports = Info;