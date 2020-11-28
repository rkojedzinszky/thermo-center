'use strict';

import DefineMap from "can-define/map/map";

const View = DefineMap.extend({
    connectedCallback(element) {
        var self = this;

        self.app.onmessage = self.onmessage;

        let listener = function() {
            if (self.app.visible) {
                self.app.loaded = false;
                self.visible().then(() => {
                    self.app.loaded = true;
                });
            } else {
                self.hidden().then(() => {
                    self.app.loaded = false;
                });
            }
        };
        self.app.listenTo('visible', listener);

        listener();

        return function() {
            self.app.onmessage = null;
            self.app.stopListening('visible', listener);
            self.stopListening();
        };
    },
    onmessage(el) {
    },
    visible() {
        return Promise.resolve();
    },
    hidden() {
        return Promise.resolve();
    },
});

export {
    View
};
