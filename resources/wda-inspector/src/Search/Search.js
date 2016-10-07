var tpl = {
    form: require('./tpl/form.html')
};

var Search = function($el, options) {

    var _isLocked = false;

    var _render = function($content) {
        $el
            .empty()
            .append($content);
    };

    var _lock = function() {
        if (_isLocked) {
            return false;
        }

        $el
            .find('input')
            .attr('readonly', 'readonly');
        _isLocked = true;

        return true;
    };

    var _unlock = function() {
        $el
            .find('input')
            .removeAttr('readonly');
        _isLocked = false;
    };

    // INIT
    (function() {
        var $form = $(tpl.form);
        $form.on('submit', function() {
            if (_lock()) {
                var locator = $(this).find('input[name="value"]').val();
                $.ajax({
                    method: 'POST',
                    url: options.url,
                    data: $(this).serialize(),
                    dataType: 'json',
                    success: options.success,
                    error: function(jqXHR) {
                        if (jqXHR.status === 400) {
                            options.notFound(locator);
                        } else {
                            options.error(locator);
                        }
                    },
                    complete: function() {
                        _unlock();
                    }
                });
            }

            return false;
        });
        _render($form);
    })();

};

module.exports = Search;