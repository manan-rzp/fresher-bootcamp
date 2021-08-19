<?php

use Illuminate\Support\Facades\Route;

use \App\Http\Controllers\MobileUserController;
/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});

Route::get('/users', [MobileUserController::class, 'getAllUsers']);
Route::post('/user', [MobileUserController::class, 'createUser']);
Route::get('/user', [MobileUserController::class, 'getUser']);
Route::delete('/user', [MobileUserController::class, 'deleteUser']);
