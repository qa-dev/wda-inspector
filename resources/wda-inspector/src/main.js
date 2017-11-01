var Screen = require('./Screen/Screen.js');
var Tree = require('./Tree/Tree.js');
var Info = require('./Info/Info.js');
var Search = require('./Search/Search.js');

$(function () {

    var screen = new Screen($('.wda_screen_container'));
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

    var info = new Info($('.wda_info_container'));

    var tree = new Tree($('.wda_tree_container'));
    tree.onElementFocus(function(rect) {
        screen.highlight(
            rect.x,
            rect.y,
            rect.width,
            rect.height
        );
    });
    tree.onElementBlur(screen.highlightSelection);
    tree.onElementSelect(function(infoData, rect) {
        screen.select(
            rect.x,
            rect.y,
            rect.width,
            rect.height
        );
        screen.highlightSelection();
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

    var search = new Search($('#navbar'), {
        url: '/find',
        success: function(data) {
            tree.select(data.value, data.type); // todo надо бы id
        },
        notFound: function(locator) {
            info.error(locator + ' not found!');
        },
        error: function(locator) {
            info.error('Wrong locator ' + locator);
        }
    });

});
