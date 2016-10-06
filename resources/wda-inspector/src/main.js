var Screen = require('./Screen/Screen.js');
var Tree = require('./Tree/Tree.js');
var Info = require('./Info/Info.js');

$(function () {

    var screen = new Screen($('#screenshot'));
    $.ajax({
        method: 'get',
        url: '/screenshot',
        dataType: 'json',
        beforeSend: function() {
            screen.lock();
        },
        complete: function() {
            screen.unlock();
        },
        success: function(data) {
            screen.update(data.img);
        },
        error: function() {
            screen.error('Не удалось загрузить экран');
        }
    });

    var info = new Info($('#element-data'));

    var tree = new Tree($('#tree'));
    tree.onElementFocus(function(rect) {
        screen.highlight(
            rect.origin.x,
            rect.origin.y,
            rect.size.width,
            rect.size.height
        );
    });
    tree.onElementBlur(screen.highlightSelection);
    tree.onElementSelect(function(infoData, rect) {
        screen.select(
            rect.origin.x,
            rect.origin.y,
            rect.size.width,
            rect.size.height
        );
        info.update(infoData);
    });
    $.ajax({
        method: 'get',
        url: '/source',
        dataType: 'json',
        beforeSend: function() {
            tree.lock();
        },
        complete: function() {
            tree.unlock();
        },
        success: function(data) {
            tree.update(data.tree);
        },
        error: function() {
            tree.error('Не удалось загрузить дерево элементов');
        }
    });

});
