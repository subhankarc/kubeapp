'use strict';

var app = angular.module('simple', ['ngMaterial']);

app.controller('appController', ['$http', '$mdDialog', function ($http, $mdDialog) {
    var vm = this;
    vm.init = init;
    vm.deleteUser = deleteUser;
    vm.addUser = addUser;
    vm.editUser = editUser;

    function init() {
        var params = {
            url: '/api/users',
            method: 'GET',
            headers: {
                'Accept': 'application/json'
            }
        };
        $http(params)
            .then(function (res) {
                vm.users = res.data.users;
            }, function (err) {
                console.log('error', err);
            });
    }

    function deleteUser(inumber) {
        var params = {
            url: `/api/users/${inumber}`,
            method: 'DELETE',
            headers: {
                'Accept': 'application/json'
            }
        };
        $http(params)
            .then(function () {
                vm.init()
            }, function (err) {
                console.log(err);
            })
    }

    function addUser() {
        $mdDialog.show({
            templateUrl: '/static/addUser.dlg.html',
            controller: dialogController,
            clickOutsideToClose: true
        })

    }

    function editUser(inumber) {
        $mdDialog.show({
            templateUrl: '/static/editUser.dlg.html',
            controller: dialogController,
            locals: {inumber: inumber},
            clickOutsideToClose: true
        })

    }

    function dialogController($scope, $mdDialog, inumber) {
        $scope.cancel = function () {
            $mdDialog.cancel();
        }

        $scope.add = function () {
            var params = {
                url: 'api/users',
                method: 'POST',
                data: {
                    inumber: $scope.inumber,
                    firstname: $scope.firstname,
                    lastname: $scope.lastname,
                    alias: $scope.alias
                },
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                }
            };
            $http(params)
                .then(function (res) {
                    init();
                    $mdDialog.hide();
                }, function (err) {
                    console.log(err);
                })
        }

        $scope.edit = function () {
            var data = {};
            data.inumber = inumber;
            if ($scope.firstname && $scope.firstname !== '') {
                data.firstname = $scope.firstname;
            }
            if ($scope.lastname && $scope.lastname !== '') {
                data.lastname = $scope.lastname;
            }
            if ($scope.alias && $scope.alias !== '') {
                data.alias = $scope.alias;
            }
            console.log(data)
            var params = {
                url: `api/users/${inumber}`,
                method: 'PUT',
                data: data,
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                }
            };
            $http(params)
                .then(function (res) {
                    init();
                    $mdDialog.hide();
                }, function (err) {
                    console.log(err);
                })
        }
    }

}]);