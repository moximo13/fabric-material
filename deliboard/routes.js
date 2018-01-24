//SPDX-License-Identifier: Apache-2.0

var tuna = require('./controller.js');

module.exports = function(app){

  app.get('/get_request/:id', function(req, res){
    tuna.get_request(req, res);
  });
  app.get('/add_request/:tuna', function(req, res){
    tuna.add_request(req, res);
  });
  app.get('/get_all_request', function(req, res){
    tuna.get_all_request(req, res);
  });
  //app.get('/change_holder/:holder', function(req, res){
  //  tuna.change_holder(req, res);
  //});
}
