//SPDX-License-Identifier: Apache-2.0

var request = require('./controller.js');

module.exports = function(app){

  app.get('/get_request/:id', function(req, res){
    request.get_request(req, res);
  });
  app.get('/add_request/:request', function(req, res){
    request.add_request(req, res);
  });
  app.get('/get_all_request', function(req, res){
    request.get_all_request(req, res);
  });
  app.get('/update_deliverer/:deliverer', function(req, res){
    request.update_deliverer(req, res);
  });
  app.get('/update_finish/:finish', function(req, res){
    request.update_finish(req, res);
  });
}
