import { Injectable } from '@angular/core';
import {Component} from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent {
  rmsg: string;
  username: string;
  password: string;

  constructor(private authService: AuthService, private router: Router){
    this.rmsg="Register here";
    this.username="";
    this.password="";
  }


  onSubmit(){
    if(this.password.length < 8){
      window.alert("Password is too short!");
    }
    /*Send login info to backend and check to make sure the new username doesn't already exist
    in the backend map. If it does, say that username already exists and navigate back to login. */
    /*If it doesn't exist send the login info to be added to the backend and redirect user back to login screen
    to enter their new credentials*/
  }
  goToLogin(){
    this.router.navigate(['']);
  }
}
