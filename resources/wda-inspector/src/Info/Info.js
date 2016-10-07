var Mustache = require('mustache');
var tpl = {
    info: require('./tpl/info.html'),
    error: require('./tpl/error.html')
};

var Info = function($el) {

    var _previousColor;

    var _render = function($content) {
        $el
            .empty()
            .append($content);
    };

    var _blink = function(color) {
        $el
            .stop()
            .animate({
                'background-color': color
            }, 100)
            .animate({
                'background-color': _previousColor
            }, 500)
    };

    this.update = function(info) {
        var $content = $(Mustache.render(tpl.info, {info: info}));
        _render($content);
        _blink('rgba(250, 255, 189, 0.8)');
    };

    this.error = function(message) {
        var $content = $(Mustache.render(tpl.error, {message: message}));
        _render($content);
        _blink('rgba(255, 150, 150, 0.8)');
    };

    // INIT
    (function() {
        _previousColor = $el.css('background-color');
    })();

};

module.exports = Info;