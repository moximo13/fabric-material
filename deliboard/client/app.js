// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_update_deliverer").hide();
	$("#success_create").hide();
	$("#error_update_deliverer").hide();
	$("#error_query").hide();
	$("#success_update_finish").hide();
	$("#error_update_finish").hide();
	
	$scope.queryAllRequest = function(){

		appFactory.queryAllRequest(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_request = array;
		});
	}

	 $scope.queryRequest = function(){

	 	var id = $scope.request_id;

	 	appFactory.queryRequest(id, function(data){
	 		$scope.query_request = data;

	 		if ($scope.query_request == "Could not locate Request"){
	 			console.log()
	 			$("#error_query").show();
	 		} else{
				$("#error_query").hide();
	 		}
	 	});
	 }

	$scope.buildRequest = function(){

		appFactory.buildRequest($scope.request, function(data){
			$scope.create_request = data;
			$("#success_create").show();
		});
	}

	$scope.updateDeliverer = function(){

		appFactory.updateDeliverer($scope.update, function(data){
			$scope.update_deliverer = data;
			if ($scope.update_deliverer == "Error: no request found"){
				$("#error_update").show();
				$("#success_update").hide();
			} else{
				$("#success_update").show();
				$("#error_update").hide();
			}
		});
	}

	$scope.updateFinish = function(){

		appFactory.updateFinish($scope.update, function(data){
			$scope.update_finish = data;
			if ($scope.update_finish == "Error: no request found"){
				$("#error_update").show();
				$("#success_update").hide();
			} else{
				$("#success_update").show();
				$("#error_update").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllRequest = function(callback){

    	$http.get('/get_all_request/').success(function(output){
			callback(output)
		});
	}

	 factory.queryRequest = function(id, callback){
     	$http.get('/get_request/'+id).success(function(output){
	 		callback(output)
	 	});
	 }

	factory.buildRequest = function(data, callback){

		// data.location = data.longitude + ", "+ data.latitude;

		var request = data.id + "-" + data.timestamp + "-" + data.sender_name + "-" + data.sender_address + "-" + data.receive_name + "-" + data.receive_address + "-" + data.deliverer_name + "-" + data.price + "-" + data.status + "-" + data.code;

    	$http.get('/add_request/'+request).success(function(output){
			callback(output)
		});
	}

	factory.updateDeliverer = function(data, callback){

		var update = data.id + "-" + data.name;

    	$http.get('/update_deliverer/'+update).success(function(output){
			callback(output)
		});
	}

	factory.updateFinish = function(data, callback){

		var update = data.id + "-" + data.code;

    	$http.get('/update_finish/'+update).success(function(output){
			callback(output)
		});
	}

	return factory;
});


