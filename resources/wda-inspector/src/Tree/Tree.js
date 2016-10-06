var Mustache = require('mustache');
var tpl = {
    item: require('./tpl/item.html'),
    lockOverlay: require('./tpl/lockOverlay.html'),
    error: require('./tpl/error.html')
};

var Tree = function($el) {

    var
        _$lockOverlay,
        _onElementSelect = function() {},
        _onElementFocus = function() {},
        _onElementBlur = function() {};

    var _buildList = function(elements) {
        var $list = $('<ul />');
        for (var i = 0; i < elements.length; ++i) {
            var item = elements[i];
            item.hasChildren = item.children? true : false;
            var
                $li = $('<li />'),
                $item = $(Mustache.render(tpl.item, item));

            $li.on('mouseenter.wda-inspector', '.el-type', item.rect, function(e) {
                _onElementFocus(e.data);

                return false;
            });
            $li.on('mouseleave.wda-inspector', '.el-type', function(e) {
                _onElementBlur(e.data);

                return false;
            });
            $li.on('click.wda-inspector', '.el-type', item, function(e) {
                _onElementSelect(e.data, e.data.rect);

                return false;
            });

            $li.append($item);

            if (item.children) {
                $li.append(_buildList(item.children));
            }

            $list.append($li);
        }

        return $list;
    };

    var _render = function($content) {
        $el
            .empty()
            .append($content, _$lockOverlay);
    };

    // Заблокировать
    this.lock = function() {
        _$lockOverlay.show();
    };

    // Разблокировать
    this.unlock = function() {
        _$lockOverlay.hide();
    };

    // Вывести сообщение об ошибке
    this.error = function(message) {
        var $content = $(Mustache.render(tpl.error, {message: message}));
        _render($content);
    };

    this.update = function(elements) {
        var $list = _buildList(elements.children);
        _render($list);
    };

    // Назначить обработчик на фокус элемента
    this.onElementFocus = function(handler) {
        _onElementFocus = handler;

    };

    // Назначить обработчик на блюр элемента
    this.onElementBlur = function(handler) {
        _onElementBlur = handler;
    };

    // Назначить обработчик на выбор элемента
    this.onElementSelect = function(handler) {
        _onElementSelect = handler;
    };

    // INIT
    (function() {
        _$lockOverlay = $(tpl.lockOverlay);
    })();

};

module.exports = Tree;