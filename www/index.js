import THSensor from 'models/Thsensor';
import 'summary/';
import $ from 'jquery';
import stache from 'can/view/stache/';
import 'tastypie';
import 'bootstrap.css!';

THSensor.findAll({'order_by': 'id'}).then(function(res) {
	$("#main").html(stache('<sensor-list />')({sensors: res}));
	can.each(res, (s) => s.startRefresh());
});
