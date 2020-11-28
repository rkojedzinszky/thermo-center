'use strict';

import DefineMap from "can-define/map/map";

const View = DefineMap.extend({
    connectedCallback(element) {
        var self = this;

        self.app.onmessage = self.onmessage;

        let listener = function() {
            if (self.app.visible) {
                self.visible();
            } else {
                self.hidden();
            }
        };
        self.app.listenTo('visible', listener);

        self.visible();

        return function() {
            self.app.onmessage = null;
            self.app.stopListening('visible', listener);
            self.stopListening();
        };
    },
    onmessage(el) {
    },
    visible() {
    },
    hidden() {
    },
});

export {
    View
};
