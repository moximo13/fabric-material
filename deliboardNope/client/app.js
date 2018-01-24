// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
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
			$scope.all_tuna = array;
		});
	}

	// $scope.queryTuna = function(){

	// 	var id = $scope.tuna_id;

	// 	appFactory.queryTuna(id, function(data){
	// 		$scope.query_tuna = data;

	// 		if ($scope.query_tuna == "Could not locate tuna"){
	// 			console.log()
	// 			$("#error_query").show();
	// 		} else{
	// 			$("#error_query").hide();
	// 		}
	// 	});
	// }

	$scope.buildRequest = function(){

		appFactory.buildRequest($scope.tuna, function(data){
			$scope.create_tuna = data;
			$("#success_create").show();
		});
	}

	// $scope.changeHolder = function(){

	// 	appFactory.changeHolder($scope.holder, function(data){
	// 		$scope.change_holder = data;
	// 		if ($scope.change_holder == "Error: no tuna catch found"){
	// 			$("#error_holder").show();
	// 			$("#success_holder").hide();
	// 		} else{
	// 			$("#success_holder").show();
	// 			$("#error_holder").hide();
	// 		}
	// 	});
	// }

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllRequest = function(callback){

    	$http.get('/get_all_request/').success(function(output){
			callback(output)
		});
	}

	// factory.queryTuna = function(id, callback){
    // 	$http.get('/get_tuna/'+id).success(function(output){
	// 		callback(output)
	// 	});
	// }

	factory.buildRequest = function(data, callback){

		// data.location = data.longitude + ", "+ data.latitude;

		var request = data.timestamp + "-" + data.sender_name + "-" + data.sender_address + "-" + data.receive_name + "-" + data.receive_address + "-" + data.deliverer_name + "-" + data.price + "-" + data.status + "-" + data.code;

    	$http.get('/add_request/'+request).success(function(output){
			callback(output)
		});
	}

	// factory.changeHolder = function(data, callback){

	// 	var holder = data.id + "-" + data.name;

    // 	$http.get('/change_holder/'+holder).success(function(output){
	// 		callback(output)
	// 	});
	// }

	return factory;
});


