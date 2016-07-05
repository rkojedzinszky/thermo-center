import 'jquery.cookie';

function sameOrigin(url) {
	var loc = window.location,
		a = document.createElement('a');
	a.href = url;

	return a.hostname == loc.hostname &&
		a.port == loc.port &&
		a.protocol == loc.protocol;
}

jQuery.ajaxSetup({
    processData: false,
    beforeSend: function(xhr, settings) {
        if (sameOrigin(settings.url) && !/^(GET|HEAD|OPTIONS|TRACE)$/.test(settings.type)) {
            xhr.setRequestHeader("X-CSRFToken", jQuery.cookie('csrftoken'));
        }
    }
});

jQuery.ajaxPrefilter(function(options) {
    if (options.data && options.type == 'GET' && typeof options.data !== "string") {
        options.data = jQuery.param(options.data, options.traditional);
        return;
    }

    if (typeof(options.data) == "object" && options.files == null) {
        options.data = JSON.stringify(options.data);
        options.contentType = 'application/json';
    }
});
