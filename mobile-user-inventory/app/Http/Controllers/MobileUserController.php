<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\DB;
use Validator;

class MobileUserController extends Controller
{
    //
    public function getAllUsers(): JsonResponse
    {
        $users = DB::table('users')->get();
        return response()->json([
            'users'=>$users
        ]);
    }

    public function createUser(Request $request): JsonResponse
    {
        $rules = [
            'name' => 'required|max:255',
            'email' => 'required|regex:/^.+@.+\..+$/i',
            'mobile' => 'required|size:10'
        ];
        $validator = Validator :: make($request->all(),$rules);
        if($validator->fails()){
            return response()->json($validator->errors(),400);
        }

        $name = \request('name');
        $email = \request('email');
        $mobile = \request('mobile');

        $data = array('name'=>$name,'email'=>$email,'mobile'=>$mobile); 
        $id = DB::table('users')->insertGetId($data);
        return response()->json([
            'id'=>$id,
            'name'=>$name,
            'email'=>$email,
            'mobile'=>$mobile]);
    }

    public function getUser(Request $request): JsonResponse
    {
        $type = $request->query('type');
        
        if($type == 'name'){
            $name = $request->query('name');
            $users = DB::table('users')->where('name',$name)->get();
            return response()->json([
                'users'=>$users,
            ]);
        }
        elseif($type == 'email'){
            $email = $request->query('email');
            $users = DB::table('users')->where('email',$email)->get();
            return response()->json([
                'users'=>$users,
            ]);
        }
        elseif($type = 'mobile'){
            $mobile = $request->query('mobile');
            $users = DB::table('users')->where('mobile',$mobile)->get();
            return response()->json([
                'users'=>$users,
            ]);
        }
        return response()->json([
            'message' => 'Invaild User Request'
        ]);
    }
    
    public function deleteUser(Request $request): JsonResponse
    {
        $type = $request->query('type');

        if($type == 'mobile'){
            $mobile = $request->query('mobile');
            DB::table('users')->where('mobile',$mobile)->delete();
            return response()->json([
                'message' => 'Successfully deleted user'
            ]);
        }
        elseif($type = 'name'){
            $name = $request->query('name');
            DB::table('users')->where('name',$name)->delete();
            return response()->json([
                'message' => 'Successfully deleted user'
            ]);
        }
        elseif($type = 'email'){
            $email  =$request->query('email');
            DB::table('users')->where('email',$email)->delete();
            return response()->json([
                'message' => 'Successfully deleted user'
            ]);
        }
        return response()->json([
            'message' => 'Invaild User Request'
        ]);
    }
}
